package convert

import (
	"time"

	"github.com/pawpawchat/core/internal/model"
	"github.com/pawpawchat/profile/api/pb"
)

func MustAvatarPb(src *pb.Avatar) *model.Avatar {
	return &model.Avatar{
		ID:      src.AvatarId,
		OrigURL: src.OrigUrl,
		AddedAt: src.AddedAt.AsTime().Format(time.RFC3339),
	}
}

func AvatarPb(src *pb.Avatar) (*model.Avatar, error) {
	if src == nil {
		return nil, nil
	}

	var avatar model.Avatar
	empty := true

	if src.AvatarId != 0 {
		empty = false
		avatar.ID = src.AvatarId
	}
	if src.OrigUrl != "" {
		empty = false
		avatar.OrigURL = src.OrigUrl
	}

	avatar.AddedAt = src.AddedAt.AsTime().Format(time.RFC3339)
	if avatar.AddedAt != "" {
		empty = false
	}

	if empty {
		return nil, nil
	}

	return &avatar, nil
}
