package form

import (
	"coven/internal/cards"
	"coven/internal/endpoint/webui"
	"coven/internal/utils"
	"fmt"
	"log/slog"
	"net/http"
)

func cardHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		cardType := r.FormValue("cardType")
		if len(cardType) == 0 {
			webui.SendFailed(w, "Нужно выбрать тип карты")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		dataFormat := r.FormValue("dataType")
		data := r.FormValue("data")
		name := r.FormValue("cardName")
		if !utils.IsInValidLenghtRange(name) {
			webui.SendFailed(w, "Название не может быть пустым или содержать больше 247 символов")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !utils.IsValidFileName(name) {
			webui.SendFailed(w, `Название карты содержит недопустимые символы: \/:*?"<>|`)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		completeName := utils.ReplaceAllSpaces(name)
		isExists, err := cards.IsCardExists(cardType, completeName)
		if err != nil {
			webui.SendFailed(w, "При сохранении произошла внутренняя ошибка сервера")
			slog.Error("failed to check if card exists", "card name", name)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if isExists {
			webui.SendFailed(w, fmt.Sprintf("карта с названием %q уже существует", name))
			webui.UIBundle.Render("override_card", w, nil)
			w.WriteHeader(http.StatusConflict)
			return
		}
		err = cards.SaveCard(cardType, completeName, dataFormat, data)
		if err != nil {
			webui.SendFailed(w, fmt.Sprintf("failed to save card %s", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		webui.SendSucces(w, "Карта успешно создана")
	case "PATCH":
		cardType := r.FormValue("cardType")
		if len(cardType) == 0 {
			webui.SendFailed(w, "Нужно выбрать тип карты")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		dataFormat := r.FormValue("dataType")
		data := r.FormValue("data")
		name := r.FormValue("cardName")
		if !utils.IsInValidLenghtRange(name) {
			webui.SendFailed(w, "Название не может быть пустым или содержать больше 247 символов")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !utils.IsValidFileName(name) {
			webui.SendFailed(w, `Название карты содержит недопустимые символы: \/:*?"<>|`)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		completeName := utils.ReplaceAllSpaces(name)
		isExists, err := cards.IsCardExists(cardType, completeName)
		if err != nil {
			webui.SendFailed(w, "При сохранении произошла внутренняя ошибка сервера")
			slog.Error("failed to check if card exists", "card name", name)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !isExists {
			webui.SendFailed(w, "Пока ты думал, блять, карту уже удалили... жми просто сохранить")
			w.WriteHeader(http.StatusConflict)
			return
		}
		err = cards.SaveCard(cardType, completeName, dataFormat, data)
		if err != nil {
			webui.SendFailed(w, fmt.Sprintf("failed to save card %s", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		webui.SendSucces(w, fmt.Sprintf("Карта успешно %q перезаписана", name))
	case "DELETE":
		w.WriteHeader(http.StatusNotImplemented)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
