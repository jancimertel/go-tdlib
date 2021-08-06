// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/Arman92/go-tdlib/tdlib"
)

// GenerateChatInviteLink Generates a new invite link for a chat; the previously generated link is revoked. Available for basic groups, supergroups, and channels. Requires administrator privileges and can_invite_users right
// @param chatID Chat identifier
func (client *Client) GenerateChatInviteLink(chatID int64) (*tdlib.ChatInviteLink, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":   "generateChatInviteLink",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatInviteLink tdlib.ChatInviteLink
	err = json.Unmarshal(result.Raw, &chatInviteLink)
	return &chatInviteLink, err

}
