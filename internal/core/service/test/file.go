package test

import (
	"github.com/paw1a/eschool/internal/core/domain"
	"io"
)

type FileBuilder struct {
	file domain.File
}

func NewFileBuilder() *FileBuilder {
	return &FileBuilder{
		file: domain.File{
			Name:   "filename",
			Path:   "media/",
			Reader: nil,
		},
	}
}

func (b *FileBuilder) WithName(name string) *FileBuilder {
	b.file.Name = name
	return b
}

func (b *FileBuilder) WithPath(path string) *FileBuilder {
	b.file.Path = path
	return b
}

func (b *FileBuilder) WithReader(reader io.Reader) *FileBuilder {
	b.file.Reader = reader
	return b
}

func (b *FileBuilder) Build() domain.File {
	return b.file
}
