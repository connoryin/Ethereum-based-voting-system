package contract

import (
	"fmt"
	"log"
	"math/big"

	"github.com/6675-voting-system/voting-system/backend/poll"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func str2byte32(s string) (a [32]byte) {
	copy(a[:], s)
	return a
}

func (adminEvents *AdminEvents) CreateInvitations(eventId int, owner *big.Int, email2InvCodes map[string]string, addr common.Address) *types.Transaction {
	auth, _ := SetAuthOptions()
	invCodes := make([][32]byte, len(email2InvCodes))
	// pollInfos := make([]poll.EventspollInfo, 0)
	pollInfo := poll.EventspollInfo{
		Owner:   owner,
		EventId: big.NewInt(int64(eventId)),
		Exists:  true,
		IsEnd:   false,
		Addr:    addr,
	}
	for _, invCode := range email2InvCodes {
		invCodes = append(invCodes, str2byte32(invCode))
		// pollInfo := poll.EventspollInfo{
		// 	Owner:   owner,
		// 	EventId: big.NewInt(int64(eventId)),
		// 	Exists:  true,
		// 	IsEnd:   false,
		// }
		// pollInfos = append(pollInfos, pollInfo)
	}
	tx, error := adminEvents.instanceEvents.UploadPollInfo(auth, invCodes, pollInfo, owner)
	if error != nil {
		log.Fatal(error)
	}
	return tx
}

func (adminEvents *AdminEvents) GetEventInfoByInvCode(invCode string) (*big.Int, int, bool, common.Address, error) {
	invCodeBytes := str2byte32(invCode)
	address, eventId, isEnd, addr, err := adminEvents.instanceEvents.GetPollInfoByInvCode(nil, invCodeBytes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("eventid: %v\n", eventId)
	return address, int(eventId.Int64()), isEnd, addr, err
}

func (adminEvents *AdminEvents) GetEventInfoByAdminID(adminID int) ([]poll.EventspollInfo, error) {
	info, err := adminEvents.instanceEvents.GetPollInfoByAdminID(nil, big.NewInt(int64(adminID)))
	return info, err
}
