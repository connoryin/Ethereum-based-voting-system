package service

import (
	"fmt"
	"net/smtp"
	"time"

	data_access "github.com/6675-voting-system/voting-system/backend/data-access"
	"github.com/6675-voting-system/voting-system/backend/model"
	"github.com/thanhpk/randstr"
)

func AdminLogin(req model.AdminLoginReq) (resp model.AdminLoginResp, err error) {
	adminId, err := adminLogin(req.Name, req.Password)
	if err != nil {
		return model.AdminLoginResp{}, err
	}

	return model.AdminLoginResp{StatusCode: model.STATUSCODE_SUCCESS, AdminId: adminId}, nil
}

func AdminDetail(adminId int) (resp model.AdminDetailResp, err error) {
	inVotingEvents, endedVotingEvents, err := getAdminEventsByAdminId(adminId)
	if err != nil {
		return model.AdminDetailResp{}, err
	}
	return model.AdminDetailResp{
		InVotingEvents:    inVotingEvents,
		EndedVotingEvents: endedVotingEvents,
	}, nil
}

func adminLogin(name string, password string) (adminId int, err error) {
	var admin *model.Admin
	admin, err = data_access.GetAdmin(name, password)
	if err != nil {
		return -1, err
	}
	return admin.Id, nil
}

func getAdminEventsByAdminId(adminId int) (inVotingEvents, endedVotingEvents []model.Event, err error) {
	events, err := data_access.GetEventsByAdminId(adminId)
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

func AdminRegister(req model.AdminRegisterReq) (resp model.AdminRegisterResp, err error) {
	err = adminRegister(req.Name, req.Password)
	if err != nil {
		return model.AdminRegisterResp{}, err
	}
	return model.AdminRegisterResp{
		IsSuccess: true,
	}, err
}

func adminRegister(name string, password string) (err error) {
	err = data_access.CreateAdmin(name, password)
	return err
}

func CreateEvent(req model.CreateEventReq) (*model.CreateEventResp, error) {
	eventId, err := data_access.CreateEvent(req.Event)
	if err != nil {
		return nil, err
	}
	email2InvCodes := generateInvCode(req.Voters)
	err = data_access.CreateInvitations(eventId, email2InvCodes)
	for email, invCode := range email2InvCodes {
		err = email2Voter(email, invCode, req.Event.Name)
		if err != nil {
			return nil, err
		}
	}

	return &model.CreateEventResp{EventId: eventId, IsSuccess: true}, nil
}

func generateInvCode(emails []string) map[string]string {
	email2InvCode := make(map[string]string)
	for _, email := range emails {
		invCode := randstr.String(8)
		email2InvCode[email] = invCode
	}
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
		fmt.Println(err)
		return err
	}
	fmt.Println("Email Sent Successfully!")

	duration := time.Since(start)
	fmt.Println("send email: ", duration)
	return nil
}

func EndEvent(req model.EndEventReq) (*model.EndEventResp, error) {
	err := data_access.EndEvent(req.EventId)
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

	return &model.EndEventResp{IsSuccess: true}, nil
}

func GetEvent(req model.GetEventReq) (*model.GetEventResp, error) {
	event, err := data_access.GetEvent(req.EventId)
	if err != nil {
		return nil, err
	}
	return &model.GetEventResp{Event: event}, nil
}

func Download(req model.DownloadReq) (*model.DownloadResp, error) {
	InvCode2Details, err := data_access.GetVoteDetailsByEventId(req.EventId)
	if err != nil {
		return nil, err
	}
	return &model.DownloadResp{InvitationCode2Details: InvCode2Details}, nil
}
