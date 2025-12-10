package projection

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
