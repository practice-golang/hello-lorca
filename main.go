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

func main() {
	go func() {
		e := echo.New()
		e.HideBanner = true
		e.Use(middleware.CORS())

		e.GET("/", func(c echo.Context) error {
			return c.HTML(http.StatusOK, string(index))
		})

		e.Logger.Fatal(e.Start("127.0.0.1:1323"))
	}()

	cwd, _ := os.Getwd()
	profilePath := cwd + `\profile`

	ui, err := lorca.New("", profilePath, 1024, 768)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	ui.Load("http://localhost:1323")
	// ui.Load("https://github.com/zserge/lorca/issues/167")
	if err != nil {
		log.Fatal(err)
	}

	<-ui.Done()
}
