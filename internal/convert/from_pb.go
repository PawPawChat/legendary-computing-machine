package convert

func MustFromPb[T any, U any](pb []*T, convert func(*T) *U) []*U {
	if len(pb) == 0 {
		return nil
	}

	var result []*U
	for _, item := range pb {
		result = append(result, convert(item))
	}
	return result
}

func FromPb[T any, U any](pb []*T, convert func(*T) (*U, error)) ([]*U, error) {
	if len(pb) == 0 {
		return nil, nil
	}

	var result []*U
	for _, item := range pb {
		iConverted, err := convert(item)
		if err != nil {
			return nil, err
		}

		result = append(result, iConverted)
	}
	return result, nil
}
