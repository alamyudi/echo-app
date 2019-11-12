package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/alamyudi/echo-app/echokit/dbmanager"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type (

	// AddContentPayload add content payload
	AddContentPayload struct {
		ContentName  string `validate:"required" json:"content_name"`
		ContentDesc  string `validate:"required" json:"content_desc"`
		ContentTags  string `validate:"required" json:"content_tags"`
		ContentImage string `validate:"required" json:"content_image"`
	}

	// FilterContentPayload to filter content list
	FilterContentPayload struct {
		Page       int      `validate:"required" json:"page" query:"page"`
		RowPerPage int      `validate:"required" json:"row_per_page" query:"row_per_page"`
		Tag        []string `validate:"required" json:"tag" query:"tag"`
	}
)

/******************* Content ********************/

// GetContents to getting Content
func GetContents(ctx echo.Context) error {
	filterContent := new(FilterContentPayload)
	if err := ctx.Bind(filterContent); err != nil {
		response := MessageResponse{
			Title:   "Failed getting contents",
			Message: "Filter param is not valid",
		}
		return ctx.JSON(http.StatusBadGateway, response)
	}

	// validate interface
	if err := ctx.Validate(filterContent); err != nil {
		response := MessageResponse{
			Title:   "Failed",
			Message: err.Error(),
		}
		return ctx.JSON(http.StatusBadGateway, response)
	}

	offset := (filterContent.Page - 1) * filterContent.RowPerPage

	count, err := iKit.DBManager.GetCountContent(filterContent.Tag)
	if err != nil {
		response := MessageResponse{
			Title:   "Failed getting contents 1",
			Message: "Query error",
		}
		return ctx.JSON(http.StatusBadGateway, response)
	}
	logrus.Info(count)

	contents, err := iKit.DBManager.GetContents(filterContent.Tag, filterContent.RowPerPage, offset)
	if err != nil {
		response := MessageResponse{
			Title:   "Failed getting contents",
			Message: "Query error",
		}
		return ctx.JSON(http.StatusBadGateway, response)
	}
	payload := make(map[string]interface{})
	payload["total_page"] = (count / filterContent.RowPerPage) + 1
	payload["contents"] = contents

	response := MessageWithPayloadResponse{
		Title:   "Success",
		Message: "Success fetch contents",
		Payload: payload,
	}

	return ctx.JSON(http.StatusOK, response)
}

// GetContentByID to getting content
func GetContentByID(ctx echo.Context) error {
	id := ctx.Param("id")
	contentID, err := strconv.Atoi(id)
	if err != nil {
		response := MessageResponse{
			Title:   "Failed",
			Message: "Cannot convert param to number",
		}
		return ctx.JSON(http.StatusBadGateway, response)
	}
	content, err := iKit.DBManager.GetContentByID(int16(contentID))
	if err != nil {
		response := MessageResponse{
			Title:   "Failed",
			Message: "Query error",
		}
		return ctx.JSON(http.StatusBadGateway, response)
	}

	response := MessageWithPayloadResponse{
		Title:   "Success",
		Message: "Success fetch content",
		Payload: content,
	}

	return ctx.JSON(http.StatusOK, response)
}

// DeleteContentByID to getting content
func DeleteContentByID(ctx echo.Context) error {
	id := ctx.Param("id")
	contentID, err := strconv.Atoi(id)
	if err != nil {
		response := MessageResponse{
			Title:   "Failed",
			Message: "Cannot convert param to number",
		}
		return ctx.JSON(http.StatusBadGateway, response)
	}

	err = iKit.DBManager.DeleteContentByID(int16(contentID))
	if err != nil {
		response := MessageResponse{
			Title:   "Failed",
			Message: "Query error",
		}
		return ctx.JSON(http.StatusBadGateway, response)
	}

	message := fmt.Sprintf("Success deleted content with id %d", contentID)
	response := MessageResponse{
		Title:   "Success",
		Message: message,
	}

	return ctx.JSON(http.StatusOK, response)
}

// PutContentByID to getting content
func PutContentByID(ctx echo.Context) error {
	id := ctx.Param("id")
	contentID, err := strconv.Atoi(id)
	if err != nil {
		response := MessageResponse{
			Title:   "Failed",
			Message: "Cannot convert param to number",
		}
		return ctx.JSON(http.StatusBadGateway, response)
	}

	payload := new(AddContentPayload)

	// serialize json body
	if err := ctx.Bind(payload); err != nil {
		response := MessageResponse{
			Title:   "Failed",
			Message: "Failed to serialize product",
		}
		return ctx.JSON(http.StatusBadGateway, response)
	}

	// validate interface
	if err := ctx.Validate(payload); err != nil {
		response := MessageResponse{
			Title:   "Failed",
			Message: err.Error(),
		}
		return ctx.JSON(http.StatusBadGateway, response)
	}

	content := dbmanager.ContentManager{
		ContentName:  payload.ContentName,
		ContentDesc:  payload.ContentDesc,
		ContentTags:  payload.ContentTags,
		ContentImage: payload.ContentImage,
		UpdatedAt:    time.Now(),
	}

	_, err = iKit.DBManager.UpdateContentByID(int16(contentID), content)
	if err != nil {
		logrus.Info(err.Error())
		response := MessageResponse{
			Title:   "Failed updating content",
			Message: "Query error",
		}
		return ctx.JSON(http.StatusBadGateway, response)
	}

	response := MessageResponse{
		Title:   "Success",
		Message: "Success updating content",
	}

	return ctx.JSON(http.StatusOK, response)
}

// PostContent to getting content
func PostContent(ctx echo.Context) error {
	payload := new(AddContentPayload)

	// serialize json body
	if err := ctx.Bind(payload); err != nil {
		response := MessageResponse{
			Title:   "Failed",
			Message: "Failed to serialize product",
		}
		return ctx.JSON(http.StatusBadGateway, response)
	}

	// validate interface
	if err := ctx.Validate(payload); err != nil {
		response := MessageResponse{
			Title:   "Failed",
			Message: err.Error(),
		}
		return ctx.JSON(http.StatusBadGateway, response)
	}

	content := dbmanager.ContentManager{
		ContentName:  payload.ContentName,
		ContentDesc:  payload.ContentDesc,
		ContentTags:  payload.ContentTags,
		ContentImage: payload.ContentImage,
	}

	_, err := iKit.DBManager.InsertContent(content)
	if err != nil {
		response := MessageResponse{
			Title:   "Failed insert content",
			Message: "Query error",
		}
		return ctx.JSON(http.StatusBadGateway, response)
	}

	response := MessageResponse{
		Title:   "Success",
		Message: "Success insert content",
	}

	return ctx.JSON(http.StatusOK, response)
}
