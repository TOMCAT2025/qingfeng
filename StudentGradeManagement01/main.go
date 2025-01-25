package main

import (
	"StudentGradeManagement01/Student"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/AddStudentHandler", Student.AddStudentHandler)
	r.DELETE("/DeleteStudentHandler", Student.DeleteStudentHandler)
	r.PUT("/UpdateStudentHandler", Student.UpdateStudentHandler)
	r.GET("/QueryStudentHandler", Student.QueryStudentHandler)
	r.POST("/AddGradeHandler", Student.AddGradeHandler)
	r.DELETE("/DeleteGradeHandler", Student.DeleteGradeHandler)
	r.PUT("/UpdateGradeHandler", Student.UpdateGradeHandler)
	r.GET("/QueryGradeHandler", Student.QueryGradeHandler)
	/*
		filePath为需要导入的csv文件地址，请替换后运行
	*/
	filePath := "C:\\Users\\moon\\Desktop\\student.csv"
	Student.AddMassivelyStudents(filePath)
	err := r.Run()
	if err != nil {
		return
	}

}
