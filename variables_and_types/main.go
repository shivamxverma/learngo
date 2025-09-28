package main

import "fmt"

const LoginToken string = "asdf4w4e32"; // This is same as putting public keyword in frontend of variables it behaves like public

func main()  {
	var username string = "shivam@12"
	fmt.Println(username);

	var isChecked bool = false;

	if(isChecked){
		fmt.Println("Yes Checked");
	} else {
		fmt.Println("Not Checked");
	}

	var choclates int = 90;
	var sumofChoc int = 0;
	var i int;

	for i = 0; i <= choclates; i++ {
		sumofChoc += i
	}
	fmt.Println(sumofChoc);

	// Implicit type 

	var yourname = "shivam"
	fmt.Println(yourname);

	// No var Style 

	noofuser := 8999;  // Inside the methods only not outside the methods
	fmt.Println(noofuser);
}

