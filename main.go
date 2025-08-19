package main

import (
	card "coven/internal/cards"
	"fmt"
	"path"
)

func main() {
	// curses := card.LoadCardsDataFromJson[card.CurseCardData]("curses")
	// ingredients := card.LoadCardsDataFromJson[card.IngredientCardPiceData]("ingredients")
	// spells := card.LoadCardsDataFromJson[card.SpellCardData]("spells")
	card.LoadAllTemplates()
	chars := card.LoadCardsDataFromJson[card.CaracterCardData]("characters")
	for i, c := range chars {
		nm := fmt.Sprintf("%d.html", i)
		relOutPath := path.Join("characters", nm)
		card.GenerateSingleCardHtml(c, "character-card", relOutPath)
	}
	spells := card.LoadCardsDataFromJson[card.SpellCardData]("spells")
	for i, s := range spells {
		nm := fmt.Sprintf("%d.html", i)
		relOutPath := path.Join("spells", nm)
		card.GenerateSingleCardHtml(s, "spell-card", relOutPath)
	}
}
