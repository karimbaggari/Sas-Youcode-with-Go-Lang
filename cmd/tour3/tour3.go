package tour3

import (
	"fmt"
	"bufio"
	"os"
)

type Vote struct {
    Elector   string
    President string
}

func getWinner(voteCounts map[string]int) string {
    var winner string
    maxVotes := 0
    
    for candidate, votes := range voteCounts {
        if votes > maxVotes {
            maxVotes = votes
            winner = candidate
        }
    }
    
    return winner
}

func Tour3(presidents, electeurs []string) ([]string, error) {
	fmt.Println("Welcome to tour3. Let the last Vote begin")
	for index, president := range presidents {
		fmt.Println(index, "-", president)
	}

	fmt.Println("Here are the electors allowed to vote:")
	for index, elector := range electeurs {
		fmt.Println(index, "-", elector)
	}

	var votes []Vote

	for _,elector := range electeurs {
		fmt.Printf("Elector %s, please pick a candidate\n", elector)
		electorScanner := bufio.NewScanner(os.Stdin)
		for{
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
            if!presidentExists {
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

	winner := getWinner(voteCounts)
	fmt.Println("The Winner is ", winner)

	return presidents, nil
}
