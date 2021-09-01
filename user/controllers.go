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
		errors := helper.APIError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("error", "Bad Request", http.StatusBadGateway, errorMessage)
		c.JSON(http.StatusBadGateway, response)
		return
	}

	checkAcc, err := s.userSerivce.CountUser("email", input.Email)
	if err != nil {
		response := helper.APIResponse("error", "Server error", http.StatusBadGateway, err)
		c.JSON(http.StatusBadGateway, response)
		return
	}
	// fmt.Println(checkAc)
	// fmt.Println(input.Email)
	if !checkAcc {
		response := helper.APIResponse("error", "Email has been registered", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := s.userSerivce.RegisterUser(input)
	if err != nil {
		errors := helper.APIError(err)
		errorMessage := gin.H{"error": errors}
		response := helper.APIResponse("error", "Server error", http.StatusBadGateway, errorMessage)
		c.JSON(http.StatusBadGateway, response)
		return
	}
	response := helper.APIResponse("success", "Account has been registered", http.StatusOK, newUser)
	c.JSON(http.StatusOK, response)
}

func (s *userHandler) GetUsers(c *gin.Context) {
	getAll, err := s.userSerivce.GetAllUser()
	if err != nil {
		response := helper.APIResponse("error", "Server error", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("success", "Users Account", http.StatusOK, getAll)
	c.JSON(http.StatusOK, response)
}

func (s *userHandler) NotFound(c *gin.Context) {
	response := helper.APIResponse("error", "Not Found Route", http.StatusNotFound, nil)
	c.JSON(http.StatusNotFound, response)
}
