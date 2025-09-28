package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "strconv"
	// "strings"
	"time"
)

func main()  {
	// welcome := "welcome to user input";

	// fmt.Println(welcome);

	// reader := bufio.NewReader(os.Stdin);

	// fmt.Println("Hey tell me your favorite cake");

	// userinput,_ := reader.ReadString('\n');

	// fmt.Println(userinput);

	// Converting string into number

	// inputData , _ := reader.ReadString('\n');

	// numrating,err := strconv.ParseFloat(strings.TrimSpace(inputData),64);

	// if err != nil{
	// 	fmt.Println(err);
	// } else {
	// 	fmt.Println("Yeah your Rating is Increased By 1",numrating+1);
	// }

	presentTime := time.Now();

	currentDate := presentTime.Format("01-02-2006");

	fmt.Println(currentDate);

	// Create new Date 

	createDate := time.Date(2026, time.August, 10, 23, 23, 0, 0, time.UTC);

	fmt.Println(createDate);

}
