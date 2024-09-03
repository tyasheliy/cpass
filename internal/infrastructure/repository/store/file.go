package store

import (
	"context"
	"github.com/tyasheliy/cpass/internal/entity"
	"github.com/tyasheliy/cpass/internal/infrastructure/file"
	"os"
)

type FileStoreRepository struct {
	manager *file.Manager
}

func NewFileStoreRepository(manager *file.Manager) *FileStoreRepository {
	return &FileStoreRepository{
		manager: manager,
	}
}

// TODO: rework when file manager will be separated.
func (r *FileStoreRepository) Create(ctx context.Context, store *entity.Store) error {
	dirPath := r.getDirPath(store)
	return os.Mkdir(dirPath, os.ModeDir)
}

func (r *FileStoreRepository) Delete(ctx context.Context, store *entity.Store) error {
	dirPath := r.getDirPath(store)
	return os.RemoveAll(dirPath)
}

func (r *FileStoreRepository) getDirPath(store *entity.Store) string {
	storePath := r.manager.GetStorePath(store)
	return r.manager.GetPathFromRoot(storePath)
}
