package controller

import (
	"net/http"
	"strconv"

	"github.com/defsky/bookstore/user-api/httputil"
	"github.com/defsky/bookstore/user-api/model"
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

// GetUserList returns user list
//  @Summary Get a list of users
//  @Description get user info list by specified page info
//  @Tags users
//  @Accept  json
//  @Produce  json
//  @Param pageindex query int false "Page index"
//  @Param pagesize query int false "Page size"
//  @Success 200 {object} model.UserList
//  @Failure 400 {object} httputil.HTTPError
//  @Failure 404 {object} httputil.HTTPError
//  @Failure 500 {object} httputil.HTTPError
//  @Router /users [get]
func (c *Controller) GetUserList(ctx *gin.Context) {
	i := ctx.Query("pageindex")
	s := ctx.Query("pagesize")

	if len(i) <= 0 {
		i = "0"
	}
	if len(s) <= 0 {
		s = "0"
	}
	pageIndex, err := strconv.Atoi(i)
	pageSize, err := strconv.Atoi(s)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	list, err := c.getUserList(pageIndex, pageSize)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, list)
}

// AddUser create a new user
//  @Summary Create user
//  @Description Create a new user
//  @Tags users
//  @Accept  json
//  @Produce  json
//  @Param user body model.User true "Add user"
//  @Success 200 {object} model.User
//  @Failure 400 {object} httputil.HTTPError
//  @Failure 404 {object} httputil.HTTPError
//  @Failure 500 {object} httputil.HTTPError
//  @Router /users [post]
func (c *Controller) AddUser(ctx *gin.Context) {
	var newUser model.User
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := newUser.Validate(); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	u, err := c.createUser(model.User{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
	})
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, u)
}
