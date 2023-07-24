package helper

import (
	"math/rand"
	"time"
)

var (
	letterRunes  = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	letterLength = 5
)

func RandStringRunes() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, letterLength)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func ThreeDaysFromToday() time.Time {
	now := time.Now()
	return now.AddDate(0, 0, 3)
}
