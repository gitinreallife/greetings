package greetings

import (
	"errors"
	"fmt"

	"math/rand"
	"time"

	"rsc.io/quote"
)

func GetQuote(name string) (string, error) {
	if name == "" {
		return "", errors.New(("empty name"))
	}
	strQuote := fmt.Sprintf(randomFormat(), name)
	return strQuote, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Quote for the day for  %v: " + quote.Go(),
		quote.Go() + " Cheers, %v",
	}
	return formats[rand.Intn(len(formats))]
}
