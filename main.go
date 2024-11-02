package main

import (
	"Ga-backend/routes"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("gamification", routes.URL)
}