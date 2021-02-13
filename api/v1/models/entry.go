package models

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx/types"
	"gitlab.devgroup.tech/shkolkovo/romanych/db"
	"math"
	"time"
)

type TimeLogEntry struct {
	Id        int64
	UserId    int64
	StartTime time.Time
	EndTime   mysql.NullTime
	IsEnded   types.BitBool
	Minutes   int `db:"-"`
	// TODO: Add payment link
}

func (model *TimeLogEntry) MarkEnd() {
	model.EndTime.Time = time.Now()
	model.EndTime.Valid = true
	model.IsEnded = true
}

func (model *TimeLogEntry) Calculate(endTime *time.Time) {
	if endTime == nil {
		endTime = &model.EndTime.Time
	}
	model.Minutes = int(math.Round(endTime.Sub(model.StartTime).Minutes()))
	// TODO: Create payment link
}


func GetLastUserEntry(userId int64) (model *TimeLogEntry, err error) {
	model = &TimeLogEntry{}
	if err = db.Database().Get(model, "SELECT \n    Id, UserId, StartTime, EndTime, IsEnded\nFROM TimeLog\n    WHERE\n        UserId = ?\n            \nORDER BY Id DESC\nLIMIT 1", userId); err != nil && err != sql.ErrNoRows {
		return
	}
	return model, nil
}