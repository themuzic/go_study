package main

import (
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/themuzic/scrapper"
)

const FileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}
func handleScrape(c echo.Context) error {
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	defer os.Remove(FileName)
	return c.Attachment(FileName, "job.csv")
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
