package routes

import (
	"github.com/alissonit/go-api-rest-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	// Handle all requests to the route /students
	r := gin.Default()
	r.GET("/students", controllers.GetStudents)
	r.GET("/students/deleted", controllers.GetStudentsDeleted)
	r.GET("/students/:id", controllers.GetStudentsByID)
	r.GET("/students/cpf/:cpf", controllers.FindStudentByCPF)
	r.GET(":name", controllers.Greeting)
	r.POST("/students", controllers.CreateNewStudent)
	r.POST("/students/:id", controllers.GetStudentsByID)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PUT("/students/:id", controllers.UpdateStudent)
	r.Run(":8000")
}
