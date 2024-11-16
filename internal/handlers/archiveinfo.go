package handlers

import (
	"net/http"

	"github.com/KarmaBeLike/doodocs_days/internal/errors"
	"github.com/KarmaBeLike/doodocs_days/internal/service"

	"github.com/gin-gonic/gin"
)

func GetArchiveInfoHandler(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		apiErr := errors.ErrFileRead
		c.JSON(apiErr.Code, gin.H{"error": apiErr.Message})
		return
	}
	defer file.Close()

	archiveInfo, err := service.GetArchiveInfo(file, header)
	if err != nil {
		if archiveErr, ok := err.(*errors.ArchiveError); ok {
			c.JSON(archiveErr.Code, gin.H{"error": archiveErr.Message})
			return
		}
		apiErr := errors.ErrInternal
		c.JSON(apiErr.Code, gin.H{"error": apiErr.Message})
		return
	}
	c.JSON(http.StatusOK, archiveInfo)
}
