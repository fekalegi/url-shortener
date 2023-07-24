package shortener

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"url-shortener/common"
	"url-shortener/config/validator"
	"url-shortener/delivery/http/shortener/model"
	"url-shortener/domain/shortener"
)

func (c *controller) CreateShortenedURL(ctx *gin.Context) {
	bodyRequest := new(model.LinkRequest)
	if err := ctx.BindJSON(bodyRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse(err.Error()))
		return
	}

	if err := validator.ValidateStruct(bodyRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, common.BadRequestResponse(err))
		return
	}

	_, err := url.ParseRequestURI(bodyRequest.URL)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.BadRequestResponse("invalid URI for request url"))
		return
	}

	l := new(shortener.Link)
	mapRequestToLink(bodyRequest, l)

	if err := c.shortenerService.CreateShortenedURL(l); err != nil {
		ctx.JSON(http.StatusInternalServerError, common.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, common.SuccessResponseWithData(l, "success"))
	return
}

func (c *controller) Get(ctx *gin.Context) {
	paramUrl := ctx.Param("url")

	data, err := c.shortenerService.GetByShortenedURL(paramUrl)
	if err != nil && errors.Is(err, common.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, common.ErrorResponse(err.Error()))
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, common.SuccessResponseWithData(data, "success"))
	return
}

func (c *controller) GetSortedURLs(ctx *gin.Context) {
	req := new(model.SortRequest)
	if err := ctx.BindQuery(req); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse(err.Error()))
		return
	}

	if err := validator.ValidateStruct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, common.BadRequestResponse(err))
		return
	}

	data, err := c.shortenerService.GetAll(req.SortBy)
	if err != nil && errors.Is(err, common.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, common.ErrorResponse(err.Error()))
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, common.SuccessResponseWithData(data, "success"))
	return
}
