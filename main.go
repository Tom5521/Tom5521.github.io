package main

import (
	"bufio"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

//go:embed pages
var pagesFS embed.FS

type Page struct {
	Title   string
	Content string
	File    string
}

func (p Page) String() string {
	var txt = `Title:%s
File:%s
Content:
%s`
	return fmt.Sprintf(txt, p.Title, p.File, p.Content)
}

var pages []Page

func LoadPages() (pages []Page, err error) {
	var dirs []os.DirEntry
	dirs, err = pagesFS.ReadDir("pages")
	if err != nil {
		return
	}

	for _, dir := range dirs {
		var (
			npage Page
			f     fs.File
		)

		npage.File = "pages/" + dir.Name()
		f, err = pagesFS.Open(npage.File)
		if err != nil {
			continue
		}

		scanner := bufio.NewScanner(f)
		var hasTitle bool
		for scanner.Scan() {
			txt := scanner.Text()
			if strings.Contains(txt, "# ") && !hasTitle {
				parts := strings.SplitN(txt, " ", 2)
				npage.Title = parts[1]
				hasTitle = true
				continue
			}
			npage.Content += txt + "\n"
		}
		pages = append(pages, npage)
		f.Close()
	}

	return
}

func main() {
	var err error
	pages, err = LoadPages()
	if err != nil {
		log.Println(err)
	}

	app := app.New()
	defer app.Run()

	w := app.NewWindow("My blog")

	content := widget.NewRichTextFromMarkdown(pages[0].Content)

	pagesList := widget.NewList(
		func() int {
			return len(pages)
		},
		func() fyne.CanvasObject {
			return &widget.Label{}
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(pages[lii].Title)
		},
	)
	pagesList.OnSelected = func(id widget.ListItemID) {
		content.ParseMarkdown(pages[id].Content)
	}

	mainBox := container.NewHSplit(pagesList, content)
	mainBox.SetOffset(0.2)
	w.SetContent(mainBox)
	w.Show()
}
