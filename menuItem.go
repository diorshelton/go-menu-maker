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

	fmt.Printf("Created %v\n", menuItem.Name)
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

func EditMenuItem(m *MenuItem, name string, price float64) {
	//find item in array first,
	//
	// then edit
	m.Name = name
	m.Price = price
}

func Save(m *MenuItem) ([]MenuItem, error) {
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
	exists := itemAlreadyExists(&menuItemsData, m)
	if exists {
		return menuItemsData, ErrAlreadyInList
	}
	return menuItemsData, nil
}

func itemAlreadyExists(items *[]MenuItem, i *MenuItem) bool {
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
