package tutorial_1

import (
	"bufio"
    "fmt"
    "os"
  
)

type Vote struct {
    Elector   string
    President string
}

func Tour1(presidents,electeurs []string) {
	fmt.Println("Welcome to tour 1. Here are the presidents:")
    for _, president := range presidents {
        fmt.Println("-", president)
    }
    fmt.Println("Here are the electors allowed to vote:")
    for _, elector := range electeurs {
        fmt.Println("-", elector)
    }

	var votes []Vote

	for _, elector := range electeurs {
        fmt.Printf("Elector %s, please pick a president:\n", elector)
		electorScanner := bufio.NewScanner(os.Stdin)
        for {
            fmt.Print("Enter the name of the president (or 'q' to quit): ")
            electorScanner.Scan()
            choice := electorScanner.Text()

            if choice == "q" {
                fmt.Println("Quitting vote selection for elector", elector)
                break
            }

            // Check if the selected president exists
            presidentExists := false
            for _, president := range presidents {
                if president == choice {
                    presidentExists = true
                    break
                }
            }
            if !presidentExists {
                fmt.Println("Invalid president name. Please pick from the list.",choice)
                continue
            }

            // Save the vote
            vote := Vote{Elector: elector, President: choice}
            votes = append(votes, vote)

            fmt.Printf("Vote saved: Elector %s voted for %s\n", elector, votes)
			break
        }
	}

	voteCounts := make(map[string]int)
    for _, vote := range votes {
        voteCounts[vote.President]++
    }

	 // Calculate the total number of votes
	 totalVotes := len(votes)

	 var remainingPresidents []string
	 for president, count := range voteCounts {
		percentage := float64(count) / float64(totalVotes)
		if percentage >= 0.15 {
			remainingPresidents = append(remainingPresidents, president)
		}
	}
	
	fmt.Println("Remaining presidents after exclusion:", remainingPresidents)
}