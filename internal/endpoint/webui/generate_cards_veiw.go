package webui

import (
	"coven/internal/cards"
	"coven/internal/projection"
)

func getChapters() projection.CardViewSkeletProj {
	proj := projection.CardViewSkeletProj{}
	proj.Chapters = make([]projection.ChapterProj, len(cards.CardTypes))
	for el := range cards.CardTypes {
		proj.Chapters = append(proj.Chapters, projection.ChapterProj{
			Title: el,
			Cards: []projection.CompletCardViewProj{
				{
					Name: "card",
				},
			},
		})
	}
	return proj
}
