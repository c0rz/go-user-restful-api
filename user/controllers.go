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

func (s *userHandler) RegisUser(c *gin.Context) {
	var input RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("Error Data", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Mantap anjeng", http.StatusOK, nil)
	c.JSON(http.StatusOK, response)
}

func (s *userHandler) GetUsers(c *gin.Context) {
	getAll, err := s.userSerivce.GetAllUser()
	if err != nil {
		response := helper.APIResponse("fail API", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("success API", http.StatusOK, getAll)
	c.JSON(http.StatusOK, response)
}

func (s *userHandler) NotFound(c *gin.Context) {
	response := helper.APIResponse("fail API", http.StatusNotFound, nil)
	c.JSON(http.StatusNotFound, response)
}
