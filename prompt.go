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
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What is the item name?")
	//Handle error
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// var cost float64
	fmt.Print("What is the item cost?")
	cost, _ := reader.ReadString('\n')
	cost = strings.TrimSpace(cost)
	price, _ := strconv.ParseFloat(cost, 64)

	item := Item{Name: name, Price: price}
	fmt.Print(item)
	return item
}
