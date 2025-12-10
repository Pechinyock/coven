package git

import (
	"os"
	"os/exec"
	"os/user"
)

var GitDirPath string
var GitOriginUrl string
var WorkingBranch string
var MergePointBranch string

func GetDefaultBranchName() string {
	user, _ := user.Current()
	userName := user.Username
	return userName
}

func RunGitInDir(command string, args ...string) error {
	dir := GitDirPath
	cmd := exec.Command("git", append([]string{command}, args...)...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
