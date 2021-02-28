package docker

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type (
	// Login defines Docker login parameters.
	Login struct {
		Registry    string // Docker registry address
		Username    string // Docker registry username
		Password    string // Docker registry password
		Email       string // Docker registry email
		IsDockerhub bool   // DockerHub
	}

	// Build defines Docker build parameters.
	Build struct {
		Repo string // Docker repo
		Name string // Docker name:tag
	}

	// Plugin defines the Docker plugin parameters.
	Plugin struct {
		Login   Login // Docker login configuration
		Build   Build // Docker build configuration
		Dryrun  bool  // Docker push is skipped
		Cleanup bool  // Docker purge is enabled
	}
)

// Pull docker images
func (p Plugin) Pull() (string, error) {
	var cmds []*exec.Cmd
	//cmds = append(cmds, commandVersion()) // docker version
	//cmds = append(cmds, commandInfo())    // docker info
	var url string
	if p.Login.IsDockerhub {
		fmt.Println("Do not need Login.")
		url = fmt.Sprintf("%s/%s", p.Build.Repo, p.Build.Name)
	} else {
		cmd := commandLogin(p.Login)
		err := cmd.Run()
		if err != nil {
			return "", fmt.Errorf("Error authenticating: %s", err)
		}
		url = fmt.Sprintf("%s/%s/%s", p.Login.Registry, p.Build.Repo, p.Build.Name)
	}

	cmds = append(cmds, commandPull(url))
	// execute all commands in batch mode.
	err := ExecCommand(cmds)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return url, fmt.Errorf("Error pull")
}

// ChangeTagAndPush : Change tag and push images
func (p Plugin) ChangeTagAndPush(url string) (string, error) {
	var pushURL string
	if p.Login.IsDockerhub {
		fmt.Println("Do not need Login.")
		pushURL = fmt.Sprintf("%s/%s", p.Build.Repo, p.Build.Name)
	} else {
		cmd := commandLogin(p.Login)
		err := cmd.Run()
		if err != nil {
			return "", fmt.Errorf("Error authenticating: %s", err)
		}
		pushURL = fmt.Sprintf("%s/%s/%s", p.Login.Registry, p.Build.Repo, p.Build.Name)
	}
	var cmds []*exec.Cmd
	// change tag
	cmds = append(cmds, changeCommandTag(url, pushURL)) // docker tag
	// push image
	cmds = append(cmds, changeCommandPush(pushURL)) // docker tag
	err := ExecCommand(cmds)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pushURL, fmt.Errorf("Error change tag or push image")
}

// CleanImages & Remove image
func (p Plugin) CleanImages(PullURL string, PushURL string) error {
	var cmds []*exec.Cmd
	if p.Cleanup {
		cmds = append(cmds, commandRmi(PullURL)) // docker pull rmi
		cmds = append(cmds, commandRmi(PushURL)) // docker push rmi
		cmds = append(cmds, commandPrune())      // docker system prune -f
	} else {
		return fmt.Errorf("Cleanup is false")
	}
	err := ExecCommand(cmds)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return nil
}

// ExecCommand : Batch run command
func ExecCommand(cmds []*exec.Cmd) error {
	for _, cmd := range cmds {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		trace(cmd)

		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("Error authenticating: %s", err)
		}
	}
	return nil
}

const dockerExe = "/usr/local/bin/docker"

//const dockerdExe = "/usr/local/bin/dockerd"

// helper function to create the docker login command.
func commandLogin(login Login) *exec.Cmd {
	if login.Email != "" {
		return commandLoginEmail(login)
	}
	return exec.Command(
		dockerExe, "login",
		"-u", login.Username,
		"-p", login.Password,
		login.Registry,
	)
}

// helper function to create the docker tag command.
func changeCommandTag(source string, target string) *exec.Cmd {
	return exec.Command(
		dockerExe, "tag", source, target,
	)
}

// remove image
func commandRmi(tag string) *exec.Cmd {
	return exec.Command(dockerExe, "rmi", tag)
}

// push image
func changeCommandPush(url string) *exec.Cmd {
	return exec.Command(dockerExe, "push", url)
}

// docker system prune -f
func commandPrune() *exec.Cmd {
	return exec.Command(dockerExe, "system", "prune", "-f")
}

// trace writes each command to stdout with the command wrapped in an xml
// tag so that it can be extracted and displayed in the logs.
func trace(cmd *exec.Cmd) {
	fmt.Fprintf(os.Stdout, "+ %s\n", strings.Join(cmd.Args, " "))
}

// helper to check if args match "docker pull <image>"
//func isCommandPull(args []string) bool {
//	return len(args) > 2 && args[1] == "pull"
//}

func commandPull(repo string) *exec.Cmd {
	return exec.Command(dockerExe, "pull", repo)
}

func commandLoginEmail(login Login) *exec.Cmd {
	return exec.Command(
		dockerExe, "login",
		"-u", login.Username,
		"-p", login.Password,
		"-e", login.Email,
		login.Registry,
	)
}

// helper function to create the docker info command.
//func commandVersion() *exec.Cmd {
//	return exec.Command(dockerExe, "version")
//}

// helper function to create the docker info command.
//func commandInfo() *exec.Cmd {
//	return exec.Command(dockerExe, "info")
//}
