package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Prompt(q []string) string {
	return q[0]
}

func ItemPrompt() Item {

	var itemName string
	var itemPrice float64

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("What is the item name? ")
		name, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("an error occurred %v", err)
		}
		itemName = strings.TrimSpace(name)

		fmt.Print("What is the item cost? ")
		stringPrice, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred while reading the item name %v \n", err)
			continue
		}
		stringPrice = strings.TrimSpace(stringPrice)
		itemPrice, err = strconv.ParseFloat(stringPrice, 64)

		if err != nil {
			fmt.Printf("Invalid cost, '%s'.Please enter a valid number.\n", stringPrice)
			continue
		}

		// Verify item details
		fmt.Printf("Is this information correct? (y/n)\n%v %.2f ?\n", itemName, itemPrice)
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred while reading confirmation %v\n", err)
			continue
		}
		response = strings.TrimSpace(strings.ToLower(response))
		if response == "y" || response == "yes" {
			break
		}
	}
	item := Item{Name: itemName, Price: itemPrice}
	fmt.Printf("Item created: %+v\n", item)
	return item
}
