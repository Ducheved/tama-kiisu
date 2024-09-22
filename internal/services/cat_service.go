package services

import (
	"errors"
	"time"

	"github.com/Ducheved/tama-kiisu/internal/models"
	"github.com/Ducheved/tama-kiisu/internal/repositories"
)

type CatService struct {
	repo *repositories.CatRepository
}

func NewCatService(repo *repositories.CatRepository) *CatService {
	return &CatService{repo: repo}
}

func (s *CatService) GetCat() (*models.Cat, error) {
	cat, err := s.repo.GetCat()
	if err != nil {
		return nil, err
	}

	s.updateStats(cat)
	return cat, s.repo.UpdateCat(cat)
}

func (s *CatService) PerformAction(action string, ip string) (*models.Cat, error) {
	cat, err := s.repo.GetCat()
	if err != nil {
		return nil, err
	}

	s.updateStats(cat)

	switch action {
	case "/givefood":
		cat.Fullness = min(cat.Fullness+10, 100)
	case "/givefun":
		cat.Happiness = min(cat.Happiness+10, 100)
	case "/givehealth":
		cat.Health = min(cat.Health+10, 100)
	default:
		return nil, errors.New("Invalid action")
	}

	return cat, s.repo.UpdateCat(cat)
}

func (s *CatService) updateStats(cat *models.Cat) {
	hoursPassed := int(time.Since(cat.UpdatedAt).Hours())
	cat.Happiness = max(cat.Happiness-hoursPassed, 0)
	cat.Fullness = max(cat.Fullness-hoursPassed, 0)
	cat.Health = max(cat.Health-hoursPassed, 0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
