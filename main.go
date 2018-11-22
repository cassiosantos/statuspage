package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/involvestecnologia/statuspage/client"
	"github.com/involvestecnologia/statuspage/component"
	"github.com/involvestecnologia/statuspage/db"
	"github.com/involvestecnologia/statuspage/incident"
	"github.com/involvestecnologia/statuspage/middleware"
)

func main() {
	mgouri := os.Getenv("MONGO_URI")
	if mgouri == "" {
		log.Panic("MongoDB URI not informed")
	}
	session := db.InitMongo(mgouri)

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	component.ComponentRouter(component.NewMongoRepository(session), router)
	incident.IncidentRouter(incident.NewMongoRepository(session), router)
	client.ClientRouter(client.NewMongoRepository(session), router)

	router.Run(":8080")
}
