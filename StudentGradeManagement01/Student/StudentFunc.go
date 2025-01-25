package Student

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
)

// studentsMap用于在内存中存储学生信息，以学号为键
var studentsMap = make(map[string]*Student)

// AddStudent 添加学生信息
func AddStudent(s *Student) {
	studentsMap[s.ID] = s
}

// DeleteStudent 删除学生信息
func DeleteStudent(id string) {
	delete(studentsMap, id)
}

// UpdateStudent 修改学生信息
func UpdateStudent(id string, updateInfo *Student) {
	s := studentsMap[id]
	s.Name = updateInfo.Name
	s.Gender = updateInfo.Gender
	s.Class = updateInfo.Class
}

// QueryStudent 根据传入的id返回需要添加成绩信息的学生
func QueryStudent(id string) *Student {
	return studentsMap[id]
}

// AddGrade 为指定学生添加成绩
func AddGrade(s *Student, courseName string, grade float64) {
	s.Grades[courseName] = grade
}

// DeleteGradeByCourse 根据学号和课程名删除学生的成绩
func DeleteGradeByCourse(s *Student, courseName string) {
	delete(s.Grades, courseName)
}

// UpdateGrade 根据学号和课程名修改学生成绩
func UpdateGrade(s *Student, courseName string, newGrade float64) {
	s.Grades[courseName] = newGrade
}

// QueryGrade 查询学生特定课程成绩
func QueryGrade(id string) *Student {
	return studentsMap[id]
}

// AddMassivelyStudents 从csv批量导入学生信息
// 邮箱附件中有个检测该方法用的Student.csv文件，接口文档的最后有测试这个方法遇到的一些小问题和注意
func AddMassivelyStudents(filePath string) {
	//定义一个互斥锁，保证对共享资源studentsMap的原子性访问
	var mutex sync.Mutex
	//创建一个通道用于在goroutine之间传递解析后的学生信息
	studentChan := make(chan Student)
	//定义一个WaitGroup，协调多个goroutine的执行
	var wg sync.WaitGroup
	//打开路径为filePath的文件，如果打开失败（`err!= nil`），则打印错误信息并返回函数。
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("打开文件失败: %v\n", err)
		return
	}
	//使用defer，在函数结束时关闭文件。
	defer file.Close()

	//创建CSV读取器，用于读取打开的CSV文件。
	reader := csv.NewReader(file)

	/*
		csv reader 会检查每一行的字段数量，如果不等于 FieldsPerRecord 就会抛出wrong number of fields的错误。
		FieldsPerRecord特性：
		为负值时，不检查。
		0，这是默认值。
		检查第一行的字段数量，然后赋值给 FieldsPerRecord变为正值
	*/
	/*将它设置为-1有两个原因：
		1.考虑到会根据第一行的字段数量进行报错，所以如果第一行数据的字段数错误，会导致后续所有的正常数据无法导入。
		2.第一行字段数正确的情况下，如果FieldsPerRecord不是负数，后续字段数出现错误，
	      只会输出wrong number of fields，触发不了验证字段数的报错机制
	*/
	reader.FieldsPerRecord = -1

	// 读取CSV文件并并发处理每一行
	go func() {
	outerLoop:
		for {
			record, err := reader.Read()
			//io.EOF是文件结尾，err == io.EOF代表读到了结尾，所以跳出循环
			if err == io.EOF {
				break
			} else if err != nil { //出现其他读取错误，运行该else if
				fmt.Printf("读取CSV行失败: %v\n", err)
				continue
			}
			// 数据验证，字段数不为四个时报错
			if len(record) != 4 {
				fmt.Printf("数据验证失败，报错的数据为: %v\n", record)
				continue
			}
			//判断每个字段是不是为空，若为空，先输出错误信息，再使用goto直接跳过该条数据
			for _, value := range record {
				if value == "" {
					fmt.Printf("出现字段为空的数据，请检查，报错的数据为: %v\n", record)
					goto outerLoop
				}
			}

			//根据读取的record创建学生实例
			student := Student{
				Name:   record[0],
				ID:     record[1],
				Gender: record[2],
				Class:  record[3],
			}
			//将WaitGroup的计数器加1，表示有一个新的goroutine任务要执行。所有goroutine执行完后再关闭通道
			wg.Add(1)
			//开启新的goroutine把建好的学生实例发送到channel中
			go func(s Student) {
				studentChan <- s
				wg.Done()
			}(student)
		}
		//wg.Wait()等待每个goroutine都发送完数据再关闭通道
		wg.Wait()
		close(studentChan)
	}()
	// 使用主goroutine处理从channel接收到的学生信息
	for student := range studentChan {
		mutex.Lock()
		if _, ok := studentsMap[student.ID]; ok {
			// 如果学生学号已存在，调用学生信息更新函数覆写修改
			UpdateStudent(student.ID, &student)
		} else {
			// 如果学生学号不存在，调用学生添加函数
			AddStudent(&student)
		}
		mutex.Unlock()
	}
	//该循环打印已经添加好的学生信息，仅用来检验有没有添加成功.
	for _, value := range studentsMap {
		fmt.Printf("姓名：%s 学号：%s 性别：%s 班级：%s\n", value.Name, value.ID, value.Gender, value.Class)
	}
}
