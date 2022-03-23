package utils

import (
	"math/rand"
	"time"
)

func PickRandomIndex(input []string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	return input[rand.Intn(len(input))]
}