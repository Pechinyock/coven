package webui

import (
	"coven/internal/cards"
	shareddirs "coven/internal/endpoint/shared_dirs"
	"coven/internal/projection"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

func getChapters() projection.CardViewSkeletProj {
	proj := projection.CardViewSkeletProj{}
	//proj.Chapters = map[string]string{"characters": "персы"}
	proj.Chapters = cards.CardTypes
	return proj
}

// func getGeneratedCards(typeName string) (*projection.ChapterProj, error) {
// 	if _, exists := cards.CardTypes[typeName]; !exists {
// 		return nil, fmt.Errorf("card type %q is not exists", typeName)
// 	}
// 	pathToCardsOut := filepath.Join(shareddirs.CompleteCardsDirPath.Path, typeName)
// 	files, err := os.ReadDir(pathToCardsOut)
// 	if err != nil {
// 		return nil, err
// 	}
// 	cards := []projection.CompletCardViewProj{}
// 	for _, f := range files {
// 		theName := f.Name()
// 		iframePath := getCompeteCardPath(typeName, theName)
// 		cards = append(cards, projection.CompletCardViewProj{
// 			Name:       theName,
// 			IFramePath: iframePath,
// 		})
// 	}
// 	result := projection.ChapterProj{
// 		Cards: cards,
// 	}
// 	return &result, nil
// }

// func getCompeteCardPath(cardType, cardName string) string {
// 	/* scheme, address, port, imagepool base uri, group name, selected image name*/
// 	fullCardHtmlRemotePath := fmt.Sprintf("%s://%s:%d/%s/%s/%s",
// 		endpoint.Scheme,
// 		endpoint.Address,
// 		endpoint.Port,
// 		shareddirs.CompleteCardsDirPath.Uri,
// 		cardType,
// 		cardName,
// 	)
// 	return fullCardHtmlRemotePath
// }

func getChapterCardsData(chapterName string) (any, error) {
	switch chapterName {
	case "characters":
		return getCharactersChapter()
	default:
		return nil, fmt.Errorf("card type %q is not exists", chapterName)
	}
}

func getCharactersChapter() (*projection.CharactersChapterProj, error) {
	pathToCharactersData := filepath.Join(shareddirs.CardsJsonDataDirPath.Path, "characters")
	files, err := os.ReadDir(pathToCharactersData)
	if err != nil {
		return nil, err
	}
	filesCount := len(files)
	if filesCount == 0 {
		return nil, nil
	}
	result := projection.CharactersChapterProj{
		Cards: []cards.Character{},
	}
	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".json") {
			slog.Warn(fmt.Sprintf("file %s has wrong extension", f.Name()))
			continue
		}
		fullPath := filepath.Join(pathToCharactersData, f.Name())
		slog.Debug(fmt.Sprintf("reading character card data from %s", fullPath))
		rawData, err := os.ReadFile(fullPath)
		if err != nil {
			slog.Error(fmt.Sprintf("failed to read character data %s", err.Error()))
			continue
		}
		var charData cards.Character
		err = json.Unmarshal(rawData, &charData)
		if err != nil {
			slog.Error(fmt.Sprintf("failed to unmarshal card data %s", err.Error()))
			continue
		}
		result.Cards = append(result.Cards, charData)
	}
	if len(result.Cards) == 0 {
		return nil, errors.New("failed to load characters json data")
	}
	return &result, nil
}
