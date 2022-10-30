package http

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/WilkerAlves/genealogy/infra/controllers/person"
	"github.com/WilkerAlves/genealogy/infra/controllers/relationship"
	"github.com/WilkerAlves/genealogy/infra/repository"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func StartServer(ctx context.Context) {
	var handlerConfig http.Handler
	gin.SetMode(gin.DebugMode)
	handlerConfig = http.TimeoutHandler(
		configureRoutes(),
		time.Minute*10,
		`{"error": "timeout"}`,
	)

	srv := &http.Server{
		Addr:        fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")),
		ReadTimeout: 30 * time.Second,
		IdleTimeout: 20 * time.Second,
		BaseContext: func(l net.Listener) context.Context { return ctx },
		Handler:     handlerConfig,
	}

	log.Info().Msg("Starting server at: " + os.Getenv("SERVER_PORT"))
	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal().Err(err).Msg("unable to start server")
	}

}

func getConnectionString() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName,
	)
}

func configureRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/alive", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"live": "yes, just yes",
		})
	})

	conn := getConnectionString()

	personRepository, err := repository.NewPersonRepository(conn)
	if err != nil {
		log.Fatal().Err(err).Msg("error for create person repository")
	}
	personController := person.NewController(personRepository)

	routerGroup := router.Group("/persons")
	routerGroup.GET("/:id", personController.FindById)
	routerGroup.POST("/", personController.Create)
	routerGroup.PUT("/:id", personController.Update)
	routerGroup.DELETE("/:id", personController.Delete)

	relationshipRepository, err := repository.NewRelationshipRepository(conn)
	if err != nil {
		log.Fatal().Err(err).Msg("error for create relationship repository")
	}
	relationshipController := relationship.NewController(relationshipRepository)

	relationshipGroup := router.Group("/relationship")
	relationshipGroup.GET("/:id", relationshipController.Genealogy)
	relationshipGroup.POST("/", relationshipController.Add)
	relationshipGroup.GET("/find", relationshipController.Find)

	return router
}
