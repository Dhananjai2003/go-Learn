package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"bufio"
)

var statusMarkers = map[string]string{"1":"Complete","0":"Incomplete"}

func showMenu(){
	fmt.Println("1 - Add Task")
	fmt.Println("2 - Mark Task Done")
	fmt.Println("3 - Delete Task")
	fmt.Println("4 - Reset Tasks")
	fmt.Println("5 - Show Task")
}

func showTasks(){

	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("Some error in opening file : ",err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	tasks, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error in reading tasks ", err)
		return
	}
	fmt.Println("[No. Task Status]")
	for i, row := range tasks {
		fmt.Println("[",i+1,row[0],statusMarkers[row[1]],"]")
	}

}

func addTask(task string) {
	file, err := os.OpenFile("data.csv",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil {
		fmt.Println("Cannot acces data ",err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	newRow := []string{task, "0"}

	err = writer.Write(newRow)
	if err != nil {
		fmt.Println("Error writing to CSV:", err)
		return
	}
}

func markTaskDone(taskNumber int) {

	// error for taskNumber below 1
	file, err := os.Open("data.csv")
	if err!= nil {
		fmt.Println("Error in reading database : ",err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	tasks, err := reader.ReadAll()
	// error for tasks out of range

	tasks[taskNumber-1][1]="1"

	file, err = os.Create("data.csv")
	if err != nil {
		fmt.Println("Error in reading database : ",err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(tasks)
	if err != nil {
		fmt.Println("Error in writing to database :",err)
		return
	}
}

func deleteTasks(taskNumber int) {

	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("Error in reading database : ", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	tasks ,err := reader.ReadAll()
	if err!=nil {
		fmt.Println("Error in reading database : ",err)
		return
	}

	tasks = append(tasks[:taskNumber-1],tasks[taskNumber:]...)

	file, err = os.Create("data.csv")
	if err != nil {
		fmt.Println("Error in writing to database : ",err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(tasks)
	if err != nil {
		fmt.Println("Error in writing to database : ",err)
		return
	}
}

func resetTasks() {

	file, err := os.Create("data.csv")
	if err != nil {
		return
	}
	defer file.Close()

}

func userLogic(){

	var userChoice int 

	fmt.Print("Enter Choice : ")

	fmt.Scan(&userChoice)

	switch userChoice {
	case 1:

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter you Task : ")

		task, _ := reader.ReadString('\n')
		task = task[:len(task)-1]
		addTask(task)

	case 2:

		var taskNumber int

		fmt.Scan(&taskNumber)

		markTaskDone(taskNumber)

	case 3:

		var taskNumber int

		fmt.Scan(&taskNumber)

		deleteTasks(taskNumber)

	case 4:
		resetTasks()

	case 5:
		showTasks()

	default:
		fmt.Println("Sorry Wrong Option ")
		userLogic()
		return
	}

}

func main(){

	
	showTasks()
	showMenu()
	userLogic()

	main()
	
}