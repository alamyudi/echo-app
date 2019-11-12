package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type (
	// Content model
	Content struct {
		ContentID    int16     `json:"content_id" sql:""`
		ContentName  string    `json:"content_name" sql:""`
		ContentDesc  string    `json:"content_desc" sql:""`
		ContentTags  string    `json:"content_tags" sql:""`
		ContentImage string    `json:"content_image" sql:""`
		CreatedAt    time.Time `json:"created_at" sql:"default:now()"`
		UpdatedAt    time.Time `json:"updated_at" sql:"default:now()"`
	}
)

// GetContent to getting content
func (m *MDL) GetContent(tags []interface{}, tagQs string) ([]Content, error) {
	db := m.MysqlClient

	tagQuery := "1=1"
	if len(tags) > 0 {
		tagQuery = tagQs
	}

	qs := "SELECT * FROM content WHERE (" + tagQuery + ") LIMIT ? OFFSET ?;"

	logrus.Infof("Filter", tags...)

	rows, err := db.Query(qs, tags...)
	defer rows.Close()
	if err != nil {
		logrus.Infof("ERR", err.Error())
		return []Content{}, err
	}
	rowIndex := 0
	var contents []Content
	for rows.Next() {
		var content Content
		err = rows.Scan(
			&content.ContentID,
			&content.ContentName,
			&content.ContentDesc,
			&content.ContentTags,
			&content.ContentImage,
			&content.CreatedAt,
			&content.UpdatedAt,
		)
		if err != nil {
			return []Content{}, err
		}
		contents = append(contents, content)
		rowIndex++
	}

	return contents, nil
}

// GetCountContent to getting content count
func (m *MDL) GetCountContent(tags []interface{}, tagQs string) (int, error) {
	db := m.MysqlClient

	tagQuery := "1=1"
	if len(tags) > 0 {
		tagQuery = tagQs
	}

	qs := "SELECT count(*) as count FROM content WHERE (" + tagQuery + ");"
	logrus.Infof("filter", qs, tags)
	rows, err := db.Query(qs, tags...)
	defer rows.Close()
	if err != nil {
		logrus.Infof("ERR", err.Error())
		return 0, err
	}
	var count int

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, err
		}
	}

	return count, nil
}

// GetContentByID to getting user content path by email and password
func (m *MDL) GetContentByID(id int16) (Content, error) {
	db := m.MysqlClient
	// Execute the query
	qs := `
	SELECT * FROM content WHERE content_id = ?
	`
	rows, err := db.Query(qs, id)
	if err != nil {
		return Content{}, err
	}
	rowIndex := 0
	var content Content
	for rows.Next() {
		err = rows.Scan(
			&content.ContentID,
			&content.ContentName,
			&content.ContentDesc,
			&content.ContentTags,
			&content.ContentImage,
			&content.CreatedAt,
			&content.UpdatedAt,
		)
		if err != nil {
			return Content{}, err
		}
		rowIndex++
	}

	if (Content{}) == content {
		msg := fmt.Sprintf("content with id %d not found", id)
		return Content{}, errors.New(msg)
	}

	return content, nil
}

// DeleteContentByID to delete
func (m *MDL) DeleteContentByID(id int16) error {
	db := m.MysqlClient

	qs := `
	DELETE FROM content WHERE content_id = ?`

	_, err := db.Exec(qs, id)
	if err != nil {
		return err
	}

	return nil
}

// InsertContent to insert content
func (m *MDL) InsertContent(content Content) (int64, error) {
	db := m.MysqlClient

	qs := `
	INSERT INTO content ( content_name, content_desc, content_tags, content_image) 
	VALUES (?, ?, ?, ?);`

	result, err := db.Exec(qs, content.ContentName, content.ContentDesc, content.ContentTags, content.ContentImage)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// UpdateContentByID to insert content
func (m *MDL) UpdateContentByID(id int16, content Content) (int64, error) {
	db := m.MysqlClient

	qs := `
	UPDATE content SET content_name = ?, content_desc = ?, content_tags = ?, content_image = ? , updated_at = ? 
	WHERE content_id = ?;`

	result, err := db.Exec(qs, content.ContentName, content.ContentDesc, content.ContentTags, content.ContentImage, content.UpdatedAt, id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
