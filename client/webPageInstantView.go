// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/Arman92/go-tdlib/tdlib"
)

// GetWebPageInstantView Returns an instant view version of a web page if available. Returns a 404 error if the web page has no instant view page
// @param uRL The web page URL
// @param forceFull If true, the full instant view for the web page will be returned
func (client *Client) GetWebPageInstantView(uRL string, forceFull bool) (*tdlib.WebPageInstantView, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":      "getWebPageInstantView",
		"url":        uRL,
		"force_full": forceFull,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var webPageInstantView tdlib.WebPageInstantView
	err = json.Unmarshal(result.Raw, &webPageInstantView)
	return &webPageInstantView, err

}
