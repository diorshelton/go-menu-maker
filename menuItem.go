package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MenuItem struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var ErrInvalidPrice = errors.New("cannot create item with price of 0 or less")
var ErrInvalidName = errors.New("cannot create item with empty string")

func CreateMenuItem(name string, price float64) (MenuItem, error) {
	inputString := strings.TrimSpace(name)

	if inputString == "" {
		return MenuItem{}, ErrInvalidName
	}

	if price <= 0 {
		return MenuItem{}, ErrInvalidPrice
	}

	menuItem := MenuItem{name, price}

	fmt.Printf("Created %v\n", menuItem.Name)
	return menuItem, nil
}

func EditMenuItem(m *MenuItem, name string, price float64) {
//find item in array first,
//
// then edit
	m.Name = name
	m.Price = price
}

func (m *MenuItem) SaveItem()  {
	//Check for existing file and read contents
		menuItemsFile, err := os.ReadFile("menuItems.json")
		if err != nil {
			fmt.Println(err.Error())
		}
	//Unmarshal JSON data
		var menuItemsData []MenuItem
		err = json.Unmarshal(menuItemsFile, &menuItemsData)
		if err != nil {
			fmt.Println(err.Error())
		}
	//Append new data
		newMenuItem := m
		menuItemsData = append(menuItemsData,*newMenuItem)

	// Marshal the struct into pretty-printed JSON
	updatedJSONData, err := json.MarshalIndent(menuItemsData, "", " ")
	if err != nil {
		fmt.Printf("Error marshaling JSON:%v", err)
	}

	//Write JSON data to a file
	err = os.WriteFile("menuItems.json", updatedJSONData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
	fmt.Printf("%v saved\n",m.Name)
}

func (i *MenuItem) PrintDetails() string {
	stringVal := strconv.FormatFloat(i.Price, 'f', -1, 64)

	printString := fmt.Sprintf("%v, $%v", i.Name, stringVal)
	fmt.Println(printString)

	return printString
}
