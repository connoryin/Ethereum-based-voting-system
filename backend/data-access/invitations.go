package data_access

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/6675-voting-system/voting-system/backend/model"
)

func CreateInvitations(eventId int, email2InvCodes map[string]string) (err error) {
	for _, invCode := range email2InvCodes {
		err := CreateInvitation(eventId, invCode)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateInvitation(eventId int, invCode string) (err error) {
	db := GetDB()
	inv := InvitationPO{
		InvitationCode: invCode,
		EventId:        eventId,
		VoteDetails:    []byte{},
	}

	err = db.Table("invitations").Create(&inv).Error
	return
}

func CreateInvitationContract(eventId int, invCode string) (err error) {
	// db := GetDB()
	inv := InvitationPO{
		InvitationCode: invCode,
		EventId:        eventId,
		VoteDetails:    []byte{},
	}

	err = db.Table("invitations").Create(&inv).Error
	return err
}

func GetEventIDByInvitationID(invitationCode string) (eventID int, err error) {
	var invitation InvitationPO
	db := GetDB()
	err = db.Table("invitations").Where("invitation_code = ?", invitationCode).First(&invitation).Error
	if err != nil {
		log.Panic(err)
		return -1, err
	}
	eventID = invitation.EventId
	return
}

type InvitationPO struct {
	Id             int    `gorm:"primaryKey;autoIncrement:true"`
	InvitationCode string `gorm:"column:invitation_code"`
	EventId        int    `gorm:"column:event_id"`
	VoteDetails    []byte `gorm:"column:vote_details"`
}

func GetVoteDetailsByEventId(eventId int) (map[string]string, error) {
	var invitations []InvitationPO
	db := GetDB()
	err := db.Table("invitations").Where("event_id = ?", eventId).Find(&invitations).Error
	if err != nil {
		return nil, err
	}

	InvCode2Details := map[string]string{}
	for _, inv := range invitations {
		var candidates []model.Candidate
		var canNames []string
		json.Unmarshal(inv.VoteDetails, &candidates)
		for _, c := range candidates {
			canNames = append(canNames, c.Name)
		}
		InvCode2Details[inv.InvitationCode] = fmt.Sprint(canNames)
	}
	return InvCode2Details, nil
}

func candidatesToJson(in []model.Candidate) string {
	json, _ := json.Marshal(in)
	return string(json)
}

func GetVoteDetailsByInvitationID(invitationCode string) ([]model.Candidate, error) {
	var invitation InvitationPO
	db := GetDB()
	err := db.Table("invitations").Where("invitation_code = ?", invitationCode).First(&invitation).Error
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	var candidatesD0 []model.Candidate
	json.Unmarshal(invitation.VoteDetails, &candidatesD0)
	return candidatesD0, nil
}

func SubmitVotes(invitationCode string, candidates_names []string) (err error) {
	db := GetDB()

	candidates := make([]model.Candidate, 0)
	for _, name := range candidates_names {
		candidates = append(candidates, model.Candidate{
			Name: name,
		})
	}

	candidatesJson, err := json.Marshal(candidates)
	if err != nil {
		return err
	}

	err = db.Table("invitations").Where("invitation_code = ?", invitationCode).Update("vote_details", candidatesJson).Error
	if err != nil {
		log.Panic(err)
		return err
	}

	return nil
}

func GetVoteNames(invitationCode string) (voteNames []int, err error) {
	db := GetDB()
	var voteIDsBin []byte
	err = db.Table("invitations").Where("invitation_code = ?", invitationCode).Select("vote_details").First(&voteIDsBin).Error
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	return byteArrayToIntArray(voteIDsBin), nil
}

func intArrayToByteArray(intArray []int) []byte {
	byteArray := make([]byte, len(intArray))
	for i, intNum := range intArray {
		byteArray[i] = byte(intNum)
	}
	return byteArray
}

func byteArrayToIntArray(byteArray []byte) []int {
	intArray := make([]int, len(byteArray))
	for i, byteNum := range byteArray {
		intArray[i] = int(byteNum)
	}
	return intArray
}
