package s3Manager

import (
	"fileServer/envParser"
	"fileServer/types"
	"fileServer/utils"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/sirupsen/logrus"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var Manager types.S3Manager

type manager struct {
	sess *session.Session
	client *s3.S3
}

func (m *manager) UploadFile(filepath string) bool {
	uploader := s3manager.NewUploader(m.sess)

	f, err  := os.Open(filepath)
	if err != nil {
		logrus.Errorf("failed to open file %q, %v", filepath, err)
		return false
	}

	genKey := utils.GenSha1(time.Now().String() + strconv.Itoa(rand.Int()) + filepath)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(envParser.GetParser().GetS3BucketPath()),
		Key:    aws.String(genKey),
		Body:   f,
	})
	if err != nil {
		logrus.Errorf("failed to upload file, %v", err)
		return false
	}
	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
	return true
}

func (m *manager) DownloadFile(fileId string) bool {
	downloader := s3manager.NewDownloader(m.sess)

	filename := "saved_sample.txt"

	f, err := os.Create(filename)
	if err != nil {
		logrus.Errorf("failed to create file %q, %v", filename, err)
		return false
	}

	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(envParser.GetParser().GetS3BucketPath()),
		Key:    aws.String(fileId),
	})
	if err != nil {
		logrus.Errorf("failed to download file, %v", err)
		return false
	}
	fmt.Printf("file downloaded, %d bytes\n", n)
	return true
}

func (m *manager) Init() {
	var err error
	yes := true

	m.sess, err = session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
		CredentialsChainVerboseErrors: &yes})
	if err != nil {
		log.Fatalf("Error creating aws s3 session: %v", err)
	}

	if _, err := m.sess.Config.Credentials.Get(); err != nil {
		log.Fatalf("Failed to find aws creds: %v", err)
	}

	m.client = s3.New(m.sess)

	Manager = &manager{}
}
