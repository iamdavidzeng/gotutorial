package utils

import "encoding/json"

// MarshalIndent parse map[string]string response from Stripe
func MarshalIndent(v interface{}) (string, error) {

	json, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}

	return string(json), nil
}
