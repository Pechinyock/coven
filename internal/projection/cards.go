package projection

import "coven/internal/cards"

type CardViewSkeletProj struct {
	Chapters map[string]string
}

type ChapterProj struct {
	Cards []CompletCardViewProj
}

type CompletCardViewProj struct {
	Name       string
	IFramePath string
}

type CharactersChapterProj struct {
	Cards []cards.Character
}
