package card

type CurseCardData struct {
	Name               string   `json:"Name"`
	DecorationText     string   `json:"DescriptionText"`
	Penalty            string   `json:"Penalty"`
	ImgPath            string   `json:"ImgPath"`
	Power              int      `json:"Power"`
	IngredientsToPruge []string `json:"IngredientsToPruge"`
}
