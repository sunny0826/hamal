package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/sunny0826/hamal/docker"
)

var repo string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start syncing mirror",
	Long: `For details, please see: https://github.com/sunny0826/hamal.git

example:
hamal run -r guoxudongdocker/drone-dingtalk:latest`,
	Run: func(cmd *cobra.Command, args []string) {

		output := viper.GetStringMapString("doutput")
		input := viper.GetStringMapString("dinput")
		// 输出仓库
		outrepo := output["registry"]
		outuser := output["user"]
		outpass := output["pass"]
		// 输入仓库
		inregistry := input["registry"]
		inuser := input["user"]
		inpass := input["pass"]

		inplugin := docker.Plugin{
			Login: docker.Login{
				Registry: inregistry,
				Username: inuser,
				Password: inpass,
			},
			Build: docker.Build{
				Repo: repo,
			},
			Cleanup: true,
		}
		outplugin := docker.Plugin{
			Login: docker.Login{
				Registry: outrepo,
				Username: outuser,
				Password: outpass,
			},
			Build: docker.Build{
				Repo: repo,
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
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&repo, "repontag", "r", "", "docker repo:tag")
}
