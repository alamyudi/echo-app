package dbmanager

import (
	"errors"
	"fmt"
	"strings"
)

type (

	// Limit for row limit
	Limit struct {
		Limit  int
		Offset int
	}

	// Order for ordering
	Order struct {
		Field string
		IsAsc bool
	}

	// Filter for filtering
	Filter struct {
		Field     string
		Condition string
		Value     string
	}
)

// GenerateOrder for ordering
func GenerateOrder(orders []Order) []string {
	queries := make([]string, len(orders))
	for i, order := range orders {
		cond := "asc"
		if !order.IsAsc {
			cond = "desc"
		}
		query := fmt.Sprintf("%s %s", order.Field, cond)
		queries[i] = query
	}
	return queries
}

// GenerateQuery to generate query from filters
func GenerateQuery(filters []*Filter) string {
	queries := make([]string, len(filters))
	for i, filter := range filters {
		rel := "or"
		if i == len(filters)-1 {
			rel = ""
		}

		query := fmt.Sprintf("%s %s ? %s", filter.Field, filter.Condition, rel)
		queries[i] = query
	}
	return strings.Join(queries, " ")
}

// GenerateParams for generate list of value
func GenerateParams(filters []*Filter) []interface{} {
	queries := make([]interface{}, len(filters))
	for i, filter := range filters {
		queries[i] = filter.Value
	}
	return queries
}

// FilterFieldIsValid to check filter file
func FilterFieldIsValid(filters []*Filter, keys []string) error {
	msg := ""
	isOK := true
	for _, filter := range filters {
		isFound := false
		for _, key := range keys {
			if filter.Field == key {
				isFound = true
				break
			}
		}

		if !isFound {
			isOK = false
			msg = fmt.Sprintf("Field %s not found for filtering", filter.Field)
			break
		}

	}

	if !isOK {
		return errors.New(msg)
	}

	return nil
}

// OrderFieldIsValid to check order field
func OrderFieldIsValid(orders []Order, keys []string) error {
	msg := ""
	isOK := true
	for _, order := range orders {
		isFound := false
		for _, key := range keys {
			if order.Field == key {
				isFound = true
				break
			}
		}

		if !isFound {
			isOK = false
			msg = fmt.Sprintf("Field %s not found for ordering", order.Field)
			break
		}

	}

	if !isOK {
		return errors.New(msg)
	}

	return nil
}

// CastUserModelToManager to cast user model to user manager
// func CastUserModelToManager(usr models.User) User {
// 	return User{
// 		AirlineID:   usr.AirlineID,
// 		AirlineName: usr.AirlineName,
// 		UserID:      usr.UserID,
// 		Email:       usr.Email,
// 		Password:    usr.Password,
// 		MacAddress:  usr.MacAddress,
// 		ValidTo:     usr.ValidTo,
// 		AppVersion:  usr.AppVersion,
// 		API:         usr.API,
// 		FolderPath:  usr.FolderPath,
// 	}
// }

// // CastUserContentModelToManager to cast user content model to user manager
// func CastUserContentModelToManager(usr models.UserContent) UserContent {
// 	return UserContent{
// 		FolderPath: usr.FolderPath,
// 		RevNum:     usr.RevNum,
// 		Type:       usr.Type,
// 	}
// }
