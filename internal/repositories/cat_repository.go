package repositories

import (
	"github.com/Ducheved/tama-kiisu/internal/models"
	"gorm.io/gorm"
)

type CatRepository struct {
	db *gorm.DB
}

func NewCatRepository(db *gorm.DB) *CatRepository {
	db.AutoMigrate(&models.Cat{})
	return &CatRepository{db: db}
}

func (r *CatRepository) GetCat() (*models.Cat, error) {
	var cat models.Cat
	result := r.db.First(&cat)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			cat = models.Cat{Happiness: 50, Fullness: 50, Health: 50}
			r.db.Create(&cat)
			return &cat, nil
		}
		return nil, result.Error
	}
	return &cat, nil
}

func (r *CatRepository) UpdateCat(cat *models.Cat) error {
	return r.db.Save(cat).Error
}
