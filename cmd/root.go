/*
Copyright © 2019 Guo Xudong

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
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"os"

	"github.com/spf13/viper"
)

var cfgFile string
var author string
var output map[string]string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hamal",
	Short: "Hamal is a tool for synchronizing images between two mirrored repositories.",
	Long: ` _   _                       _ 
| | | | __ _ _ __ ___   __ _| |
| |_| |/ _\ | '_ \ _ \ / _\ | |
|  _  | (_| | | | | | | (_| | |
|_| |_|\__,_|_| |_| |_|\__,_|_|

Hamal is a tool for synchronizing images between two mirrored repositories. 
You can synchronize mirrors between two private image repositories.

WARN:The docker must be installed locally.
Currently only Linux and MacOS are supported.`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		//greeting := "Hello"
		if viper.GetString("author") != "" {
			author = viper.GetString("author")
		}
		if viper.GetStringMap("output") != nil {
			output = viper.GetStringMapString("output")
		}
		fmt.Println(author)
		fmt.Println(output)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hamal/config.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	viper.AddConfigPath(home)
	viper.SetConfigName(".hamal/config")
	err = viper.ReadInConfig() // 搜索路径，并读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

