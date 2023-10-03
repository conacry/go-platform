package mongoModel

type Collection string

func CollectionFrom(str string) (Collection, error) {
	if len(str) == 0 {
		return "", ErrCollectionValueIsRequired
	}

	return Collection(str), nil
}

func (c Collection) String() string {
	return string(c)
}
