package app

import (
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
	c.JSON(http.StatusOK, userCollection.Users[userID])
}

func CreatePhoto(c *gin.Context) {
	//need to check if user exists
	photo := photo.NewPhoto()

	if err := c.Bind(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// for _, v := range userCollection.Users {
	// 	if photo.UserID == _ {
	// 		errors.New("Area calculation failed, radius is less than zero")
	// 	}
	// }
	fmt.Println(photo.UserID)
	photoCollection.AddPhoto(photo)
	c.JSON(http.StatusOK, photoCollection)
}

func GetAllPhoto(c *gin.Context) {
	c.JSON(http.StatusOK, photoCollection)
}

func Start() {
	r := gin.Default()
	r.POST("/user", CreateUser)
	r.GET("/user", GetAllUser)
	r.GET("/user/:userID", GetUser)

	r.POST("/photo", CreatePhoto)
	r.GET("/photo", GetAllPhoto)
	r.Run() // listen and serve on 0.0.0.0:8080
}
