package applicant

import (
	"cv_service/config"
	"cv_service/internal/usecase/profile"
	"cv_service/pkg/logger"

	"github.com/gorilla/mux"
)

type ApplicantRoutes struct {
	l   *logger.Logger
	cfg *config.Config
	pu  profile.IProfileUsecase
}

func NewwApplicantRoutes(r *mux.Router, l *logger.Logger, cfg *config.Config, pu profile.IProfileUsecase) {
	a := &ApplicantRoutes{l, cfg, pu}

}
