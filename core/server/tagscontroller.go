package server

import (
	. "core/db/query"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) getTags() {
	s.engine.GET("/tags", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"tags": SelectAllTags(s.sqliteDb.Db),
		})
	})
}

func (s *Server) getTagsWithId() {
	s.engine.GET("/tags/:id", func(c *gin.Context) {
		// Check if id is a number
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("id <%s> is not a number", c.Param("id")),
			})
			return
		}

		data, err := SelectTag(s.sqliteDb.Db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"tags": data,
		})
	})
}
