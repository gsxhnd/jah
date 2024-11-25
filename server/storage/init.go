package storage

import (
	"errors"

	"github.com/google/wire"
	"github.com/gsxhnd/jaha/utils"
)

type Storage interface {
	Ping() error
	GetImage(cover string, id uint, filename string) ([]byte, string, error)
	SaveImage(data []byte, id uint, filename string) error
}

func NewStorage(cfg *utils.Config) (Storage, error) {
	if cfg.Storage.Type == "minio" {
		return NewMinioStorage(cfg)
	}

	if cfg.Storage.Type == "local" {
		return NewLocalStorage(cfg.Storage)
	}

	return nil, errors.New("no storage type")
}

var StorageSet = wire.NewSet(NewStorage)
