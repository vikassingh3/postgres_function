package handler

import (
	"net/http"
	"path/config"
	"path/model"

	"github.com/gin-gonic/gin"
)

func GetAllStudent(c *gin.Context) {

	db := config.DB

	var data []model.Student

	db.Find(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func CretateStudent(c *gin.Context) {
	db := config.DB

	var data model.Student
	c.BindJSON(&data)

	db.Exec("CALL AddStudent($1 ,$2 ,$3,$4,$5,$6) ", data.ID, data.Name, data.Dateofbirth, data.City, data.Designation, data.Joiningdate).Scan(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func UpdateStudent(c *gin.Context) {
	db := config.DB

	var data model.Student
	c.BindJSON(&data)

	id := c.Param("id")

	db.Exec("CALL UpdateStudent($1 ,$2 ,$3,$4,$5,$6) ", id, data.Name, data.Dateofbirth, data.City, data.Designation, data.Joiningdate).Scan(&data)

	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})

}

func DeleteStudent(c *gin.Context) {
	db := config.DB

	id := c.Param("id")

	db.Exec("CALL DeleteStudent($1)", id)

	c.JSON(http.StatusOK, gin.H{
		"message": "deletion succressfully",
	})

}
