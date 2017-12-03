package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// CountryData is a simple model struct.
// It can be populated with the Name and IsoCode of a country
type CountryData struct {
	Name    string `json:"Name"`
	IsoCode string `json:"IsoCode"`
}

var sourceCountries = []CountryData{
	{Name: "United Kingdom", IsoCode: "GB"},
	{Name: "Ireland", IsoCode: "IE"},
	{Name: "France", IsoCode: "FR"},
}
var destinationCountries = []CountryData{
	{Name: "Spain", IsoCode: "ES"},
	{Name: "Germany", IsoCode: "DE"},
	{Name: "Belgium", IsoCode: "BE"},
	{Name: "Austria", IsoCode: "AT"},
}


func requestAuthorizerMiddleware(c *gin.Context) {
	if c.ContentType() != "application/json" {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

func handleCountryRequest(c *gin.Context) {
	if c.Query("target") == "source" {
		c.AbortWithStatusJSON(http.StatusOK, sourceCountries)
	} else if c.Query("target") == "destination" {
		c.AbortWithStatusJSON(http.StatusOK, destinationCountries)
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Page not found"})
	}
}