package handler

import (
	"Invalytics/app/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ListProfitability struct {
	Data []model.ProfitInfo `json:"data"`
}

func (h *Handler) GetAllSharesProfitability(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var term Term
	term.AmountOfMonths, err = strconv.Atoi(c.Query("term"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid term param")
		return
	}

	var needSort bool
	needSort, err = strconv.ParseBool(c.Query("sort"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid sort param")
		return
	}

	list, err := h.services.Profit.AllShareProfitability(userId, term.AmountOfMonths, needSort)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": list,
	})
}

type Term struct {
	AmountOfMonths int `json:"term"`
}

func (h *Handler) GetShareProfitabilityById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	shareId, err := getId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var term Term
	term.AmountOfMonths, err = strconv.Atoi(c.Query("term"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid term param")
		return
	}

	result, err := h.services.Profit.ShareProfitabilityById(userId, shareId, term.AmountOfMonths)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ticker":                result.Ticker,
		"termInMonths":          result.AmountOfMonths,
		"annualPercentageYield": result.Profit,
	})
}

func (h *Handler) GetBondProfitabilityById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	bondId, err := getId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	result, err := h.services.Profit.BondProfitabilityById(userId, bondId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ticker":                result.Ticker,
		"termInMonths":          result.AmountOfMonths,
		"annualPercentageYield": result.Profit,
	})
}

func (h *Handler) GetAllBondsProfitability(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var needSort bool
	needSort, err = strconv.ParseBool(c.Query("sort"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid sort param")
		return
	}

	list, err := h.services.Profit.AllBondProfitability(userId, needSort)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": list,
	})
}

func (h *Handler) GetDepositProfitabilityById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	depositId, err := getId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	result, err := h.services.Profit.DepositProfitabilityById(userId, depositId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"percentageRate:": result.PercentageRate,
		"termInMonths":    result.AmountOfMonths,
		"profit":          result.Profit,
	})
}

func (h *Handler) GetAllDepositsProfitability(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var needSort bool
	needSort, err = strconv.ParseBool(c.Query("sort"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid sort param")
		return
	}

	list, err := h.services.Profit.AllDepositProfitability(userId, needSort)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": list,
	})
}
