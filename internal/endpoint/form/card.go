package form

import (
	"coven/internal/cards"
	"coven/internal/endpoint/webui"
	"fmt"
	"net/http"
)

func cardHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		cardType := r.FormValue("cardType")
		dataFormat := r.FormValue("dataType")
		data := r.FormValue("data")
		err := cards.SaveCard(cardType, dataFormat, data)
		if err != nil {
			webui.SendFailed(w, fmt.Sprintf("failed to save card %s", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		webui.SendSucces(w, "Карта успешно создана")
	case "PUT":
		w.WriteHeader(http.StatusNotImplemented)
	case "DELETE":
		w.WriteHeader(http.StatusNotImplemented)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
