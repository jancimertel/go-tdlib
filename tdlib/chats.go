// AUTOGENERATED - DO NOT EDIT

package tdlib

// Chats Represents a list of chats
type Chats struct {
	tdCommon
	TotalCount int32   `json:"total_count"` // Approximate total count of chats found
	ChatIDs    []int64 `json:"chat_ids"`    // List of chat identifiers
}

// MessageType return the string telegram-type of Chats
func (chats *Chats) MessageType() string {
	return "chats"
}

// NewChats creates a new Chats
//
// @param totalCount Approximate total count of chats found
// @param chatIDs List of chat identifiers
func NewChats(totalCount int32, chatIDs []int64) *Chats {
	chatsTemp := Chats{
		tdCommon:   tdCommon{Type: "chats"},
		TotalCount: totalCount,
		ChatIDs:    chatIDs,
	}

	return &chatsTemp
}
