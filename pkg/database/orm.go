package database

import (
	"gorm.io/gorm"
	"math"
	"time"
)

// Model Define base model
type Model struct {
	Id uint64 `gorm:"primaryKey;"`
}

type Timestamps struct {
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type SoftDelete struct {
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func GetDBWithModel(defDB *gorm.DB, m interface{}) *gorm.DB {
	return defDB.Model(m)
}

func Exists(db *gorm.DB) (bool, error) {
	var exists bool
	result := db.Select("count(*) > 0").Find(&exists)
	if err := result.Error; err != nil {
		return false, err
	}
	return exists, nil
}

func FindOne(db *gorm.DB, out interface{}) (bool, error) {
	result := db.First(out)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func FindAll(db *gorm.DB, out interface{}) error {
	result := db.Find(out)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func FindByPaginate(db *gorm.DB, page PaginationParams, out interface{}) (*PageMeta, error) {
	meta := getPageMeta(db, page)
	// 如果总数没有值，那么就不查真实数据
	if meta.Total == 0 {
		return meta, nil
	}
	offset := (page.GetPage() - 1) * page.GetPageSize()
	result := db.Limit(page.GetPageSize()).Offset(offset).Find(out)
	if err := result.Error; err != nil {
		return nil, err
	}
	return meta, nil
}

func getPageMeta(db *gorm.DB, pageOption PaginationParams) *PageMeta {
	var count int64
	err := db.Count(&count).Error
	if err != nil {
		count = 0
	}
	return &PageMeta{
		Current:  pageOption.GetPage(),
		Last:     int(math.Ceil(float64(count) / float64(pageOption.GetPageSize()))),
		PageSize: pageOption.GetPageSize(),
		Total:    count,
	}
}
