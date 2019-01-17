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
	mgouri, exist := os.LookupEnv("MONGO_URI")
	if !exist {
		log.Panic("MongoDB URI not informed")
	}
	session := db.InitMongo(mgouri)

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	v1 := router.Group("/v1")

	component.ComponentRouter(component.NewMongoRepository(session), v1)
	incident.IncidentRouter(incident.NewMongoRepository(session), v1)
	client.ClientRouter(client.NewMongoRepository(session), v1)

	router.Run(":8080")
}
