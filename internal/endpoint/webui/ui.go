package webui

import (
	"coven/internal/cards"
	"coven/internal/endpoint"
	shareddirs "coven/internal/endpoint/shared_dirs"
	"coven/internal/projection"
	"coven/internal/ui"
	"fmt"
	"log/slog"
	"net/http"
	"path"
	"strings"
)

const UIPrefix = "/ui"

var UIBundle ui.WebUIBundle

func SetUIBundle(newBundle ui.WebUIBundle) {
	slog.Info("ui builde has been set")
	UIBundle = newBundle
}

func GetUIEndpoints() []endpoint.Endpoint {
	return []endpoint.Endpoint{
		{
			Path:    "/",
			Methods: []string{"GET"},
			Secure:  true,
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				err := UIBundle.Render("main", w, nil)
				if err != nil {
					SendFailed(w, fmt.Sprintf("failed to load %q", "main"))
				}
			},
		},
		{
			Path:    "/login",
			Methods: []string{"GET"},
			Secure:  false,
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				UIBundle.Render("login_screen", w, nil)
			},
		},
		{
			Path:    path.Join(UIPrefix, "main-menu"),
			Methods: []string{"GET"},
			Secure:  true,
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				UIBundle.Render("menu", w, nil)
			},
		},
		{
			Path:    path.Join(UIPrefix, "coven"),
			Methods: []string{"GET"},
			Secure:  true,
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				err := UIBundle.Render("coven", w, nil)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			},
		},
		{
			Path:    path.Join(UIPrefix, "modal-window"),
			Methods: []string{"GET"},
			Secure:  true,
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				err := UIBundle.Render("modal_window", w, nil)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			},
		},
		{
			Path:    path.Join(UIPrefix, "modal-body", "{modalName}"),
			Methods: []string{"GET"},
			Secure:  true,
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				modalName := r.PathValue("modalName")
				if modalName == "" {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				templName := strings.ReplaceAll(modalName, "-", "_")
				renderFunc, defined := modalsMap[templName]
				if !defined {
					message := fmt.Sprintf("modal winodw not found %q", modalName)
					SendFailed(w, message)
					slog.Error(message)
					w.WriteHeader(http.StatusNotFound)
					return
				}
				renderFunc(templName, w)
			},
		},
		{
			Path:    path.Join(UIPrefix, "image-pool", "{poolName}"),
			Methods: []string{"GET"},
			Secure:  true,
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				poolName := r.PathValue("poolName")
				if poolName == "" {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				poolNameLower := strings.ToLower(poolName)
				_, defined := cards.CardTypes[poolNameLower]
				if !defined {
					slog.Error("unknown type of image pool")
					w.WriteHeader(http.StatusNotFound)
					return
				}
				prewiewData, err := loadImagesPrewiewData(poolName)
				if err != nil {
					slog.Error("failed to load image pool view", "error message", err.Error())
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				if prewiewData == nil {
					slog.Info("there's no images to display")
					w.WriteHeader(http.StatusNoContent)
					w.Write([]byte("Картинки данного типа еще не загружены"))
					return
				}
				baseUriPath := shareddirs.ImagePoolDirPath.Uri
				poolView := projection.ImageViewProj{
					BasePath:  baseUriPath,
					FileGroup: poolNameLower,
					Images:    prewiewData,
				}

				err = UIBundle.Render("select_image", w, poolView)
				if err != nil {
					slog.Error("failed to render image pool", "group", poolName,
						"error message", err.Error(),
					)
					errMsg := fmt.Sprintf("Не удалось отобразить картинки для группы: %s", poolName)
					UIBundle.Render("alert", w, projection.AlertProj{
						Message: errMsg,
						Type:    "danger",
					})
					return
				}
			},
		},
		{
			Path:    "/ui/generated-cards-view",
			Methods: []string{"GET"},
			Secure:  true,
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				ch := projection.CardViewSkeletProj{
					Chapters: cards.CardTypes,
				}
				err := UIBundle.Render("generated_cards_view", w, ch)
				if err != nil {
					SendFailed(w, fmt.Sprintf("failed to load %q", "generated_cards_view"))
				}
			},
		},
		{
			Path:    "/ui/chapter/{chapterName}",
			Methods: []string{"GET"},
			Secure:  true,
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				chapterName := r.PathValue("chapterName")
				_, exists := cards.CardTypes[chapterName]
				if !exists {
					SendFailed(w, fmt.Sprintf("неизвестный тип карт %s", chapterName))
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				pathToCards := path.Join(shareddirs.CompleteCardsDirPath.Path, chapterName)
				fileNames, err := cards.GetCardsFileNames(pathToCards, "png")
				if err != nil {
					SendFailed(w, fmt.Sprintf("при загрузке карт %s произошла ошибка сервера", chapterName))
					slog.Error("failed to load card file names", "card type name", chapterName)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				if len(fileNames) == 0 {
					SendFailed(w, "Таких карт ещё нет")
					w.WriteHeader(http.StatusNoContent)
					return
				}
				uriPaths := make([]string, len(fileNames))
				uriRoot := path.Join(shareddirs.CompleteCardsDirPath.Uri, chapterName)
				for i, c := range fileNames {
					uriPaths[i] = path.Join(uriRoot, c)
				}

				tmp := struct {
					Cards []string
				}{
					Cards: uriPaths,
				}
				w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
				err = UIBundle.Render("complete_cards_chapter", w, tmp)
				if err != nil {
					SendFailed(w, err.Error())
					return
				}
			},
		},
		{
			Path:    "/ui/remote-repo",
			Methods: []string{"GET"},
			Secure:  true,
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				switch r.Method {
				case "GET":
					getChanges(w)
				case "POST":
					postChanges(w)
				default:
					w.WriteHeader(http.StatusMethodNotAllowed)
				}
			},
		},
		{
			Path:    "/ui/pull-data",
			Methods: []string{"GET"},
			Secure:  true,
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				switch r.Method {
				case "GET":
					pullChanges(w)
				default:
					w.WriteHeader(http.StatusMethodNotAllowed)
				}
			},
		},
		{
			Path:        "/editor",
			Methods:     []string{"GET"},
			Secure:      true,
			HandlerFunc: loadEditorFunc,
		},
		{
			Path:    "/ui/card-types",
			Methods: []string{"GET"},
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				UIBundle.Render("card_types_selection", w, cards.CardTypes)
			},
		},
	}
}
