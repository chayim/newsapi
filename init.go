package newsapi

import "os"

var token string

// Assuming the NEWSAPI environment variable is set, use that as the newsapi token.
func init() {
	a := os.Getenv("NEWSAPI")
	if a != "" {
		token = a
	}
}
