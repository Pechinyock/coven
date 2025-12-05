package form

import (
	"coven/internal/cards"
	"coven/internal/endpoint"
	shareddirs "coven/internal/endpoint/shared_dirs"
	"coven/internal/endpoint/webui"
	"fmt"
	"log/slog"
	"net/http"
)

const (
	characterGroupName = "characters"
)

func cardHandleFunc(w http.ResponseWriter, r *http.Request) {
	selectedCardType := r.FormValue("creating-card-type")
	if selectedCardType == "" {
		webui.SendFailed(w, "Тип карты не может быть пустым")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch r.Method {
	case "POST":
		creationMap := map[string]func(*http.Request) error{
			characterGroupName: prepareCharacterCreation,
		}
		err := creationMap[characterGroupName](r)
		if err != nil {
			webui.SendFailed(w, fmt.Sprintf("Не удалось создать карту %s", err.Error()))
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

func prepareCharacterCreation(r *http.Request) error {
	charName := r.FormValue("character-name")
	decorTxt := r.FormValue("character-description")
	role := r.FormValue("character-role")
	selectedImageName := r.FormValue("selected-character-image")

	/* scheme, address, port, imagepool base uri, group name, selected image name*/
	imgFullRemotePath := fmt.Sprintf("%s://%s:%d/%s/%s/%s",
		endpoint.Scheme,
		endpoint.Address,
		endpoint.Port,
		shareddirs.ImagePoolDirPath.Uri,
		characterGroupName,
		selectedImageName,
	)
	data := cards.Character{
		Name:           charName,
		DecorationText: decorTxt,
		Role:           role,
		ImgPath:        imgFullRemotePath,
	}
	err := cards.GenerateCard(characterGroupName,
		charName,
		shareddirs.CompleteCardsDirPath.Path,
		shareddirs.CardTemplatesDirPath.Path,
		data,
	)
	if err != nil {
		slog.Error("failed to create card", "card name", charName)
		return err
	}
	slog.Info("ready to create character")
	return nil
}
