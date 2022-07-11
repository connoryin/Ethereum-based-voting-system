package data_access

import (
	"encoding/json"

	"github.com/6675-voting-system/voting-system/backend/model"
	"gorm.io/gorm"
)

func UpdateCandidatesOnSubmit(eventID int, candidatesNames []string) (err error) {
	db := GetDB()
	err = db.Table("events").Where("id = ?", eventID).Update("actual_voter_num", gorm.Expr("actual_voter_num + ?", 1)).Error
	if err != nil {
		return err
	}

	var po EventPO
	err = db.Table("events").Where("id = ?", eventID).First(&po).Error
	if err != nil {
		return err
	}

	var candidatesD0 []*model.Candidate
	err = json.Unmarshal(po.Candidates, &candidatesD0)
	if err != nil {
		return err
	}

	for _, cand := range candidatesD0 {
		for _, name := range candidatesNames {
			if cand.Name == name {
				cand.VoteNumGet++
				break
			}
		}
	}

	candidatesP0, err := json.Marshal(candidatesD0)
	if err != nil {
		return err
	}

	err = db.Table("events").Where("id = ?", eventID).Update("candidates", candidatesP0).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateEventWinner(eventId int, candidate []model.Candidate) (err error) {
	candidates, _ := json.Marshal(candidate)
	db := GetDB()
	err = db.Table("events").Where("id = ?", eventId).Update("candidates", candidates).Error
	return err
}

func GetEventIdsByAdminId(adminId int) ([]int, error) {
	var eventIds []int
	var pos []EventPO
	db := GetDB()
	err := db.Table("events").Where("admin_id = ?", adminId).Find(&pos).Error
	if err != nil {
		return nil, err
	}
	for _, po := range pos {
		eventIds = append(eventIds, po.Id)
	}
	return eventIds, nil
}
