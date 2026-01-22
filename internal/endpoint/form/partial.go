package form

import (
	shareddirs "coven/internal/endpoint/shared_dirs"
	"coven/internal/endpoint/webui"
	"coven/internal/utils"
	"encoding/base64"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func HandlePartial(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		{
			saveNew(w, r)
		}
	case http.MethodPatch:
		{
			override(w, r)
		}
	case http.MethodGet:
		{
			getPartial(w, r)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getPartial(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		slog.Error("failed to get partial the name is empty")
		return
	}
	partialJsonPath := path.Join(shareddirs.PartialsDirPath.Path,
		"json",
		fmt.Sprintf("%s.json", name),
	)

	if !utils.IsFileExists(partialJsonPath) {
		w.WriteHeader(http.StatusNotFound)
		slog.Error("failed to find partial json", "name:", name)
		return
	}
	data, err := os.ReadFile(partialJsonPath)
	if err != nil {
		slog.Error("failed to read partial json file", "error message:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func override(w http.ResponseWriter, r *http.Request) {
	pngData := r.FormValue("pngData")
	if pngData == "" {
		w.WriteHeader(http.StatusBadRequest)
		webui.SendFailed(w, "Картинка пустая")
		return
	}
	jsonData := r.FormValue("jsonData")
	if jsonData == "" {
		w.WriteHeader(http.StatusBadRequest)
		webui.SendFailed(w, "json пустой")
		return
	}
	name := r.FormValue("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		webui.SendFailed(w, "Название не может быть пустым")
		return
	}
	if !utils.IsValidFileName(name) || !utils.IsInValidLenghtRange(name) {
		webui.SendFailed(w, `Название содержит недопустимые символы: \/:*?"<>|`)
		webui.SendFailed(w, "Название содержать больше 247 символов")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	completeName := utils.ReplaceAllSpaces(name)
	err := savePreview(completeName, pngData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("failed to save partial png data", "name", name)
		webui.SendFailed(w, err.Error())
		return
	}
	err = saveJson(completeName, jsonData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("failed to save partial json data", "name", name)
		webui.SendFailed(w, err.Error())
		return
	}
	webui.SendSucces(w, "Перезаписалось вроде")
}

func saveNew(w http.ResponseWriter, r *http.Request) {
	pngData := r.FormValue("pngData")
	if pngData == "" {
		w.WriteHeader(http.StatusBadRequest)
		webui.SendFailed(w, "Картинка пустая")
		return
	}
	jsonData := r.FormValue("jsonData")
	if jsonData == "" {
		w.WriteHeader(http.StatusBadRequest)
		webui.SendFailed(w, "json пустой")
		return
	}
	name := r.FormValue("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		webui.SendFailed(w, "Название не может быть пустым")
		return
	}
	if !utils.IsValidFileName(name) || !utils.IsInValidLenghtRange(name) {
		webui.SendFailed(w, `Название содержит недопустимые символы: \/:*?"<>|`)
		webui.SendFailed(w, "Название содержать больше 247 символов")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	completeName := utils.ReplaceAllSpaces(name)
	isAlreadyExists, err := isExists(completeName)
	if err != nil {
		webui.SendFailed(w, "Внутреняя ошибка сервеа")
		slog.Error("failed to save partial", "error message", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isAlreadyExists {
		webui.SendFailed(w, fmt.Sprintf("Элемент с таким название уже существует: %s", name))
		w.WriteHeader(http.StatusConflict)
		webui.UIBundle.Render("override_partial", w, nil)
		return
	}
	err = savePreview(completeName, pngData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("failed to save partial png data", "name", name)
		webui.SendFailed(w, err.Error())
		return
	}
	err = saveJson(completeName, jsonData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error("failed to save partial json data", "name", name)
		webui.SendFailed(w, err.Error())
		return
	}
	webui.SendSucces(w, "Сохранилось вроде")
}

func isExists(name string) (bool, error) {
	pathTo, err := getPartialDirPath(shareddirs.PartialsDirPath.Path, "json")
	if err != nil {
		return true, err
	}
	fileFullPath := filepath.Join(pathTo, fmt.Sprintf("%s.json", name))
	files, err := filepath.Glob(fileFullPath)
	if err != nil {
		return true, nil
	}
	return len(files) > 0, nil
}

func getPartialDirPath(root, sub string) (string, error) {
	completePath := path.Join(root, sub)
	if !utils.IsDirExists(completePath) {
		err := utils.CreateDirIfNotExists(completePath)
		if err != nil {
			return "", err
		}
	}
	return completePath, nil
}

func savePreview(name, data string) error {
	partialsPreview, err := getPartialDirPath(shareddirs.PartialsDirPath.Path, "preview")
	if err != nil {
		return err
	}
	pngData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return err
	}
	fullPath := path.Join(partialsPreview, fmt.Sprintf("%s.png", name))
	return os.WriteFile(fullPath, pngData, 0644)
}

func saveJson(name, data string) error {
	partialsJson, err := getPartialDirPath(shareddirs.PartialsDirPath.Path, "json")
	if err != nil {
		return err
	}
	fullPath := path.Join(partialsJson, fmt.Sprintf("%s.json", name))
	return os.WriteFile(fullPath, []byte(data), 0644)
}
