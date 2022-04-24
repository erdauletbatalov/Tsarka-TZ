package web

import (
	"log"
	"net/http"

	"github.com/erdauletbatalov/tsarka/domain"
	"github.com/jackc/pgx/v4"
)

func getStatusCode(err error) int {
	log.Println(err)
	switch err {
	case http.ErrNoCookie:
		return http.StatusUnauthorized
	case pgx.ErrNoRows:
		return http.StatusBadRequest
	case domain.ErrWrongPassword:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
