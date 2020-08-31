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
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	bucket = "enchat-files"
	downloadDir = "files_download"
	uploadDir = "files_upload"
)

var Manager types.S3Manager

type manager struct {
	sess *session.Session
	client *s3.S3
}

func (m *manager) UploadFile(filepath string) (string, bool) {
	uploader := s3manager.NewUploader(m.sess)

	f, err  := os.Open(filepath)
	if err != nil {
		logrus.Errorf("failed to open file %q, %v", filepath, err)
		return "", false
	}

	genKey := utils.GenSha1(time.Now().String() + strconv.Itoa(rand.Int()) + filepath)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(envParser.GetParser().GetS3BucketPath()),
		Key:    aws.String(genKey),
		Body:   f,
	})
	if err != nil {
		logrus.Errorf("failed to upload file, %v", err)
		return "", false
	}
	logrus.Infof("file uploaded to, %s\n", aws.StringValue(&result.Location))
	return genKey, true
}

func (m *manager) DownloadFile(fileId string) (*os.File, error) {
	downloader := s3manager.NewDownloader(m.sess)

	f, err := ioutil.TempFile(downloadDir, "")
	if err != nil {
		logrus.Errorf("failed to create file %v", err)
		return nil, err
	}

	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(envParser.GetParser().GetS3BucketPath()),
		Key:    aws.String(fileId),
	})
	if err != nil {
		logrus.Errorf("failed to download file, %v", err)
		return nil, err
	}
	fmt.Printf("file downloaded, %d bytes\n", n)
	return f, nil
}

func Init() {
	var err error
	yes := true

	m := manager{}

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

	Manager = &m

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	// Create S3 service client
	svc := s3.New(sess)
	result, err := svc.ListBuckets(nil)
	if err != nil {
		logrus.Errorf("Unable to list buckets, %v", err)
	}

	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}
