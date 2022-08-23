package utils

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/baiest/myblog/myblog-backend/models"
)

func GetIdUserRequest(r *http.Request) models.IdUser {
	idUser, err := strconv.ParseUint(r.Header.Get("IdUser"), 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	return uint(idUser)
}
