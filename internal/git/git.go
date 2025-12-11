package git

import (
	"errors"
	"fmt"
	"log/slog"
	"os/exec"
	"os/user"
	"regexp"
	"strings"
	"time"
)

var GitDirPath string
var GitOriginUrl string
var WorkingBranchName string
var MainBranchName string

const originName = "origin"

const (
	newFilePrefix  = "new file:"
	modefiedPrefix = "modified:"
	deletedPrefix  = "deleted:"
)

func GetDefaultBranchName() string {
	user, _ := user.Current()
	userName := user.Username
	re := regexp.MustCompile(`[^a-zA-Z0-9_-]`)
	userName = re.ReplaceAllString(userName, "")
	return userName
}

func PushToRemote() (string, error) {
	newDateTime := time.Now().Format("2006-01-02 15:04:05")
	user, _ := user.Current()
	userName := user.Username
	comment := fmt.Sprintf("%s. %s", newDateTime, userName)
	_, err := runGitInDir("commit", "-a", "-m", comment)
	if err != nil {
		return "", err
	}
	_, err = runGitInDir("push", originName, WorkingBranchName)
	if err != nil {
		return "", err
	}
	return "", nil
}

func CheckoutToMyBranch() (string, error) {
	out, err := runGitInDir("checkout", "-b", WorkingBranchName)
	if err != nil {
		return "", err
	}
	return out, err
}

func PullFromMain() (string, error) {
	out, err := runGitInDir("pull", originName, MainBranchName)
	if err != nil {
		return "", err
	}
	return out, err
}

func CheckStatus() ([]string, error) {
	output, err := runGitInDir("status")
	pretty := getStatusPretty(output)
	return pretty, err
}

func AddAll() (string, error) {
	output, err := runGitInDir("add", ".")
	return output, err
}

func getStatusPretty(rawOutput string) []string {
	var result []string
	lines := strings.Split(rawOutput, "\n")
	for _, line := range lines {
		if strings.Contains(line, newFilePrefix) {
			parts := strings.SplitN(line, newFilePrefix, 2)
			if len(parts) == 2 {
				fileName := strings.TrimSpace(parts[1])
				result = append(result, fmt.Sprintf("новый файл: %s", fileName))
			}
		} else if strings.Contains(line, modefiedPrefix) {
			parts := strings.SplitN(line, modefiedPrefix, 2)
			if len(parts) == 2 {
				fileName := strings.TrimSpace(parts[1])
				result = append(result, fmt.Sprintf("изменился: %s", fileName))
			}
		} else if strings.Contains(line, deletedPrefix) {
			parts := strings.SplitN(line, deletedPrefix, 2)
			if len(parts) == 2 {
				fileName := strings.TrimSpace(parts[1])
				result = append(result, fmt.Sprintf("удалён: %s", fileName))
			}
		}
	}
	return result
}

func runGitInDir(command string, args ...string) (string, error) {
	if WorkingBranchName == "" {
		return "[ERROR]", errors.New("failed to push data, the working branch is not specified")
	}
	if MainBranchName == "" {
		return "[ERROR]", errors.New("failed to push data, the main branch is not specified")
	}
	dir := GitDirPath
	cmd := exec.Command("git", append([]string{command}, args...)...)
	cmd.Dir = dir
	out, err := cmd.Output()
	if err != nil {
		slog.Error("failed to execute git cmd", "error message", err.Error())
	}
	output := string(out)
	return output, nil
}
