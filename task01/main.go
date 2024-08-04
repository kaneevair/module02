package main

import "fmt"

func main() {
	week := []string{"понедельник", "вторник", "среда", "четверг", "пятница", "суббота", "воскресенье"}
	workWeek := make([]string, 5, 5)
	copyWeek := copy(workWeek, week)

	fmt.Println(workWeek, copyWeek)

	week = week[copyWeek:]
	fmt.Println(week)

}
