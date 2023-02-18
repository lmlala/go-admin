package dto

import (
	"go-admin/app/ops/models"
	"time"

	"go-admin/common/dto"
	common "go-admin/common/models"
)

type OpsOnPremiseGetPageReq struct {
	dto.Pagination `search:"-"`
	OnPremiseId    int `form:"onPremiseId" search:"type:exact;column:on_premise_id;table:ops_on_premise" comment:"私有化项目id"`
	// Name           string    `form:"name" comment:"私有化项目名称" search:"type:exact;column:name;table:ops_on_premise"`
	// PmId           int       `form:"pmId" comment:"负责项目经理" search:"type:contains;column:pm_id;table:ops_on_premise"`
	// StartDate      string    `form:"startDate" comment:"启动时间" search:"type:gt;column:start_date;table:ops_on_premise"`
	// PocDate        time.Time `form:"pocDate" comment:"poc时间" search:"type:lt;column:poc_date;table:ops_on_premise"`
	// UatDate        time.Time `form:"uatDate" comment:"uat时间" search:"type:lt;column:uat_date;table:ops_on_premise"`
	// GoLiveDate     time.Time `form:"goLiveDate" comment:"产线交付时间" search:"type:lt;column:go_live_date;table:ops_on_premise"`
	// Remark         string    `form:"remark" comment:"备注" search:"type:isnull;column:remark;table:ops_on_premise"`
	// Status         string    `form:"status" comment:"项目状态" search:"type:exact;column:status;table:ops_on_premise"`
	// common.ControlBy
	// common.ModelTime

	// DeptJoin       `search:"type:left;on:dept_id:dept_id;table:sys_user;join:sys_dept"`
	OpsOnPremiseOrder
}

type OpsOnPremiseOrder struct {
	OnPremiseOrder string `search:"type:order;column:on_premise_id;table:ops_on_premise" form:"onPremiseId"`
	// UsernameOrder  string `search:"type:order;column:username;table:sys_user" form:"usernameOrder"`
	// StatusOrder    string `search:"type:order;column:status;table:sys_user" form:"statusOrder"`
	// CreatedAtOrder string `search:"type:order;column:created_at;table:sys_user" form:"createdAtOrder"`
}

// type DeptJoin struct {
// 	DeptId string `search:"type:contains;column:dept_path;table:sys_dept" form:"deptId"`
// }

func (m *OpsOnPremiseGetPageReq) GetNeedSearch() interface{} {
	return *m
}

// type ResetSysUserPwdReq struct {
// 	UserId   int    `json:"userId" comment:"用户ID" vd:"$>0"` // 用户ID
// 	Password string `json:"password" comment:"密码" vd:"len($)>0"`
// 	common.ControlBy
// }

// func (s *ResetSysUserPwdReq) GetId() interface{} {
// 	return s.UserId
// }

// func (s *ResetSysUserPwdReq) Generate(model *models.SysUser) {
// 	if s.UserId != 0 {
// 		model.UserId = s.UserId
// 	}
// 	model.Password = s.Password
// }

// type UpdateSysUserAvatarReq struct {
// 	UserId int    `json:"userId" comment:"用户ID" vd:"len($)>0"` // 用户ID
// 	Avatar string `json:"avatar" comment:"头像" vd:"len($)>0"`
// 	common.ControlBy
// }

// func (s *UpdateSysUserAvatarReq) GetId() interface{} {
// 	return s.UserId
// }

// func (s *UpdateSysUserAvatarReq) Generate(model *models.SysUser) {
// 	if s.UserId != 0 {
// 		model.UserId = s.UserId
// 	}
// 	model.Avatar = s.Avatar
// }

// type UpdateSysUserStatusReq struct {
// 	UserId int    `json:"userId" comment:"用户ID" vd:"$>0"` // 用户ID
// 	Status string `json:"status" comment:"状态" vd:"len($)>0"`
// 	common.ControlBy
// }

// func (s *UpdateSysUserStatusReq) GetId() interface{} {
// 	return s.UserId
// }

// func (s *UpdateSysUserStatusReq) Generate(model *models.SysUser) {
// 	if s.UserId != 0 {
// 		model.UserId = s.UserId
// 	}
// 	model.Status = s.Status
// }

type OnPremiseInsertReq struct {
	OnPremiseId int    `json:"onPremiseId" comment:"私有化项目id"`
	Name        string `json:"name" comment:"私有化项目名称" vd:"len($)>0"`
	PmId        int    `json:"pmId" comment:"负责项目经理"`
	StartDate   string `json:"startDate" comment:"启动时间" default:"2001-01-01"`
	// PocDate     time.Time `json:"pocDate" comment:"poc时间" default:"2001-02-01"`
	// UatDate     time.Time `json:"uatDate" comment:"uat时间" default:"2001-03-01"`
	// GoLiveDate  time.Time `json:"goLiveDate" comment:"产线交付时间" default:"2001-04-01"`
	Remark string `json:"remark" comment:"备注"`
	Status string `json:"status" comment:"项目状态" default:"unknown"`
	common.ControlBy
	common.ModelTime
}

func (o *OnPremiseInsertReq) Generate(model *models.OpsOnPremise) {
	if o.OnPremiseId != 0 {
		model.OnPremiseId = o.OnPremiseId
	}

	stime, _ := time.Parse("2006-01-02", o.StartDate)

	model.Name = o.Name
	model.PmId = o.PmId

	model.StartDate = stime
	model.PocDate = stime.AddDate(0, 2, 0)
	model.UatDate = stime.AddDate(0, 4, 0)
	model.GoLiveDate = stime.AddDate(0, 6, 0)

	model.Remark = o.Remark
	model.Status = "start"

	model.CreateBy = o.CreateBy
	model.CreatedAt = o.CreatedAt
}

func (o *OnPremiseInsertReq) GetId() interface{} {
	return o.OnPremiseId
}

type OnPremiseUpdateReq struct {
	OnPremiseId int    `json:"onPremiseId" comment:"私有化项目id"`
	Name        string `json:"name" comment:"私有化项目名称"`
	PmId        int    `json:"pmId" comment:"负责项目经理"`
	StartDate   string `json:"startDate" comment:"启动时间"`
	PocDate     string `json:"pocDate" comment:"poc时间"`
	UatDate     string `json:"uatDate" comment:"uat时间"`
	GoLiveDate  string `json:"goLiveDate" comment:"产线交付时间"`
	Remark      string `json:"remark" comment:"备注"`
	Status      string `json:"status" comment:"项目状态"`
	common.ControlBy
}

func (o *OnPremiseUpdateReq) Generate(model *models.OpsOnPremise) {
	if o.OnPremiseId != 0 {
		model.OnPremiseId = o.OnPremiseId
	}

	stime, _ := time.Parse("2006-01-02", o.StartDate)
	ptime, _ := time.Parse("2006-01-02", o.PocDate)
	utime, _ := time.Parse("2006-01-02", o.UatDate)
	gtime, _ := time.Parse("2006-01-02", o.GoLiveDate)

	model.Name = o.Name
	model.PmId = o.PmId
	model.StartDate = stime
	model.PocDate = ptime
	model.UatDate = utime
	model.GoLiveDate = gtime
	model.Remark = o.Remark
	model.Status = o.Status
}

func (o *OnPremiseUpdateReq) GetId() interface{} {
	return o.OnPremiseId
}

type OnPremiseById struct {
	dto.ObjectById
	common.ControlBy
}

func (o *OnPremiseById) GetId() interface{} {
	if len(o.Ids) > 0 {
		o.Ids = append(o.Ids, o.Id)
		return o.Ids
	}
	return o.Id
}

func (s *OnPremiseById) GenerateM() (common.ActiveRecord, error) {
	return &models.OpsOnPremise{}, nil
}

// PassWord 密码
// type PassWord struct {
// 	NewPassword string `json:"newPassword" vd:"len($)>0"`
// 	OldPassword string `json:"oldPassword" vd:"len($)>0"`
// }
