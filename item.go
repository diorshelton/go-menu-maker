package main

import (
	"fmt"
	"strconv"
)

type Item struct {
	Name  string
	Price float64
}

func (i Item) PrintDetails() string {
	stringVal := strconv.FormatFloat(i.Price, 'f', -1, 64)

	printString := fmt.Sprintf("%v, $%v", i.Name, stringVal)
	fmt.Println(printString)

	return printString
}

func (i Item) CreateItem(name string, price float64) Item {
	menuItem := Item{name, price}
	fmt.Println("Created", menuItem.Name)
	return menuItem
}
