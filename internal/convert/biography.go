package convert

import (
	"log"
	"time"

	"github.com/pawpawchat/core/internal/model"
	"github.com/pawpawchat/profile/api/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MustBiographyPb(src *pb.Biography) *model.Biography {
	return &model.Biography{
		FirstName:  src.FirstName,
		SecondName: src.SecondName,
		Birthday:   src.Birthday.AsTime().Format(time.RFC3339),
	}
}

func BiographyPb(src *pb.Biography) (*model.Biography, error) {
	if src == nil {
		return nil, nil
	}

	var dst model.Biography
	empty := true

	if src.FirstName != "" {
		empty = false
		dst.FirstName = src.FirstName
	}
	if src.SecondName != "" {
		empty = false
		dst.SecondName = src.SecondName
	}
	if src.Birthday != nil {
		dst.Birthday = src.Birthday.AsTime().Format(time.RFC3339)
		if dst.Birthday != "" {
			empty = false
		}
	}

	if empty {
		return nil, nil
	}

	return &dst, nil
}

func MustBiography(src *model.Biography) *pb.Biography {
	birthday, err := time.Parse(time.RFC3339, src.Birthday)
	if err != nil {
		log.Fatal(err)
	}

	return &pb.Biography{
		FirstName:  src.FirstName,
		SecondName: src.SecondName,
		Birthday:   timestamppb.New(birthday),
	}
}

func Biography(src *model.Biography) (*pb.Biography, error) {
	if src == nil {
		return nil, nil
	}

	var biographyPb pb.Biography
	empty := true

	if src.FirstName != "" {
		empty = false
		biographyPb.FirstName = src.FirstName
	}
	if src.SecondName != "" {
		empty = false
		biographyPb.SecondName = src.SecondName
	}

	if src.Birthday != "" {
		birthday, err := time.Parse(time.RFC3339, src.Birthday)
		if err != nil {
			return nil, err
		}
		empty = false
		biographyPb.Birthday = timestamppb.New(birthday)
	}

	if empty {
		return nil, nil
	}

	return &biographyPb, nil
}
