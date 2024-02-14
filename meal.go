package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Meal struct {
	Day  string
	Date string
	Type string
	Menu []string
}

var meals []Meal

func main() {

	jsonData, err := os.ReadFile("json_data.json")
	if err != nil {
		log.Fatal("Failed to read JSON file:", err)
	}
	var data [][]string
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatal("Failed to unmarshal JSON:", err)
	}

	for _, dayData := range data {
		var meal Meal
		for i, item := range dayData {
			switch i {
			case 0:
				meal.Day = item
			case 1:
				meal.Date = item
			case 2:
				meal.Type = item
			default:
				if item == "BREAKFAST" || item == "LUNCH" || item == "DINNER" {
					if len(meal.Menu) > 0 {
						meals = append(meals, meal)
						meal = Meal{Day: meal.Day, Date: meal.Date, Type: item}
					}
				} else if item != "" && item != "Feb" && item != "MONDAY" && item != "TUESDAY" && item != "WEDNESDAY" && item != "THURSDAY" && item != "FRIDAY" && item != "SATURDAY" && item != "SUNDAY" {

					meal.Menu = append(meal.Menu, item)
				}
			}
		}

		if len(meal.Menu) > 0 {
			meals = append(meals, meal)
		}
	}
	fmt.Println("Select an option: ")
	fmt.Println("1. Get items")
	fmt.Println("2. Get number of items")
	fmt.Println("3. Check item")
	fmt.Println("4. Exit")
	var option int
	fmt.Scan(&option)
	switch option {
	case 1:
		var day, Type string
		fmt.Println("Enter day: ")
		fmt.Scan(&day)
		fmt.Println("Enter type: ")
		fmt.Scan(&Type)
		fmt.Println(getItem(strings.ToUpper(day), strings.ToUpper(Type)))
	case 2:
		var day, Type string
		fmt.Println("Enter day: ")
		fmt.Scan(&day)
		fmt.Println("Enter type: ")
		fmt.Scan(&Type)
		fmt.Println(getNom(strings.ToUpper(day), strings.ToUpper(Type)))
	case 3:
		reader := bufio.NewReader(os.Stdin)
		var day, Type, item string
		fmt.Println("Enter day: ")
		fmt.Scan(&day)
		fmt.Println("Enter type: ")
		fmt.Scan(&Type)
		fmt.Println("Enter item(Please note if there is space in item name you need to try entering  space once and twice because in Excel sheet there are two spaces in some items ): ")
		item, _ = reader.ReadString('\n')
		item = strings.TrimSpace(item)
		fmt.Println(checkItem(strings.ToUpper(day), strings.ToUpper(Type), strings.ToUpper(item)))
	case 4:
		return
	default:
		fmt.Println("Invalid option")
	}

}

func getItem(day string, mealType string) string {
	for _, meal := range meals {
		if meal.Day == day && meal.Type == mealType {
			return strings.Join(meal.Menu, ", ")
		}
	}
	return "No menu found"
}

func getNom(day string, mealType string) int {

	for _, meal := range meals {
		if meal.Day == day && meal.Type == mealType {
			return len(meal.Menu)
		}
	}
	return 0
}
func checkItem(day string, mealType string, item string) bool {

	for _, meal := range meals {
		if meal.Day == day && meal.Type == mealType {
			for _, menu := range meal.Menu {
				if menu == item {
					return true
				}
			}
		}
	}
	return false

}
