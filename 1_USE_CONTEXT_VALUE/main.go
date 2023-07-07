package main

import (
	"context"
	"fmt"
)

const ConnectedUserKey = "CONNECTED_USER_KEY"

// handleAnAccountCredit a trivial example of function that use value carried by context
func handleAnAccountCredit(ctx context.Context, amount int) {
	fmt.Printf("A credit of %v dollars for user %v", amount, ctx.Value(ConnectedUserKey))
}

func main() {
	ctx := context.WithValue(context.Background(), ConnectedUserKey, "Bobby")
	handleAnAccountCredit(ctx, 666)
}
