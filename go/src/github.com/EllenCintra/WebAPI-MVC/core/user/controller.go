package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webapi-with-go/database/entity"
)

type ControllerUser struct {
	service IUserService
}

func NewControllerUser(service IUserService) ControllerUser {
	return ControllerUser{service: service}
}

func (s *ControllerUser) GetUsers(c *gin.Context) {
	users, err := s.service.GetUsers()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, &users)
}

func (s *ControllerUser) GetUser(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Id must be integer",
		})
		return
	}

	user, err := s.service.GetUser(int64(newid))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Can not find user by id: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (s *ControllerUser) CreateUser(c *gin.Context) {
	var user entity.User
	errJson := c.ShouldBindJSON(&user)

	if errJson != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Can not bind JSON: " + errJson.Error(),
		})
		return
	}

	newUser, err := s.service.CreateUser(&user)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Can not create user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func (s *ControllerUser) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	newid, integerErr := strconv.Atoi(id)

	if integerErr != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Id must be integer",
		})
		return
	}

	user, getErr := s.service.GetUser(int64(newid))

	if getErr != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Can not find user by id: " + getErr.Error(),
		})
		return
	}

	deleteErr := s.service.DeleteUser(user)

	if deleteErr != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Can not delete user: " + deleteErr.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (s *ControllerUser) EditUser(c *gin.Context) {
	id := c.Param("id")
	newid, integerErr := strconv.Atoi(id)

	if integerErr != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Id must be integer",
		})
		return
	}

	user, getErr := s.service.GetUser(int64(newid))

	if getErr != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Can not find user by id: " + getErr.Error(),
		})
		return
	}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Can not bind JSON: " + err.Error(),
		})
		return
	}

	userSaved, err := s.service.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Can not edit user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, userSaved)
}

func (s *ControllerUser) ValidationUserId(c *gin.Context) *entity.User {
	id := c.Param("id")
	newid, integerErr := strconv.Atoi(id)

	if integerErr != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Id must be integer",
		})
		c.Abort() //como parar a solicitação da mesma maneira que o return?
	}

	user, getErr := s.service.GetUser(int64(newid))

	if getErr != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Can not find book by id: " + getErr.Error(),
		})
		c.Abort() //como parar a solicitação da mesma maneira que o return?
	}
	return user
}
