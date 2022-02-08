package user

import "fmt"

var (
	Fname string
	Lname string
)

func User_Info() {

	fmt.Print("Enter Your First Name :", "")
	fmt.Scanf("%s\n", &Fname)
	fmt.Print("Enter Your Last  Name :", "")
	fmt.Scanf("%s\n", &Lname)
	fmt.Println("Welcome To The Quiz Mr./Ms./Mrs. ", Fname, Lname)

}
