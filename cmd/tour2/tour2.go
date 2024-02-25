package tour2

import(
	"fmt"
	"os"
	"bufio"
	"math"
)
type Vote struct {
    Elector   string
    President string
}
func Tour2(presidents,electeurs []string) {
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

	 var minVotes int = math.MaxInt64
	 var leastVotedPresident string

	 for president, votes := range voteCounts {
        if votes < minVotes {
            minVotes = votes
            leastVotedPresident = president
        }
	}
		if leastVotedPresident != "" {
			fmt.Printf("The least voted president is %s with %d votes.\n", leastVotedPresident, minVotes)
		} else {
			fmt.Println("No votes recorded.")
		}

		for i, president := range presidents {
			if president == leastVotedPresident {
				presidents = append(presidents[:i], presidents[i+1:]...)
				break
			}
		}
	
		fmt.Println("Presidents after excluding the least voted president:", presidents)	
}