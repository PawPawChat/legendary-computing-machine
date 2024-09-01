package convert

import (
	"time"

	"github.com/pawpawchat/core/internal/model"
	"github.com/pawpawchat/profile/api/pb"
)

func MustProfilePb(src *pb.Profile) *model.Profile {
	return &model.Profile{
		ID:          src.Id,
		Username:    src.Username,
		Description: src.Description,
		LastSeen:    src.LastSeen.AsTime().Format(time.RFC3339),
		CreatedAt:   src.CreatedAt.AsTime().Format(time.RFC3339),
		Biography:   *MustBiographyPb(src.Biography),
	}
}

func ProfilePb(src *pb.Profile) (*model.Profile, error) {
	if src == nil {
		return nil, nil
	}

	var dst model.Profile
	empty := true

	if src.Id != 0 {
		empty = false
		dst.ID = src.Id
	}
	if src.Username != "" {
		empty = false
		dst.Username = src.Username
	}
	if src.Description != "" {
		empty = false
		dst.Description = src.Description
	}

	dst.LastSeen = src.LastSeen.AsTime().Format(time.RFC3339)
	if dst.LastSeen != "" {
		empty = false
	}
	dst.CreatedAt = src.LastSeen.AsTime().Format(time.RFC3339)
	if dst.LastSeen != "" {
		empty = false
	}

	bio, err := BiographyPb(src.Biography)
	if err != nil {
		return nil, err
	}
	dst.Biography = *bio

	if empty {
		return nil, nil
	}

	return &dst, nil
}
