package app

import (
	"fmt"
	"net/http"
	"shade/config"
	"shade/controllers"
	"shade/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	err error
)

func CreatePhoto(c *gin.Context) {
	// photo := photo.NewPhoto()
	// if err := c.Bind(&photo); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// for _, v := range userCollection.Users {
	// 	if photo.UserID == v.ID {
	// 		photoCollection.AddPhoto(photo)
	// 		c.JSON(http.StatusOK, photoCollection)
	// 		return
	// 	}
	// }
	// c.JSON(http.StatusBadRequest, errors.New("ID does not exist"))
	return
}

// func GetAllPhoto(c *gin.Context) {
// 	c.JSON(http.StatusOK, photoCollection)
// }

// func GetPhoto(c *gin.Context) {
// 	photoID, err := strconv.Atoi(c.Param("photoID"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	for _, v := range photoCollection.Photos {
// 		if photoID == v.ID {
// 			c.JSON(http.StatusOK, v)
// 			return
// 		}
// 	}
// 	c.JSON(http.StatusBadRequest, errors.New("ID does not exist"))
// 	return
// }

func AddUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := models.AddUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func Start() {
	r := gin.Default()
	r.POST("/user", controllers.AddUser)
	r.GET("/user", controllers.GetUsers)
	r.GET("/user/:id", controllers.GetUser)

	r.POST("/photo", controllers.AddPhoto)
	// r.GET("/photo", GetAllPhoto)
	// r.GET("photo/:photoID", GetPhoto)

	config.DB, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=shade password=admin sslmode=disable")
	if err != nil {
		fmt.Println("status: ", err)
	}

	defer config.DB.Close()

	config.DB.AutoMigrate(&models.User{}, &models.Photo{})

	r.Run() // listen and serve on 0.0.0.0:8080
}
