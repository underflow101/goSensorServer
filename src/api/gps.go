package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/underflow101/src/db"
)

var GDb *db.Db
var DbDoc *db.Doc

// GetGps
func GetGps(c echo.Context) error {
	GDb = db.NewDatabase()

	response := &ApiGpsResponse{
		StatusCode: 200,
		Message:    "",
		Body:       GDb.Read()[0].Sensor.Gps,
	}

	return c.JSON(http.StatusOK, response)
}

// PostGps
func PostGps(c echo.Context) error {
	GDb = db.NewDatabase()

	params := make(map[string]string)
	c.Bind(&params)

	rawLat := params["latitude"]
	rawLong := params["longitude"]

	lat, _ := strconv.ParseFloat(rawLat, 64)
	long, _ := strconv.ParseFloat(rawLong, 64)

	log.Println("[SERVER] latitude RAW Input: ", rawLat)
	log.Println("[SERVER] longitude RAW Input: ", rawLong)
	log.Println("[SERVER] latitude Input: ", lat)
	log.Println("[SERVER] longitude Input: ", long)

	gpsDoc := &db.Gps{
		Latitude:  lat,
		Longitude: long,
	}

	sensorValDoc := &db.SensorClass{
		Gps: *gpsDoc,
	}

	gpsClientInput := &db.Doc{
		UserId:  "root",
		Created: GetTimeNowUnix(),
		Sensor:  *sensorValDoc,
	}

	GDb.Write("root", *gpsClientInput)

	response := &ApiResponse{
		StatusCode: 200,
		Message:    "Ok",
	}

	return c.JSON(http.StatusOK, response)
}
