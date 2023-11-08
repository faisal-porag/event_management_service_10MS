package controllers

import (
	"event_management_service_10MS/state"
	"event_management_service_10MS/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func HandleWorkshopList(s *state.State) func(c *gin.Context) {
	type SuccessHealthCheck struct {
		Status string `json:"status"`
	}

	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
		language := utils.GetLanguageFromRequest(c.Request)
		logger := log.With().
			Str("handler", "HandleWorkshopList").
			Logger()

		logger.Info().Msg(language)

		utils.ShowAppCommonSuccessResponse(c, "request success", SuccessHealthCheck{
			Status: "Ok",
		})

		return
	}
}
