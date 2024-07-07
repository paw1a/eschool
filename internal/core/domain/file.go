package domain

import (
	"github.com/paw1a/eschool/internal/core/errs"
	"io"
)

type Url string

func (url Url) String() string {
	return string(url)
}

type File struct {
	Name   string
	Path   string
	Reader io.Reader
}

func (f *File) Validate() error {
	if f.Name == "" {
		return errs.ErrFilenameEmpty
	}
	if f.Path == "" {
		return errs.ErrFilepathEmpty
	}
	if f.Reader == nil {
		return errs.ErrFileReaderEmpty
	}
	return nil
}
