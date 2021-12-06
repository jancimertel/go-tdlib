// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/jancimertel/go-tdlib/v2/tdlib"
)

// GetTextEntities Returns all entities (mentions, hashtags, cashtags, bot commands, bank card numbers, URLs, and email addresses) contained in the text. Can be called synchronously
// @param text The text in which to look for entites
func (client *Client) GetTextEntities(text string) (*tdlib.TextEntities, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getTextEntities",
		"text":  text,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var textEntities tdlib.TextEntities
	err = json.Unmarshal(result.Raw, &textEntities)
	return &textEntities, err

}
