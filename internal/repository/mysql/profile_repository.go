package mysql

import (
	"cv_service/commons"
	"cv_service/config"
	"cv_service/internal/entity"
	"cv_service/pkg/logger"
	"database/sql"
	"net/http"
)

type IProfileRepository interface {
	Create(data *entity.Profile) (int64, error)
	// GetDetail(id int) (*entity.Profile, error)
}

type ProfileRepository struct {
	l   *logger.Logger
	cfg *config.Config
	db  *sql.DB
}

func NewProfileRepository(l *logger.Logger, cfg *config.Config, db *sql.DB) *ProfileRepository {
	return &ProfileRepository{l, cfg, db}
}

func (p *ProfileRepository) Create(data *entity.Profile) (int64, error) {

	trx, err := p.db.Begin()
	if err != nil {
		return 0, err
	}

	query := ` INSERT INTO profiles (
		full_name,
		gender,
		birth_place,
		birth_date,
		email,
		created_at,
		updated_at
	)
	values (?,?,?,?,?,?,?)`

	stmt, err := trx.Prepare(query)

	if err != nil {
		defer p.l.CreateLog(&logger.Log{
			Event:      commons.REPO_MYSQL_PROFILE + "|prepareQuery",
			Method:     "",
			StatusCode: http.StatusInternalServerError,
			Request:    data,
			Query:      query,
			Response:   err,
			Message:    commons.ErrPrepareQuery,
		}, logger.LVL_ERROR)
		return 0, commons.ErrPrepareQuery
	}

	defer stmt.Close()

	result, err := stmt.Exec(
		data.Name,
		data.Gender,
		data.BirthPlace,
		data.BirthDate,
		data.Email,
		data.CreatedAt,
		data.UpdatedAt,
	)

	if err != nil {
		defer p.l.CreateLog(&logger.Log{
			Event:      commons.REPO_MYSQL_PROFILE + "|execQuery",
			Method:     "",
			StatusCode: http.StatusInternalServerError,
			Request:    data,
			Query:      query,
			Response:   err,
			Message:    commons.ErrQuery,
		}, logger.LVL_ERROR)
		return 0, commons.ErrPrepareQuery
	}

	res, err := result.LastInsertId()
	if err != nil {
		defer p.l.CreateLog(&logger.Log{
			Event:      commons.REPO_MYSQL_PROFILE + "|lastInsertId",
			Method:     "",
			StatusCode: http.StatusInternalServerError,
			Request:    data,
			Query:      query,
			Response:   err,
			Message:    commons.ErrQuery,
		}, logger.LVL_ERROR)
		return 0, commons.ErrPrepareQuery
	}

	return res, nil
}
