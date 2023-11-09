package controllers

import (
	"errors"
	"event_management_service_10MS/data_responses"
	"event_management_service_10MS/db_repo"
	"event_management_service_10MS/services"
	"event_management_service_10MS/state"
	"event_management_service_10MS/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
	"strings"
)

func HandleWorkshopListForSingleEvent(s *state.State) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
		logger := log.With().
			Str("handler", "HandleWorkshopListForSingleEvent").
			Logger()

		strEventId := c.Param("event_id")
		if strings.TrimSpace(strEventId) == "" {
			utils.ShowAppCommonErrorResponse(
				c,
				http.StatusBadRequest,
				"event_id required",
				nil,
			)
			return
		}

		eventIdInt, err := strconv.ParseInt(strEventId, 10, 64)
		if err != nil {
			utils.ShowAppCommonErrorResponse(
				c,
				http.StatusBadRequest,
				err.Error(),
				nil,
			)
			return
		}

		event, workshops, err := services.GetWorkshopListForEventServices(eventIdInt, s)
		if err != nil {
			logger.Error().Err(err).Msg("services.GetWorkshopListForEventServices:err")
			if errors.Is(err, db_repo.NoDataFound) {
				utils.ShowAppCommonErrorResponse(c, http.StatusNotFound, "no event found", nil)
			} else {
				utils.ShowAppCommonErrorResponse(c, http.StatusInternalServerError, err.Error(), nil)
			}
			return
		}

		finalResponse := data_responses.WorkshopsDataListResponses{
			Id:        event.Id,
			Title:     event.Title,
			StartAt:   event.StartAt,
			EndAt:     event.EndAt,
			Workshops: workshops,
		}

		c.JSON(http.StatusOK, finalResponse)
		return
	}
}
