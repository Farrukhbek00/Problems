package tasklistpgmy

import (
	"time"
)

type Task struct {
	Id         int       `db: "id"`
	Name       string    `db: "name"`
	Note       string    `db: "note"`
	Worker     string    `db: "worker"`
	ExpireDate time.Time `db: "expiredate"`
}

type TaskManagerI interface {
	Add(t Task) error
	Update(id int, t Task) error
	Delete(id int) error
	GetAll() ([]Task, error)
	GetFreshTask() ([]Task, error)
}
