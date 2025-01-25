package Student

// Student 结构体定义学生的基本信息
type Student struct {
	Name   string
	ID     string
	Gender string
	Class  string
	Grades map[string]float64 // 存储课程名和成绩的映射
}
