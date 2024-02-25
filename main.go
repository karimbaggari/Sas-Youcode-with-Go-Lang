package main
import (
	"bufio"
    "fmt"
    "os"
	"learn-go/cmd/tour1"
	"learn-go/cmd/tour2"
)

func main() {
	var presidents []string
	var electeurs []string
	
	fmt.Println("Enter the presidents name please : (enter q when you complete):")
	
	scanner := bufio.NewScanner(os.Stdin)
	for len(presidents) < 5 {
		for scanner.Scan() {
        	str := scanner.Text()
			
			if str == "q" {
				if len(presidents) >= 5 {
					fmt.Println("Quitting input loop...")
					break // Exit the loop if enough elements are entered
				} else {
					fmt.Println("You cannot quit. Please enter at least 5 elements.")
					continue // Restart the loop to prompt for input again
				}
			}

			if scanner.Err() != nil {
				fmt.Println("Error reading input:", scanner.Err())
			}

        	presidents = append(presidents, str) // Adding the new string to the presidents slice
    }
}


fmt.Println("Enter the electeurs cin please : (enter q when you complete):")

	
for len(electeurs) < 10 {
	for scanner.Scan() {
		str := scanner.Text()
		
		if str == "q" {
			if len(electeurs) >= 5 {
				fmt.Println("Quitting input loop...")
				break // Exit the loop if enough elements are entered
			} else {
				fmt.Println("You cannot quit. Please enter at least 10 elements.")
				continue // Restart the loop to prompt for input again
			}
		}

		if scanner.Err() != nil {
			fmt.Println("Error reading input:", scanner.Err())
		}

		electeurs = append(electeurs, str) // Adding the new string to the presidents slice
}
}

    fmt.Println("Entered presidents:", presidents)
	fmt.Println("Entered electeurs:", electeurs)
	tour1Results := tour1.Tour1(presidents, electeurs)
	fmt.Println("--------------------------------------------------------" , tour1Results)
	tour2.Tour2(tour1Results,electeurs)
}