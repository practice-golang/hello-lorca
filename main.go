package main // import "hello-lorca"

import (
	"log"
	"net/url"

	"github.com/ncruces/zenity"
	"github.com/zserge/lorca"
)

const defaultPath = ``

func dlgInfo() {
	// exec.Command("zenity", "--question", "--title", "WTFG", "--text", "WTWTWTWT").Run()
	zenity.Info("All updates are complete.",
		zenity.Title("Information"),
		zenity.InfoIcon)

	// dlg, _ := zenity.Progress()
	// time.Sleep(time.Second * 1)
	// dlg.Complete()
	// dlg.Close()

}

func dlgBrowse() {
	zenity.SelectFile(
		zenity.Filename(defaultPath),
		zenity.FileFilters{
			{Name: "Go files", Patterns: []string{"*.go"}},
			{Name: "Web files", Patterns: []string{"*.html", "*.js", "*.css"}},
			{Name: "Image files", Patterns: []string{"*.png", "*.gif", "*.ico", "*.jpg", "*.webp"}},
		})
}

func main() {
	// Create UI with basic HTML passed via data URI
	ui, err := lorca.New("data:text/html,"+url.PathEscape(`
	<html>
		<head><title>Hello</title></head>
		<body>
			<h1>Hello, world!</h1>
			<a href="data:application/xml;charset=utf-8,your code here" download="filename.html">Save</a>
			<input type="file">
			<button onclick="save()">Save</button>
		</body>
	</html>
	`), "", 1024, 768)
	// `), "", 480, 320)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	ui.Bind("save", func() {
		dlgBrowse()
	})

	dlgInfo()

	// Wait until UI window is closed
	<-ui.Done()
}
