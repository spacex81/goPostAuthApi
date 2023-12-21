package controller

import (
	"blog/api/service"
	"blog/models"
	"blog/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

type UserController struct {
	service service.UserService
}

func NewUserController(s service.UserService) UserController {
	return UserController{
		service: s,
	}
}

func (u *UserController) CreateUser(c *gin.Context) {
	var user models.UserRegister
	if err := c.ShouldBindJSON(&user); err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Invalid Json Provided")
		return
	}

	fmt.Println(user)

	hashPassword, _ := util.HashPassword(user.Password)
	user.Password = hashPassword

	err := u.service.CreateUser(user)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Failed to create user")
		return
	}

	util.SuccessJSON(c, http.StatusOK, "Successfully Created user")
}

func (u *UserController) LoginUser(c *gin.Context) {
	var user models.UserLogin
	var hmacSampleSecret []byte
	if err := c.ShouldBindJSON(&user); err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Invalid LoginUser Credentials")
		return
	}
	dbUser, err := u.service.LoginUser(user)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Invalid LoginUser Credentials")
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": dbUser,
		"exp":  time.Now().Add(time.Second * 30).Unix(),
	})

	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Failed to get token")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Token generated successfully",
		Data:    tokenString,
	}

	c.JSON(http.StatusOK, response)
}
