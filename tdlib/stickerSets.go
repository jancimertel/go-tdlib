// AUTOGENERATED - DO NOT EDIT

package tdlib

// StickerSets Represents a list of sticker sets
type StickerSets struct {
	tdCommon
	TotalCount int32            `json:"total_count"` // Approximate total number of sticker sets found
	Sets       []StickerSetInfo `json:"sets"`        // List of sticker sets
}

// MessageType return the string telegram-type of StickerSets
func (stickerSets *StickerSets) MessageType() string {
	return "stickerSets"
}

// NewStickerSets creates a new StickerSets
//
// @param totalCount Approximate total number of sticker sets found
// @param sets List of sticker sets
func NewStickerSets(totalCount int32, sets []StickerSetInfo) *StickerSets {
	stickerSetsTemp := StickerSets{
		tdCommon:   tdCommon{Type: "stickerSets"},
		TotalCount: totalCount,
		Sets:       sets,
	}

	return &stickerSetsTemp
}