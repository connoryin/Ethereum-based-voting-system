package contract

import (
	"fmt"
	"log"
	"math/big"
	"net/smtp"
	"time"

	data_access "github.com/6675-voting-system/voting-system/backend/data-access"
	"github.com/6675-voting-system/voting-system/backend/model"
	"github.com/6675-voting-system/voting-system/backend/poll"
	"github.com/ethereum/go-ethereum/common"
	"github.com/thanhpk/randstr"
)

func (adminEvents *AdminEvents) GetEventsByAdminId(adminId int) (events []model.Event, err error) {
	if adminEvents.instanceEvents == nil {
		adminEvents.loadEventInstances()
	}

	// infos, _ := adminEvents.GetEventInfoByAdminID(adminId)
	var pos []data_access.EventPO
	db := data_access.GetDB()
	err = db.Table("events").Where("admin_id = ?", adminId).Find(&pos).Error
	if err != nil {
		return nil, err
	}
	// fmt.Printf("events: %v\n", pos)

	for _, po := range pos {
		b := Ballot{
			address:      common.HexToAddress(po.Address),
			instancePoll: nil,
			event:        po,
		}

		fmt.Printf("loading address: %v\n", po.Address)

		b.LoadPoll()
		fmt.Printf("poll instance: %v\n", b.instancePoll)

		adminEvents.events[po.Id] = b

		do := data_access.EventPO2DO(po)
		b.GetPollInfo(&do)
		fmt.Printf("got poll info\n")

		events = append(events, do)
	}
	return events, nil
}

func (adminEvents *AdminEvents) getAdminEventsByAdminId(adminId int) (inVotingEvents, endedVotingEvents []model.Event, err error) {
	events, err := adminEvents.GetEventsByAdminId(adminId)
	if err != nil {
		return nil, nil, err
	}

	for _, event := range events {
		if !event.IsEnd {
			inVotingEvents = append(inVotingEvents, event)
		} else {
			endedVotingEvents = append(endedVotingEvents, event)
		}
	}
	return inVotingEvents, endedVotingEvents, nil
}

func (adminEvents *AdminEvents) AdminDetail(adminId int) (resp model.AdminDetailResp, err error) {
	inVotingEvents, endedVotingEvents, err := adminEvents.getAdminEventsByAdminId(adminId)
	if err != nil {
		return model.AdminDetailResp{}, err
	}
	return model.AdminDetailResp{
		InVotingEvents:    inVotingEvents,
		EndedVotingEvents: endedVotingEvents,
	}, nil
}

func generateInvCode(emails []string) map[string]string {
	email2InvCode := make(map[string]string)
	invcodes := make([]string, 0)
	for _, email := range emails {
		invCode := randstr.String(8)
		email2InvCode[email] = invCode
		invcodes = append(invcodes, invCode)
	}

	fmt.Printf("invcodes: %v\n", invcodes)
	return email2InvCode
}

func email2Voter(toEmail string, InvCode string, eventName string) (err error) {
	start := time.Now()

	from := "voting_system@yahoo.com"
	password := "idqlogcoxsjfcact"

	msg := []byte("To: " + toEmail + "\r\n" +
		"Subject: Vote for " + eventName + "!\r\n" +
		"\r\n" +
		"Your invitation code is " + InvCode + ".\r\n")

	// smtp server configuration.
	smtpHost := "smtp.mail.yahoo.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{toEmail}, msg)
	if err != nil {
		println("In send email")
		fmt.Println(err)
		return err
	}
	fmt.Println("Email Sent Successfully!")

	duration := time.Since(start)
	fmt.Println("send email: ", duration)
	return nil
}

func (adminEvents *AdminEvents) CreateEvent(req model.CreateEventReq, adminID int) (*model.CreateEventResp, error) {
	// eventId, err := data_access.CreateEvent(req.Event)
	start := time.Now()

	_eventId, err := data_access.CreateEvent(req.Event)
	if err != nil {
		return nil, err
	}
	req.Event.Id = _eventId

	ballot := Ballot{
		event: data_access.EventDO2PO(req.Event),
	}
	auth, _ := SetAuthOptions()
	// var _eventId *big.Int = big.NewInt(int64(req.Event.Id))
	var _totalNumVoter *big.Int = big.NewInt(int64(req.Event.TotalVoteNum))
	var _numBallotPerVoter *big.Int = big.NewInt(int64(req.Event.MaxVoteNumPerPerson))
	var _adminID *big.Int = big.NewInt(int64(adminID))
	_isAnonymous := req.Event.IsAnonymous
	_voters := make([]poll.PollVoter, 0)
	_candidates := make([]poll.PollCandidate, 0)
	email2InvCodes := generateInvCode(req.Voters)
	err = data_access.CreateInvitations(_eventId, email2InvCodes)
	if err != nil {
		return nil, err
	}

	for email, invCode := range email2InvCodes {
		err := email2Voter(email, invCode, req.Event.Name)
		if err != nil {
			return nil, err
		}
	}

	for _, candidate := range req.Event.Candidates {
		pollCandidate := poll.PollCandidate{
			Id:         big.NewInt(int64(candidate.Id)),
			Name:       str2byte32(candidate.Name),
			VoteNumGet: big.NewInt(0),
			IsWinner:   false,
			Exists:     true,
		}
		_candidates = append(_candidates, pollCandidate)
	}

	for _, voterCode := range email2InvCodes {
		pollVoter := poll.PollVoter{
			Id:       str2byte32(voterCode),
			Voted:    false,
			VotedFor: make([]*big.Int, 0),
			Exists:   true,
		}
		_voters = append(_voters, pollVoter)
	}
	address := ballot.DeployPoll(auth, _adminID, big.NewInt(int64(_eventId)), _totalNumVoter, _numBallotPerVoter, _voters, _candidates, _isAnonymous)
	db := data_access.GetDB()
	err = db.Table("events").Where("id = ?", _eventId).Update("address", address.String()).Error

	adminEvents.CreateInvitations(ballot.event.Id, _adminID, email2InvCodes, ballot.address)

	duration := time.Since(start)
	fmt.Println(duration)

	return &model.CreateEventResp{EventId: ballot.event.Id, IsSuccess: true}, nil
}

func (adminEvents *AdminEvents) EndEvent(req model.EndEventReq, adminID int) (*model.EndEventResp, error) {
	ballot := adminEvents.events[req.EventId]
	auth, _ := SetAuthOptions()
	tx, err := ballot.EndVote(auth, adminID)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Printf("tx %v", tx)

	err = data_access.EndEvent(req.EventId)
	if err != nil {
		return nil, err
	}

	event, err := data_access.GetEvent(req.EventId)
	if err != nil {
		return nil, err
	}

	max := 0
	for _, candidate := range event.Candidates {
		if candidate.VoteNumGet > max {
			max = candidate.VoteNumGet
		}
	}

	for index, candidate := range event.Candidates {
		if candidate.VoteNumGet == max {
			event.Candidates[index].IsWinner = true
		}
	}

	err = data_access.UpdateEventWinner(req.EventId, event.Candidates)
	if err != nil {
		return nil, err
	}
	// winnerId, winnerCandidate := ballot.GetWinner()
	return &model.EndEventResp{IsSuccess: true}, nil
}

func (adminEvents *AdminEvents) GetEvent(eventID int) (*model.GetEventResp, error) {
	if adminEvents.instanceEvents == nil {
		adminEvents.loadEventInstances()
	}
	var po data_access.EventPO

	db := data_access.GetDB()
	err := db.Table("events").Where("id = ?", eventID).First(&po).Error
	if err != nil {
		return nil, err
	}

	b := Ballot{
		address:      common.HexToAddress(po.Address),
		instancePoll: nil,
		event:        po,
	}

	b.LoadPoll()

	adminEvents.events[eventID] = b

	do := data_access.EventPO2DO(po)
	b.GetPollInfo(&do)

	return &model.GetEventResp{Event: do}, nil
}

func Download(req model.DownloadReq, adminEvents *AdminEvents) (*model.DownloadResp, error) {
	//get voter details
	InvitationCode2Details := make(map[string]string)
	ballot := adminEvents.events[req.EventId]
	// if !ballot.event.IsAnonymous {
	voterDetails, candidateNames, err := ballot.instancePoll.GetAllVoterDetails(nil)
	if err != nil {
		return nil, err
	}
	for i, voterDetail := range voterDetails {
		invCode := fmt.Sprintf("%v", voterDetail.Id)
		candidateNamesPerVoter := candidateNames[i]
		candidateNameStrArray := make([]string, 0)
		for _, candidateName := range candidateNamesPerVoter {
			candidateNameStr := fmt.Sprint(candidateName)
			candidateNameStrArray = append(candidateNameStrArray, candidateNameStr)
		}
		InvitationCode2Details[invCode] = fmt.Sprint(candidateNameStrArray)
	}
	// }
	return &model.DownloadResp{InvitationCode2Details: InvitationCode2Details}, nil
}
