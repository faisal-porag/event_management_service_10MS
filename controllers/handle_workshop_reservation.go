package controllers

import (
	"errors"
	"event_management_service_10MS/data_params"
	"event_management_service_10MS/db_repo"
	"event_management_service_10MS/services"
	"event_management_service_10MS/state"
	"event_management_service_10MS/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func HandleWorkshopReservation(s *state.State) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
		logger := log.With().
			Str("handler", "HandleWorkshopReservation").
			Logger()

		var reqPayload data_params.GetWorkshopReservationDataPayload
		if err := c.ShouldBindJSON(&reqPayload); err != nil {
			logger.Error().Err(err).Msg("requestPayload decode error!")
			utils.ShowAppCommonErrorResponse(
				c,
				http.StatusBadRequest,
				"The provided information is invalid. Please recheck and try again.",
				nil,
			)
			return
		}

		if s.Cfg.DebugMode {
			logger.Debug().Msgf("HandleWorkshopReservation payload: %+v", reqPayload)
		}

		dataErr := utils.ValidateStructData(reqPayload)
		if dataErr != nil {
			utils.ShowAppCommonErrorResponse(
				c,
				http.StatusBadRequest,
				dataErr.Error(),
				nil,
			)
			return
		}

		data, err := services.GetWorkshopReservationServices(reqPayload.Name, reqPayload.Email, s)
		if err != nil {
			logger.Error().Err(err)
			if errors.Is(err, db_repo.NoDataFound) {
				utils.ShowAppCommonErrorResponse(c, http.StatusNotFound, "no reservation found", nil)
			} else {
				utils.ShowAppCommonErrorResponse(c, http.StatusInternalServerError, err.Error(), nil)
			}
			return
		}

		c.JSON(http.StatusOK, data)
		return
	}
}
