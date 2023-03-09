package api

import (
	"Zxun2/OneCV-Govtech/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config     utils.Config
	store      *gorm.DB
	router     *gin.Engine
}


// NewServer creates a new HTTP server and set up routing.
func NewServer(config utils.Config, store *gorm.DB) (*Server, error) {
	server := &Server{
		config:     config,
		store:      store,
	}

	server.setupRouter()
	return server, nil
}


func (server *Server) setupRouter() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	api := router.Group("/api")

	api.POST("/register", server.Register)
	api.GET("/commonstudents", server.GetCommonStudents)
	api.POST("/suspend", server.Suspend)
	api.POST("/retrievefornotifications", server.RetrieveNotifications)

	router.Run(server.config.HTTPServerAddress)

	server.router = router
}


// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}