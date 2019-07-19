package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/sunny0826/hamal/docker"
)

var name string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start syncing mirror",
	Long: `For details, please see: https://github.com/sunny0826/hamal.git

example:
hamal run -n drone-dingtalk:latest`,
	Run: func(cmd *cobra.Command, args []string) {

		output := viper.GetStringMapString("doutput")
		input := viper.GetStringMapString("dinput")
		// 输出仓库
		outregistry := output["registry"]
		outrepo := output["repo"]
		if outrepo == "" {
			fmt.Println("Please enter the <doutput><repo> field in the configuration file(default is $HOME/.hamal/config.yaml)!")
			os.Exit(1)
		}
		outuser := output["user"]
		outpass := output["pass"]
		outhub, _ := strconv.ParseBool(output["isdockerhub"])

		// 输入仓库
		inregistry := input["registry"]
		inrepo := input["repo"]
		if inrepo == "" {
			fmt.Println("Please enter the <dinput><repo> field in the configuration file(default is $HOME/.hamal/config.yaml)!")
			os.Exit(1)
		}
		inuser := input["user"]
		inpass := input["pass"]
		inhub, _ := strconv.ParseBool(input["isdockerhub"])

		inplugin := docker.Plugin{
			Login: docker.Login{
				Registry:    inregistry,
				Username:    inuser,
				Password:    inpass,
				IsDockerhub: inhub,
			},
			Build: docker.Build{
				Repo: inrepo,
				Name: name,
			},
			Cleanup: true,
		}
		outplugin := docker.Plugin{
			Login: docker.Login{
				Registry:    outregistry,
				Username:    outuser,
				Password:    outpass,
				IsDockerhub: outhub,
			},
			Build: docker.Build{
				Repo: outrepo,
				Name: name,
			},
			Cleanup: true,
		}
		PullURL, _ := inplugin.Pull()
		PushURL, _ := outplugin.ChangeTagAndPush(PullURL)
		err := inplugin.CleanImages(PullURL, PushURL)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Sync success！")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&name, "name", "n", "", "docker name:tag")
}
