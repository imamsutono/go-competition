// This file contains the implementation of the ServerInterface.
package handler

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kompit-recruitment/backend/generated/api"
	"github.com/kompit-recruitment/backend/initializers"
	"github.com/kompit-recruitment/backend/models"
	"github.com/kompit-recruitment/backend/utils"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	var req models.CompetitionPostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{
			Message: null.StringFrom(err.Error()).Ptr(),
		})
		return
	}

	newCompetition := models.Competition{
		Name:      req.Name,
		StartDate: req.StartDate,
	}

	initializers.DB.Create(&newCompetition)

	c.JSON(http.StatusCreated, models.CompetitionsPost201Response{
		Id:        newCompetition.ID,
		Name:      newCompetition.Name,
		StartDate: newCompetition.StartDate,
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

	c.JSON(http.StatusOK, models.CompetitionsIdGet200Response{
		Id:           &competition.ID,
		Name:         &competition.Name,
		StartDate:    &competition.StartDate,
		Participants: &[]models.CompetitionsIdGet200ResponseParticipantsItem{},
	})
}

// TODO: ASSIGNMENT 3
// Joins a competition with the given id.
// (POST /competitions/{id}/join)
func (h *Handler) CompetitionsIdJoinPost(c *gin.Context, id int64) {
	competitionID, _ := strconv.Atoi(c.Param("id"))
	userID := 1

	data := models.UserCompetition{
		CompetitionID: int64(competitionID),
		UserID:        int64(userID),
	}

	result := initializers.DB.Clauses(clause.Returning{}).Create(&data)
	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{
			Message: null.StringFrom(err.Error()).Ptr(),
		})
		return
	}

	c.JSON(http.StatusCreated, api.CompetitionsIdJoinPost201Response{
		RegistrationId: data.RegistrationID,
	})
}

// TODO: ASSIGNMENT 4
// Registers a new user.
// (POST /register)
func (h *Handler) RegisterPost(c *gin.Context) {
	var req api.RegisterPostJSONRequestBody

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{
			Message: null.StringFrom(err.Error()).Ptr(),
		})
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusBadGateway, api.ErrorResponse{
			Message: null.StringFrom(err.Error()).Ptr(),
		})
		return
	}

	newUser := models.User{
		Username: req.Username,
		Email:    string(req.Email),
		Password: hashedPassword,
	}

	result := initializers.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{
			Message: null.StringFrom(result.Error.Error()).Ptr(),
		})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse{
			Message: null.StringFrom(result.Error.Error()).Ptr(),
		})
		return
	}

	c.JSON(http.StatusCreated, api.RegisterPost201Response{
		JoinDate: newUser.CreatedAt,
	})
}

// TODO: ASSIGNMENT 5
// Returns a user with the given username.
// (GET /users/{username})
func (h *Handler) UsersUsernameGet(c *gin.Context, username string) {
	var user models.User

	data := initializers.DB.Preload("Competitions", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Name", "StartDate")
	}).First(&user, "username = ?", username)

	if err := data.Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, api.ErrorResponse{
				Message: null.StringFrom("User not found").Ptr(),
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, api.ErrorResponse{
				Message: null.StringFrom(err.Error()).Ptr(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"username":     user.Username,
		"join_date":    user.CreatedAt,
		"competitions": user.Competitions,
	})
}
