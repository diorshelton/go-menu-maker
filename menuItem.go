package main

import (
	"errors"
	"fmt"
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
	fmt.Printf("Created %v", menuItem.Name)
	return menuItem, nil
}


func (i MenuItem) SaveItem() string {
	// // Marshal the struct into pretty-printed JSON
	// jsonData, err := json.MarshalIndent(i, "", " ")
	// if err != nil {
	// 	fmt.Println("Error marshaling JSON:", err)
	// }
	// // read file
	// menuItemsContent, err := os.ReadFile("menuItem.json")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// //Unmarshal JSON data
	// var itemData []I
	// //Write JSON data to a file
	// err = os.WriteFile("menuItems.json", jsonData, 0644)
	// if err != nil {
	// 	fmt.Println("Error writing to file:", err)
	// 	return err.Error()
	// }
	return i.Name + " " + "saved"
}

func (i MenuItem) PrintDetails() string {
	stringVal := strconv.FormatFloat(i.Price, 'f', -1, 64)

	printString := fmt.Sprintf("%v, $%v", i.Name, stringVal)
	fmt.Println(printString)

	return printString
}
