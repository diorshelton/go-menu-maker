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
	stringVal := strconv.FormatFloat(i.Price, 'f',-1,64)
	details := fmt.Sprintf("%v, $%v", i.Name, stringVal)
	return details
}

func (i Item) CreateItem(name string, price float64) (string, *Item) {
	menuItem := Item{name, price}
	msg := fmt.Sprintf("%v was created", name)
	return msg, &menuItem
}