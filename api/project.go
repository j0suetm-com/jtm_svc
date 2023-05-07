package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	Title        string             `json:"title" bson:"title"`
	ExternalLink string             `json:"external_link" bson:"external_link"`
	Tags         []string           `json:"tags" bson:"tags"`
	HeaderId     primitive.ObjectID `json:"header_id" bson"header_id"`
	Content      []byte             `json:"content" bson:"content"`
}

func (dbSrv *DBServer) GetAllProjects(c *gin.Context) {
	coll := dbSrv.DB.Collection("projects")
	cursor, err := coll.Find(context.Background(), bson.D{{}})
	if err != nil {
		c.AbortWithError(http.StatusFailedDependency, errors.New("failed to retrieve projects"))
	}

	var projects []Project
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

	var projects []Project
	err = cursor.All(context.Background(), &projects)
	if err != nil {
		c.AbortWithError(http.StatusFailedDependency, errors.New("failed to decode projects"))
	}

	if projects == nil {
		projects = make([]Project, 0)
	}
	c.JSON(http.StatusOK, projects)
}
