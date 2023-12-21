package main

import "fmt"

// User with id and first name
type user struct {
	id        int
	firstName string
}

// Temporary store that stores retrievable data
type mockDataStore struct {
	users map[int]user
}

func (md mockDataStore) GetUser(id int) (user, error) {
	usr, ok := md.users[id]
	if !ok {
		return user{}, fmt.Errorf("User not founf: %d", id)
	}
	return usr, nil
}

func (md mockDataStore) SaveUser(u user) error {
	_, ok := md.users[u.id]
	if ok {
		return fmt.Errorf("User already exists: %d", u.id)
	}
	md.users[u.id] = u
	return nil
}

/*
DataStore defines an interface for storing retrievable data
Any TYPE that implements this interface (has two methods) is also of type DataStore
This means any value of type mockDataStore is also type DataStore
This means we could have a value of type *sql.DB and it can also be of type DataStore
This means we can write functions to take type DataStore and use either of these:

	mockDataStore
	*sql.DB
*/
type DataStore interface {
	GetUser(id int) (user, error)
	SaveUser(u user) error
}

/*
This represents a service that stores retrievable data.
We will attach the method to service so that we cna use either of these:

	mockDataStore
	*sql.DB
*/
type Service struct {
	ds DataStore
}

func (s Service) GetUser(id int) (user, error) {
	return s.ds.GetUser(id)
}

func (s Service) SaveUser(u user) error {
	return s.ds.SaveUser(u)
}

func main() {
	db := mockDataStore{
		users: make(map[int]user),
	}

	srvc := Service{
		ds: db,
	}

	u1 := user{
		id:        1,
		firstName: "A",
	}

	err := srvc.SaveUser(u1)
	if err != nil {
		fmt.Println("Error while adding user:", err)
	}

	u1Ret, err := srvc.GetUser(1)
	if err != nil {
		fmt.Println("Error while getting user:", err)
	}

	fmt.Println(u1Ret)
	fmt.Println(u1)

}
