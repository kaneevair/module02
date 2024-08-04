package main

import "fmt"

func main() {
	workWeek := []string{"понедельник", "вторник", "среда", "четверг", "пятница"}
	weekend := []string{"суббота", "воскресенье"}

	var week = append(workWeek, weekend...)
	fmt.Print(week)
}
