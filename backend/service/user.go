package service

import (
	data_access "github.com/6675-voting-system/voting-system/backend/data-access"
	"github.com/6675-voting-system/voting-system/backend/model"
)

func Login(req model.LoginReq) (*model.LoginResp, error) {
	resp := &model.LoginResp{
		UserStatus: 0,
		EventId:    -1,
	}

	eventID, err := data_access.GetEventIDByInvitationID(req.InvCode)
	if err != nil {
		return resp, err
	}

	voteDetails, err := data_access.GetVoteDetailsByInvitationID(req.InvCode)
	if err != nil {
		return resp, err
	}

	resp.EventId = eventID

	eventDetails, err := data_access.GetEvent(eventID)
	if err != nil {
		return resp, err
	}

	isOver := eventDetails.IsEnd

	if len(voteDetails) > 0 {
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

func GetVoteDetails(req model.GetVoteDetailsReq) (*model.GetVoteDetailsResp, error) {
	resp := &model.GetVoteDetailsResp{
		Event:    model.Event{},
		IsVoted:  false,
		SelfVote: nil,
	}

	voteDetails, err := data_access.GetVoteDetailsByInvitationID(req.InvCode)
	if err != nil {
		return resp, err
	}

	if len(voteDetails) > 0 {
		resp.IsVoted = true
		resp.SelfVote = voteDetails
	}

	eventID, err := data_access.GetEventIDByInvitationID(req.InvCode)
	if err != nil {
		return resp, err
	}

	eventDetails, err := data_access.GetEvent(eventID)
	if err != nil {
		return resp, err
	}

	resp.Event = eventDetails

	return resp, nil
}

func SubmitVote(req model.VoteReq) (resp *model.VoteResp, err error) {
	resp = &model.VoteResp{IsSuccess: false}

	err = data_access.SubmitVotes(req.InvCode, req.VotedCandidateNames)
	if err != nil {
		return
	}

	err = data_access.UpdateCandidatesOnSubmit(req.EventId, req.VotedCandidateNames)
	if err != nil {
		return
	}

	resp.IsSuccess = true
	return
}
