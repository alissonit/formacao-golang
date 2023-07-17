package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/alissonit/go-api-rest-gin/controllers"
	"github.com/alissonit/go-api-rest-gin/database"
	"github.com/alissonit/go-api-rest-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRoutesTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func RunAssertions(t *testing.T, resp *httptest.ResponseRecorder, expectedStatusCode int, expectedResponse string) {
	assert.Equal(t, expectedStatusCode, resp.Code, "OK response is expected")
	responseBody, _ := io.ReadAll(resp.Body)
	assert.Equal(t, expectedResponse, string(responseBody), "Response body is not expected")
}

func CreateStudentMock() {
	student := models.Student{Name: "Alisson Test", CPF: "12345678910", RG: "123456789"}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteStudentMock() {
	database.DB.Delete(&models.Student{}, ID)
}

func TestVerifyStatusCodeGreeting(t *testing.T) {
	r := SetupRoutesTest()
	r.GET("/:name", controllers.Greeting)
	// make a request to the /:name route
	req, _ := http.NewRequest("GET", "/Alisson", nil)
	// create a response recorder
	resp := httptest.NewRecorder()
	// serve the request to the recorder
	r.ServeHTTP(resp, req)
	//  Assert that the status code is 200
	RunAssertions(t, resp, http.StatusOK, "{\"message\":\"Hello Alisson\"}")
}

func MakeRequest(t *testing.T, r *gin.Engine, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	return resp
}
func TestGetStudentHandler(t *testing.T) {
	database.ConnectToDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutesTest()
	r.GET("/students", controllers.GetStudents)
	resp := MakeRequest(t, r, "GET", "/students", nil)
	assert.Equal(t, http.StatusOK, resp.Code, "OK response is expected")
}

func TestGetStudentByIDHandler(t *testing.T) {
	database.ConnectToDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutesTest()
	r.GET("/students/:id", controllers.GetStudentsByID)
	resp := MakeRequest(t, r, "GET", "/students/"+strconv.Itoa(ID), nil)
	assert.Equal(t, http.StatusOK, resp.Code, "OK response is expected")
}

func TestDeleteStudentHandler(t *testing.T) {
	database.ConnectToDatabase()
	CreateStudentMock()
	r := SetupRoutesTest()
	r.DELETE("/students/:id", controllers.DeleteStudent)
	resp := MakeRequest(t, r, "DELETE", "/students/"+strconv.Itoa(ID), nil)
	assert.Equal(t, http.StatusOK, resp.Code, "OK response is expected")
}

func TestUpdateStudentHandler(t *testing.T) {
	database.ConnectToDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()
	r := SetupRoutesTest()
	r.PATCH("/students/:id", controllers.UpdateStudent)
	student := models.Student{Name: "Alisson Test", CPF: "12345678910", RG: "123456789"}
	studentJSON, _ := json.Marshal(student)
	resp := MakeRequest(t, r, "PATCH", "/students/"+strconv.Itoa(ID), bytes.NewBuffer(studentJSON))
	var studentResponse models.Student
	json.Unmarshal(resp.Body.Bytes(), &studentResponse)

	assert.Equal(t, "12345678910", studentResponse.CPF, "OK response is expected")
	assert.Equal(t, "123456789", studentResponse.RG, "OK response is expected")
	assert.Equal(t, "Alisson Test", studentResponse.Name, "OK response is expected")
}
