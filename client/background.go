// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/jancimertel/go-tdlib/v2/tdlib"
)

// SearchBackground Searches for a background by its name
// @param name The name of the background
func (client *Client) SearchBackground(name string) (*tdlib.Background, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "searchBackground",
		"name":  name,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var background tdlib.Background
	err = json.Unmarshal(result.Raw, &background)
	return &background, err

}

// SetBackground Changes the background selected by the user; adds background to the list of installed backgrounds
// @param background The input background to use. Pass null to create a new filled backgrounds. Pass null to remove the current background
// @param typeParam Background type. Pass null to use default type of the remote background. Pass null to remove the current background
// @param forDarkTheme True, if the background is chosen for dark theme
func (client *Client) SetBackground(background tdlib.InputBackground, typeParam tdlib.BackgroundType, forDarkTheme bool) (*tdlib.Background, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":          "setBackground",
		"background":     background,
		"type":           typeParam,
		"for_dark_theme": forDarkTheme,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var backgroundDummy tdlib.Background
	err = json.Unmarshal(result.Raw, &backgroundDummy)
	return &backgroundDummy, err

}
