package controllers

import (
	"net/http"

	"github.com/alissonit/go-api-rest-gin/database"
	"github.com/alissonit/go-api-rest-gin/models"
	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

func GetStudentsDeleted(c *gin.Context) {
	var students []models.Student
	database.DB.Unscoped().Where("deleted_at IS NOT NULL").Find(&students)
	c.JSON(http.StatusOK, students)
}

func Greeting(c *gin.Context) {
	name := c.Params.ByName("name")

	c.JSON(200, gin.H{
		"message": "Hello " + name,
	})
}

func CreateNewStudent(c *gin.Context) {
	var newStudent models.Student
	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidateStudent(newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&newStudent)
	c.JSON(http.StatusOK, newStudent)
}

func GetStudentsByID(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found!",
		})
		return
	}

	if err := models.ValidateStudent(student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&student, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Student deleted successfully!",
	})
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	// check if the student exists
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found!",
		})
		return
	}

	// validate the input
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidateStudent(student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&student).UpdateColumns(student).Where("id = ?", id)
	c.JSON(http.StatusOK, student)
}

func FindStudentByCPF(c *gin.Context) {
	var student models.Student
	cpf := c.Params.ByName("cpf")
	// another way to do the same thing
	// database.DB.Where("cpf = ?", cpf).First(&student)
	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found!",
		})
		return
	}
	c.JSON(http.StatusOK, student)
}

func ShowPageIndex(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"students": students,
	})
}

func ShowPageNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", gin.H{})
}
