package main

import (
	"fmt"
)

func main() {
	leaguesTitle := make(map[string]int)
	leaguesTitle["Sunderland"] = 6
	leaguesTitle["Newcastle"] = 4

	fmt.Println(leaguesTitle)

	recentHead2Head := map[string]int{
		"League 1": 1,
		"League 2": 2,
	}

	fmt.Println(recentHead2Head)

	for key, value := range recentHead2Head {
		fmt.Printf("\nKey is %v, Value is %v", key, value)
	}

	fmt.Println("\nAfter map delete")
	delete(recentHead2Head, "League 1")
	for key, value := range recentHead2Head {
		fmt.Printf("\nKey is %v, Value is %v", key, value)
	}
}
