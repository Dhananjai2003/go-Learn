package main

import (
	"fmt"
	"os"
	"encoding/csv"
)

var statusMarkers = map[string]string{"1":"Complete","0":"Incomplete"}

func showMenu(){
	fmt.Println("1 - Add Task")
	fmt.Println("2 - Mark Task Done")
	fmt.Println("3 - Delete Task")
	fmt.Println("4 - Reset Tasks")
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

	reader = csv.NewReader(file)

	tasks, err := reader.ReadAll()

	// error for tasks out of range


}

func main(){

	
	showTasks()
	showMenu()

	addTask("test2")
	showTasks()
}