package server

import (
	"core/db/data"
	"core/db/query"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) addTagToFile() {

	s.engine.POST("/link_file_tag", func(c *gin.Context) {
		var link data.LinkFileTagCreate

		if err := c.BindJSON(&link); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := query.AddTagToFile(s.sqliteDb.Db, link.FileId, link.TagId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
	})

	s.engine.DELETE("/link_file_tag", func(c *gin.Context) {
		var link data.LinkFileTagCreate

		if err := c.BindJSON(&link); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := query.RemoveTagFromFile(s.sqliteDb.Db, link.FileId, link.TagId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
	})
}
