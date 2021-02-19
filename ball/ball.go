package ball

import "math/rand"

// Ball is a Magic 8 Ball object
type Ball struct {
	Responses []string
}

// GetResponse returns a random response from a Ball
func (b Ball) GetResponse() string {
	return b.Responses[rand.Intn(len(b.Responses))]
}
