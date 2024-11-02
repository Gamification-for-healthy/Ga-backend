package config

import (
	"Ga-backend/helper/apdb" // untuk MongoConnect
	"Ga-backend/model"       // untuk DBinfo
	"os"
)

var mongoURI = os.Getenv("MONGO_URI")

var mongoInfo = model.DBinfo{ // menggunakan model.DBinfo
    DBString: mongoURI,
    DBName:   "gamification",
}

var DB, ErrorMongoconn = apdb.MongoConnect(mongoInfo) // menggunakan apdb.MongoConnect
