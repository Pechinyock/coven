package webui

import (
	"coven/internal/cards"
	"net/http"
)

var modalsMap = map[string]func(string, http.ResponseWriter){
	"add_image":   addImageModal,
	"create_card": createCardModal,
}

func addImageModal(templName string, w http.ResponseWriter) {
	err := uiBundle.Render(templName, w, cards.CardTypes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func createCardModal(templName string, w http.ResponseWriter) {
	err := uiBundle.Render(templName, w, cards.CardTypes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
