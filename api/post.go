package api

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (dbSrv *DBServer) GetAllPosts(c *gin.Context) {
	coll := dbSrv.DB.Collection("posts")
	cursor, err := coll.Find(context.Background(), bson.D{{}})
	if err != nil {
		c.AbortWithError(http.StatusFailedDependency, errors.New("failed to retrieve posts"))
	}

	var posts []LiteArtifact
	err = cursor.All(context.Background(), &posts)
	if err != nil {
		c.AbortWithError(http.StatusFailedDependency, errors.New("failed to decode posts"))
	}

	c.JSON(http.StatusOK, posts)
}

func (dbSrv *DBServer) GetPostsByTitle(c *gin.Context) {
	postTitle := c.Param("title")
	postTitle = strings.ReplaceAll(postTitle, "-", " ")

	coll := dbSrv.DB.Collection("posts")
	filter := bson.D{{"title", postTitle}}
	cursor, err := coll.Find(context.Background(), filter)
	if err != nil {
		c.AbortWithError(http.StatusFailedDependency, errors.New("failed to retrieve posts"))
	}

	var posts []Artifact
	err = cursor.All(context.Background(), &posts)
	if err != nil {
		c.AbortWithError(http.StatusFailedDependency, errors.New("failed to decode posts"))
	}

	c.JSON(http.StatusOK, posts)
}
