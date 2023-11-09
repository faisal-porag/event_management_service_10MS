package httpserver

import (
	"event_management_service_10MS/controllers"
	"event_management_service_10MS/state"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Serve(s *state.State) {
	// TODO:: UNCOMMENT THIS LINE IN PRODUCTION
	//gin.SetMode(gin.ReleaseMode)

	// Create a Gin router instance
	router := gin.Default()

	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Accept-Language"}
	router.Use(cors.New(config))

	// Define routes
	v1 := router.Group("/api/v1")
	{
		// Health Check
		v1.GET("/health-check", healthCheck(s))

		// APIS FOR EVENT MANAGEMENT SERVICE
		v1.GET("/get-event-list", controllers.HandleEventList(s))
		v1.GET("/get-event-details/:event_id", controllers.HandleEventDetails(s))
		v1.GET("/get-workshop-list/:event_id", controllers.HandleWorkshopListForSingleEvent(s))
		v1.GET("/get-workshop-details/:workshop_id", controllers.HandleWorkshopDetails(s))
		v1.POST("/workshop-reservation", controllers.HandleWorkshopReservation(s))

	}

	// Start the Gin server
	port := s.Cfg.ApplicationPort
	log.Info().Int("port", port).Msg("Starting >>> event management service >>> http server")
	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal().Err(err).Msg("router.Run err")
	}
}
