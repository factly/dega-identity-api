package keto

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/factly/kavach-server/config"
	"github.com/factly/kavach-server/model"
)

// UpdatePolicy PUT request to keto server to update the policy
func UpdatePolicy(uri string, body *model.Policy) error {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(&body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", config.KetoURL+uri, buf)

	if err != nil {
		return err
	}

	client := &http.Client{}
	_, err = client.Do(req)

	if err != nil {
		return err
	}
	return nil
}

// DeletePolicy DELETE request to keto server to delete policy
func DeletePolicy(uri string) error {
	req, err := http.NewRequest("DELETE", config.KetoURL+uri, nil)

	if err != nil {
		return err
	}

	client := &http.Client{}
	_, err = client.Do(req)

	if err != nil {
		return err
	}
	return nil
}