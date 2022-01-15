package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// DB connection
const connectionUrl = "mongodb://localhost:27017"

// DB name
const dbName = "SensorServer"

// Collection name
const collecName = "root"

// Db struct for MongoDB control
type Db struct {
	DbName     string
	CollecName string
	Collection *mongo.Collection
}

// Doc struct for maintaining MongoDB Document
type Doc struct {
	UserId  string      `bson:"userId"`
	Created int64       `bson:"created"`
	Sensor  SensorClass `bson:"sensorClass"`
}

// SensorClass
type SensorClass struct {
	Gps Gps `bson:"gps"`
}

// Gps
type Gps struct {
	Latitude  float64 `bson:"latitude"`
	Longitude float64 `bson:"longitude"`
}
