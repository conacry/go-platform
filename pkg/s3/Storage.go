package s3Storage

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
	"github.com/conacry/go-platform/pkg/storage"
	storageModel "github.com/conacry/go-platform/pkg/storage/model"
	"io"
)

type Logger interface {
	LogError(ctx context.Context, errs ...error)
	LogInfo(ctx context.Context, messages ...string)
}

type Config struct {
	Region    string
	Host      string
	AccessKey string
	SecretKey string
}

type Storage struct {
	config *Config
	logger Logger
	client *s3.Client
}

func (s *Storage) Start(ctx context.Context) error {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:           s.config.Host,
			SigningRegion: s.config.Region,
		}, nil
	})

	customCredentialProvider := aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
		return aws.Credentials{
			AccessKeyID:     s.config.AccessKey,
			SecretAccessKey: s.config.SecretKey,
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithCredentialsProvider(customCredentialProvider),
	)
	if err != nil {
		s.logger.LogError(ctx, err)
		return err
	}

	s.client = s3.NewFromConfig(cfg)

	s.logger.LogInfo(ctx, "S3 connection is initialized")

	result, err := s.client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		s.logger.LogError(ctx, err)
		return err
	}

	for _, bucket := range result.Buckets {
		msg := fmt.Sprintf("backet=%s creation time=%s", *bucket.Name, bucket.CreationDate.Format("2006-01-02 15:04:05 Monday"))
		s.logger.LogInfo(ctx, msg)
	}

	return nil
}

func (s *Storage) UploadFile(ctx context.Context, file *storageModel.File) error {
	bufferReader := bytes.NewReader(file.Content)

	params := &s3.PutObjectInput{
		Bucket:        aws.String(file.Scope),
		Key:           aws.String(file.FilePath),
		Body:          bufferReader,
		ContentLength: bufferReader.Size(),
		ContentType:   aws.String(file.MIME),
		ACL:           types.ObjectCannedACLPublicRead,
	}

	if _, err := s.client.PutObject(ctx, params); err != nil {
		s.logger.LogError(ctx, err)
		return err
	}

	return nil
}

func (s *Storage) GetFile(ctx context.Context, bucket, path string) (*storageModel.File, error) {
	result, err := s.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(path),
	})

	if err != nil {
		s.logger.LogError(ctx, err)

		if isNotFoundErr(err) {
			return nil, storage.ErrFileNotFound
		}

		return nil, err
	}

	defer result.Body.Close()
	fileBytes, err := io.ReadAll(result.Body)
	if err != nil {
		s.logger.LogError(ctx, err)
		return nil, err
	}

	file := &storageModel.File{
		Content:  fileBytes,
		Scope:    bucket,
		FilePath: path,
	}

	if result.ContentType != nil {
		file.MIME = *result.ContentType
	}

	return file, nil
}

func (s *Storage) RemoveFile(ctx context.Context, bucket, path string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(path),
	})

	if err != nil {
		s.logger.LogError(ctx, err)
		return err
	}

	return nil
}

func (s *Storage) GetFileMetaData(ctx context.Context, bucket, path string) (*storageModel.FileMetaData, error) {
	result, err := s.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(path),
	})

	if err != nil {
		s.logger.LogError(ctx, err)

		if isNotFoundErr(err) {
			return nil, storage.ErrFileNotFound
		}

		return nil, err
	}

	fileInfo := &storageModel.FileMetaData{
		Metadata: make(map[string]string),
	}
	for key, val := range result.Metadata {
		fileInfo.Metadata[key] = val
	}

	return fileInfo, nil
}

func isNotFoundErr(err error) bool {
	var apiError smithy.APIError
	if errors.As(err, &apiError) {
		var notFoundErr *types.NotFound
		if errors.As(apiError, &notFoundErr) {
			return true
		}
	}

	return false
}
