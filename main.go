package main

import (
	"flag"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	httpport = flag.Uint("port", 8000, "starting http port number")
)

func main() {
	flag.Parse()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", indexHandler)
	e.POST("/", uploadHandler)
	e.HideBanner = true
	e.Start(fmt.Sprintf(":%d", *httpport))
}

func indexHandler(c echo.Context) error {
	return c.File("templates/index.html")
}

func uploadHandler(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintln("multipart form error:", err))
	}
	for _, fhs := range form.File {
		for _, fh := range fhs {
			if err = saveFile(fh); err != nil {
				c.Logger().Error(err)
				continue
			}
		}
	}
	return c.File("templates/index.html")
}

func saveFile(fh *multipart.FileHeader) error {
	// satisfy this func
	fmt.Println("saving:", fh.Filename)
	return nil
}
