package api

import (
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (dbSrv *DBServer) GetMediaById(c *gin.Context) {
	mediaId := c.Param("id")

	objId, err := primitive.ObjectIDFromHex(mediaId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.New("invalid media id"))
		return
	}

	dbFile, err := dbSrv.Bucket.OpenDownloadStream(objId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusFailedDependency, errors.New("failed to retrieve media"))
		return
	}
	defer dbFile.Close()

	byteData, err := io.ReadAll(dbFile)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusFailedDependency, errors.New("failed to parse media"))
		return
	}

	c.Header("Content-Disposition", "attachment; filename=media.jpeg")
	c.Data(http.StatusOK, "application/octet-stream", byteData)
}
