package webui

import (
	"coven/internal/cards"
	"coven/internal/endpoint"
	shareddirs "coven/internal/endpoint/shared_dirs"
	"coven/internal/projection"
	"fmt"
	"os"
	"path/filepath"
)

func getChapters() projection.CardViewSkeletProj {
	proj := projection.CardViewSkeletProj{}
	proj.Chapters = cards.CardTypes
	return proj
}

func getGeneratedCards(typeName string) (*projection.ChapterProj, error) {
	if _, exists := cards.CardTypes[typeName]; !exists {
		return nil, fmt.Errorf("card type %q is not exists", typeName)
	}
	pathToCardsOut := filepath.Join(shareddirs.CompleteCardsDirPath.Path, typeName)
	files, err := os.ReadDir(pathToCardsOut)
	if err != nil {
		return nil, err
	}
	cards := []projection.CompletCardViewProj{}
	for _, f := range files {
		theName := f.Name()
		iframePath := getCompeteCardPath(typeName, theName)
		cards = append(cards, projection.CompletCardViewProj{
			Name:       theName,
			IFramePath: iframePath,
		})
	}
	result := projection.ChapterProj{
		Cards: cards,
	}
	return &result, nil
}

func getCompeteCardPath(cardType, cardName string) string {
	/* scheme, address, port, imagepool base uri, group name, selected image name*/
	fullCardHtmlRemotePath := fmt.Sprintf("%s://%s:%d/%s/%s/%s",
		endpoint.Scheme,
		endpoint.Address,
		endpoint.Port,
		shareddirs.CompleteCardsDirPath.Uri,
		cardType,
		cardName,
	)
	return fullCardHtmlRemotePath
}
