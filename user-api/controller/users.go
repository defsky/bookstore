package controller

import (
	"net/http"
	"strconv"

	"github.com/defsky/bookstore/user-api/httputil"
	"github.com/gin-gonic/gin"
)

// GetUser return the user info
//  @Summary Get an user info
//  @Description get user info by ID
//  @Tags users
//  @Accept  json
//  @Produce  json
//  @Param id path int true "User ID"
//  @Success 200 {object} model.User
//  @Failure 400 {object} httputil.HTTPError
//  @Failure 404 {object} httputil.HTTPError
//  @Failure 500 {object} httputil.HTTPError
//  @Router /users/{id} [get]
func (c *Controller) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	u, err := c.getUserByID(aid)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, u)
}
