package main // import "hello-lorca"

import (
	"log"
	"net/http"
	"os"

	_ "embed"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/practice-golang/lorca"
)

//go:embed index.html
var index []byte

//go:embed logo.png
var logo []byte

func initEcho() {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, string(index))
	})
	e.GET("/logo.png", func(c echo.Context) error {
		return c.HTML(http.StatusOK, string(logo))
	})

	e.Logger.Fatal(e.Start("127.0.0.1:1323"))
}

func initLorca() {
	cwd, _ := os.Getwd()
	profilePath := cwd + `\profile`

	lorca.DefaultChromeArgs = []string{
		"--disable-background-networking",
		"--disable-background-timer-throttling",
		"--disable-backgrounding-occluded-windows",
		"--disable-breakpad",
		"--disable-client-side-phishing-detection",
		"--disable-default-apps",
		"--disable-dev-shm-usage",
		"--disable-infobars",
		"--disable-extensions",
		"--disable-features=site-per-process",
		"--disable-hang-monitor",
		"--disable-ipc-flooding-protection",
		"--disable-popup-blocking",
		"--disable-prompt-on-repost",
		"--disable-renderer-backgrounding",
		"--disable-sync",
		"--disable-translate",
		// "--disable-windows10-custom-titlebar",
		"--metrics-recording-only",
		// "--no-first-run",
		"--no-default-browser-check",
		"--safebrowsing-disable-auto-update",
		// "--enable-automation",
		"--password-store=basic",
		"--use-mock-keychain",
	}

	// args := []string{"--ash-force-desktop", "--force-app-mode"}
	args := []string{}

	ui, err := lorca.New("http://localhost:1323", profilePath, 1024, 768, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	<-ui.Done()

	os.RemoveAll(cwd + `/profile`)
}

func main() {
	go func() { initEcho() }()
	initLorca()
}
