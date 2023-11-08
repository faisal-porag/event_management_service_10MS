package controllers

import (
	"event_management_service_10MS/data_responses"
	"event_management_service_10MS/services"
	"event_management_service_10MS/state"
	"event_management_service_10MS/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
	"strings"
)

func HandleEventList(s *state.State) func(c *gin.Context) {
	type EventListResponse struct {
		Events     []*data_responses.EventListResponses `json:"events"`
		Pagination *data_responses.Pagination           `json:"pagination"`
	}

	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
		logger := log.With().
			Str("handler", "HandleEventList").
			Logger()

		currentPage := c.Query("current_page")
		itemPerPage := c.Query("item_per_page")

		var currentPageInt, itemPerPageInt int
		var parseErr error

		if strings.TrimSpace(currentPage) != "" {
			currentPageInt, parseErr = strconv.Atoi(currentPage)
			if parseErr != nil {
				utils.ShowAppCommonErrorResponse(c, http.StatusBadRequest, parseErr.Error(), nil)
				return
			}
		}
		if strings.TrimSpace(itemPerPage) != "" {
			itemPerPageInt, parseErr = strconv.Atoi(itemPerPage)
			if parseErr != nil {
				utils.ShowAppCommonErrorResponse(c, http.StatusBadRequest, parseErr.Error(), nil)
				return
			}
		}

		if currentPageInt <= 0 {
			currentPageInt = 1
		}

		if itemPerPageInt <= 0 {
			itemPerPageInt = 10
		}

		data, pageInfo, dtErr := services.GetEventsServices(currentPageInt, itemPerPageInt, s)
		if dtErr != nil {
			logger.Error().Err(dtErr)
			utils.ShowAppCommonErrorResponse(c, http.StatusBadRequest, dtErr.Error(), nil)
			return
		}

		c.JSON(http.StatusOK, EventListResponse{
			Events:     data,
			Pagination: pageInfo,
		})

		return
	}
}
