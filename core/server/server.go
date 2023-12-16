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
