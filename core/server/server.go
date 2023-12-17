package server

import (
	"core/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	sqliteDb *database.SqliteDb
	engine   *gin.Engine
}

func NewServer() *Server {
	sqliteDb := &database.SqliteDb{}
	sqliteDb.Connect()
	engine := gin.Default()
	// Add CORS middleware
	engine.Use(CORSMiddleware())

	return &Server{sqliteDb, engine}
}

func (s *Server) Run() {
	s.engine.Run()
}

func (s *Server) Close() {
	s.sqliteDb.Close()
}

func (s *Server) Init() {
	s.ping()

	s.addTag()
	s.getTags()
	s.getTagsWithId()

	s.addTagToFile()

	s.addFile()
	s.getFile()
}

func (s *Server) ping() {
	s.engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
