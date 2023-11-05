package store

import (
	"fmt"
	"github.com/godev111222333/shoe-backend/src/model"
	"gorm.io/gorm"
)

type FileStore struct {
	Db *gorm.DB
}

func NewFileStore(db *gorm.DB) *FileStore {
	return &FileStore{
		Db: db,
	}
}

func (s *FileStore) Create(file *model.File) error {
	if err := s.Db.Create(file).Error; err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
func (s *FileStore) GetByUUID(uuid string) (*model.File, error) {
	res := &model.File{}
	if err := s.Db.Model(&model.File{}).Where("uuid = ?", uuid).First(res).Error; err != nil {
		fmt.Println("error when get file by uuid", err)
		return nil, err
	}

	return res, nil
}
