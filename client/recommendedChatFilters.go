// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/jancimertel/go-tdlib/v2/tdlib"
)

// GetRecommendedChatFilters Returns recommended chat filters for the current user
func (client *Client) GetRecommendedChatFilters() (*tdlib.RecommendedChatFilters, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getRecommendedChatFilters",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var recommendedChatFilters tdlib.RecommendedChatFilters
	err = json.Unmarshal(result.Raw, &recommendedChatFilters)
	return &recommendedChatFilters, err

}
