package card

type IngredientCardData struct {
	UpSideIngredient   IngredientCardPiceData
	DownSideIngredient IngredientCardPiceData
}

type IngredientCardPiceData struct {
	Name           string `json:"Name"`
	DecorationText string `json:"DescriptionText"`
	ImgPath        string `json:"ImgPath"`
}
