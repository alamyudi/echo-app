package dbmanager

import (
	"time"

	"github.com/alamyudi/echo-app/echokit/helpers"

	"github.com/alamyudi/echo-app/echokit/models"
)

type (
	// ContentManager model
	ContentManager struct {
		ContentID    int16     `json:"content_id"`
		ContentName  string    `json:"content_name"`
		ContentDesc  string    `json:"content_desc"`
		ContentTags  string    `json:"content_tags"`
		ContentImage string    `json:"content_image"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
)

/********************* Content ***********************/

// GetContents to getting contents
func (m *DBManager) GetContents(tags []string, limit int, offset int) ([]models.Content, error) {
	tagQR := helpers.GetTagsQueryFromList(tags)
	tagInterface := helpers.GetTagsInterface(tags)
	tagInterface = append(tagInterface, limit)
	tagInterface = append(tagInterface, offset)
	contents, err := m.MDL.GetContent(tagInterface, tagQR)
	if err != nil {
		return []models.Content{}, err
	}
	return contents, nil
}

// GetCountContent to getting content count
func (m *DBManager) GetCountContent(tags []string) (int, error) {
	tagQR := helpers.GetTagsQueryFromList(tags)
	tagInterface := helpers.GetTagsInterface(tags)
	count, err := m.MDL.GetCountContent(tagInterface, tagQR)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetContentByID to getting content by id
func (m *DBManager) GetContentByID(id int16) (models.Content, error) {
	content, err := m.MDL.GetContentByID(id)
	if err != nil {
		return models.Content{}, err
	}
	return content, nil
}

// DeleteContentByID to getting content by id
func (m *DBManager) DeleteContentByID(id int16) error {
	err := m.MDL.DeleteContentByID(id)
	return err
}

// UpdateContentByID to update content by id
func (m *DBManager) UpdateContentByID(id int16, content ContentManager) (int64, error) {
	contentModel := models.Content{
		ContentName:  content.ContentName,
		ContentDesc:  content.ContentDesc,
		ContentTags:  content.ContentTags,
		ContentImage: content.ContentImage,
		UpdatedAt:    content.UpdatedAt,
	}
	row, err := m.MDL.UpdateContentByID(id, contentModel)
	return row, err
}

// InsertContent to add content
func (m *DBManager) InsertContent(content ContentManager) (int64, error) {
	contentModel := models.Content{
		ContentName:  content.ContentName,
		ContentDesc:  content.ContentDesc,
		ContentTags:  content.ContentTags,
		ContentImage: content.ContentImage,
	}
	row, err := m.MDL.InsertContent(contentModel)
	return row, err
}
