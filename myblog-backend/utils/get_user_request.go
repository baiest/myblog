package utils

import (
	"net/http"
	"strconv"

	"github.com/baiest/myblog/myblog-backend/models"
)

func GetIdUserRequest(r *http.Request) models.IdUser {
	idUser, _ := strconv.ParseUint(r.Header.Get("userId"), 10, 64)
	return uint(idUser)
}
