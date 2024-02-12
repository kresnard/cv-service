package profile

import (
	"cv_service/commons"
	"cv_service/config"
	"cv_service/internal/entity"
	"cv_service/internal/repository/mysql"
	"cv_service/pkg/logger"
	"net/http"
	"sync"
)

type IProfileUsecase interface {
	CreateProfile(data *entity.Profile) (int64, error)
}

type ProfileUsecase struct {
	l   *logger.Logger
	cfg *config.Config
	pr  mysql.IProfileRepository
}

func NewProfileUsecase(l *logger.Logger, cfg *config.Config, pr mysql.IProfileRepository) *ProfileUsecase {
	return &ProfileUsecase{l, cfg, pr}
}

func (p *ProfileUsecase) CreateProfile(data *entity.Profile) (int64, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errCh error

	if data == nil || data.Name == "" || data.Gender == "" || data.Email == "" {
		return 0, commons.ErrInvalidPayload
	}

	var id int64
	mu.Lock()
	wg.Add(1)
	go func() {
		defer wg.Done()

		res, err := p.pr.Create(data)
		if err != nil {
			errCh = err
			defer p.l.CreateLog(&logger.Log{
				Event:      commons.USECASE_PROFILE + "|CreateProfile",
				Method:     "POST",
				StatusCode: http.StatusInternalServerError,
				Request:    data,
				Query:      "",
				Response:   err,
				Message:    commons.ErrUsecaseProfile,
			}, logger.LVL_ERROR)
		}
		id = res
	}()
	mu.Unlock()

	wg.Wait()
	if errCh != nil {
		return 0, commons.ErrUsecaseProfile
	}

	return id, nil
}
