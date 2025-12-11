package webui

import (
	"coven/internal/git"
	"fmt"
	"net/http"
)

func getChanges(w http.ResponseWriter) {
	git.AddAll()
	gitStatusOutput, err := git.CheckStatus()
	if err != nil {
		SendFailed(w, fmt.Sprintf("failed to check data: %s", err.Error()))
		return
	}
	uiBundle.Render("storage_status", w, gitStatusOutput)
}

func postChanges(w http.ResponseWriter) {
	git.AddAll()
	_, err := git.PushToRemote()
	if err != nil {
		SendFailed(w, fmt.Sprintf("failed to push data: %s", err.Error()))
		return
	}
	SendSucces(w, "succesfully push data")
}

func pullChanges(w http.ResponseWriter) {
	_, err := git.PullFromMain()
	if err != nil {
		SendFailed(w, fmt.Sprintf("failed to pull data: %s", err.Error()))
		return
	}
	SendSucces(w, "succesfully pull data")
}
