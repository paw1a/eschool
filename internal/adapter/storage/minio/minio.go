package storage

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/paw1a/eschool/internal/core/domain"
	"github.com/paw1a/eschool/internal/core/errs"
	minio2 "github.com/paw1a/eschool/pkg/minio"
	"github.com/pkg/errors"
	"mime"
	"net/url"
	"path/filepath"
)

type MinioObjectStorage struct {
	minioConfig *minio2.Config
	minioClient *minio.Client
}

func NewObjectStorage(minioClient *minio.Client, cfg *minio2.Config) *MinioObjectStorage {
	return &MinioObjectStorage{
		minioClient: minioClient,
		minioConfig: cfg,
	}
}

func (m *MinioObjectStorage) SaveFile(ctx context.Context, file domain.File) (domain.Url, error) {
	minioFilename := filepath.Join(file.Path, file.Name)
	contentType := mime.TypeByExtension(filepath.Ext(minioFilename))
	_, err := m.minioClient.PutObject(ctx, m.minioConfig.BucketName, minioFilename,
		file.Reader, -1, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", errors.Wrap(errs.ErrSaveFileError, err.Error())
	}
	fileUrl := url.URL{
		Scheme: "http",
		Host:   m.minioConfig.Endpoint,
		Path:   filepath.Join(m.minioConfig.BucketName, minioFilename),
	}
	return domain.Url(fileUrl.String()), nil
}
