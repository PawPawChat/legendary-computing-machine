package convert

import (
	"time"

	"github.com/pawpawchat/chat/api/pb"
	"github.com/pawpawchat/core/internal/model"
)

func MustMessagePb(src *pb.Message) *model.Message {
	return &model.Message{
		ChatID:         src.ChatId,
		SenderID:       src.SenderId,
		SenderUsername: src.Username,
		Body:           src.Body,
		SentAt:         src.SentAt.AsTime().Format(time.RFC3339),
	}
}

func MessagePb(src *pb.Message) (*model.Message, error) {
	if src == nil {
		return nil, nil
	}

	var message model.Message
	empty := true

	if src.ChatId != 0 {
		message.ChatID = src.ChatId
		empty = false
	}
	if src.SenderId != 0 {
		message.SenderID = src.SenderId
		empty = false
	}
	if src.Username != "" {
		message.SenderUsername = src.Username
		empty = false
	}
	if src.Body != "" {
		message.Body = src.Body
		empty = false
	}

	message.SentAt = src.SentAt.AsTime().Format(time.RFC3339)
	if message.SentAt != "" {
		empty = false
	}

	if empty {
		return nil, nil
	}

	return &message, nil
}
