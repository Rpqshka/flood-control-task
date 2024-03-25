package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type paramInput struct {
	Id string `json:"id" binding:"required"`
	N  string `json:"n"`
	K  string `json:"k"`
}

func (h *Handler) check(c *gin.Context) {
	var input paramInput

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

	n, k, err := checkParam(input.N, input.K)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if n == 0 {
		n = int64(1) // Количество вызовов метода Check
	}

	if k == 0 {
		k = time.Now().Add(-5 * time.Second).Unix() // Время отчета
	}

	ctx := context.Background()

	deadline := time.Now().Add(1 * time.Second)

	ctx, cancel := context.WithDeadline(ctx, deadline)
	defer cancel()

	currentTime := time.Now().Unix()
	ctx = context.WithValue(ctx, "time", currentTime)

	ctx = context.WithValue(ctx, "n", n)
	ctx = context.WithValue(ctx, "k", k)

	isReady, err := h.services.Check(ctx, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": isReady,
	})
}
