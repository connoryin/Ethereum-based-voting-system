package handler

import (
	"fmt"
	"net/http"

	contract "github.com/6675-voting-system/voting-system/backend/contract"
	"github.com/6675-voting-system/voting-system/backend/model"
	"github.com/6675-voting-system/voting-system/backend/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var AdminEventsInstance contract.AdminEvents

func InitAdminEvents() {
	auth, _ := contract.SetAuthOptions()
	AdminEventsInstance.DeployEvents(auth)
	AdminEventsInstance.Init()
}

func LoginHandler(c *gin.Context) {
	var req model.LoginReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request. Illegal JSON")
		return
	}
	resp, err := service.Login(req)
	if err != nil {
		c.String(http.StatusInternalServerError, "err: %s", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func UserVoteDetailsHandler(c *gin.Context) {
	var req model.GetVoteDetailsReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request. Illegal JSON")
		return
	}
	resp, err := service.GetVoteDetails(req)
	if err != nil {
		c.String(http.StatusInternalServerError, "err: %s", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func UserVoteDetailsContractHandler(c *gin.Context) {
	var req model.GetVoteDetailsReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request. Illegal JSON")
		return
	}
	resp, err := contract.GetVoteDetails(req, &AdminEventsInstance)

	if err != nil {
		c.String(http.StatusInternalServerError, "err: %s", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func UserVoteSubmitHandler(c *gin.Context) {
	var req model.VoteReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request. Illegal JSON")
		return
	}
	resp, err := service.SubmitVote(req)
	if err != nil {
		c.String(http.StatusInternalServerError, "err: %s", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func UserVoteSubmitContractHandler(c *gin.Context) {
	var req model.VoteReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request. Illegal JSON")
		return
	}
	// resp, err := service.SubmitVote(req)
	resp, err := AdminEventsInstance.SubmitVote(req)
	if err != nil {
		c.String(http.StatusInternalServerError, "err: %s", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func AdminLoginHandler(c *gin.Context) {
	var req model.AdminLoginReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	resp, err := service.AdminLogin(req)
	if err != nil {
		c.String(http.StatusInternalServerError, "err: %s", err)
		return
	}

	session := sessions.Default(c)
	session.Set("adminid", resp.AdminId)
	session.Save()
	c.JSON(http.StatusOK, resp)
}

func AdminDetailHandler(c *gin.Context) {
	session := sessions.Default(c)
	adminId := session.Get("adminid").(int)
	resp, err := service.AdminDetail(adminId)
	if err != nil {
		c.String(http.StatusInternalServerError, "err: %s", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func AdminDetailContractHandler(c *gin.Context) {
	session := sessions.Default(c)
	adminId := session.Get("adminid").(int)
	// resp, err := service.AdminDetail(adminId)
	resp, err := AdminEventsInstance.AdminDetail(adminId)
	if err != nil {
		c.String(http.StatusInternalServerError, "err: %s", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func AdminRegisterHandler(c *gin.Context) {
	var req model.AdminRegisterReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	resp, err := service.AdminRegister(req)
	if err != nil {
		c.String(http.StatusInternalServerError, "err: %s", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func AdminCreateEventContractHandler(c *gin.Context) {
	var req model.CreateEventReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	session := sessions.Default(c)
	adminId := session.Get("adminid").(int)
	if adminId != req.Event.AdminId {
		c.String(http.StatusUnauthorized, "unauthorized")
		return
	}

	// resp, err := service.CreateEvent(req)
	resp, err := AdminEventsInstance.CreateEvent(req, adminId)
	if err != nil {
		c.String(http.StatusInternalServerError, "err: %s", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func AdminEndEventContractHandler(c *gin.Context) {
	var req model.EndEventReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	session := sessions.Default(c)
	adminId := session.Get("adminid").(int)
	// eventIds, err := data_access.GetEventIdsByAdminId(adminId)

	// event, err:= AdminEventsInstance.GetEvent(req.EventId)

	events, err := AdminEventsInstance.GetEventsByAdminId(adminId)
	if err != nil {
		c.String(http.StatusUnauthorized, "bad request")
		return
	}
	eventIds := make([]int, 0)
	for _, event := range events {
		eventIds = append(eventIds, event.Id)
	}

	if !checkReqEventIdValid(req.EventId, eventIds) {
		c.String(http.StatusUnauthorized, "unauthorized")
		return
	}

	// resp, err := service.EndEvent(req)
	resp, err := AdminEventsInstance.EndEvent(req, adminId)
	if err != nil {
		c.String(http.StatusInternalServerError, "err: %s", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func checkReqEventIdValid(reqEventId int, eventIds []int) bool {
	IsValid := false
	for _, id := range eventIds {
		if reqEventId == id {
			IsValid = true
			break
		}
	}
	return IsValid
}

func AdminGetEventContractHandler(c *gin.Context) {
	var req model.GetEventReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	// session := sessions.Default(c)
	// adminId := session.Get("adminid").(int)
	// eventIds, err := data_access.GetEventIdsByAdminId(adminId)
	// events, err := AdminEventsInstance.GetEventsByAdminId(adminId)

	event, err := AdminEventsInstance.GetEvent(req.EventId)

	if err != nil {
		c.String(http.StatusUnauthorized, "bad request")
		return
	}

	fmt.Printf("events: %v", event)

	// eventIds := make([]int, 0)
	// for _, event := range events {
	// 	eventIds = append(eventIds, event.Id)
	// }

	// if !checkReqEventIdValid(req.EventId, eventIds) {
	// 	c.String(http.StatusUnauthorized, "unauthorized")
	// 	return
	// }

	// resp, err := service.GetEvent(req)
	// if err != nil {
	// 	c.String(http.StatusInternalServerError, "err: %s", err)
	// 	return
	// }
	c.JSON(http.StatusOK, event)
}

func DownloadHandler(c *gin.Context) {
	var req model.DownloadReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	session := sessions.Default(c)
	adminId := session.Get("adminid").(int)
	// eventIds, err := data_access.GetEventIdsByAdminId(adminId)
	events, err := AdminEventsInstance.GetEventsByAdminId(adminId)
	if err != nil {
		c.String(http.StatusUnauthorized, "bad request")
		return
	}

	eventIds := make([]int, 0)
	for _, event := range events {
		eventIds = append(eventIds, event.Id)
	}

	if err != nil {
		c.String(http.StatusUnauthorized, "bad request")
		return
	}

	if !checkReqEventIdValid(req.EventId, eventIds) {
		c.String(http.StatusUnauthorized, "unauthorized")
		return
	}

	resp, err := service.Download(req)
	if err != nil {
		c.String(http.StatusInternalServerError, "err: %s", err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
