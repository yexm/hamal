package cmd

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

import (
	"fmt"
	"github.com/spf13/cobra"
)

// setInputCmd represents the setInput command
var setInputCmd = &cobra.Command{
	Use:   "set-input",
	Short: "Set config dinput",
	Example: `
# set config dinput
hamal init set-input -u test -p test123 -n name -r registry.test.com

# set dockerhub
hamal init set-input -u test -p test123 -n name -d
`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := Config{}
		conf.ReadYaml(cfgFile)
		conf.Dinput = Dinput{}
		conf.Dinput.User = iuser
		conf.Dinput.Pass = ipass
		conf.Dinput.Repo = irepo
		conf.Dinput.Registry = iregistry
		isdocker, _ := cmd.Flags().GetBool("iisdohub")
		if isdocker {
			conf.Dinput.IsDockerHub = isdocker
		}
		conf.WriteYaml()
		fmt.Println("hamal config set dinput success! ")
	},
}

var iuser string
var ipass string
var irepo string
var iregistry string

func init() {
	initCmd.AddCommand(setInputCmd)
	setInputCmd.Flags().StringVarP(&iuser, "user", "u", "", "dinput docker user")
	setInputCmd.Flags().StringVarP(&ipass, "pass", "p", "", "dinput docker pass")
	setInputCmd.Flags().StringVarP(&irepo, "repo", "n", "", "dinput docker repo")
	setInputCmd.Flags().StringVarP(&iregistry, "registry", "r", "", "dinput docker registry")
	setInputCmd.Flags().BoolP("iisdohub", "d", false, "dinput docker isDockerhub")
	setInputCmd.MarkFlagRequired("iuser")
	setInputCmd.MarkFlagRequired("ipass")
	setInputCmd.MarkFlagRequired("irepo")
}
