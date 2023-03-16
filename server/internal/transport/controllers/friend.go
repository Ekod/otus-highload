package controllers

import (
	"net/http"
	"strconv"

	"github.com/Ekod/otus-highload/internal/dto"
	"github.com/Ekod/otus-highload/internal/transport/services"
	"github.com/Ekod/otus-highload/utils/errors"
	"github.com/Ekod/otus-highload/utils/security"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FriendHandlers struct {
	Service *services.Services
	Logger  *zap.SugaredLogger
}

func (h *FriendHandlers) RemoveFriend(c *gin.Context) {
	ctx := c.Request.Context()

	userId, err := security.GetUserIdFromToken(c)
	if err != nil {
		err := errors.ParseError(err)

		h.Logger.Error(err.DebugMessage)

		c.JSON(err.Status, err)

		return
	}

	friendId := c.Param("id")

	parsedFriendId, err := strconv.Atoi(friendId)
	if err != nil {
		h.Logger.Error("[ERROR] Controllers_RemoveFriend - Error parsing incoming JSON")

		err := errors.NewHandlerBadRequestError("Invalid params")
		c.JSON(err.Status, err)

		return
	}

	err = h.Service.FriendService.RemoveFriend(ctx, userId, parsedFriendId)

	c.Status(http.StatusOK)
}

func (h *FriendHandlers) GetFriends(c *gin.Context) {
	ctx := c.Request.Context()

	userId, err := security.GetUserIdFromToken(c)
	if err != nil {
		h.Logger.Error(err)

		err := errors.ParseError(err)

		c.JSON(err.Status, err)
		return
	}

	response, err := h.Service.FriendService.GetFriends(ctx, userId)
	if err != nil {
		err := errors.ParseError(err)

		h.Logger.Error(err.DebugMessage)

		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *FriendHandlers) MakeFriends(c *gin.Context) {
	ctx := c.Request.Context()

	var friendRequest dto.FriendRequest
	if err := c.ShouldBindJSON(&friendRequest); err != nil {
		h.Logger.Errorf("[ERROR] Controllers_RegisterUser - Error parsing incoming JSON with %d and err %s", friendRequest, err)

		err := errors.NewHandlerBadRequestError("Invalid json request")
		c.JSON(err.Status, err)

		return
	}

	userId, err := security.GetUserIdFromToken(c)
	if err != nil {
		err := errors.ParseError(err)

		h.Logger.Error(err.DebugMessage)

		c.JSON(err.Status, err)
		return
	}

	response, err := h.Service.FriendService.MakeFriends(ctx, userId, friendRequest.ID)
	if err != nil {
		err := errors.ParseError(err)

		h.Logger.Error(err.DebugMessage)

		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
