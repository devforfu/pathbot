package renderer

import (
	"fmt"
	"pathbot/locations"
)

func ShowLocation(location locations.PathBotLocation) {
	fmt.Println()
	fmt.Println(location.Message)
	fmt.Println(location.Description)
	fmt.Printf("Location exits: %v", location.Exits)
}
