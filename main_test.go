package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

var expectedSources = []CountryData{
	{Name: "United Kingdom", IsoCode: "GB"},
	{Name: "Ireland", IsoCode: "IE"},
	{Name: "France", IsoCode: "FR"},
}

var expectedTargets = []CountryData{
	{Name: "Spain", IsoCode: "ES"},
	{Name: "Germany", IsoCode: "DE"},
	{Name: "Belgium", IsoCode: "BE"},
	{Name: "Austria", IsoCode: "AT"},
}

func TestExpectedSourceData(t *testing.T) {

	for i, data := range sourceCountries {
		v := expectedSources[i]
		if v.Name != data.Name {
			t.Error(
				"For", v.Name,
				"expected", v.Name,
				"got", data.Name,
			)
		}
		if v.IsoCode != data.IsoCode {
			t.Error(
				"For", v.IsoCode,
				"expected", v.IsoCode,
				"got", data.IsoCode,
			)
		}

	}
}

func TestExpectedDestinationData(t *testing.T) {

	for i, data := range destinationCountries {
		v := expectedTargets[i]
		if v.Name != data.Name {
			t.Error(
				"For", v.Name,
				"expected", data.Name,
				"got", data.Name,
			)
		}
		if v.IsoCode != data.IsoCode {
			t.Error(
				"For", v.IsoCode,
				"expected", data.IsoCode,
				"got", data.IsoCode,
			)
		}

	}
}

func TestUnsupportedContentTypeReturnsBadRequest(t *testing.T) {
	// reduce output noise
	gin.SetMode(gin.TestMode)

	// Setup your router
	// register your routes
	r := gin.Default()
	// use requestAuthorizerMiddleware middleware to handle unsupported requests
	v1 := r.Group("/v1")
	{
		v1.GET("/countries", requestAuthorizerMiddleware, handleCountryRequest)
	}

	req, err := http.NewRequest(http.MethodGet, "/v1/countries?target=source", nil)
	// add xml content type not json
	req.Header.Add("content-type", "application/xml")
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	if http.StatusBadRequest != w.Code {
		t.Error(
			"For", http.StatusBadRequest,
			"expected", http.StatusBadRequest,
			"got", w.Code,
		)
	}
}

func TestUnsupportedAcceptReturnsBadRequest(t *testing.T) {
	// reduce output noise
	gin.SetMode(gin.TestMode)

	// Setup your router
	// register your routes
	r := gin.Default()
	// use requestAuthorizerMiddleware middleware to handle unsupported requests
	v1 := r.Group("/v1")
	{
		v1.GET("/countries", requestAuthorizerMiddleware, handleCountryRequest)
	}

	req, err := http.NewRequest(http.MethodGet, "/v1/countries?target=source", nil)
	// add xml content type not json
	req.Header.Add("content-type", "application/xml")
	req.Header.Add("accept", "application/xml")
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	if http.StatusBadRequest != w.Code {
		t.Error(
			"For", http.StatusBadRequest,
			"expected", http.StatusBadRequest,
			"got", w.Code,
		)
	}
}

func TestExpectedResponseCodes(t *testing.T) {
	testUrlData := []string{
		"/v1/countries?target=source|200",
		"/v1/countries?target=destination|200",
		"/v1/countries?target=foo|404",
	}

	// reduce output noise
	gin.SetMode(gin.TestMode)

	// Setup router
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/countries", handleCountryRequest)
	}

	for _, data := range testUrlData {
		// Create the mock request.
		splitData := strings.Split(data, "|")
		url := splitData[0]
		expectedCode, _ := strconv.Atoi(splitData[1])

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			t.Fatalf("Couldn't create request: %v\n", err)
		}

		// Create a response recorder so we can inspect the response
		w := httptest.NewRecorder()

		// Perform the request
		r.ServeHTTP(w, req)

		if expectedCode != w.Code {
			t.Error(
				"For", expectedCode,
				"expected", expectedCode,
				"got", w.Code,
			)
		}
	}
}

func TestSourceCountriesRequestData(t *testing.T) {
	// reduce output noise
	gin.SetMode(gin.TestMode)

	// Setup router
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/countries", requestAuthorizerMiddleware, handleCountryRequest)
	}

	req, err := http.NewRequest(http.MethodGet, "/v1/countries?target=source", nil)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("accept", "application/json")
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so we inspect the response
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	//decode the response data from mock request to make sure its expected
	countryData := [3]CountryData{}
	json.NewDecoder(w.Body).Decode(&countryData)

	for i, data := range countryData {

		v := expectedSources[i]
		if v.Name != data.Name {
			t.Error(
				"For", "/v1/countries?target=source response",
				"expected", v.Name,
				"got", data.Name,
			)
		}
		if v.IsoCode != data.IsoCode {
			t.Error(
				"For", "/v1/countries?target=source response",
				"expected", v.IsoCode,
				"got", data.IsoCode,
			)
		}
	}
}

