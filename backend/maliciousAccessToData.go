// package main

// import (
// 	"fmt"

// 	"github.com/6675-voting-system/voting-system/backend/contract"
// )

// func str2byte32(s string) (a [32]byte) {
// 	copy(a[:], s)
// 	return a
// }

// func main() {
// 	var ballot contract.Ballot

// 	ballot.SetUpAddress("")
// 	ballot.LoadPoll()

// 	voter := ballot.GetVoterInfo("")
// 	fmt.Printf("%v\n", voter)
// }
