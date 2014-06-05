package hipchat

import (
	"fmt"

	hipchat "."
)

func ExampleNewClient() {
	authToken := "le api token here"
	client, err := hipchat.NewClient(authToken)
	_ = client
	fmt.Println(err)
	// output:
	// <nil>
}
