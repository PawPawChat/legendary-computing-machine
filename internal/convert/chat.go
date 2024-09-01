package convert

import (
	"time"

	"github.com/pawpawchat/chat/api/pb"
	"github.com/pawpawchat/core/internal/model"
)

func MustChatPb(src *pb.Chat) *model.Chat {
	return &model.Chat{
		ID:         src.ChatId,
		Title:      src.Title,
		NumMembers: src.NumberMembers,
		CreatedAt:  src.CreatedAt.AsTime().Format(time.RFC3339),
	}
}

func ChatPb(src *pb.Chat) (*model.Chat, error) {
	if src == nil {
		return nil, nil
	}

	var chat model.Chat
	empty := true

	if src.ChatId != 0 {
		chat.ID = src.ChatId
		empty = false
	}
	if src.Title != "" {
		chat.Title = src.Title
		empty = false
	}

	chat.NumMembers = src.NumberMembers

	chat.CreatedAt = src.CreatedAt.AsTime().Format(time.RFC3339)
	if chat.CreatedAt != "" {
		empty = false
	}

	if empty {
		return nil, nil
	}

	return &chat, nil
}
