package card

import (
	"context"
	"encoding/json"
	"html/template"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

const cardsDirName = "cards"
const generatedHtmlsDirName = "generated-htmls"

const cardPixelsWidth = 720
const cardPixelsHeight = 1280

var rootTempl *template.Template = template.New("root")

func LoadCardsDataFromJson[T any](typeName string) []T {
	appHome, err := os.Getwd()
	if err != nil {
		panic("failed to get app's current dir: " + err.Error())
	}
	pathToCards := path.Join(appHome, cardsDirName, typeName)
	searchPattern := path.Join(pathToCards, "*.json")
	jsons, err := filepath.Glob(searchPattern)
	if err != nil {
		panic("failed to get " + typeName + " jsons: " + err.Error())
	}

	var result []T = make([]T, len(jsons))

	for idx, pathToCard := range jsons {
		cureCardsBytes, err := os.ReadFile(pathToCard)
		if err != nil {
			panic("failed to load bytes:" + err.Error())
		}
		var cardData *T
		err = json.Unmarshal(cureCardsBytes, &cardData)
		if err != nil {
			panic("failed to parse json:" + err.Error())
		}
		result[idx] = *cardData
	}
	return result

}

func LoadAllTemplates() {
	appHome, err := os.Getwd()
	if err != nil {
		panic("failed to get app's current dir: " + err.Error())
	}
	templHomePath := path.Join(appHome, "templates", "*.html")
	_, err = rootTempl.ParseGlob(templHomePath)
	if err != nil {
		panic("failed to parse templates")
	}
}

func GenerateSingleCardHtml[T any](cardData T, templName, outFileName string) {
	appHome, err := os.Getwd()
	if err != nil {
		panic("failed to get app's current dir: " + err.Error())
	}

	genHtmls := path.Join(appHome, generatedHtmlsDirName, outFileName)
	outFile, err := os.Create(genHtmls)
	if err != nil {
		panic("failed to create file for template " + templName + ": " + err.Error())
	}
	err = rootTempl.ExecuteTemplate(outFile, templName, cardData)
	if err != nil {
		panic("failed to build template: " + err.Error())
	}
}

func GeneratePNGFromHtml(cardHtml, outputPath string) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buffer []byte
	err := chromedp.Run(ctx,
		emulation.SetDeviceMetricsOverride(cardPixelsWidth,
			cardPixelsHeight,
			1.0,
			false,
		),
		chromedp.Navigate("about:blank"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			frameTree, err := page.GetFrameTree().Do(ctx)
			if err != nil {
				return err
			}
			return page.SetDocumentContent(frameTree.Frame.ID, cardHtml).Do(ctx)
		}),
		chromedp.Sleep(2*time.Second),
		chromedp.FullScreenshot(&buffer, 100),
	)
	if err != nil {
		panic("falied to convert from html to png: " + err.Error())
	}

	err = os.WriteFile(outputPath, buffer, fs.ModePerm)
	if err != nil {
		panic("failed to wirte to file:" + err.Error())
	}
}

func GetCardsHtmlsPaths(typeName string) []string {
	appHome, err := os.Getwd()
	if err != nil {
		panic("failed to get app's current dir: " + err.Error())
	}
	pathToCards := path.Join(appHome, generatedHtmlsDirName, typeName)
	searchPattern := path.Join(pathToCards, "*.html")
	htmls, err := filepath.Glob(searchPattern)
	if err != nil {
		panic("failed to get generated htmls" + err.Error())
	}
	return htmls
}
