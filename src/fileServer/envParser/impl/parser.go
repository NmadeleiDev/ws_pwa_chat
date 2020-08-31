package impl

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

type Parser struct {
}

func (p *Parser) GetDashboardStorageDbDsn() string {
	user := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")
	addr := os.Getenv("MONGO_ADDRESS")

	if user == "" || password == "" || addr == "" {
		logrus.Fatal("Env is empty: ", user, password, addr)
	}
	return fmt.Sprintf("mongodb://%v:%v@%v", user, password, addr)
}

func (p *Parser) GetS3BucketPath() string {
	path := os.Getenv("S3_BUCKET_PATH")
	if len(path) == 0 {
		logrus.Errorf("bucket path is empty!!!")
	}
	return path
}

func (p *Parser) GetApiUrl() string {
	return os.Getenv("API_URL")
}

func (p *Parser) IsDevMode() bool {
	return os.Getenv("DEV_MODE") == "on"
}

func (p *Parser) GetServerPort() string {
	port := os.Getenv("FILE_BACKEND_PORT")
	if len(port) == 0 {
		return ":8080"
	} else {
		return ":" + port
	}
}
