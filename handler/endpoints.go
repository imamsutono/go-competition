// This file contains the implementation of the ServerInterface.
package handler

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/kompit-recruitment/backend/generated/api"
	"github.com/kompit-recruitment/backend/initializers"
	"github.com/kompit-recruitment/backend/models"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

// Sample endpoint.
// (POST /ping)
func (h *Handler) PingPost(c *gin.Context) {
	var req api.PingPostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{
			Message: null.StringFrom(err.Error()).Ptr(),
		})
		return
	}

	if req.Ping == nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{
			Message: null.StringFrom("message is required").Ptr(),
		})
		return
	}

	// The message should start with 'pong'
	var regex = regexp.MustCompile(`^pong`)
	if !regex.MatchString(*req.Ping) {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{
			Message: null.StringFrom("message should start with 'pong'").Ptr(),
		})
		return
	}

	c.JSON(http.StatusOK, api.PingPost200Response{
		Pong: req.Ping,
	})
}

// TODO: ASSIGNMENT 1
// Creates competition.
// (POST /competitions)
func (h *Handler) CompetitionsPost(c *gin.Context) {
	// gin.BasicAuth(gin.Accounts{"admin": "admin"})

	var req api.CompetitionsPostJSONRequestBody

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{
			Message: null.StringFrom(err.Error()).Ptr(),
		})
		return
	}

	newCompetition := models.Competition{
		Name:      *req.Name,
		StartDate: *req.StartDate,
	}

	initializers.DB.Create(&newCompetition)

	c.JSON(http.StatusCreated, api.CompetitionsPost201Response{
		Id:        &newCompetition.ID,
		Name:      &newCompetition.Name,
		StartDate: &newCompetition.StartDate,
	})
}

// TODO: ASSIGNMENT 2
// Returns a competition with the given id.
// (GET /competitions/{id})
func (h *Handler) CompetitionsIdGet(c *gin.Context, id int64) {
	var competition models.Competition

	if err := initializers.DB.Preload("Users").First(&competition, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, api.CompetitionsIdGet200Response{
		Id:           &competition.ID,
		Name:         &competition.Name,
		StartDate:    &competition.StartDate,
		Participants: &[]api.CompetitionsIdGet200ResponseParticipantsItem{},
	})
}

// TODO: ASSIGNMENT 3
// Joins a competition with the given id.
// (POST /competitions/{id}/join)
func (h *Handler) CompetitionsIdJoinPost(c *gin.Context, id int64) {
	c.JSON(http.StatusNotImplemented, "TODO: Implement me")
}

// TODO: ASSIGNMENT 4
// Registers a new user.
// (POST /register)
func (h *Handler) RegisterPost(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "TODO: Implement me")
}

// TODO: ASSIGNMENT 5
// Returns a user with the given username.
// (GET /users/{username})
func (h *Handler) UsersUsernameGet(c *gin.Context, username string) {
	c.JSON(http.StatusNotImplemented, "TODO: Implement me")
}
