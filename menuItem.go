package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"os"
	"strconv"
	"strings"
)

type Category string

const (
	Appetizer Category = "appetizer"
	Entree    Category = "entree"
	Dessert   Category = "dessert"
	Drink     Category = "drink"
)

type MenuItem struct {
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	Category Category  `json:"category"`
	ID       uuid.UUID `json:"id"`
}

var ErrInvalidPrice = errors.New("cannot create item with price of 0 or less")
var ErrInvalidName = errors.New("cannot create item with empty string")
var ErrAlreadyInList = errors.New("item already in list")

func CreateMenuItem(name string, price float64, cat Category) (MenuItem, error) {
	id := uuid.New()
	inputString := strings.TrimSpace(name)

	if inputString == "" {
		return MenuItem{}, ErrInvalidName
	}

	if price <= 0 {
		return MenuItem{}, ErrInvalidPrice
	}

	menuItem := MenuItem{name, price, cat, id}

	Save(&menuItem)
	return menuItem, nil
}

var categoryName = map[string]Category{
	"0": Appetizer,
	"1": Entree,
	"2": Dessert,
	"3": Drink,
}

func ReturnCategory(s string) Category {
	return categoryName[s]
}

func EditMenuItem(m *MenuItem) {
	//find item in array first,

}

func Save(m *MenuItem) error {
	//Check for existing file and read contents
	menuItemsFile, err := os.ReadFile("menuItems.json")
	if err != nil {
		fmt.Println(err)
	}
	//Unmarshal JSON data
	var menuItemsData []MenuItem
	err = json.Unmarshal(menuItemsFile, &menuItemsData)
	if err != nil {
		fmt.Println(err)
	}
	// Check for menu item name in data
	exists := Find(m, &menuItemsData)
	if exists {
		fmt.Printf("Error occurred, %v", ErrAlreadyInList)
		return ErrAlreadyInList
	}
	//Append new data
	newMenuItem := m
	menuItemsData = append(menuItemsData, *newMenuItem)

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
	fmt.Printf("%v saved\n", m.Name)
	return nil
}

func Find(i *MenuItem, items *[]MenuItem) bool {
	for _, value := range *items {
		if value.Name == i.Name {
			return true
		}
	}
	return false
}

func (i *MenuItem) PrintDetails() string {
	stringVal := strconv.FormatFloat(i.Price, 'f', -1, 64)

	printString := fmt.Sprintf("%v, $%v", i.Name, stringVal)
	fmt.Println(printString)

	return printString
}
