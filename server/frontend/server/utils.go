package server

import (
	"net/http"

	"github.com/assafvayner/fs3/protos/fs3"
)

func GetFs3StatusHttpCode(status fs3.Status) int {
	switch status {
	case fs3.Status_GREAT_SUCCESS:
		return http.StatusAccepted
	case fs3.Status_INTERNAL_ERROR:
		return http.StatusInternalServerError
	case fs3.Status_NOT_FOUND:
		return http.StatusNotFound
	case fs3.Status_ILLEGAL_PATH:
		return http.StatusForbidden
	default:
		return http.StatusBadRequest
	}
}
