package creature

import "math/rand"

import "time"

var creatures = []string{"shark", "jellyfish", "squid", "octopus", "dolphin"}

func init() {
	rand.Seed(time.Now().UnixNano())
}

//Random send rendom element in array
func Random() string {
	i := rand.Intn(len(creatures))
	return creatures[i]
}
