package http

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/WilkerAlves/genealogy/infra/controllers"
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

	// criando um servidor http customizado
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

func configureRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/alive", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"live": "yes, just yes",
		})
	})

	routerGroup := router.Group("/persons")
	routerGroup.GET("/:id", controllers.GetPerson)
	routerGroup.GET("/", controllers.GetPersons)
	routerGroup.POST("/", controllers.CreatePerson)
	routerGroup.PUT("/:id", controllers.UpdatePerson)
	routerGroup.DELETE("/:id", controllers.DeletePerson)

	return router
}
