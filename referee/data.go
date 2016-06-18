package referee

type Match struct {
	HomeTeam  string `gorethink:"HomeTeam"`
	AwayTeam  string `gorethink:"AwayTeam"`
	HomeScore int    `gorethink:"HomeScore"`
	AwayScore int    `gorethink:"AwayScore"`
}

func newMatch(home string, away string) *Match {
	return &Match{home, away, 0, 0}
}
