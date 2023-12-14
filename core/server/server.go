package server

import (
	"core/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	sqliteDb *db.SqliteDb
	engine   *gin.Engine
}

func NewServer() *Server {
	sqliteDb := &db.SqliteDb{}
	sqliteDb.Connect()
	return &Server{sqliteDb, gin.Default()}
}

func (s *Server) Run() {
	s.engine.Run()
}

func (s *Server) Close() {
	s.sqliteDb.Close()
}

func (s *Server) Ping() {
	s.engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

func (s *Server) GetTags() {
	s.engine.GET("/tags", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"tags": s.sqliteDb.SelectAllTags(),
		})
	})
}
