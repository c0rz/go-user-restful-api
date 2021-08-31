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

	err := c.Bind(&input)

	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := s.userSerivce.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Error Data 2", http.StatusBadRequest, input)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Mantap anjeng", http.StatusOK, newUser)
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

func (s *userHandler) Erorpokoe(c *gin.Context) {
	c.Next()
	errorToPrint := c.Errors.ByType(gin.ErrorTypePublic).Last()
	if errorToPrint != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": errorToPrint.Error(),
		})
	}
}
