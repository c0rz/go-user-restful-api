package user

import (
	"net/http"
	"simple-api-go-c0rz/helper"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userSerivce Service
}

func Controllers(userSerivce Service) *userHandler {
	return &userHandler{userSerivce}

}

func (s *userHandler) GetUsers(c *gin.Context) {
	getAll, err := s.userSerivce.GetAllUser()
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Mantap anjeng", http.StatusOK, getAll)
	c.JSON(http.StatusOK, response)
}
