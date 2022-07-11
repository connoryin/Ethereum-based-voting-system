// package main

// import (
// 	"fmt"

// 	"github.com/6675-voting-system/voting-system/backend/contract"
// 	"github.com/6675-voting-system/voting-system/backend/poll"
// )

// func main() {
// 	contract.SetupContract()
// 	auth, Client := contract.SetAuthOptions()
// 	addressEvents, _, instanceEvents, err := poll.DeployEvents(auth, Client)
// 	fmt.Printf("address: %v, instance: %v, err: %v\n", addressEvents, instanceEvents, err)
// }
