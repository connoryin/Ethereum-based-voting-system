package contract

import (
	"fmt"
	"log"
	"math/big"
	"sync"

	data_access "github.com/6675-voting-system/voting-system/backend/data-access"
	"github.com/6675-voting-system/voting-system/backend/model"
)

var mu sync.Mutex

func Login(req model.LoginReq, adminEvents *AdminEvents) (*model.LoginResp, error) {
	resp := &model.LoginResp{
		UserStatus: 0,
		EventId:    -1,
	}

	owner, eventID, _, eventAddr, err := adminEvents.GetEventInfoByInvCode(req.InvCode)
	log.Printf("event %v address %v", eventID, eventAddr)
	if err != nil {
		return resp, err
	}

	eventDetails, err := data_access.GetEvent(eventID)
	if err != nil {
		return resp, err
	}
	var ballot Ballot
	var ok bool
	if ballot, ok = adminEvents.events[eventID]; !ok {
		ballot = Ballot{
			address: eventAddr,
			owner:   owner,
		}
		ballot.LoadPoll()
		ballot.GetPollInfo(&eventDetails)
		ballot.event = data_access.EventDO2PO(eventDetails)
	}

	voter, _, err := ballot.instancePoll.GetVoterInfo(nil, str2byte32(req.InvCode))

	if err != nil {
		return resp, err
	}

	resp.EventId = eventID

	isOver := eventDetails.IsEnd

	if voter.Voted {
		resp.UserStatus = model.USERSTATUS_VOTED
	} else {
		if isOver {
			resp.UserStatus = model.USERSTATUS_UNVOTE
		} else {
			resp.UserStatus = model.USERSTATUS_VOTING
		}
	}

	return resp, nil
}

func GetVoteDetails(req model.GetVoteDetailsReq, adminEvents *AdminEvents) (*model.GetVoteDetailsResp, error) {
	resp := &model.GetVoteDetailsResp{
		Event:    model.Event{},
		IsVoted:  false,
		SelfVote: nil,
	}

	println(req.InvCode)

	owner, eventID, _, eventAddr, err := adminEvents.GetEventInfoByInvCode(req.InvCode)
	log.Printf("event %v address %v", eventID, eventAddr)
	if err != nil {
		return resp, err
	}

	eventDetails, err := data_access.GetEvent(eventID)
	if err != nil {
		return resp, err
	}

	var ballot Ballot
	var ok bool
	if ballot, ok = adminEvents.events[eventID]; !ok {
		ballot = Ballot{
			address: eventAddr,
			owner:   owner,
		}
		ballot.LoadPoll()
		ballot.GetPollInfo(&eventDetails)
		ballot.event = data_access.EventDO2PO(eventDetails)
	}

	voter, candidates, err := ballot.instancePoll.GetVoterInfo(nil, str2byte32(req.InvCode))

	for _, _candidate := range candidates {
		candidate := model.Candidate{
			Name:       string(_candidate.Name[:]),
			Id:         int(_candidate.Id.Int64()),
			VoteNumGet: int(_candidate.VoteNumGet.Int64()),
			IsWinner:   _candidate.IsWinner,
		}
		resp.SelfVote = append(resp.SelfVote, candidate)
	}
	resp.IsVoted = voter.Voted

	resp.Event = data_access.EventPO2DO(ballot.event)

	fmt.Printf("event: %v\n", resp.Event)
	fmt.Printf("SelfVote: %v\n", resp.SelfVote)

	return resp, nil
}

func (adminEvents *AdminEvents) SubmitVote(req model.VoteReq) (resp *model.VoteResp, err error) {
	err = data_access.SubmitVotes(req.InvCode, req.VotedCandidateNames)
	if err != nil {
		return
	}

	err = data_access.UpdateCandidatesOnSubmit(req.EventId, req.VotedCandidateNames)
	if err != nil {
		return
	}

	resp = &model.VoteResp{IsSuccess: false}
	owner, eventID, _, eventAddr, err := adminEvents.GetEventInfoByInvCode(req.InvCode)
	log.Printf("event %v address %v", eventID, eventAddr)
	if err != nil {
		return resp, err
	}

	eventDetails, err := data_access.GetEvent(eventID)
	if err != nil {
		return resp, err
	}

	var ballot Ballot
	var ok bool
	if ballot, ok = adminEvents.events[eventID]; !ok {
		ballot = Ballot{
			address: eventAddr,
			owner:   owner,
		}
		ballot.LoadPoll()
		ballot.GetPollInfo(&eventDetails)
		ballot.event = data_access.EventDO2PO(eventDetails)
	}

	Event := data_access.EventPO2DO(ballot.event)

	fmt.Printf("event: %v\n", Event)
	// fmt.Printf("InvCode: %v\n", str2byte32(req.InvCode))

	voter, candidates, err := ballot.instancePoll.GetVoterInfo(nil, str2byte32(req.InvCode))
	fmt.Printf("candidates: %v\n", candidates)
	fmt.Printf("voter: %v\n", voter)

	if err != nil {
		log.Fatal(err)
	}

	voterId := voter.Id
	// candidates[0].Id
	// candidates[0].Name
	candidateMap := make(map[string]*big.Int)
	for _, candidate := range Event.Candidates {
		candidateMap[fmt.Sprint(candidate.Name)] = big.NewInt(int64(candidate.Id))
	}
	fmt.Printf("candidateMap: %v\n", candidateMap)

	votedCandidateIds := make([]*big.Int, 0)
	for _, candidateName := range req.VotedCandidateNames {
		votedCandidateIds = append(votedCandidateIds, candidateMap[candidateName])
	}
	fmt.Printf("votedCandidateIds: %v\n", votedCandidateIds)
	mu.Lock()
	auth, _ := SetAuthOptions()
	_, err = ballot.instancePoll.Vote(auth, votedCandidateIds, voterId)
	mu.Unlock()

	if err != nil {
		log.Fatal(err)
	}
	resp.IsSuccess = true
	return
}
