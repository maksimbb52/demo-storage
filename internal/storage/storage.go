package storage

import "github.com/maksimbb52/storage/internal/file"

type Storage struct {
}

func NewStorage() *Storage {
	return &Storage{}
}

func (S *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	return file.NewFile(filename, blob)

	//if err != nil {
	//	return nil, err
	//}
}
