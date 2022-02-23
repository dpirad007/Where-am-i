package Routes

import (
    "net/http"
    "where-am-i-server/Models"
    "github.com/gin-gonic/gin"
)

// Sample student data
var students = []Models.Student {
    { ID: "1", Name:"dion", RollNo:"501854", PhoneNo:[]string{""}, Email:[]string{""}, ParentPhoneNo:[]string{""}},
}

func getStudent(c* gin.Context) {
    c.IndentedJSON(http.StatusOK, students)
}

func SetupRouter() *gin.Engine {
    router := gin.Default()
    router.GET("/student/:rollNo", getStudent)
    router.GET("/teacher/:teacherId")

    return router
}
