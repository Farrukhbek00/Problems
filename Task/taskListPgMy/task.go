package tasklistpgmy

import (
	_ "database/sql"
	"fmt"

	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type TaskManager struct {
	db *sqlx.DB
}

func NewTaskManager() (TaskManagerI, error) {
	tm := TaskManager{}
	var err error
	tm.db, err = sqlx.Connect("postgres", "user=postgres password=1234 dbname=test sslmode=disable")

	if err != nil {
		return nil, err
	}

	return &tm, nil
}

func (tm *TaskManager) Add(t Task) error {
	query := "INSERT INTO taskmy(id, name, note, worker, expiredate) VALUES($1, $2, $3, $4, $5)"
	tm.db.MustExec(query, t.Id, t.Name, t.Note, t.Worker, t.ExpireDate)

	return nil
}

func (tm *TaskManager) Update(id int, t Task) error {
	query := "UPDATE taskmy SET name=$1, note=$2, worker=$3, expiredate=$4 WHERE id=$5"
	tm.db.MustExec(query, t.Name, t.Note, t.Worker, t.ExpireDate, id)

	return nil
}

func (tm *TaskManager) Delete(id int) error {
	query := "DELETE FROM taskmy WHERE id=$1"
	tm.db.MustExec(query, id)

	return nil
}

func (tm *TaskManager) GetAll() ([]Task, error) {
	var tasks []Task
	rows, err := tm.db.Queryx("Select id, name, note, worker, expiredate FROM taskmy")

	for rows.Next() {
		var t Task
		err = rows.Scan(&t.Id, &t.Name, &t.Note, &t.Worker, &t.ExpireDate)
		tasks = append(tasks, t)
	}

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (tm *TaskManager) GetFreshTask() ([]Task, error) {
	now := time.Now()
	fmt.Println(now, "sddddddddddddddddddddddddsdsdssddsds")

	var tasks []Task
	rows, err := tm.db.Queryx("Select id, name, note, worker, expiredate FROM taskmy WHERE CURRENT_DATE <= expiredate")
	fmt.Println(rows)
	fmt.Println(err)
	for rows.Next() {
		var t Task
		err = rows.Scan(&t.Id, &t.Name, &t.Note, &t.Worker, &t.ExpireDate)
		tasks = append(tasks, t)
	}

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
