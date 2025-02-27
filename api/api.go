package api

import (
	"log"
	"net/http"
	"time"

	"github.com/PhuPhuoc/curanest-appointment-service/config"
	"github.com/PhuPhuoc/curanest-appointment-service/docs"
	"github.com/PhuPhuoc/curanest-appointment-service/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type server struct {
	port string
	db   *sqlx.DB
}

func InitServer(port string, db *sqlx.DB) *server {
	return &server{
		port: port,
		db:   db,
	}
}

const (
	env_local = "local"
	env_vps   = "vps"
)

// @Summary		ping server
// @Description	ping server
// @Tags			ping
// @Accept			json
// @Produce		json
// @Success		200	{object}	map[string]any	"message success"
// @Failure		400	{object}	error			"Bad request error"
// @Router			/ping [get]
func (sv *server) RunApp() error {
	envDevlopment := config.AppConfig.EnvDev
	if envDevlopment == env_local {
		// gin.SetMode(gin.ReleaseMode)
		docs.SwaggerInfo.BasePath = "/"
	}

	if envDevlopment == env_vps {
		gin.SetMode(gin.ReleaseMode)
		docs.SwaggerInfo.BasePath = "/appointment"
	}

	router := gin.New()

	configcors := cors.DefaultConfig()
	configcors.AllowAllOrigins = true
	configcors.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	configcors.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	configcors.ExposeHeaders = []string{"Content-Length"}
	configcors.AllowCredentials = true
	configcors.MaxAge = 12 * time.Hour

	router.Use(cors.New(configcors))
	router.Use(middleware.SkipSwaggerLog(), gin.Recovery())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// tokenProvider := common.NewJWTx(config.AppConfig.Key, 65*60*24*7)

	router.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "curanest-appointment-service - pong"}) })

	// api := router.Group("/api/v1")
	// {
	// }

	// rpc := router.Group("/external/rpc")
	// {
	// }

	log.Println("server start listening at port: ", sv.port)
	return router.Run(sv.port)
}
