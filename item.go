package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Item struct {
	Name  string
	Price float64
}

var ErrInvalidPrice = errors.New("cannot create item with price of 0 or less")

func (i Item) CreateItem(name string, price float64) (Item, error) {
	if price <= 0 {
		return Item{}, ErrInvalidPrice
	}
	menuItem := Item{name, price}
	fmt.Printf("Created %v", menuItem.Name)
	return menuItem, nil
}

func (i Item) PrintDetails() string {
	stringVal := strconv.FormatFloat(i.Price, 'f', -1, 64)

	printString := fmt.Sprintf("%v, $%v", i.Name, stringVal)
	fmt.Println(printString)

	return printString
}
