package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/involvestecnologia/statuspage/client"
	"github.com/involvestecnologia/statuspage/component"
	"github.com/involvestecnologia/statuspage/prometheus"
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

	// Initialize repositories
	componentRepository := component.NewMongoRepository(session)
	incidentRepository := incident.NewMongoRepository(session)
	clientRepository := client.NewMongoRepository(session)

	// Initialize services
	componentService := component.NewService(componentRepository)
	incidentService := incident.NewService(incidentRepository,componentService)

	// Initialize routers
	component.ComponentRouter(componentRepository, v1)
	incident.IncidentRouter(incidentRepository, componentService, v1)
	client.ClientRouter(clientRepository, v1)
	prometheus.PrometheusRouter(incidentService,componentService, v1)

	router.Run(":8080")
}
