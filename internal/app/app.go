package app

import (
	"cv_service/config"
	profile_repo_mysql "cv_service/internal/repository/mysql"
	"cv_service/internal/usecase/profile"
	"cv_service/pkg/logger"
	"cv_service/pkg/mysql"
)

func Run(cfg *config.Config) {
	var err error

	l := logger.New(cfg)

	//db
	db := mysql.New(cfg.MysqlDriverName, cfg, l)
	defer db.Close()

	//repo
	profileRepo := profile_repo_mysql.NewProfileRepository(l, cfg, db)

	//usecase
	profileUsecase := profile.NewProfileUsecase(l, cfg, profileRepo)

}
