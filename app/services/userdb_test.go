package services

import "testing"

var testDB = NewUsersDB()

func TestInsertRetrieve(t *testing.T) {
	u, err := testDB.Insert("test@test.com", "Test Me")
	if err != nil {
		t.Fatalf("Unable to insert: %s", err)
	}

	if u.Email != "test@test.com" || u.Fullname != "Test Me" {
		t.Fatalf("Insert returned record with unexpected values: %#v", u)
	}

	u, err = testDB.GetByKey("email", "test@test.com")
	if err != nil {
		t.Fatalf("Unable to lookup user by email: %s", err)
	}

	if u.Email != "test@test.com" || u.Fullname != "Test Me" {
		t.Fatalf("Lookup returned wrong record: %#v", u)
	}

	u, err = testDB.GetByKey("fullname", "Test Me")
	if err != nil {
		t.Fatalf("Unable to lookup user by fullname: %s", err)
	}

	if u.Email != "test@test.com" || u.Fullname != "Test Me" {
		t.Fatalf("Lookup returned wrong record: %#v", u)
	}
}
