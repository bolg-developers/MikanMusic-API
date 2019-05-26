package infra

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/bolg-developers/MikanMusic-API/config"
	"github.com/pkg/errors"
)

var _bucket = &bucket{
	name: aws.String(config.Env().S3Bucket),
	svc: s3.New(session.New(), &aws.Config{
		Credentials: credentials.NewStaticCredentials(config.Env().S3AccessKeyID, config.Env().S3SecretAccessKey, ""),
		Region:      aws.String(endpoints.ApNortheast1RegionID),
	}),
}

func Bucket() *bucket {
	return _bucket
}

type bucket struct {
	name *string
	svc  *s3.S3
}

func (b *bucket) Create(rs io.ReadSeeker, key string) error {
	_, err := b.svc.PutObject(&s3.PutObjectInput{
		Bucket: b.name,
		Key:    aws.String(key),
		ACL:    aws.String("public-read"),
		Body:   rs,
	})
	return errors.WithStack(err)
}

func (b *bucket) URL(key string) string {
	return "https://s3-" + *b.svc.Config.Region + ".amazonaws.com/" + *b.name + "/" + key
}
