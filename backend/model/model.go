package model

import "fmt"

type UserStatus int

const (
	USERSTATUS_UNDEFINED UserStatus = 0
	USERSTATUS_VOTING               = 1 // event not end, user not vote
	USERSTATUS_VOTED                = 2 // voted (no matter event end or not)
	USERSTATUS_UNVOTE               = 3 // event ends, user not vote
)

type StatusCode int

const (
	STATUSCODE_UNDEFINED StatusCode = 0
	STATUSCODE_SUCCESS              = 1
)

type AdminDetailReq struct {
	AdminId int `json:"admin_id"`
}

type AdminDetailResp struct {
	InVotingEvents    []Event `json:"in_voting_events"`
	EndedVotingEvents []Event `json:"ended_voting_events"`
}

type AdminLoginReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type AdminLoginResp struct {
	StatusCode int `json:"status_code"`
	AdminId    int `json:"admin_id"`
}

type AdminRegisterReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type AdminRegisterResp struct {
	IsSuccess bool `json:"is_success"`
}

type Event struct {
	Id                  int         `json:"id"`
	AdminId             int         `json:"admin_id"`
	Name                string      `json:"name"`
	Description         string      `json:"description"`
	MaxVoteNumPerPerson int         `json:"max_vote_num_per_person" gorm:"column:vote_bound"`
	Candidates          []Candidate `json:"candidates"`
	IsEnd               bool        `json:"is_end" gorm:"column:is_over"`
	TotalVoteNum        int         `json:"total_vote_num" gorm:"column:suppose_voter_num"`
	ReceivedVoteNum     int         `json:"received_vote_num" gorm:"column:actual_voter_num"`
	// AddressHex          string
	IsAnonymous bool
}

type Candidate struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	VoteNumGet  int    `json:"vote_num_get"`
	IsWinner    bool   `json:"is_winner"`
}

func (c *Candidate) GetStr() string {
	return fmt.Sprintf("%v,%v,%v,%v", c.Name, c.Description, c.VoteNumGet, c.IsWinner)
}

type Admin struct {
	Id       int `gorm:"primaryKey"`
	Name     string
	Password string `gorm:"column:password_hash"`
}

type LoginReq struct {
	InvCode string `json:"inv_code"`
}

type LoginResp struct {
	UserStatus UserStatus `json:"user_status"`
	EventId    int        `json:"event_id"`
}

type AdminDetailsReq struct {
}

type AdminDetailsResp struct {
	votingEvents  []Event `json:"voting_events"`
	VoteEndEvents []Event `json:"vote_end_events"`
}

type GetVoteDetailsReq struct {
	InvCode string `json:"inv_code"`
}

type GetVoteDetailsResp struct {
	Event    Event       `json:"event"`
	IsVoted  bool        `json:"is_voted"`
	SelfVote []Candidate `json:"self_vote"`
}

type VoteReq struct {
	InvCode             string   `json:"inv_code"`
	EventId             int      `json:"event_id"`
	VotedCandidateNames []string `json:"voted_candidate_names"`
}

type VoteResp struct {
	IsSuccess bool `json:"is_success"`
}

type CreateEventReq struct {
	Event Event `json:"event"`
	// Can Voters be merged into Event?
	Voters []string `json:"voters"`
}

type CreateEventResp struct {
	EventId   int  `json:"event_id"`
	IsSuccess bool `json:"is_success"`
}

type EndEventReq struct {
	EventId int `json:"event_id"`
}

type EndEventResp struct {
	IsSuccess bool `json:"is_success"`
}

type GetEventReq struct {
	EventId int `json:"event_id"`
}

type GetEventResp struct {
	Event Event `json:"event"`
}

type DownloadReq struct {
	EventId int `json:"event_id"`
}

type DownloadResp struct {
	InvitationCode2Details map[string]string `json:"invitation_code_2_details"`
}

type Invitation struct {
	Id             int         `gorm:"primaryKey;autoIncrement:true"`
	InvitationCode string      `json:"invitation_code" gorm:"column:invitation_code"`
	EventId        int         `json:"event_id" gorm:"column:event_id"`
	VoteDetails    []Candidate `json:"vote_details" gorm:"column:vote_details"`
}
