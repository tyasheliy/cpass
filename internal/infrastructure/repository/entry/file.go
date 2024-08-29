package entry

import (
	"context"
	"fmt"
	"github.com/tyasheliy/cpass/internal/entity"
	"github.com/tyasheliy/cpass/internal/infrastructure/file"
	"github.com/tyasheliy/cpass/internal/passcl"
)

const (
	OTP_DIR  = "otp"
	TODO_DIR = "todo"
)

type FileEntryRepository struct {
	manager *file.Manager
	client  passcl.Client
}

func NewFileEntryRepository(
	manager *file.Manager,
	client passcl.Client,
) *FileEntryRepository {
	return &FileEntryRepository{
		manager: manager,
		client:  client,
	}
}

func (r *FileEntryRepository) CreatePassword(ctx context.Context, store *entity.Store, name string, password string) (*entity.Entry, error) {
	storeDir := r.manager.GetStoreFullName(store)
	passName := fmt.Sprintf("%s/%s", storeDir, name)
	options := passcl.InsertOptions{
		Force:     false,
		MultiLine: false,
	}

	err := r.client.Insert(ctx, passName, []string{password}, options)
	if err != nil {
		return nil, err
	}

	entry := entity.Entry{
		Store:     store,
		Name:      name,
		EntryType: entity.PasswordEntryType,
	}

	return &entry, nil
}

func (r *FileEntryRepository) GeneratePassword(ctx context.Context, store *entity.Store, name string, gen entity.PasswordGeneration) (*entity.Entry, error) {
	storeDir := r.manager.GetStoreFullName(store)
	passName := fmt.Sprintf("%s/%s", storeDir, name)
	options := passcl.GenerateOptions{
		Force:     false,
		NoSymbols: gen.NoSymbols,
		Length:    gen.Length,
	}

	err := r.client.Generate(ctx, passName, options)
	if err != nil {
		return nil, err
	}

	entry := entity.Entry{
		Store:     store,
		Name:      name,
		EntryType: entity.PasswordEntryType,
	}

	return &entry, nil
}

func (r *FileEntryRepository) CreateOtp(ctx context.Context, store *entity.Store, name string, uri string) (*entity.Entry, error) {
	storeDir := r.manager.GetStoreFullName(store)
	passName := fmt.Sprintf("%s/%s/%s", storeDir, OTP_DIR, name)
	options := passcl.InsertOtpOptions{
		Force: false,
	}

	err := r.client.InsertOtp(ctx, passName, uri, options)
	if err != nil {
		return nil, err
	}

	entry := entity.Entry{
		Store:     store,
		Name:      name,
		EntryType: entity.OtpEntryType,
	}

	return &entry, nil
}

func (r *FileEntryRepository) CreateTodo(ctx context.Context, store *entity.Store, name string, lines []string) (*entity.Entry, error) {
	storeDir := r.manager.GetStoreFullName(store)
	passName := fmt.Sprintf("%s/%s/%s", storeDir, TODO_DIR, name)
	options := passcl.InsertOptions{
		Force:     false,
		MultiLine: true,
	}

	err := r.client.Insert(ctx, passName, lines, options)
	if err != nil {
		return nil, err
	}

	entry := entity.Entry{
		Store:     store,
		Name:      name,
		EntryType: entity.TodoEntryType,
	}

	return &entry, nil
}

func (f *FileEntryRepository) Get(ctx context.Context) ([]*entity.Entry, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FileEntryRepository) GetByStore(ctx context.Context, store *entity.Store) ([]*entity.Entry, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FileEntryRepository) GetByType(ctx context.Context, store *entity.Store, t entity.EntryType) ([]*entity.Entry, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FileEntryRepository) GetByName(ctx context.Context, store *entity.Store, name string) (*entity.Entry, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FileEntryRepository) Delete(ctx context.Context, store *entity.Store, name string) error {
	//TODO implement me
	panic("implement me")
}
