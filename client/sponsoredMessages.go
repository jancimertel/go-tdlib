// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/jancimertel/go-tdlib/v2/tdlib"
)

// GetChatSponsoredMessages Returns sponsored messages to be shown in a chat; for channel chats only
// @param chatID Identifier of the chat
func (client *Client) GetChatSponsoredMessages(chatID int64) (*tdlib.SponsoredMessages, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":   "getChatSponsoredMessages",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var sponsoredMessages tdlib.SponsoredMessages
	err = json.Unmarshal(result.Raw, &sponsoredMessages)
	return &sponsoredMessages, err

}
