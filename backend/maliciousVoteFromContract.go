// package main

// import (
// 	"math/big"

// 	"github.com/6675-voting-system/voting-system/backend/contract"
// )

// func str2byte32(s string) (a [32]byte) {
// 	copy(a[:], s)
// 	return a
// }

// func main() {
// 	auth, _ := contract.SetAuthOptions()
// 	votedCandidateIds := make([]*big.Int, 0)
// 	var ballot contract.Ballot

// 	ballot.SetUpAddress("")
// 	ballot.LoadPoll()

// 	ballot.Vote(votedCandidateIds, str2byte32(""), auth)

// }
