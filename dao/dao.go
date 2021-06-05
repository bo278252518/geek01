package dao

import (
	"github.com/bo278252518/geek01/types"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Dao interface {
	Close()
	FindById(id uint) (string, error)
}

// dao dao.
type dao struct {
	db *gorm.DB
}

// New new a dao and return.
func New(db *gorm.DB) (d Dao, err error) {
	return newDao(db)
}

func newDao(db *gorm.DB) (d *dao, err error) {
	d = &dao{
		db: db,
	}
	return
}

func (d *dao) FindById(id uint) (string, error) {
	var test types.Test
	err := d.db.First(&test, id).Error
	if err != nil {
		return "", errors.Wrap(err, "find test failed.")
	}
	return test.Message, nil
}

// Close close the resource.
func (d *dao) Close() {
}
