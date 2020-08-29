package contactListPg

import "testing"

var cm ContactManagerI

func TestNewContactManager(t *testing.T) {
	var err error
	cm, err = NewContactManager()

	if err != nil {
		t.Error("Cannot connect to database: ", err)
	}
}

// func TestContactManager_Add(t *testing.T) {
// 	contact := Contact{Id: 2, Name: "Farrukh", Age: 20, Gender: "male", Phone: "233432"}

// 	err := cm.Add(contact)

// 	if err != nil {
// 		t.Error("Cannot add contact: ", err)
// 	}
// }

// func TestContactManager_Update(t *testing.T) {
// 	contact := Contact{Name: "Zokirov", Age: 21, Gender: "male", Phone: "22323"}

// 	err := cm.Update(1, contact)

// 	if err != nil {
// 		t.Error("Cannot update contact: ", err)
// 	}
// }

// func TestContactManager_Delete(t *testing.T) {

// 	err := cm.Delete(1)

// 	if err != nil {
// 		t.Error("Cannot delete: ", err)
// 	}
// }

func TestContactManager_GetAll(t *testing.T) {

	_, err := cm.GetAll()

	if err != nil {
		t.Error("Cannot get all contacts", err)
	}
}

func TestContactManager_ListAll(t *testing.T) {
	err := cm.ListAll()

	if err != nil {
		t.Error("Cannot list all contacts", err)
	}
}
