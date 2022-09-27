package models

import (
	"time"
)

type Trip struct {
	Id                 int       `json:"id" gorm:"primaryKey;autoIncrement;notNull"`
	Distance           int32     `json:"distance" gorm:"notNull;check:time >= 10"`
	Time               int32     `json:"time" gorm:"notNull;check:time >= 10"`
	Date               time.Time `json:"date" gorm:"notNull"`
	ReturnStationId    int       `json:"returnStationId" gorm:"notNull"`
	ReturnStation      Station   `json:"returnStation" gorm:"foreignKey:ReturnStationId"`
	DepartureStationId int       `json:"departureStationId" gorm:"notNull"`
	DepartureStation   Station   `json:"departureStation" gorm:"foreignKey:DepartureStationId"`
}

type Station struct {
	Id             int    `json:"id" gorm:"primaryKey;notNull"`
	Name           string `json:"name" gorm:"notNull"`
	Address        string `json:"address" gorm:"notNull"`
	CoordinateX    string `json:"coordinateX" gorm:"notNull"`
	CoordinateY    string `json:"coordinateY" gorm:"notNull"`
	ReturnTrips    []Trip `json:"returnTrips" gorm:"foreignKey:ReturnStationId"`
	DepartureTrips []Trip `json:"departureTrips" gorm:"foreignKey:DepartureStationId"`
}
