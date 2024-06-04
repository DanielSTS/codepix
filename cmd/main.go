package main

import (
	"os"

	"github.com/DanielSTS/codepix/application/grpc"
	"github.com/DanielSTS/codepix/infra/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
