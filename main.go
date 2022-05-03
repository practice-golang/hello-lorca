package main // import "hello-lorca"

import (
	"log"
	"net/url"
	"os"

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
	cwd, _ := os.Getwd()
	profilePath := cwd + `\profile`

	// Create UI with basic HTML passed via data URI
	ui, err := lorca.New("data:text/html,"+url.PathEscape(`
	<html>

	<head>
		<title>Hello</title>
	</head>

	<body>
		<h1>Hello, world!</h1>
		<a href="data:application/xml;charset=utf-8,your code here" download="filename.html">Save</a>
		<input type="file">
		<button onclick="save()">Save</button>

		<button onclick="openFileBrowser()">Open file picker</button>
		<button onclick="save()">Save</button>
		<button onclick="saveAS()">Save as</button>

		<div id="text-area">This is default text</div>
	</body>
	<script>
		let fileHandle

		async function openFileBrowser() {
			[fileHandle] = await window.showOpenFilePicker()
			const f = await fileHandle.getFile()
			const contents = await f.text()

			document.getElementById('text-area').innerText = contents
		}

		async function save() {
			if (fileHandle == undefined) {
				alert("File is not opened")
				return false
			}

			const writer = await fileHandle.createWritable()
			await writer.write(document.getElementById("text-area").innerText)
			await writer.close()
		}

		async function saveAS() {
			const handle = await getNewFileHandle()
			const contents = document.getElementById("text-area").innerText
			await writeFile(handle, contents)
		}

		async function writeFile(handle, contents) {

			const writable = await handle.createWritable()
			await writable.write(contents)
			await writable.close()
		}

		async function getNewFileHandle() {
			const options = {
				types: [{
					description: "Text Files",
					accept: { "text/plain": [".txt"] },
				}],
			};
			const handle = await window.showSaveFilePicker(options);
			return handle;
		}

		async function writeFile(fileHandle, contents) {
			const writable = await fileHandle.createWritable()
			await writable.write(contents)
			await writable.close()
		}
	</script>

	</html>
	`), profilePath, 1024, 768)
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
