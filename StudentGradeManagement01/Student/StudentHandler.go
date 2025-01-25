package Student

/***
	该文件的函数用于处理从接口传来的信息，同时使用Gin框架的错误处理机制返回适当的HTTP状态码和错误信息
	同时调用基本增删查改函数完成对应操作
***/
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddStudentHandler 添加学生信息
func AddStudentHandler(c *gin.Context) {
	var newStudent Student
	/*
		c.BindJSON：Gin框架方法，用于将请求体中的JSON数据解析并绑定到newStudent变量上。
		如果解析过程中出现错误（例如请求体中的数据格式不符合Student类型的定义），则err不为nil，返回错误信息
	*/
	if err := c.BindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	//调用添加函数
	AddStudent(&newStudent)
	fmt.Println(newStudent)
	c.JSON(http.StatusCreated, gin.H{"message": "Student added successfully"})
}

// DeleteStudentHandler 删除学生信息
func DeleteStudentHandler(c *gin.Context) {
	id := c.Query("id")
	if _, ok := studentsMap[id]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	//调用删除函数
	DeleteStudent(id)
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}

// UpdateStudentHandler 修改学生信息
func UpdateStudentHandler(c *gin.Context) {
	id := c.Query("id")
	if _, ok := studentsMap[id]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	var updateInfo Student
	if err := c.BindJSON(&updateInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	//调用更新函数
	UpdateStudent(id, &updateInfo)
	c.JSON(http.StatusOK, gin.H{"message": "Student updated successfully"})
}

// QueryStudentHandler 查询学生基本信息和所有成绩
func QueryStudentHandler(c *gin.Context) {
	id := c.Query("id")
	if _, ok := studentsMap[id]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	//调用查询函数
	s := QueryStudent(id)
	c.JSON(http.StatusOK, gin.H{"student": s})
}

// AddGradeHandler 为学生添加成绩
func AddGradeHandler(c *gin.Context) {
	id := c.Query("id")
	println(id)
	if _, ok := studentsMap[id]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	var gradeInfo struct {
		CourseName string  `json:"course_name"`
		Grade      float64 `json:"grade"`
	}
	if err := c.BindJSON(&gradeInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	s := studentsMap[id] //根据id找到要添加成绩的学生
	//调用添加函数
	AddGrade(s, gradeInfo.CourseName, gradeInfo.Grade)
	c.JSON(http.StatusOK, gin.H{"message": "Grade added successfully"})
}

// DeleteGradeHandler 删除学生特定课程成绩
func DeleteGradeHandler(c *gin.Context) {
	id := c.Query("id")
	if _, ok := studentsMap[id]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	courseName := c.Query("course_name")
	s := studentsMap[id]
	DeleteGradeByCourse(s, courseName)
	c.JSON(http.StatusOK, gin.H{"message": "Grade deleted successfully"})
}

// UpdateGradeHandler 修改学生特定课程成绩
func UpdateGradeHandler(c *gin.Context) {
	id := c.Query("id")
	if _, ok := studentsMap[id]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	courseName := c.Query("course_name")
	newGradeStr := c.Query("new_grade")
	newGrade, _ := strconv.ParseFloat(newGradeStr, 64)
	if newGradeStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}
	s := studentsMap[id]
	UpdateGrade(s, courseName, newGrade)
	c.JSON(http.StatusOK, gin.H{"message": "Grade updated successfully"})
}

// QueryGradeHandler 查询学生特定课程成绩
func QueryGradeHandler(c *gin.Context) {
	id := c.Query("id")
	if _, ok := studentsMap[id]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	courseName := c.Query("course_name")
	s := QueryGrade(id)
	if grade, ok := s.Grades[courseName]; ok {
		c.JSON(http.StatusOK, gin.H{"grade": grade})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Grade not found"})
	}
}
