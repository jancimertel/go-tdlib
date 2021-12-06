// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/jancimertel/go-tdlib/v2/tdlib"
)

// GetSavedOrderInfo Returns saved order info, if any
func (client *Client) GetSavedOrderInfo() (*tdlib.OrderInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getSavedOrderInfo",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var orderInfo tdlib.OrderInfo
	err = json.Unmarshal(result.Raw, &orderInfo)
	return &orderInfo, err

}
