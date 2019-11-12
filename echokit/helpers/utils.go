package helpers

import "fmt"

// GetTagsQueryFromList serializer to get string query
func GetTagsQueryFromList(tags []string) string {
	strQuery := ""
	for i := range tags {
		if i == 0 {
			strQuery += "content_tags like ?"
		} else {
			strQuery += " OR content_tags like ?"
		}
	}
	return strQuery
}

// GetTagsInterface serializer to get interface
func GetTagsInterface(tags []string) []interface{} {
	var stuff []interface{}
	for _, val := range tags {
		valFormat := fmt.Sprintf("%%%s%%", val)
		stuff = append(stuff, valFormat)
	}
	return stuff
}
