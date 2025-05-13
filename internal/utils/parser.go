package utils

import (
	"encoding/json"

	"gorm.io/datatypes"
)

// JSONToMap convert datatypes.JSON to map[string]string
func JSONToMap(jsonData datatypes.JSON) (map[string]string, error) {
	var result map[string]string
	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MapToJSON convert map[string]string to datatypes.JSON
func MapToJSON(data map[string]string) (datatypes.JSON, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
