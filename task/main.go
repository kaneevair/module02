package main

import (
	"fmt"
)

func main() {
	books := map[string]map[string][]string{
		"0 читатель": {
			"Книги": {},
		},
		"1 читатель": {
			"Книги": {
				"Гарри Поттер", "Мы", "Вий", "Черчиль и Оруэлл", "Мифы Ктулху", "Дом в котором",
			},
		},
		"2 читатель": {
			"Журналы": {
				"Cosmopolitan", "GQ", "Esquire", "Glamour", "Vogue",
			},
		},
		"3 читатель": {
			"Книги": {
				"Война и мир",
			},
		},
		"4 читатель": {
			"Газеты": {
				"Известия", "Коммерсантъ", "Ведомости",
			},
		},
		"5 читатель": {},
	}

	books["6 читатель"] = map[string][]string{
		"Книги": {
			"Властелин колец", "Винни Пух", "Унесённые ветром", "Столпы Земли", "Дюна", "Эмма", "1984",
		},
		"Журналы": {
			"Власть",
		},
	}
	fmt.Println(getCountOfReadersWithBooks(books))
	getCountOfBooksPerReader(books)
}

func getCountOfReadersWithBooks(books map[string]map[string][]string) int {
	//вывести количество читателей у кого издания на руках
	var count int = 0
	for _, v := range books {
		for _, v1 := range v {
			if len(v1) != 0 {
				count = count + 1
				break
			}
		}
	}
	return count
}

func getCountOfBooksPerReader(books map[string]map[string][]string) {
	//вывести сколько изданий у каждого читателя
	for k, v := range books {
		var count int = 0
		fmt.Print(k, ": ")
		for _, v1 := range v {
			count += len(v1)
		}
		fmt.Println("количество изданий = ", count)
	}
}
