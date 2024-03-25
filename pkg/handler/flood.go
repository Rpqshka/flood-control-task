package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type checkInput struct {
	Id string `json:"id" binding:"required"`
}

func (h *Handler) check(c *gin.Context) {
	var input checkInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.ParseInt(input.Id, 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if id <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid userId value")
		return
	}

	n, k, err := h.services.Checker.GetParam()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	ctx := context.Background()

	deadline := time.Now().Add(1 * time.Second)

	ctx, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()

	currentTime := time.Now().Unix()

	ctx = context.WithValue(ctx, "time", currentTime)

	ctx = context.WithValue(ctx, "n", n)
	ctx = context.WithValue(ctx, "k", k)

	if err = h.services.Checker.SetParam(ctx); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	isReady, err := h.services.Check(ctx, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": isReady,
	})
}

type paramInput struct {
	N string `json:"n"`
	K string `json:"k"`
}

func (h *Handler) setParam(c *gin.Context) {
	var input paramInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	nInt, kInt, err := checkParam(input.N, input.K)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	ctx := context.Background()

	deadline := time.Now().Add(1 * time.Second)

	ctx, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()

	ctx = context.WithValue(ctx, "n", nInt)
	ctx = context.WithValue(ctx, "k", kInt)

	if err = h.services.Checker.SetParam(ctx); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": true,
	})
}
