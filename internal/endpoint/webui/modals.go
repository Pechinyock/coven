package webui

import (
	"coven/internal/cards"
	"net/http"
)

var modalsMap = map[string]func(string, http.ResponseWriter){
	"add_image":      addImageModal,
	"remote_storage": remoteStorageModal,
}

func addImageModal(templName string, w http.ResponseWriter) {
	err := UIBundle.Render(templName, w, cards.CardTypes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func remoteStorageModal(templName string, w http.ResponseWriter) {
	err := UIBundle.Render(templName, w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
