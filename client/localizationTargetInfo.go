// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/jancimertel/go-tdlib/v2/tdlib"
)

// GetLocalizationTargetInfo Returns information about the current localization target. This is an offline request if only_local is true. Can be called before authorization
// @param onlyLocal If true, returns only locally available information without sending network requests
func (client *Client) GetLocalizationTargetInfo(onlyLocal bool) (*tdlib.LocalizationTargetInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":      "getLocalizationTargetInfo",
		"only_local": onlyLocal,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var localizationTargetInfo tdlib.LocalizationTargetInfo
	err = json.Unmarshal(result.Raw, &localizationTargetInfo)
	return &localizationTargetInfo, err

}
