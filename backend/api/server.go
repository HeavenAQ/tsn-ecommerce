package api

import (
	"net/http"
	db "tsn-ecommerce/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func (server *Server) Start() error {
	return server.router.Run()
}

func (server *Server) healthCheck(ctx *gin.Context) {
	err := server.store.HealthCheck()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

func (server *Server) setupRouter() {
	v1 := server.router.Group("/api/v1")
	{
		v1.GET("/healthcheck", server.healthCheck)
		products := v1.Group("/products")
		{
			products.GET("/", server.ListProducts)
		}

	}
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	server.router = gin.Default()
	server.setupRouter()
	return server
}
