package main

import (
	"fmt"
	"github.com/google/uuid"
	"os"
)

func main() {
	ItemPrompt(os.Stdin)
	uuid := uuid.New()

	fmt.Println("Generated UUIDuuid:", uuid.String())
}
