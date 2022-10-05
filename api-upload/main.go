package main

import (
	"citybiker-go-api/upload/csvparser"
)

func main() {
	//db := database.InitDB()
	csvparser.ParseStations("./testData/testStations.csv", "parsedStations.csv")
}
