package data_access

import (
	"encoding/json"

	"github.com/6675-voting-system/voting-system/backend/model"
)

func GetAdmin(name, password string) (admin *model.Admin, err error) {
	db := GetDB()
	err = db.Table("admins").Where("name = ? AND password_hash = ?", name, password).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func CreateAdmin(name, password string) (err error) {
	admin := model.Admin{
		Name:     name,
		Password: password,
	}
	db := GetDB()
	err = db.Table("admins").Create(&admin).Error
	return err
}

func GetEventsByAdminId(adminId int) (events []model.Event, err error) {
	var pos []EventPO
	db := GetDB()
	err = db.Table("events").Where("admin_id = ?", adminId).Find(&pos).Error
	if err != nil {
		return nil, err
	}
	for _, po := range pos {
		do := EventPO2DO(po)
		events = append(events, do)
	}
	return events, nil
}

type EventPO struct {
	Id                  int
	AdminId             int
	Name                string
	Description         string
	MaxVoteNumPerPerson int `gorm:"column:vote_bound"`
	Candidates          []byte
	IsEnd               bool `gorm:"column:is_over"`
	TotalVoteNum        int  `gorm:"column:suppose_voter_num"`
	ReceivedVoteNum     int  `gorm:"column:actual_voter_num"`
	Address             string
	// IsAnonymous         bool
}

func EventDO2PO(do model.Event) (po EventPO) {
	candidates, _ := json.Marshal(do.Candidates)
	po = EventPO{
		Id:                  do.Id,
		AdminId:             do.AdminId,
		Name:                do.Name,
		Description:         do.Description,
		MaxVoteNumPerPerson: do.MaxVoteNumPerPerson,
		Candidates:          candidates,
		IsEnd:               do.IsEnd,
		TotalVoteNum:        do.TotalVoteNum,
		ReceivedVoteNum:     do.ReceivedVoteNum,
		// IsAnonymous:         do.IsAnonymous,
	}
	return po
}

func EventPO2DO(po EventPO) (do model.Event) {
	var candidates []model.Candidate
	_ = json.Unmarshal(po.Candidates, &candidates)
	do = model.Event{
		Id:                  po.Id,
		AdminId:             po.AdminId,
		Name:                po.Name,
		Description:         po.Description,
		MaxVoteNumPerPerson: po.MaxVoteNumPerPerson,
		Candidates:          candidates,
		IsEnd:               po.IsEnd,
		TotalVoteNum:        po.TotalVoteNum,
		// IsAnonymous:         po.IsAnonymous,
		ReceivedVoteNum: po.ReceivedVoteNum,
	}
	return do
}

func CreateEvent(event model.Event) (eventId int, err error) {
	po := EventDO2PO(event)
	db := GetDB()
	err = db.Table("events").Create(&po).Error
	if err != nil {
		return -1, err
	}
	return po.Id, nil
}

func EndEvent(eventId int) (err error) {
	db := GetDB()
	err = db.Table("events").Where("id = ?", eventId).Update("is_over", true).Error
	return err
}

func GetEvent(eventId int) (event model.Event, err error) {
	var po EventPO
	db := GetDB()
	err = db.Table("events").Where("id = ?", eventId).First(&po).Error
	if err != nil {
		return model.Event{}, err
	}
	do := EventPO2DO(po)
	return do, nil
}
