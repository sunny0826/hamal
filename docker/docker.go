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
		Registry string // Docker registry address
		Username string // Docker registry username
		Password string // Docker registry password
		Email    string // Docker registry email
	}

	// Build defines Docker build parameters.
	Build struct {
		Repo        string   // Docker repo:tag
	}

	// Plugin defines the Docker plugin parameters.
	Plugin struct {
		Login   Login  // Docker login configuration
		Build   Build  // Docker build configuration
		Dryrun  bool   // Docker push is skipped
		Cleanup bool   // Docker purge is enabled
	}
)

// pull docker images
func (p Plugin) Pull() (string, error) {
	if p.Login.Password != "" {
		cmd := commandLogin(p.Login)
		err := cmd.Run()
		if err != nil {
			return "", fmt.Errorf("Error authenticating: %s", err)
		}
	} else {
		fmt.Println("Registry credentials not provided. Guest mode enabled.")
	}
	var cmds []*exec.Cmd
	//cmds = append(cmds, commandVersion()) // docker version
	//cmds = append(cmds, commandInfo())    // docker info
	url := fmt.Sprintf("%s/%s", p.Login.Registry, p.Build.Repo)
	cmds = append(cmds, commandPull(url))
	// execute all commands in batch mode.
	ExecCommand(cmds)
	return url, fmt.Errorf("Error pull")
}

// change tag and push images
func (p Plugin) ChangeTagAndPush(url string) (string, error) {
	pushUrl := fmt.Sprintf("%s/%s", p.Login.Registry, p.Build.Repo)
	var cmds []*exec.Cmd
	// change tag
	cmds = append(cmds, changeCommandTag(url, pushUrl)) // docker tag
	// push image
	cmds = append(cmds, changeCommandPush(pushUrl)) // docker tag
	ExecCommand(cmds)
	return pushUrl, fmt.Errorf("Error change tag or push image")
}

// remove image
func (p Plugin) CleanImages(PullUrl string, PushUrl string) error {
	var cmds []*exec.Cmd
	if p.Cleanup {
		cmds = append(cmds, commandRmi(PullUrl)) // docker pull rmi
		cmds = append(cmds, commandRmi(PushUrl)) // docker push rmi
		cmds = append(cmds, commandPrune())      // docker system prune -f
	} else {
		return fmt.Errorf("Cleanup is false")
	}
	ExecCommand(cmds)
	return nil
}

// batch run command
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
func isCommandPull(args []string) bool {
	return len(args) > 2 && args[1] == "pull"
}

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
func commandVersion() *exec.Cmd {
	return exec.Command(dockerExe, "version")
}

// helper function to create the docker info command.
func commandInfo() *exec.Cmd {
	return exec.Command(dockerExe, "info")
}