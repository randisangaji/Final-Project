package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Review struct {
	Rating  int
	Comment string
}

type CoworkingSpace struct {
	Name       string
	Location   string
	Facilities []string
	Price      float64
	Reviews    []Review
}

var spaces []CoworkingSpace
var reader = bufio.NewReader(os.Stdin)

func main() {
	for {
		fmt.Println("--- Co-Working Space App ---")
		fmt.Println("1. Add coworking space")
		fmt.Println("2. View all spaces")
		fmt.Println("3. Add review")
		fmt.Println("4. Search by name or location")
		fmt.Println("5. Sort by price or rating")
		fmt.Println("6. Filter by facility")
		fmt.Println("7. Binary search by name")
		fmt.Println("8. Exit")
		fmt.Print("Choose: ")

		var choice int
		fmt.Scan(&choice)
		reader.ReadString('\n')
		switch choice {
		case 1:
			addSpace()
		case 2:
			viewSpaces()
		case 3:
			addReview()
		case 4:
			searchSpace()
		case 5:
			sortSpaces()
		case 6:
			filterFacility()
		case 7:
			binarySearch()
		case 8:
			fmt.Println("Goodbye.")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func readLine() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func addSpace() {
	fmt.Print("Name: ")
	name := readLine()
	fmt.Print("Location: ")
	location := readLine()
	fmt.Print("Rental Price: ")
	var price float64
	priceInput := readLine()
	fmt.Sscan(priceInput, &price)

	fmt.Println("Enter facilities (type 'done' to finish):")
	var facilities []string
	for {
		fmt.Print("Facility: ")
		facility := readLine()
		if strings.ToLower(facility) == "done" {
			break
		}
		facilities = append(facilities, facility)
	}

	spaces = append(spaces, CoworkingSpace{
		Name:       name,
		Location:   location,
		Price:      price,
		Facilities: facilities,
	})

	fmt.Println("Coworking space added.")
}

func viewSpaces() {
	if len(spaces) == 0 {
		fmt.Println("No coworking spaces available.")
		return
	}
	for i := 0; i < len(spaces); i++ {
		s := spaces[i]
		fmt.Println("---")
		fmt.Printf("%s (%s) - %.2f\n", s.Name, s.Location, s.Price)
		fmt.Println("Facilities:", strings.Join(s.Facilities, ", "))
		fmt.Printf("Reviews: %d\n", len(s.Reviews))
	}
}

func addReview() {
	fmt.Print("Enter coworking space name: ")
	name := readLine()
	for i := 0; i < len(spaces); i++ {
		if strings.EqualFold(spaces[i].Name, name) {
			fmt.Print("Rating (1-5): ")
			var rating int
			fmt.Scan(&rating)
			reader.ReadString('\n')
			fmt.Print("Comment: ")
			comment := readLine()
			spaces[i].Reviews = append(spaces[i].Reviews, Review{rating, comment})
			fmt.Println("Review added.")
			return
		}
	}
	fmt.Println("Space not found.")
}

func searchSpace() {
	fmt.Print("Search keyword: ")
	keyword := strings.ToLower(readLine())
	found := false
	for i := 0; i < len(spaces); i++ {
		s := spaces[i]
		if strings.Contains(strings.ToLower(s.Name), keyword) || strings.Contains(strings.ToLower(s.Location), keyword) {
			fmt.Printf("%s - %s | Price: %.2f\n", s.Name, s.Location, s.Price)
			found = true
		}
	}
	if !found {
		fmt.Println("No results.")
	}
}

func sortSpaces() {
	fmt.Println("1. Sort by price")
	fmt.Println("2. Sort by average rating")
	fmt.Print("Choose: ")
	var choice int
	fmt.Scan(&choice)

	if choice == 1 {
		for i := 0; i < len(spaces)-1; i++ {
			min := i
			for j := i + 1; j < len(spaces); j++ {
				if spaces[j].Price < spaces[min].Price {
					min = j
				}
			}
			spaces[i], spaces[min] = spaces[min], spaces[i]
		}
		fmt.Println("Sorted by price.")
		viewSpaces()
	} else if choice == 2 {
		for i := 1; i < len(spaces); i++ {
			key := spaces[i]
			j := i - 1
			for j >= 0 && avgRating(spaces[j]) < avgRating(key) {
				spaces[j+1] = spaces[j]
				j--
			}
			spaces[j+1] = key
		}
		fmt.Println("Sorted by rating.")
		viewSpaces()
	} else {
		fmt.Println("Invalid.")
	}
}

func avgRating(s CoworkingSpace) float64 {
	if len(s.Reviews) == 0 {
		return 0
	}
	total := 0
	for i := 0; i < len(s.Reviews); i++ {
		total += s.Reviews[i].Rating
	}
	return float64(total) / float64(len(s.Reviews))
}

func filterFacility() {
	fmt.Print("Filter by facility: ")
	query := strings.ToLower(readLine())
	found := false
	for i := 0; i < len(spaces); i++ {
		s := spaces[i]
		for j := 0; j < len(s.Facilities); j++ {
			if strings.EqualFold(s.Facilities[j], query) {
				fmt.Printf("%s - %s\n", s.Name, s.Location)
				found = true
				break
			}
		}
	}
	if !found {
		fmt.Println("No spaces matched.")
	}
}

func binarySearch() {
	if len(spaces) == 0 {
		fmt.Println("No spaces.")
		return
	}

	for i := 0; i < len(spaces)-1; i++ {
		min := i
		for j := i + 1; j < len(spaces); j++ {
			if strings.ToLower(spaces[j].Name) < strings.ToLower(spaces[min].Name) {
				min = j
			}
		}
		spaces[i], spaces[min] = spaces[min], spaces[i]
	}

	fmt.Print("Enter name to search: ")
	target := strings.ToLower(readLine())
	low := 0
	high := len(spaces) - 1
	for low <= high {
		mid := (low + high) / 2
		name := strings.ToLower(spaces[mid].Name)
		if name == target {
			fmt.Printf("Found: %s (%s) - %.2f\n", spaces[mid].Name, spaces[mid].Location, spaces[mid].Price)
			return
		} else if name < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("Not found.")
}
