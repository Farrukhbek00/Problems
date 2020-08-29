package taskList

import (
	"testing"
	"time"
)

var tm TaskManagerI

func TestNewTaskManager(t *testing.T) {
	var err error
	tm, err = NewTaskManager()
	if err != nil {
		t.Error("Can not connect a database: ", err)
	}
}

func TestTaskManager_Add(t *testing.T) {
	time, _ := time.Parse("2006-02-01", "2020-02-02")
	task := Task{Id: 1, Assignee: "Farruh", Title: "Qale", Deadline: time}
	err := tm.Add(task)
	if err != nil {
		t.Error("Can not create a Task: ", err)
	}
}

func TestTaskManager_Update(t *testing.T) {
	time, _ := time.Parse("2020-08-28", "2022-22-22")
	task := Task{Assignee: "Sherzod1", Title: "Begin Bos", Deadline: time}

	err := tm.Update(1, task)
	if err != nil {
		t.Error("Cannot update a Task: ", err)
	}
}

func TestTaskManager_MakeDone(t *testing.T) {
	err := tm.MakeDone(1)

	if err != nil {
		t.Error("Cannot make true: ", err)
	}
}

func TestTaskManager_GetAll(t *testing.T) {
	_, err := tm.GetAll()

	if err != nil {
		t.Error("Cannot get all data: ", err)
	}
}

func TestTaskManager_Finished(t *testing.T) {
	_, err := tm.GetFinished()

	if err != nil {
		t.Error("Cannot get all finished tasks: ", err)
	}
}

func TestTaskManager_NotFinished(t *testing.T) {
	_, err := tm.GetNotFinished()

	if err != nil {
		t.Error("Cannot get not finished tasks: ", err)
	}
}

func TestTaskManager_ListAll(t *testing.T) {
	err := tm.ListAll()

	if err != nil {
		t.Error("Cannot lsit all", err)
	}
}

func TestTaskManager_Delete(t *testing.T) {
	err := tm.Delete(1)

	if err != nil {
		t.Error("Cannot delete", err)
	}
}
