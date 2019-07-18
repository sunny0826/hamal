/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
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
hamal run -r keking/test:2`,
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
		PullUrl, _ := inplugin.Pull()
		PushUrl, _ := outplugin.ChangeTagAndPush(PullUrl)
		inplugin.CleanImages(PullUrl, PushUrl)

	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&repo, "repontag", "r", "", "docker repo:tag")
}
