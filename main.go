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

func initEcho() {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, string(index))
	})

	e.Logger.Fatal(e.Start("127.0.0.1:1323"))
}

func initLorca() {
	cwd, _ := os.Getwd()
	profilePath := cwd + `\profile`

	// args := []string{"--ash-force-desktop"}
	args := []string{}

	ui, err := lorca.New("http://localhost:1323", profilePath, 1024, 768, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	<-ui.Done()
}

func main() {
	go func() { initEcho() }()
	initLorca()
}
