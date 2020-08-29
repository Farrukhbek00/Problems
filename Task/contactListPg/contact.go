package contactListPg

import (
	_ "database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ContactManager struct {
	db *sqlx.DB
}

func NewContactManager() (ContactManagerI, error) {
	cm := ContactManager{}
	var err error

	cm.db, err = sqlx.Connect("postgres", "user=postgres password=1234 dbname=test sslmode=disable")

	if err != nil {
		return nil, err
	}

	return &cm, nil
}

func (cm *ContactManager) Add(c Contact) error {
	query := "INSERT INTO contact(id, name, age, gender, phone) VALUES($1, $2, $3, $4, $5)"
	cm.db.MustExec(query, c.Id, c.Name, c.Age, c.Gender, c.Phone)

	return nil
}

func (cm *ContactManager) Update(id int, c Contact) error {
	query := "UPDATE contact SET name=$1, age=$2, gender=$3, phone=$4 WHERE id=$5"
	cm.db.MustExec(query, c.Name, c.Age, c.Gender, c.Phone, id)

	return nil
}

func (cm *ContactManager) Delete(id int) error {
	query := "DELETE FROM contact WHERE id=$1"
	cm.db.MustExec(query, id)

	return nil
}

func (cm *ContactManager) GetAll() ([]Contact, error) {
	contacts := []Contact{}
	rows, err := cm.db.Queryx("Select id, name, age, gender, phone FROM contact")

	for rows.Next() {
		var c Contact
		err = rows.Scan(&c.Id, &c.Name, &c.Age, &c.Gender, &c.Phone)
		contacts = append(contacts, c)
	}

	if err != nil {
		return nil, err
	}

	return contacts, err
}

func (cm *ContactManager) ListAll() error {
	contacts := []Contact{}
	rows, err := cm.db.Queryx("Select id, name, age, gender, phone FROM contact")

	for rows.Next() {
		var c Contact
		err = rows.Scan(&c.Id, &c.Name, &c.Age, &c.Gender, &c.Phone)
		contacts = append(contacts, c)
	}

	if err != nil {
		return err
	}

	fmt.Println(contacts)
	return nil
}
