package openxml

import "iter"

// First returns the first child element of type T from the parent.
// Returns the zero value of T if not found.
func First[T Element](
	parent CompositeElement,
) (result T) {
	if parent == nil {
		return result
	}

	for child := range parent.Children() {
		if typed, ok := child.(T); ok {
			return typed
		}
	}

	return result
}

// All returns an iterator over all child elements of type T from the parent.
func All[T Element](
	parent CompositeElement,
) iter.Seq[T] {
	return func(yield func(T) bool) {
		if parent == nil {
			return
		}

		for child := range parent.Children() {
			if typed, ok := child.(T); ok {
				if !yield(typed) {
					return
				}
			}
		}
	}
}

// OfType filters an element sequence to only include elements of type T.
func OfType[T Element](
	elements iter.Seq[Element],
) iter.Seq[T] {
	return func(yield func(T) bool) {
		for elem := range elements {
			if typed, ok := elem.(T); ok {
				if !yield(typed) {
					return
				}
			}
		}
	}
}

// Descendants returns an iterator over all descendant elements (depth-first).
func Descendants(
	el CompositeElement,
) iter.Seq[Element] {
	return func(yield func(Element) bool) {
		if el == nil {
			return
		}
		descendantsRecursive(el, yield)
	}
}

// descendantsRecursive is the recursive helper for Descendants.
func descendantsRecursive(
	parent CompositeElement,
	yield func(Element) bool,
) bool {
	for child := range parent.Children() {
		if !yield(child) {
			return false
		}
		if childComp, ok := child.(CompositeElement); ok {
			if !descendantsRecursive(
				childComp,
				yield,
			) {
				return false
			}
		}
	}

	return true
}

// DescendantsOfType returns an iterator over all descendant elements of type T.
func DescendantsOfType[T Element](
	el CompositeElement,
) iter.Seq[T] {
	return func(yield func(T) bool) {
		if el == nil {
			return
		}
		descendantsOfTypeRecursive[T](el, yield)
	}
}

// descendantsOfTypeRecursive is the recursive helper for DescendantsOfType.
func descendantsOfTypeRecursive[T Element](
	parent CompositeElement,
	yield func(T) bool,
) bool {
	for child := range parent.Children() {
		if typed, ok := child.(T); ok {
			if !yield(typed) {
				return false
			}
		}
		if childComp, ok := child.(CompositeElement); ok {
			if !descendantsOfTypeRecursive[T](
				childComp,
				yield,
			) {
				return false
			}
		}
	}

	return true
}

// Ancestors returns an iterator over all ancestor elements (from parent to root).
func Ancestors(el Element) iter.Seq[Element] {
	return func(yield func(Element) bool) {
		if el == nil {
			return
		}

		for parent := el.Parent(); parent != nil; parent = parent.Parent() {
			if !yield(parent) {
				return
			}
		}
	}
}

// AncestorsOfType returns an iterator over all ancestor elements of type T.
func AncestorsOfType[T Element](
	el Element,
) iter.Seq[T] {
	return func(yield func(T) bool) {
		if el == nil {
			return
		}

		for parent := el.Parent(); parent != nil; parent = parent.Parent() {
			if typed, ok := parent.(T); ok {
				if !yield(typed) {
					return
				}
			}
		}
	}
}

// FindAncestor returns the first ancestor of type T.
func FindAncestor[T Element](
	el Element,
) (result T, found bool) {
	for ancestor := range AncestorsOfType[T](el) {
		return ancestor, true
	}

	return result, false
}

// FindDescendant returns the first descendant of type T.
func FindDescendant[T Element](
	el CompositeElement,
) (result T, found bool) {
	for desc := range DescendantsOfType[T](el) {
		return desc, true
	}

	return result, false
}

// SiblingsAfter returns an iterator over siblings after the given element.
func SiblingsAfter(el Element) iter.Seq[Element] {
	return func(yield func(Element) bool) {
		if el == nil {
			return
		}

		parent := el.Parent()
		if parent == nil {
			return
		}

		parentComp, ok := parent.(CompositeElement)
		if !ok {
			return
		}

		found := false
		for child := range parentComp.Children() {
			if found {
				if !yield(child) {
					return
				}
			} else if child == el {
				found = true
			}
		}
	}
}

// SiblingsBefore returns an iterator over siblings before the given element.
func SiblingsBefore(
	el Element,
) iter.Seq[Element] {
	return func(yield func(Element) bool) {
		if el == nil {
			return
		}

		parent := el.Parent()
		if parent == nil {
			return
		}

		parentComp, ok := parent.(CompositeElement)
		if !ok {
			return
		}

		// Collect siblings before the element
		var siblings []Element
		for child := range parentComp.Children() {
			if child == el {
				break
			}
			siblings = append(siblings, child)
		}

		// Yield in reverse order (closest first)
		for i := len(siblings) - 1; i >= 0; i-- {
			if !yield(siblings[i]) {
				return
			}
		}
	}
}

// Count returns the number of elements in the sequence.
func Count(elements iter.Seq[Element]) int {
	count := 0
	for range elements {
		count++
	}

	return count
}

// CountOfType returns the number of elements of type T in the sequence.
func CountOfType[T Element](
	elements iter.Seq[Element],
) int {
	count := 0
	for elem := range elements {
		if _, ok := elem.(T); ok {
			count++
		}
	}

	return count
}

// ToSlice converts an element iterator to a slice.
func ToSlice(
	elements iter.Seq[Element],
) []Element {
	var result []Element
	for elem := range elements {
		result = append(result, elem)
	}

	return result
}

// ToSliceOfType converts an element iterator to a typed slice.
func ToSliceOfType[T Element](
	elements iter.Seq[Element],
) []T {
	var result []T
	for elem := range elements {
		if typed, ok := elem.(T); ok {
			result = append(result, typed)
		}
	}

	return result
}
