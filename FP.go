package main

import "fmt"

type Review struct {
	Rating  int
	Comment string
}

type CoworkingSpace struct {
	Name          string
	Location      string
	Facilities    [10]string
	FacilityCount int
	Price         float64
	Reviews       [100]Review
	ReviewCount   int
	Active        bool
}

const maxSpaces = 100

var spaces [maxSpaces]CoworkingSpace
var nSpaces int

func main() {
	for {
		fmt.Println("--- Co-Working Space App ---")
		fmt.Println("1. Add coworking space")
		fmt.Println("2. View all spaces")
		fmt.Println("3. Add review")
		fmt.Println("4. Search by name or location")
		fmt.Println("5. Sort by price or rating")
		fmt.Println("6. Filter by facility")
		fmt.Println("7. Edit coworking space")
		fmt.Println("8. Delete coworking space")
		fmt.Println("9. Binary search by name")
		fmt.Println("10. Exit")
		fmt.Print("Choose: ")

		var choice int
		fmt.Scan(&choice)

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
			editSpace()
		case 8:
			deleteSpace()
		case 9:
			binarySearchByName()
		case 10:
			fmt.Println("Goodbye.")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func addSpace() {
	var name, location string
	var price float64

	fmt.Print("Name: ")
	fmt.Scan(&name)
	fmt.Print("Location: ")
	fmt.Scan(&location)
	fmt.Print("Price: ")
	fmt.Scan(&price)

	space := CoworkingSpace{Name: name, Location: location, Price: price, Active: true}

	fmt.Println("Enter facilities one by one (type 'done' to finish):")
	for space.FacilityCount < 10 {
		fmt.Print("Facility: ")
		var fac string
		fmt.Scan(&fac)
		if fac == "done" {
			break
		}
		space.Facilities[space.FacilityCount] = fac
		space.FacilityCount++
	}

	spaces[nSpaces] = space
	nSpaces++
	fmt.Println("Coworking space added.")
}

func viewSpaces() {
	for i := 0; i < nSpaces; i++ {
		if !spaces[i].Active {
			continue
		}
		fmt.Printf("%d. %s (%s) - %.2f\n", i+1, spaces[i].Name, spaces[i].Location, spaces[i].Price)
		fmt.Print("Facilities: ")
		for j := 0; j < spaces[i].FacilityCount; j++ {
			fmt.Print(spaces[i].Facilities[j])
			if j < spaces[i].FacilityCount-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println()
		if spaces[i].ReviewCount > 0 {
			total := 0
			for j := 0; j < spaces[i].ReviewCount; j++ {
				total += spaces[i].Reviews[j].Rating
			}
			avg := float64(total) / float64(spaces[i].ReviewCount)
			fmt.Printf("Average Rating: %.2f\n", avg)
		} else {
			fmt.Println("Average Rating: No reviews yet")
		}
	}
}

func addReview() {
	var name string
	fmt.Print("Enter coworking space name: ")
	fmt.Scan(&name)

	for i := 0; i < nSpaces; i++ {
		if spaces[i].Name == name && spaces[i].Active {
			fmt.Print("Rating (1-5): ")
			var rate int
			fmt.Scan(&rate)
			fmt.Print("Comment (one word only): ")
			var comment string
			fmt.Scan(&comment)
			spaces[i].Reviews[spaces[i].ReviewCount] = Review{Rating: rate, Comment: comment}
			spaces[i].ReviewCount++
			fmt.Println("Review added.")
			return
		}
	}
	fmt.Println("Coworking space not found.")
}

func searchSpace() {
	var keyword string
	fmt.Print("Enter name or location: ")
	fmt.Scan(&keyword)
	for i := 0; i < nSpaces; i++ {
		if spaces[i].Active && (spaces[i].Name == keyword || spaces[i].Location == keyword) {
			fmt.Printf("%s (%s) - %.2f\n", spaces[i].Name, spaces[i].Location, spaces[i].Price)
			fmt.Print("Facilities: ")
			for j := 0; j < spaces[i].FacilityCount; j++ {
				fmt.Print(spaces[i].Facilities[j])
				if j < spaces[i].FacilityCount-1 {
					fmt.Print(", ")
				}
			}
			fmt.Println()
			return
		}
	}
	fmt.Println("No matching space found.")
}

func sortSpaces() {
	fmt.Println("1. Sort by price (lowest to highest)")
	fmt.Println("2. Sort by average rating (highest to lowest)")
	fmt.Print("Choose: ")
	var opt int
	fmt.Scan(&opt)

	if opt == 1 {
		for i := 0; i < nSpaces-1; i++ {
			for j := i + 1; j < nSpaces; j++ {
				if spaces[j].Price < spaces[i].Price {
					spaces[i], spaces[j] = spaces[j], spaces[i]
				}
			}
		}
	} else if opt == 2 {
		for i := 0; i < nSpaces-1; i++ {
			for j := i + 1; j < nSpaces; j++ {
				avgI := 0
				avgJ := 0
				if spaces[i].ReviewCount > 0 {
					for r := 0; r < spaces[i].ReviewCount; r++ {
						avgI += spaces[i].Reviews[r].Rating
					}
					avgI /= spaces[i].ReviewCount
				}
				if spaces[j].ReviewCount > 0 {
					for r := 0; r < spaces[j].ReviewCount; r++ {
						avgJ += spaces[j].Reviews[r].Rating
					}
					avgJ /= spaces[j].ReviewCount
				}
				if avgJ > avgI {
					spaces[i], spaces[j] = spaces[j], spaces[i]
				}
			}
		}
	}
	fmt.Println("Sorted.")
}

func filterFacility() {
	var fac string
	fmt.Print("Enter facility to filter: ")
	fmt.Scan(&fac)
	for i := 0; i < nSpaces; i++ {
		if !spaces[i].Active {
			continue
		}
		for j := 0; j < spaces[i].FacilityCount; j++ {
			if spaces[i].Facilities[j] == fac {
				fmt.Printf("%s (%s) - %.2f\n", spaces[i].Name, spaces[i].Location, spaces[i].Price)
				break
			}
		}
	}
}

func editSpace() {
	var name string
	fmt.Print("Enter name of space to edit: ")
	fmt.Scan(&name)
	for i := 0; i < nSpaces; i++ {
		if spaces[i].Active && spaces[i].Name == name {
			fmt.Print("New location: ")
			fmt.Scan(&spaces[i].Location)
			fmt.Print("New price: ")
			fmt.Scan(&spaces[i].Price)

			fmt.Println("Enter new facilities (type 'done' to finish):")
			spaces[i].FacilityCount = 0
			for spaces[i].FacilityCount < 10 {
				fmt.Print("Facility: ")
				var f string
				fmt.Scan(&f)
				if f == "done" {
					break
				}
				spaces[i].Facilities[spaces[i].FacilityCount] = f
				spaces[i].FacilityCount++
			}

			fmt.Println("Updated.")
			return
		}
	}
	fmt.Println("Space not found.")
}

func deleteSpace() {
	var name string
	fmt.Print("Enter name to delete: ")
	fmt.Scan(&name)
	for i := 0; i < nSpaces; i++ {
		if spaces[i].Active && spaces[i].Name == name {
			spaces[i].Active = false
			fmt.Println("Deleted.")
			return
		}
	}
	fmt.Println("Space not found.")
}

func binarySearchByName() {
	fmt.Print("Enter name to search: ")
	var name string
	fmt.Scan(&name)

	sortByName()

	low, high := 0, nSpaces-1
	for low <= high {
		mid := (low + high) / 2
		if !spaces[mid].Active {
			mid++
			continue
		}
		if spaces[mid].Name == name {
			fmt.Printf("Found: %s (%s) - %.2f\n", spaces[mid].Name, spaces[mid].Location, spaces[mid].Price)
			fmt.Print("Facilities: ")
			for j := 0; j < spaces[mid].FacilityCount; j++ {
				fmt.Print(spaces[mid].Facilities[j])
				if j < spaces[mid].FacilityCount-1 {
					fmt.Print(", ")
				}
			}
			fmt.Println()
			return
		} else if spaces[mid].Name < name {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("Space not found.")
}

func sortByName() {
	for i := 0; i < nSpaces-1; i++ {
		min := i
		for j := i + 1; j < nSpaces; j++ {
			if spaces[j].Name < spaces[min].Name {
				min = j
			}
		}
		spaces[i], spaces[min] = spaces[min], spaces[i]
	}
}
