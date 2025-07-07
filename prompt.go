package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func ItemPrompt(reader io.Reader) MenuItem {
	buf := bufio.NewReader(reader)
	var itemName string
	var itemPrice float64
	var itemCategory Category

	for {
		fmt.Print("What is the item name?")
		name, err := buf.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred while reading the item name %v", err)
			continue
		}
		itemName = strings.TrimSpace(name)

		fmt.Printf("What is the item cost?")
		stringPrice, err := buf.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred while reading the item cost %v \n", err)
			continue
		}
		stringPrice = strings.TrimSpace(stringPrice)
		itemPrice, err = strconv.ParseFloat(stringPrice, 64)

		if err != nil {
			fmt.Printf("Invalid cost, '%s'.Please enter a valid price.\n", stringPrice)
			continue
		}
		//Input menuItem category
		fmt.Printf("What is the item category? \n 0:Appetizer, 1:Entree, 2:Dessert, 3:Drink")
		category, err := buf.ReadString('\n')

		if err != nil {
			fmt.Printf("An error occurred with item category %v", err)
			continue
		}

		itemCategory = ReturnCategory(strings.TrimSpace(category))

		// Verify item details
		fmt.Printf("Is this information correct? (y/n)\n%v %.2f %v?\n", itemName, itemPrice, itemCategory)
		response, err := buf.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred while reading confirmation %v\n", err)
			continue
		}
		response = strings.TrimSpace(strings.ToLower(response))
		if response == "y" || response == "yes" {
			break
		}
	}

	item := MenuItem{Name: itemName, Price: itemPrice, Category: itemCategory}
	fmt.Printf("Item created: %+v\n", item)
	return item
}