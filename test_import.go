package main

import (
	"fmt"

	gamev1 "github.com/ssargent/aether-core-editor/gen/game/v1"
)

func main() {
	fmt.Printf("WorldResponse type: %T\n", &gamev1.WorldResponse{})
}
