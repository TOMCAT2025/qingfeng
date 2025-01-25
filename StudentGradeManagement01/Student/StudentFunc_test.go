package Student

import "testing"

// 定义一个全局的学生实例用于测试
var testStudent *Student

func init() {
	// 初始化学生实例
	testStudent = &Student{
		ID:     "12345",
		Name:   "John",
		Gender: "Male",
		Class:  "Class1",
		Grades: make(map[string]float64),
	}
}

// 测试AddStudent方法
func TestAddStudent(t *testing.T) {
	AddStudent(testStudent)
	// 检查是否成功添加到studentsMap中
	if _, ok := studentsMap[testStudent.ID]; !ok {
		t.Errorf("AddStudent failed, student with ID %s not added to the map", testStudent.ID)
	}
}

// 测试DeleteStudent方法
func TestDeleteStudent(t *testing.T) {
	AddStudent(testStudent)
	DeleteStudent(testStudent.ID)
	// 检查是否从studentsMap中成功删除
	if _, ok := studentsMap[testStudent.ID]; ok {
		t.Errorf("DeleteStudent failed, student with ID %s still exists in the map", testStudent.ID)
	}
}

// 测试UpdateStudent方法
func TestUpdateStudent(t *testing.T) {
	//先重新添加一遍，再更新
	AddStudent(testStudent)

	updateInfo := &Student{
		ID:     "12345",
		Name:   "Robert",
		Gender: "Male",
		Class:  "Class1",
		Grades: make(map[string]float64),
	}
	UpdateStudent(testStudent.ID, updateInfo)

	updatedStudent := studentsMap[testStudent.ID]
	if updatedStudent.Name != "Robert" {
		t.Errorf("UpdateStudent failed, name not updated correctly, got %s, expected Robert", updatedStudent.Name)
	}
}

// 测试QueryStudent方法
func TestQueryStudent(t *testing.T) {
	AddStudent(testStudent)

	queriedStudent := QueryStudent(testStudent.ID)
	if queriedStudent == nil {
		t.Errorf("QueryStudent failed, returned nil for existing student with ID %s", testStudent.ID)
	}
	if queriedStudent.ID != testStudent.ID {
		t.Errorf("QueryStudent failed, returned student with wrong ID, got %s, expected %s", queriedStudent.ID, testStudent.ID)
	}
}

// 测试AddGrade方法
func TestAddGrade(t *testing.T) {
	AddStudent(testStudent)

	AddGrade(testStudent, "Math", 90.0)
	if testStudent.Grades["Math"] != 90.0 {
		t.Errorf("AddGrade failed, grade not added correctly, got %f, expected 90.0", testStudent.Grades["Math"])
	}
}

// 测试DeleteGradeByCourse方法
func TestDeleteGradeByCourse(t *testing.T) {
	AddStudent(testStudent)
	testStudent.Grades = map[string]float64{"English": 80.0}

	DeleteGradeByCourse(testStudent, "English")
	if _, ok := testStudent.Grades["English"]; ok {
		t.Errorf("DeleteGradeByCourse failed, grade still exists after deletion for course English")
	}
}

// 测试UpdateGrade方法
func TestUpdateGrade(t *testing.T) {
	AddStudent(testStudent)
	testStudent.Grades = map[string]float64{"Physics": 70.0}

	UpdateGrade(testStudent, "Physics", 75.0)
	if testStudent.Grades["Physics"] != 75.0 {
		t.Errorf("UpdateGrade failed, grade not updated correctly, got %f, expected 75.0", testStudent.Grades["Physics"])
	}
}

// 测试QueryGrade方法
func TestQueryGrade(t *testing.T) {
	AddStudent(testStudent)
	queriedStudent := QueryGrade(testStudent.ID)
	if queriedStudent == nil {
		t.Errorf("QueryGrade failed, returned nil for existing student with ID %s", testStudent.ID)
	}
	if queriedStudent.ID != testStudent.ID {
		t.Errorf("QueryGrade failed, returned student with wrong ID, got %s, expected %s", queriedStudent.ID, testStudent.ID)
	}
}
