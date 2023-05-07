package api

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/j0suetm-com/jtm_svc/util"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DBServer struct {
	Client *mongo.Client
	DB     *mongo.Database
	Bucket *gridfs.Bucket
}

func connectToMongoDB(cfg *util.DBCfg) (*DBServer, error) {
	cmdMonitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			logrus.Info("mongodb -- ", evt.Command.String())
		},
	}

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?&authSource=admin", cfg.User, cfg.Password, cfg.Host, cfg.Port)
	clientOptions := options.Client().ApplyURI(uri).SetMonitor(cmdMonitor)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, err
	}

	logrus.Info("Connected to mongodb")

	db := client.Database("jtm_svc_db")
	bucketOptions := options.GridFSBucket().SetName("media")
	bucket, err := gridfs.NewBucket(db, bucketOptions)
	if err != nil {
		return nil, err
	}

	return &DBServer{
		Client: client,
		DB:     db,
		Bucket: bucket,
	}, nil
}

func New(cfg util.Cfg) (*gin.Engine, error) {
	if cfg.Server.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	dbSrv, err := connectToMongoDB(&cfg.DB)
	if err != nil {
		return nil, err
	}

	rtr := gin.Default()
	rtr.GET("/media/:id", dbSrv.GetMediaById)
	rtr.GET("/projects", dbSrv.GetAllProjects)
	rtr.GET("/projects/:title", dbSrv.GetProjectsByTitle)
	rtr.GET("/posts", dbSrv.GetAllPosts)
	rtr.GET("/posts/:title", dbSrv.GetPostsByTitle)

	return rtr, nil
}
