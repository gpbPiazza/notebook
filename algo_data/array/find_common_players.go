package array

type Players struct {
	FirstName string
	LastName  string
	Team      string
}

func (p Players) FullName() string {
	return p.FirstName + " " + p.LastName
}

func FindWhoPlaysInBothSports(footBallPlayers, basquetBallPlayers []Players) []string {
	var result []string
	playersName := make(map[string]bool)
	for _, p := range footBallPlayers {
		playersName[p.FullName()] = true
	}

	for _, p := range basquetBallPlayers {
		if playersName[p.FullName()] {
			result = append(result, p.FullName())
		}
	}

	return result
}
