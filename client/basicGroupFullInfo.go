// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/Arman92/go-tdlib/tdlib"
)

// GetBasicGroupFullInfo Returns full information about a basic group by its identifier
// @param basicGroupID Basic group identifier
func (client *Client) GetBasicGroupFullInfo(basicGroupID int32) (*tdlib.BasicGroupFullInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":          "getBasicGroupFullInfo",
		"basic_group_id": basicGroupID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var basicGroupFullInfo tdlib.BasicGroupFullInfo
	err = json.Unmarshal(result.Raw, &basicGroupFullInfo)
	return &basicGroupFullInfo, err

}
