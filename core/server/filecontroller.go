package server

import (
	"core/database/data"
	"core/database/query"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TagsIdList struct {
	TagsId []int `json:"tagsId"`
}

func (s *Server) addFile() {
	s.engine.POST("/file", func(c *gin.Context) {
		var file data.File
		// todo id sert a rien

		if err := c.BindJSON(&file); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := s.sqliteDb.AddFile(file.Name, file.FolderId, file.Label)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "received"})
	})
}

func (s *Server) getFile() {
	s.engine.POST("/file/query", func(c *gin.Context) {
		var tagsId TagsIdList

		if err := c.BindJSON(&tagsId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		files, err := query.QueryFileWithTagsIdOr(s.sqliteDb.Db, tagsId.TagsId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		// Now you can use nums.Numbers
		c.JSON(http.StatusOK, gin.H{"status": "received", "numbers": files})
	})
}
