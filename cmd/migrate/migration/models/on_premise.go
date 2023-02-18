package models

import "time"

type OpsOnPremise struct {
	OnPremiseId int       `gorm:"primaryKey;autoIncrement;comment:私有化项目id"  json:"onPremiseId"`
	Name        string    `json:"name" gorm:"type:varchar(64);not null;unique;comment:私有化项目名称"`
	PmId        int       `json:"pmId" gorm:"type:bigint;default null;comment:负责项目经理"`
	StartDate   time.Time `json:"startDate" gorm:"type:datetime(3);default 2001-01-01;comment:启动时间"`
	PocDate     time.Time `json:"pocDate" gorm:"type:datetime(3);default 2001-02-01;comment:poc时间"`
	UatDate     time.Time `json:"uatDate" gorm:"type:datetime(3);default 2001-03-01;comment:uat时间"`
	GoLiveDate  time.Time `json:"goLiveDate" gorm:"type:datetime(3);default 2001-04-01;comment:产线时间"`
	Remark      string    `json:"remark" gorm:"type:varchar(255);default null;comment:备注"`
	Status      string    `json:"status" gorm:"type:varchar(6);default unknown;comment:项目状态"`
	ControlBy
	ModelTime
}

func (OpsOnPremise) TableName() string {
	return "ops_on_premise"
}
