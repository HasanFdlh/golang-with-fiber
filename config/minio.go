package config

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

var MinioClient *minio.Client

func InitMinio() {
	endpoint := viper.GetString("MINIO_ENDPOINT")
	accessKey := viper.GetString("MINIO_ACCESS_KEY")
	secretKey := viper.GetString("MINIO_SECRET_KEY")
	region := viper.GetString("MINIO_REGION")

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Region: region,
		Secure: false, // ⚠️ set true kalau endpoint HTTPS
	})
	if err != nil {
		log.Fatalf("❌ Minio init failed: %v", err)
	}
	MinioClient = client
	log.Println("✅ Minio connected")
}
