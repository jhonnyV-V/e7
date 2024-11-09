package main

import (
	"math/rand/v2"
	"net/http"

	"github.com/a-h/templ"
	"github.com/jhonnyV-V/altscore-e7/view"
	"github.com/labstack/echo/v4"
)

type DamagedSystem struct {
	Damaged string `json:"damaged_system"`
}

var lastDamagedSystem string
var toRepairSystem string

var potentialDamagedSystems []string = []string{
	"navigation",
	"communications",
	"life_support",
	"engines",
	"deflector_shield",
}

var damagedSystemToRepairName map[string]string = map[string]string{
	"navigation":       "NAV-01",
	"communications":   "COM-02",
	"life_support":     "LIFE-03",
	"engines":          "ENG-04",
	"deflector_shield": "SHLD-05",
}

func main() {
	e := echo.New()

	e.GET("/status", func(c echo.Context) error {
		lastDamagedSystem = potentialDamagedSystems[rand.IntN(len(potentialDamagedSystems))]
		return c.JSON(http.StatusOK, DamagedSystem{Damaged: lastDamagedSystem})
	})

	e.GET("/repair-bay", func(c echo.Context) error {
		buf := templ.GetBuffer()
		defer templ.ReleaseBuffer(buf)

		if lastDamagedSystem == "" {
			lastDamagedSystem = potentialDamagedSystems[0]
		}

		repairBay := components.RepairBay(
			damagedSystemToRepairName[lastDamagedSystem],
		)
		if err := repairBay.Render(c.Request().Context(), buf); err != nil {
			return err
		}
		return c.HTML(http.StatusOK, buf.String())
	})

	e.POST("/teapot", func(c echo.Context) error {
		return c.String(http.StatusTeapot, "")
	})

	e.Start(":8080")
}
