package main

import (
	"context"
	"fmt"
)

const ConnectedUserKey = "CONNECTED_USER_KEY"

func main() {
	ctx := context.WithValue(context.Background(), ConnectedUserKey, "Bobby")
	handleAnAccountCredit(ctx, 666)
}

// handleAnAccountCredit a trivial example of function that use value stored in context
func handleAnAccountCredit(ctx context.Context, amount int) {
	fmt.Printf("A credit of %v dollars for user %v", amount, ctx.Value(ConnectedUserKey))
}
