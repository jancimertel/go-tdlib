// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/Arman92/go-tdlib/tdlib"
)

// GetCountries Returns information about existing countries. Can be called before authorization
func (client *Client) GetCountries() (*tdlib.Countries, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "getCountries",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var countries tdlib.Countries
	err = json.Unmarshal(result.Raw, &countries)
	return &countries, err

}
