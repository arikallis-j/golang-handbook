package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func test_recursive() {
	PrintAllFilesWithFilter(".", ".go")
}

func PrintAllFiles(path string) {
	// получаем список всех элементов в папке (и файлов, и директорий)
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}
	//  проходим по списку
	for _, f := range files {
		// получаем имя элемента
		// filepath.Join — функция, которая собирает путь к элементу с разделителями
		filename := filepath.Join(path, f.Name())
		// печатаем имя элемента
		fmt.Println(filename)
		// если элемент — директория, то вызываем для него рекурсивно ту же функцию
		if f.IsDir() {
			PrintAllFiles(filename)
		}
	}
}

func filter_check(name string, filter string) (check bool) {
	for i := range name {
		for j := range filter {
			if string(name[i+j]) != string(filter[j]) {
				check = false
				break
			} else {
				check = true
			}
		}
		if check {
			return true
		}
	}
	return false
}

func PrintAllFilesWithFilter(path string, filter string) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}
	for _, f := range files {
		filename := filepath.Join(path, f.Name())
		if strings.Contains(filename, filter) {
			fmt.Println(filename)
		}

		if f.IsDir() {
			PrintAllFilesWithFilter(filename, filter)
		}
	}
}
