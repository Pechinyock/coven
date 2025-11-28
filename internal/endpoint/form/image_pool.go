package form

import (
	"coven/internal/cards"
	"coven/internal/endpoint/webui"
	"coven/internal/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func imagePoolFileFunc(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		uploadImage(w, r)
		return
	case "GET":
		w.WriteHeader(http.StatusNotImplemented)
		return
	case "DELETE":
		w.WriteHeader(http.StatusNotImplemented)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func uploadImage(w http.ResponseWriter, r *http.Request) {
	groupName := r.FormValue("group")
	if groupName == "" {
		webui.SendFailed(w, "Выбрите группу изображения")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, isDefined := cards.CardTypes[groupName]
	if !isDefined {
		webui.SendFailed(w, "Неизвестная группа изображений")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		webui.SendFailed(w,
			fmt.Sprintf("При загрузке файла возникла ошибка на стороне сервера: %s",
				err.Error(),
			))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	if file == nil {
		webui.SendFailed(w, "Выбирите файл")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	overrideName := r.FormValue("file-name")
	ext := filepath.Ext(handler.Filename)
	if !utils.IsExtension(ext, ".png") {
		webui.SendFailed(w, "Поддерживаются только png файлы")
		return
	}
	var fileName string
	if overrideName == "" {
		fileName = handler.Filename
	} else {
		fileName = fmt.Sprintf("%s.%s", overrideName, ext)
	}
	if !utils.IsValidPath(fileName) {
		webui.SendFailed(w, `Недопустимые символы в имени файла. Убедитесь, что название не содержит: < > : \" | ? *`)
		return
	}
	fullPath := filepath.Join(cards.ImagePool, groupName, fileName)
	dst, err := os.Create(fullPath)
	if err != nil {
		webui.SendFailed(w,
			fmt.Sprintf("При загрузке файла возникла ошибка на стороне сервера: %s",
				err.Error(),
			))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		webui.SendFailed(w,
			fmt.Sprintf("При загрузке файла возникла ошибка на стороне сервера: %s",
				err.Error(),
			))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	webui.SendSucces(w, fmt.Sprintf("Файл %s успешно загружен", fileName))
}
