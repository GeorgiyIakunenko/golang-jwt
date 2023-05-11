package main

import (
	"fmt"
	"homework2/internal/models/user"
	"homework2/internal/repos/file_repository"
	"os"
)

func main() {

	filepath := "internal/repos/file_repository/datastore/userDataStorage.txt"
	os.Remove(filepath)

	Repository := file_repository.NewUserFileRepository()

	GoshaUser := user.User{ID: 1, FirstName: "Georgiy", LastName: "Iakunenko", Email: "goshanyakunenko@gmail.com", Password: "2004"}
	MishaUser := user.User{ID: 2, FirstName: "Misha", LastName: "Kovach", Email: "misha@gmail.com", Password: "5060"}
	MishaUser2 := user.User{ID: 2, FirstName: "Misha2", LastName: "Kovash2", Email: "misha2@gmail.com", Password: "50602"}

	fmt.Println("\n-------------------Print added user---------------\n")

	user, _ := Repository.Create(&MishaUser)
	fmt.Println(user)
	user, _ = Repository.Create(&GoshaUser)
	fmt.Println(user)

	fmt.Println("\n-------------------Print user found by email goshanyakunenko@gmail.com---------\n")

	s := "goshanyakunenko@gmail.com"
	user_by_email, _ := Repository.GetByEmail(&s)
	fmt.Println(*user_by_email)

	fmt.Println("\n------------------GetAll users in repository----\n")

	all_users, _ := Repository.GetAll()
	fmt.Println(*all_users)

	fmt.Println("\n------------------Delete userGoshan users in repository----")
	Repository.Delete(1)

	fmt.Println("\n------------------GetAll users in repository----\n")
	all_users, _ = Repository.GetAll()
	fmt.Println(*all_users)

	fmt.Println("\n------------------Update user with id 2----\n")

	Repository.Update(&MishaUser2)

	all_users, _ = Repository.GetAll()
	fmt.Println(*all_users)

	/*var w io.Writer = os.Stdout

	// Call the Write method on the Writer
	n, err := w.Write([]byte("hello,ld\n"))
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("wrote %d bytes\n", n)*/

}
