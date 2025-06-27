package Set

func Union[T comparable](sets ...*Set[T]) *Set[T] {
	result := NewSet[T]()

	for _, set := range sets {
		for item := range set.items {
			result.Add(item)
		}
	}

	return result
}

func Intersection[T comparable](sets ...*Set[T]) *Set[T] {
	result := NewSet[T]()

	if len(sets) == 0 {
		return result
	}

	smallestSet := sets[0]

	for _, set := range sets[1:] {
		if set.Size() < smallestSet.Size() {
			smallestSet = set
		}
	}

	for item := range smallestSet.items {
		isCommon := true

		for _, set := range sets {
			if set != smallestSet && !set.Contains(item) {
				isCommon = false
				break
			}
		}

		if isCommon {
			result.Add(item)
		}

	}
	return result
}
