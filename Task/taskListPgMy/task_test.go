package tasklistpgmy

import (
	"fmt"
	"testing"
)

var tm TaskManagerI

func TestNewTaskManager(t *testing.T) {
	var err error
	tm, err = NewTaskManager()
	if err != nil {
		t.Error("Cannot connect to database: ", err)
	}
}

// func TestTaskManager_Add(t *testing.T) {
// 	time, _ := time.Parse("2011-11-11", "2030-09-09")
// 	task := Task{Id: 3, Name: "Qale", Note: "Task For rebasing", Worker: "Jamshid", ExpireDate: time}

// 	err := tm.Add(task)

// 	if err != nil {
// 		t.Error("Cannot add task: ", err)
// 	}
// }

// func TestTaskManager_Update(t *testing.T) {
// 	time, _ := time.Parse("2011-11-11", "2010-09-09")
// 	task := Task{Name: "Bola", Note: "Task For rebasing", Worker: "Jamshid", ExpireDate: time}

// 	err := tm.Update(3, task)

// 	if err != nil {
// 		t.Error("Cannot update task: ", err)
// 	}

// }

// func TestTaskManager_Delete(t *testing.T) {
// 	err := tm.Delete(2)

// 	if err != nil {
// 		t.Error("Cannot delete task: ", err)
// 	}
// }

// func TestTaskManager_GetAll(t *testing.T) {

// 	tasks, err := tm.GetAll()
// 	fmt.Println(tasks)
// 	if err != nil {
// 		t.Error("Cannot get All: ", err)
// 	}
// }

func TestTaskManager_GetFreshTask(t *testing.T) {
	tasks, err := tm.GetFreshTask()
	fmt.Println(tasks)
	if err != nil {
		t.Error("Cannot get fresh tasks: ", err)
	}
}
