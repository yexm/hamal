package cmd

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/yexm/hamal/docker"
)

var name string
var rename string

type (
	Config struct {
		Author  string  `yaml:"author"`
		License string  `yaml:"license"`
		Dinput  Dinput  `yaml:"dinput"`
		Doutput Doutput `yaml:"doutput"`
	}
	Dinput struct {
		Registry    string `yaml:"registry,omitempty"`
		Repo        string `yaml:"repo"`
		User        string `yaml:"user"`
		Pass        string `yaml:"pass"`
		IsDockerHub bool   `yaml:"isDockerhub,omitempty"`
	}
	Doutput struct {
		Registry    string `yaml:"registry,omitempty"`
		Repo        string `yaml:"repo"`
		User        string `yaml:"user"`
		Pass        string `yaml:"pass"`
		IsDockerHub bool   `yaml:"isDockerhub,omitempty"`
		ImageName   string `yaml:"imageName,omitempty"`
	}
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start syncing mirror",
	Long: `For details, please see: https://github.com/yexm/hamal.git
`,
	Example:`
# sync docker image
hamal run -n drone-dingtalk:latest

# sync docker image and rename that image name or tag
hamal run -n drone-dingtalk:latest -r drone-test:v1.0
`,
	Run: func(cmd *cobra.Command, args []string) {

		hamalconfig := Config{}
		hamalconfig.ReadYaml(cfgFile)

		// 输出仓库
		outregistry := hamalconfig.Doutput.Registry
		outrepo := hamalconfig.Doutput.Repo
		if outrepo == "" {
			fmt.Println("Please enter the <doutput><repo> field in the configuration file(default is $HOME/.hamal/config)!")
			os.Exit(1)
		}
		outuser := hamalconfig.Doutput.User
		outpass := hamalconfig.Doutput.Pass
		outhub := hamalconfig.Doutput.IsDockerHub

		// 输入仓库
		inregistry := hamalconfig.Dinput.Registry
		inrepo := hamalconfig.Dinput.Repo
		if inrepo == "" {
			fmt.Println("Please enter the <dinput><repo> field in the configuration file(default is $HOME/.hamal/config)!")
			os.Exit(1)
		}
		inuser := hamalconfig.Dinput.User
		inpass := hamalconfig.Dinput.Pass
		inhub := hamalconfig.Dinput.IsDockerHub

		// 组合参数
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

		var outname string
		if rename != "" {
			fmt.Printf("rename <<%s>> to <<%s>>\n", name, rename)
			outname = rename
		} else {
			outname = name
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
				Name: outname,
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
	runCmd.Flags().StringVarP(&name, "name", "n", "", "docker image name:tag")
	runCmd.Flags().StringVarP(&rename, "rename", "r", "", "rename docker image name:tag")
	runCmd.MarkFlagRequired("name")
}

func (c *Config) ReadYaml(f string) {
	buffer, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = yaml.Unmarshal(buffer, &c)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func (c *Config) WriteYaml() {
	buffer, err := yaml.Marshal(&c)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = ioutil.WriteFile(cfgFile, buffer, 0777)
	if err != nil {
		fmt.Println(err.Error())
	}
}
