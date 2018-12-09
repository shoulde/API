package app

import (
	"errors"
	"fmt"
	"net/http"
	"shade/photo"
	"shade/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	userCollection  = user.NewUserCollection()
	photoCollection = photo.NewPhotoCollection()
)

func CreateUser(c *gin.Context) {
	user := user.NewUser()
	fmt.Println(user)
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(user)
	fmt.Println(userCollection)
	userCollection.AddUser(user)
	c.JSON(http.StatusOK, userCollection)
}

func GetAllUser(c *gin.Context) {
	c.JSON(http.StatusOK, userCollection)
}

func GetUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, v := range userCollection.Users {
		if userID == v.ID {
			c.JSON(http.StatusOK, v)
			return
		}
	}
	c.JSON(http.StatusBadRequest, errors.New("ID does not exist"))
	return
}

func CreatePhoto(c *gin.Context) {
	photo := photo.NewPhoto()
	if err := c.Bind(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, v := range userCollection.Users {
		if photo.UserID == v.ID {
			photoCollection.AddPhoto(photo)
			c.JSON(http.StatusOK, photoCollection)
			return
		}
	}
	c.JSON(http.StatusBadRequest, errors.New("ID does not exist"))
	return
}

func GetAllPhoto(c *gin.Context) {
	c.JSON(http.StatusOK, photoCollection)
}

func GetPhoto(c *gin.Context) {
	photoID, err := strconv.Atoi(c.Param("photoID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, v := range photoCollection.Photos {
		if photoID == v.ID {
			c.JSON(http.StatusOK, v)
			return
		}
	}
	c.JSON(http.StatusBadRequest, errors.New("ID does not exist"))
	return
}

func Start() {
	r := gin.Default()
	r.POST("/user", CreateUser)
	r.GET("/user", GetAllUser)
	r.GET("/user/:userID", GetUser)

	r.POST("/photo", CreatePhoto)
	r.GET("/photo", GetAllPhoto)
	r.GET("photo/:photoID", GetPhoto)
	r.Run() // listen and serve on 0.0.0.0:8080
}
