package card

type SpellCardData struct {
	Name           string `json:"Name"`
	Description    string `json:"Description"`
	DecorationText string `json:"DecorationText"`
	ImgPath        string `json:"ImgPath"`
	Power          int    `json:"Power"`
}
