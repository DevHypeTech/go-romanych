package v1

import (
	"github.com/gin-gonic/gin"
	"gitlab.devgroup.tech/shkolkovo/romanych/api/v1/middlewares"
	"gitlab.devgroup.tech/shkolkovo/romanych/api/v1/models"
	"gitlab.devgroup.tech/shkolkovo/romanych/api/v1/response"
	"gitlab.devgroup.tech/shkolkovo/romanych/db"
	"time"
)

func AllApi(r *gin.Engine) {
	group := r.Group("v1")

	journal := group.Group("journal")
	journal.Use(middlewares.UserMiddleware)

	journal.POST("start", func(c *gin.Context) {
		userId := c.GetInt64("UserId")

		conn := db.Database()
		res, err := conn.Exec("INSERT INTO TimeLog (UserId, StartTime, EndTime) VALUES (?, ?, NULL)", userId, time.Now())
		if err != nil {
			c.AbortWithStatusJSON(500, response.NewErrorResponse(err.Error(), 100))
			return
		}

		var timeLogId int64
		if timeLogId, err = res.LastInsertId(); err != nil {
			c.AbortWithStatusJSON(500, response.NewErrorResponse(err.Error(), 101))
			return
		}

		c.JSON(200, timeLogId)
	})

	journal.POST("end", func(c *gin.Context) {
		userId := c.GetInt64("UserId")

		var record *models.TimeLogEntry
		var err error
		if record, err = models.GetLastUserEntry(userId); err != nil {
			c.AbortWithStatusJSON(500, response.NewErrorResponse(err.Error(), 100))
			return
		}

		if record.IsEnded == true || record.Id == 0 || record.EndTime.Valid == true {
			c.AbortWithStatusJSON(404, response.NewErrorResponse("no TimeLog entry", 101))
			return
		}

		record.MarkEnd()

		if _, err = db.Database().NamedExec("UPDATE TimeLog SET EndTime=:EndTime, IsEnded=:IsEnded WHERE Id = :Id", record); err != nil {
			c.AbortWithStatusJSON(500, response.NewErrorResponse(err.Error(), 102))
			return
		}

		record.Calculate(nil)
		c.JSON(200, record)
	})

	group.GET("entry/:userId", func(c *gin.Context) {
		var req struct {
			UserId int64 `binding:"required" uri:"userId"`
		}
		var err error
		if err = c.ShouldBindUri(&req); err != nil {
			c.AbortWithStatusJSON(400, response.NewErrorResponse(err.Error(), 100))
			return
		}

		var model *models.TimeLogEntry
		if model, err = models.GetLastUserEntry(req.UserId); err != nil {
			c.AbortWithStatusJSON(500, response.NewErrorResponse(err.Error(), 101))
			return
		}

		if model.Id == 0 {
			c.AbortWithStatusJSON(404, response.NewErrorResponse("not found", 102))
			return
		}

		t := time.Now()
		model.Calculate(&t)

		c.JSON(200, model)
	})
}
