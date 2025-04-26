package main

import (
	"fmt"
)

type StackPromotion struct {
	ID    int64
	Title string
}

func main() {
	// Sample map with some data
	stackPromotions := map[int64][]StackPromotion{
		101: {
			{ID: 1, Title: "Promo 1"},
			{ID: 2, Title: "Promo 2"},
		},
		202: {
			{ID: 3, Title: "Promo 3"},
		},
	}
	fmt.Println("Start testing")
	var userID int64 = 101 // key that exists
	promos, ok := stackPromotions[userID]
	fmt.Println(promos, ok)

	// userID = 203 // key that doesn't exist

	// // Safe access with fallback to empty slice
	// promos, ok = stackPromotions[userID]
	// if !ok {
	// 	fmt.Printf("No promotions found for user %d, using empty list.\n", userID)
	// 	promos = []StackPromotion{}
	// }

	// Loop over the promotions safely
	fmt.Println("Promotions for user", userID)
	for _, promo := range promos {
		fmt.Printf("Promo ID: %d, Title: %s\n", promo.ID, promo.Title)
	}

	// Optional one-liner for fallback
	fmt.Println(append([]StackPromotion{}, stackPromotions[userID]...))
}
