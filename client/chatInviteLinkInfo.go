// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/jancimertel/go-tdlib/v2/tdlib"
)

// CheckChatInviteLink Checks the validity of an invite link for a chat and returns information about the corresponding chat
// @param inviteLink Invite link to be checked
func (client *Client) CheckChatInviteLink(inviteLink string) (*tdlib.ChatInviteLinkInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":       "checkChatInviteLink",
		"invite_link": inviteLink,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chatInviteLinkInfo tdlib.ChatInviteLinkInfo
	err = json.Unmarshal(result.Raw, &chatInviteLinkInfo)
	return &chatInviteLinkInfo, err

}
