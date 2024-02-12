package commons

import "errors"

var (
	// UsecaseProfile
	USECASE_PROFILE = "USECASE|PROFILE"

	// RepositoryMysql
	REPO_MYSQL_PROFILE    = "REPOSITORY|MYSQL|PROFILE"
	REPO_MYSQL_ADDRESS    = "REPOSITORY|MYSQL|ADDRESS"
	REPO_MYSQL_EXPERIENCE = "REPOSITORY|MYSQL|EXPERIENCE"

	// RepositoryMysql ERROR
	ErrQuery        = errors.New("error - execute query")
	ErrPrepareQuery = errors.New("error - preparing query statement")
	ErrRowScan      = errors.New("error - scanning rows repository")

	// Profile usecase  ERROR
	ErrUsecaseProfile = errors.New("error - usecase pofile")

	// COMMON ERROR
	ErrInvalidPayload = errors.New("error - invalid request payload")
	ErrInternalServer = errors.New("error - internal server error")
)
