package api

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	if posts == nil {
		posts = make([]Artifact, 0)
	}

	c.JSON(http.StatusOK, posts)
}

func (dbSrv *DBServer) GetRecentPost(c *gin.Context) {
	coll := dbSrv.DB.Collection("posts")
	opts := &options.FindOneOptions{
		Projection: bson.M{
			"title": 1,
			"tags": 1,
			"external_link": 1,
			"created_at": 1,
			"header_id": 1,
		},
		Sort: bson.M{"created_at": -1},
	}

	var post LiteArtifact
	err := coll.FindOne(context.Background(), bson.M{}, opts).Decode(&post)
	if err != nil {
		c.AbortWithError(http.StatusFailedDependency, errors.New("failed to retrieve projects"))
	}

	c.JSON(http.StatusOK, post)
}
