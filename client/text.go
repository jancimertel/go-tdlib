// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/Arman92/go-tdlib/tdlib"
)

// GetMessageEmbeddingCode Returns an HTML code for embedding the message. Available only for messages in supergroups and channels with a username
// @param chatID Identifier of the chat to which the message belongs
// @param messageID Identifier of the message
// @param forAlbum Pass true to return an HTML code for embedding of the whole media album
func (client *Client) GetMessageEmbeddingCode(chatID int64, messageID int64, forAlbum bool) (*tdlib.Text, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":      "getMessageEmbeddingCode",
		"chat_id":    chatID,
		"message_id": messageID,
		"for_album":  forAlbum,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text tdlib.Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// GetFileMimeType Returns the MIME type of a file, guessed by its extension. Returns an empty string on failure. Can be called synchronously
// @param fileName The name of the file or path to the file
func (client *Client) GetFileMimeType(fileName string) (*tdlib.Text, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":     "getFileMimeType",
		"file_name": fileName,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text tdlib.Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// GetFileExtension Returns the extension of a file, guessed by its MIME type. Returns an empty string on failure. Can be called synchronously
// @param mimeType The MIME type of the file
func (client *Client) GetFileExtension(mimeType string) (*tdlib.Text, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":     "getFileExtension",
		"mime_type": mimeType,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text tdlib.Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// CleanFileName Removes potentially dangerous characters from the name of a file. The encoding of the file name is supposed to be UTF-8. Returns an empty string on failure. Can be called synchronously
// @param fileName File name or path to the file
func (client *Client) CleanFileName(fileName string) (*tdlib.Text, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":     "cleanFileName",
		"file_name": fileName,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text tdlib.Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// GetJsonString Converts a JsonValue object to corresponding JSON-serialized string. Can be called synchronously
// @param jsonStringValue The JsonValue object
func (client *Client) GetJsonString(jsonStringValue tdlib.JsonValue) (*tdlib.Text, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":      "getJsonString",
		"json_value": jsonStringValue,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text tdlib.Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// GetChatFilterDefaultIconName Returns default icon name for a filter. Can be called synchronously
// @param filter Chat filter
func (client *Client) GetChatFilterDefaultIconName(filter *tdlib.ChatFilter) (*tdlib.Text, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":  "getChatFilterDefaultIconName",
		"filter": filter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text tdlib.Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// GetPreferredCountryLanguage Returns an IETF language tag of the language preferred in the country, which should be used to fill native fields in Telegram Passport personal details. Returns a 404 error if unknown
// @param countryCode A two-letter ISO 3166-1 alpha-2 country code
func (client *Client) GetPreferredCountryLanguage(countryCode string) (*tdlib.Text, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":        "getPreferredCountryLanguage",
		"country_code": countryCode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text tdlib.Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// GetCountryCode Uses current user IP address to find their country. Returns two-letter ISO 3166-1 alpha-2 country code. Can be called before authorization
func (client *Client) GetCountryCode() (*tdlib.Text, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getCountryCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text tdlib.Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// GetInviteText Returns the default text for invitation messages to be used as a placeholder when the current user invites friends to Telegram
func (client *Client) GetInviteText() (*tdlib.Text, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getInviteText",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text tdlib.Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// GetProxyLink Returns an HTTPS link, which can be used to add a proxy. Available only for SOCKS5 and MTProto proxies. Can be called before authorization
// @param proxyID Proxy identifier
func (client *Client) GetProxyLink(proxyID int32) (*tdlib.Text, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":    "getProxyLink",
		"proxy_id": proxyID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text tdlib.Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}
