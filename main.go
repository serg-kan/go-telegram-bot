package main

import (
	"flag"
	"log"
)

func main() {
	token := mustToken()

	// tgClient = telegram.New(token)

	// fetcher = fetcher.New()

	// processor = processor.New()

	// consumer.Start(fetcher, processor)
}

// "must" -- means that the function doesn't return an error, but somtehing unexpected can occur
// so a developer should pay attention to this function
func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
