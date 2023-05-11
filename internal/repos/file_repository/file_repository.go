package file_repository

import (
	"bufio"
	"encoding/json"
	"fmt"
	"homework2/internal/helpers"
	"homework2/internal/models/user"
	"os"
)

type UserRepositoryI interface {
	Create(user *user.User) (int32, error)
	GetByEmail(email *string) (*user.User, error)
	GetAll() (*[]user.User, error)
	Update(UpdateUser *user.User) (*user.User, error)
	Delete(id int32) error
}

type UserFileRepository struct{}

func NewUserFileRepository() *UserFileRepository {
	return &UserFileRepository{}
}

func (ufr *UserFileRepository) Create(user *user.User) (user.User, error) {
	err := helpers.Create("userDataStorage", user)
	if err != nil {
		return *user, err
	}

	return *user, nil
}

func (ufe *UserFileRepository) GetByEmail(email *string) (*user.User, error) {

	file, err := os.OpenFile("internal/repos/file_repository/datastore/userDataStorage.txt", os.O_RDONLY, 0666)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	users := []*user.User{}
	for _, line := range lines {
		curr := new(user.User)
		err := json.Unmarshal([]byte(line), &curr)

		if err != nil {
			return nil, err
		}

		users = append(users, curr)
	}

	for _, curr := range users {
		if curr.Email == *email {
			return curr, nil // returning User with same email;
		}
	}

	return nil, fmt.Errorf("No user with such email: %s", *email)
}

func (ufe *UserFileRepository) GetAll() (*[]user.User, error) {

	file, err := os.OpenFile("internal/repos/file_repository/datastore/userDataStorage.txt", os.O_RDONLY, 0666)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	users := []user.User{}

	for _, line := range lines {
		curr := new(user.User)

		err := json.Unmarshal([]byte(line), &curr)

		if err != nil {
			return nil, err
		}

		users = append(users, *curr)
	}

	return &users, nil
}

func (ufe *UserFileRepository) Update(UpdateUser *user.User) (*user.User, error) {
	file, err := os.OpenFile("internal/repos/file_repository/datastore/userDataStorage.txt", os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	updatedUsers := []user.User{}

	for scanner.Scan() {
		line := scanner.Bytes()

		curr := &user.User{}
		err := json.Unmarshal(line, curr)
		if err != nil {
			return nil, err
		}

		if curr.ID == UpdateUser.ID {
			curr.FirstName = UpdateUser.FirstName
			curr.LastName = UpdateUser.LastName
			curr.Email = UpdateUser.Email
			curr.Password = UpdateUser.Password
		}
		updatedUsers = append(updatedUsers, *curr)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	err = file.Truncate(0)
	if err != nil {
		return nil, err
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	for _, u := range updatedUsers {
		updatedData, err := json.Marshal(u)
		if err != nil {
			return nil, err
		}
		_, err = file.Write(append(updatedData, []byte("\n")...))
		if err != nil {
			return nil, err
		}
	}

	return UpdateUser, nil
}

func (ufe *UserFileRepository) Delete(id int) error {
	file, err := os.OpenFile("internal/repos/file_repository/datastore/userDataStorage.txt", os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	deleted := false
	for scanner.Scan() {
		u := new(user.User)
		if err := json.Unmarshal([]byte(scanner.Text()), &u); err != nil {
			return err
		}
		if u.ID != id {
			lines = append(lines, scanner.Text())
		} else {
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("no such user in database")
	}

	if _, err := file.Seek(0, 0); err != nil {
		return err
	}
	if err := file.Truncate(0); err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	if err := writer.Flush(); err != nil {
		return err
	}

	return nil
}
