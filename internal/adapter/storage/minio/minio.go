package storage

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	"github.com/pkg/errors"
	"mime"
	"net/url"
	"path/filepath"
)

type Config struct {
	Endpoint   string
	User       string
	Password   string
	BucketName string
}

type MinioObjectStorage struct {
	config      *Config
	minioClient *minio.Client
}

func NewObjectStorage(minioClient *minio.Client, config *Config) *MinioObjectStorage {
	return &MinioObjectStorage{
		minioClient: minioClient,
		config:      config,
	}
}

func (m *MinioObjectStorage) SaveFile(ctx context.Context, file domain.File) (domain.Url, error) {
	minioFilename := filepath.Join(file.Path, file.Name)
	contentType := mime.TypeByExtension(filepath.Ext(minioFilename))
	_, err := m.minioClient.PutObject(ctx, m.config.BucketName, minioFilename,
		file.Reader, -1, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", errors.Wrap(errs.ErrSaveFileError, err.Error())
	}
	fileUrl := url.URL{
		Scheme: "http",
		Host:   m.config.Endpoint,
		Path:   filepath.Join(m.config.BucketName, minioFilename),
	}
	return domain.Url(fileUrl.String()), nil
}
