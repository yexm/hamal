package cmd

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

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var cfgFile string

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
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	home, err := homedir.Dir()
	haconfig := flag.String("haconfig", filepath.Join(home, ".hamal", "config"), "(optional) absolute path to the haconfig file")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", *haconfig, "config.yaml file (default is $HOME/.hamal)")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
