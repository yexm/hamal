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

// setOutputCmd represents the setOutput command
var setOutputCmd = &cobra.Command{
	Use:   "set-output",
	Short: "Set config doutput",
	Example: `
# set config dinput
hamal set-output -u gxd -p 123 -n name -r registry-vpc.gxd.com

# set dockerhub
hamal set-output -u gxd -p 123 -n name -r registry-vpc.gxd.com -d
`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := Config{}
		conf.ReadYaml(cfgFile)
		conf.Doutput = Doutput{}
		conf.Doutput.User = ouser
		conf.Doutput.Pass = opass
		conf.Doutput.Repo = orepo
		conf.Doutput.Registry = oregistry
		isdocker, _ := cmd.Flags().GetBool("oisdohub")
		if isdocker {
			conf.Doutput.IsDockerHub = isdocker
		}
		conf.WriteYaml()
		fmt.Println("hamal config set doutput success! ")
	},
}

var ouser string
var opass string
var orepo string
var oregistry string

func init() {
	initCmd.AddCommand(setOutputCmd)
	setOutputCmd.Flags().StringVarP(&ouser, "user", "u", "", "dinput docker user")
	setOutputCmd.Flags().StringVarP(&opass, "pass", "p", "", "dinput docker pass")
	setOutputCmd.Flags().StringVarP(&orepo, "repo", "n", "", "dinput docker repo")
	setOutputCmd.Flags().StringVarP(&oregistry, "registry", "r", "", "dinput docker registry")
	setOutputCmd.Flags().BoolP("oisdohub", "d", false, "dinput docker isDockerhub")
	setOutputCmd.MarkFlagRequired("ouser")
	setOutputCmd.MarkFlagRequired("opass")
	setOutputCmd.MarkFlagRequired("orepo")
}
