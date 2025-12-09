package projection

type CardViewSkeletProj struct {
	Chapters []ChapterProj
}

type ChapterProj struct {
	Title string
	Cards []CompletCardViewProj
}

type CompletCardViewProj struct {
	Name string
}
