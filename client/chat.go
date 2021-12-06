// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/jancimertel/go-tdlib/v2/tdlib"
)

// GetChat Returns information about a chat by its identifier, this is an offline request if the current user is not a bot
// @param chatID Chat identifier
func (client *Client) GetChat(chatID int64) (*tdlib.Chat, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":   "getChat",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chatDummy tdlib.Chat
	err = json.Unmarshal(result.Raw, &chatDummy)
	return &chatDummy, err

}

// SearchPublicChat Searches a public chat by its username. Currently only private chats, supergroups and channels can be public. Returns the chat if found; otherwise an error is returned
// @param username Username to be resolved
func (client *Client) SearchPublicChat(username string) (*tdlib.Chat, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":    "searchPublicChat",
		"username": username,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chat tdlib.Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// CreatePrivateChat Returns an existing chat corresponding to a given user
// @param userID User identifier
// @param force If true, the chat will be created without network request. In this case all information about the chat except its type, title and photo can be incorrect
func (client *Client) CreatePrivateChat(userID int64, force bool) (*tdlib.Chat, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":   "createPrivateChat",
		"user_id": userID,
		"force":   force,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chat tdlib.Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// CreateBasicGroupChat Returns an existing chat corresponding to a known basic group
// @param basicGroupID Basic group identifier
// @param force If true, the chat will be created without network request. In this case all information about the chat except its type, title and photo can be incorrect
func (client *Client) CreateBasicGroupChat(basicGroupID int64, force bool) (*tdlib.Chat, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":          "createBasicGroupChat",
		"basic_group_id": basicGroupID,
		"force":          force,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chat tdlib.Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// CreateSupergroupChat Returns an existing chat corresponding to a known supergroup or channel
// @param supergroupID Supergroup or channel identifier
// @param force If true, the chat will be created without network request. In this case all information about the chat except its type, title and photo can be incorrect
func (client *Client) CreateSupergroupChat(supergroupID int64, force bool) (*tdlib.Chat, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":         "createSupergroupChat",
		"supergroup_id": supergroupID,
		"force":         force,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chat tdlib.Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// CreateSecretChat Returns an existing chat corresponding to a known secret chat
// @param secretChatID Secret chat identifier
func (client *Client) CreateSecretChat(secretChatID int32) (*tdlib.Chat, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":          "createSecretChat",
		"secret_chat_id": secretChatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chatDummy tdlib.Chat
	err = json.Unmarshal(result.Raw, &chatDummy)
	return &chatDummy, err

}

// CreateNewBasicGroupChat Creates a new basic group and sends a corresponding messageBasicGroupChatCreate. Returns the newly created chat
// @param userIDs Identifiers of users to be added to the basic group
// @param title Title of the new basic group; 1-128 characters
func (client *Client) CreateNewBasicGroupChat(userIDs []int64, title string) (*tdlib.Chat, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":    "createNewBasicGroupChat",
		"user_ids": userIDs,
		"title":    title,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chat tdlib.Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// CreateNewSupergroupChat Creates a new supergroup or channel and sends a corresponding messageSupergroupChatCreate. Returns the newly created chat
// @param title Title of the new chat; 1-128 characters
// @param isChannel True, if a channel chat needs to be created
// @param description Chat description; 0-255 characters
// @param location Chat location if a location-based supergroup is being created
// @param forImport True, if the supergroup is created for importing messages using importMessage
func (client *Client) CreateNewSupergroupChat(title string, isChannel bool, description string, location *tdlib.ChatLocation, forImport bool) (*tdlib.Chat, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":       "createNewSupergroupChat",
		"title":       title,
		"is_channel":  isChannel,
		"description": description,
		"location":    location,
		"for_import":  forImport,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chat tdlib.Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// CreateNewSecretChat Creates a new secret chat. Returns the newly created chat
// @param userID Identifier of the target user
func (client *Client) CreateNewSecretChat(userID int64) (*tdlib.Chat, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":   "createNewSecretChat",
		"user_id": userID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chat tdlib.Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// UpgradeBasicGroupChatToSupergroupChat Creates a new supergroup from an existing basic group and sends a corresponding messageChatUpgradeTo and messageChatUpgradeFrom; requires creator privileges. Deactivates the original basic group
// @param chatID Identifier of the chat to upgrade
func (client *Client) UpgradeBasicGroupChatToSupergroupChat(chatID int64) (*tdlib.Chat, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":   "upgradeBasicGroupChatToSupergroupChat",
		"chat_id": chatID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chatDummy tdlib.Chat
	err = json.Unmarshal(result.Raw, &chatDummy)
	return &chatDummy, err

}

// JoinChatByInviteLink Uses an invite link to add the current user to the chat if possible
// @param inviteLink Invite link to use
func (client *Client) JoinChatByInviteLink(inviteLink string) (*tdlib.Chat, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":       "joinChatByInviteLink",
		"invite_link": inviteLink,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var chat tdlib.Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}
