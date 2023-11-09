package controllers

import (
	"errors"
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

func HandleWorkshopDetails(s *state.State) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
		logger := log.With().
			Str("handler", "HandleWorkshopDetails").
			Logger()

		strWorkshopId := c.Param("workshop_id")
		if strings.TrimSpace(strWorkshopId) == "" {
			utils.ShowAppCommonErrorResponse(
				c,
				http.StatusBadRequest,
				"workshop_id required",
				nil,
			)
			return
		}

		workshopIdInt, err := strconv.ParseInt(strWorkshopId, 10, 64)
		if err != nil {
			utils.ShowAppCommonErrorResponse(
				c,
				http.StatusBadRequest,
				err.Error(),
				nil,
			)
			return
		}

		workshopDetails, err := services.GetWorkshopDetailsServices(workshopIdInt, s)
		if err != nil {
			logger.Error().Err(err)
			if errors.Is(err, db_repo.NoDataFound) {
				utils.ShowAppCommonErrorResponse(c, http.StatusNotFound, "no workshop found", nil)
			} else {
				utils.ShowAppCommonErrorResponse(c, http.StatusInternalServerError, err.Error(), nil)
			}
			return
		}

		c.JSON(http.StatusOK, workshopDetails)
		return
	}
}
