package form

import (
	"coven/internal/endpoint/webui"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

func cardHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		createCard(w, r)
	case "PUT":
		w.WriteHeader(http.StatusNotImplemented)
	case "DELETE":
		w.WriteHeader(http.StatusNotImplemented)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func createCard(w http.ResponseWriter, r *http.Request) {
	selectedImageName := r.FormValue("selected-character-image")
	selectedCardType := r.FormValue("creating-card-type")
	err := distribute(selectedCardType, selectedImageName, w, r)
	if err != nil {
		webui.SendFailed(w, fmt.Sprintf("Не удалось создать карту: %s", err.Error()))
	}
}

func distribute(cardType, imageName string, w http.ResponseWriter, r *http.Request) error {
	if cardType == "" {
		return errors.New("card type is empty")
	}
	if imageName == "" {
		return errors.New("image name is empty")
	}
	switch cardType {
	case "characters":
		webui.SendFailed(w, "UNIMPLEMENTED")
		charName := r.FormValue("character-name")
		decorTxt := r.FormValue("character-description")
		role := r.FormValue("character-role")

		slog.Info(charName, decorTxt, role)

	case "spells":
		return errors.New("unimplemented spells")
	case "secrets":
		return errors.New("unimplemented secrets")
	case "curses":
		return errors.New("unimplemented curses")
	case "ingredients":
		return errors.New("unimplemented ingredients")
	case "potions":
		return errors.New("unimplemented potions")
	default:
		return errors.New("unknown card type")
	}

	return nil
}
