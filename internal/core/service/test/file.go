package test

import (
	"github.com/paw1a/eschool/internal/core/domain"
	"io"
)

// FileBuilder is a builder for the File struct.
type FileBuilder struct {
	file domain.File
}

// NewFileBuilder creates a new instance of FileBuilder.
func NewFileBuilder() *FileBuilder {
	return &FileBuilder{
		file: domain.File{
			Name:   "filename",
			Path:   "media/",
			Reader: nil,
		},
	}
}

// WithName sets the name of the file.
func (b *FileBuilder) WithName(name string) *FileBuilder {
	b.file.Name = name
	return b
}

// WithPath sets the path of the file.
func (b *FileBuilder) WithPath(path string) *FileBuilder {
	b.file.Path = path
	return b
}

// WithReader sets the reader of the file.
func (b *FileBuilder) WithReader(reader io.Reader) *FileBuilder {
	b.file.Reader = reader
	return b
}

// Build constructs the File object and validates it.
func (b *FileBuilder) Build() domain.File {
	return b.file
}
