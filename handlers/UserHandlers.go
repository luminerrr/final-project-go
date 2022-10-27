package handlers

import (
	"final-project-go/database"
	"final-project-go/helpers"
	"final-project-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var appJSON = "application/json"

func UserRegister(c *gin.Context) {
	
	db := database.GetDB()
	user := models.User{}
	contentType := helpers.GetContentType(c)

	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	

	err := db.Debug().Create(&user).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"error message" : "error while creating user data",
			"error" : err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"age" : user.Age,
		"email": user.Email,
		"id" : user.ID,
		"username": user.Username,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	user := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	password = user.Password
	err := db.Debug().Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error message" : "unauthorized",
			"error" : err,
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(user.Password), []byte(password))
	if !comparePass {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error message" : "invalid password",
			"error" : err,
		})
		return
	}

	token := helpers.GenerateToken(user.ID, user.Email)
	c.JSON(http.StatusOK, gin.H{
		"token" : token,
	})
}
