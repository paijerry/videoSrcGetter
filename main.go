package main

import (
	"net/http"
	"test/videoSrcGetter/vimeo"
	"test/videoSrcGetter/youtube"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/vimeo/:vid", vimeoHandler)
	e.GET("/youtube/:vid", youtubeHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

func vimeoHandler(c echo.Context) error {
	result, err := vimeo.GetInfo(c.Param("vid"))
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
func youtubeHandler(c echo.Context) error {
	result, err := youtube.GetInfo(c.Param("vid"))
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
