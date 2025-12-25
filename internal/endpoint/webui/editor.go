package webui

import (
	"fmt"
	"net/http"
)

func loadEditorFunc(w http.ResponseWriter, r *http.Request) {
	err := UIBundle.Render("editor", w, nil)
	if err != nil {
		SendFailed(w, fmt.Sprintf("failed to load %q", "editor"))
	}
}
