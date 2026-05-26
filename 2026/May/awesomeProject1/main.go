package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func getCats(c echo.Context) error {
	dataType := c.Param("data")
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")
	dataType = strings.ToLower(dataType)
	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": catName,
			"type": catType,
		})
	}
	if dataType == "string" {
		return c.String(http.StatusOK, "Cat name: "+catName+"Cat type: "+catType)
	}
	return c.String(http.StatusBadRequest, "Invalid data type")
}

func addCat(c echo.Context) error {
	cat := Cat{}
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}
	err = json.Unmarshal(b, &cat)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid JSON format")
	}
	return c.JSON(http.StatusOK, cat)
}

func main() {
	e := echo.New()
	e.GET("/cats/:data", getCats)
	e.POST("/cats", addCat)
	if err := e.Start(":8080"); err != nil {
		panic(err)
	}
}
