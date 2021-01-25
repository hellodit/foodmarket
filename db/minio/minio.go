package minio

import (
	"bytes"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
	"log"
	"strings"
)

var (
	client *minio.Client
	err    error
	accessKey,
	secKey,
	endpoint,
	bucket,
	region,
	cdn string
	ssl = false
)

func InitS3() {
	accessKey = viper.GetString("MINIO_KEY")
	secKey = viper.GetString("MINIO_SECRET")
	endpoint = viper.GetString("MINIO_ENDPOINT")
	bucket = viper.GetString("MINIO_BUCKET")
	region = viper.GetString("MINIO_REGION")
	cdn = viper.GetString("MINIO_CDN")
	ssl = viper.GetBool("MINIO_SSL")

	// Initialize minio client object.
	client, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secKey, ""),
		Secure: ssl,
	})

	if err != nil {
		log.Fatalf("Could not initialize S3 %v", err.Error())
	}

	fmt.Println("S3 Initialized")
}

func GetClient() *minio.Client {
	if client == nil {
		InitS3()
	}
	return client
}

func Upload(tempFileName string, buffer []byte, size int64, contentType string) (name string, err error) {
	userMetaData := map[string]string{"x-amz-acl": "public-read"}
	_, err = GetClient().PutObject(
		context.TODO(),
		bucket,
		strings.TrimPrefix(tempFileName, "/"),
		bytes.NewReader(buffer),
		size,
		minio.PutObjectOptions{
			ContentType:  contentType,
			UserMetadata: userMetaData,
		})

	if err != nil {
		return "", err
	}

	return tempFileName, nil
}
