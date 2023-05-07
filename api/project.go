package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func (dbSrv *DBServer) GetAllProjects(c *gin.Context) {
	coll := dbSrv.DB.Collection("projects")
	cursor, err := coll.Find(context.Background(), bson.D{{}})
	if err != nil {
		c.AbortWithError(http.StatusFailedDependency, errors.New("failed to retrieve projects"))
	}

	var projects []LiteArtifact
	err = cursor.All(context.Background(), &projects)
	if err != nil {
		c.AbortWithError(http.StatusFailedDependency, errors.New("failed to decode projects"))
	}

	c.JSON(http.StatusOK, projects)
}

func (dbSrv *DBServer) GetProjectsByTitle(c *gin.Context) {
	projTitle := c.Param("title")

	coll := dbSrv.DB.Collection("projects")
	filter := bson.D{{"title", projTitle}}
	cursor, err := coll.Find(context.Background(), filter)
	if err != nil {
		c.AbortWithError(http.StatusFailedDependency, errors.New("failed to retrieve projects"))
	}

	var projects []Artifact
	err = cursor.All(context.Background(), &projects)
	if err != nil {
		c.AbortWithError(http.StatusFailedDependency, errors.New("failed to decode projects"))
	}

	if projects == nil {
		projects = make([]Artifact, 0)
	}

	c.JSON(http.StatusOK, projects)
}
