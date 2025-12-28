package features

import (
	"reflect"
	"sync"
	"testing"
)

const (
	testValueString = "test"
)

// testFeature is a simple feature for testing.
type testFeature struct {
	featureBase
	value string
}

// testFeature2 is another feature type for testing.
type testFeature2 struct {
	featureBase
	count int
}

func TestFeatureCollection(t *testing.T) {
	t.Run(
		"NewFeatureCollection",
		func(t *testing.T) {
			fc := NewFeatureCollection()
			if fc == nil {
				t.Error(
					"expected non-nil feature collection",
				)
			}
			if fc.Parent() != nil {
				t.Error(
					"expected nil parent for new collection",
				)
			}
		},
	)

	t.Run(
		"NewFeatureCollectionWithParent",
		func(t *testing.T) {
			parent := NewFeatureCollection()
			child := NewFeatureCollectionWithParent(
				parent,
			)

			if child.Parent() != parent {
				t.Error(
					"expected parent to be set",
				)
			}
		},
	)

	t.Run("Set and Get", func(t *testing.T) {
		fc := NewFeatureCollection()
		f := &testFeature{value: testValueString}

		fc.SetByType(
			reflect.TypeFor[*testFeature](),
			f,
		)

		got := Get[*testFeature](fc)
		if got == nil {
			t.Fatal("expected to get feature")
		}
		if got.value != testValueString {
			t.Errorf(
				"expected value 'test', got '%s'",
				got.value,
			)
		}
	})

	t.Run(
		"Get nil collection",
		func(t *testing.T) {
			var fc *FeatureCollection
			got := Get[*testFeature](fc)
			if got != nil {
				t.Error(
					"expected nil for nil collection",
				)
			}
		},
	)

	t.Run("Get not found", func(t *testing.T) {
		fc := NewFeatureCollection()
		got := Get[*testFeature](fc)
		if got != nil {
			t.Error(
				"expected nil for not found feature",
			)
		}
	})

	t.Run("Replace feature", func(t *testing.T) {
		fc := NewFeatureCollection()
		f1 := &testFeature{value: "first"}
		f2 := &testFeature{value: "second"}

		fc.SetByType(
			reflect.TypeFor[*testFeature](),
			f1,
		)
		fc.SetByType(
			reflect.TypeFor[*testFeature](),
			f2,
		)

		got := Get[*testFeature](fc)
		if got.value != "second" {
			t.Errorf(
				"expected 'second', got '%s'",
				got.value,
			)
		}
	})

	t.Run(
		"Multiple feature types",
		func(t *testing.T) {
			fc := NewFeatureCollection()
			f1 := &testFeature{value: "feature1"}
			f2 := &testFeature2{count: 42}

			fc.SetByType(
				reflect.TypeFor[*testFeature](),
				f1,
			)
			fc.SetByType(
				reflect.TypeFor[*testFeature2](),
				f2,
			)

			got1 := Get[*testFeature](fc)
			got2 := Get[*testFeature2](fc)

			if got1 == nil ||
				got1.value != "feature1" {
				t.Error("expected feature1")
			}
			if got2 == nil || got2.count != 42 {
				t.Error("expected feature2")
			}
		},
	)
}

func TestFeatureInheritance(t *testing.T) {
	t.Run(
		"Inherit from parent",
		func(t *testing.T) {
			parent := NewFeatureCollection()
			child := NewFeatureCollectionWithParent(
				parent,
			)

			f := &testFeature{value: "parent"}
			parent.SetByType(
				reflect.TypeFor[*testFeature](),
				f,
			)

			got := Get[*testFeature](child)
			if got == nil {
				t.Fatal(
					"expected to inherit from parent",
				)
			}
			if got.value != "parent" {
				t.Errorf(
					"expected 'parent', got '%s'",
					got.value,
				)
			}
		},
	)

	t.Run(
		"Child overrides parent",
		func(t *testing.T) {
			parent := NewFeatureCollection()
			child := NewFeatureCollectionWithParent(
				parent,
			)

			parentF := &testFeature{
				value: "parent",
			}
			childF := &testFeature{value: "child"}

			parent.SetByType(
				reflect.TypeFor[*testFeature](),
				parentF,
			)
			child.SetByType(
				reflect.TypeFor[*testFeature](),
				childF,
			)

			got := Get[*testFeature](child)
			if got.value != "child" {
				t.Errorf(
					"expected 'child', got '%s'",
					got.value,
				)
			}
		},
	)

	t.Run(
		"Deep inheritance chain",
		func(t *testing.T) {
			grandparent := NewFeatureCollection()
			parent := NewFeatureCollectionWithParent(
				grandparent,
			)
			child := NewFeatureCollectionWithParent(
				parent,
			)

			f := &testFeature{
				value: "grandparent",
			}
			grandparent.SetByType(
				reflect.TypeFor[*testFeature](),
				f,
			)

			got := Get[*testFeature](child)
			if got == nil ||
				got.value != "grandparent" {
				t.Error(
					"expected to inherit from grandparent",
				)
			}
		},
	)

	t.Run(
		"GetLocal does not check parent",
		func(t *testing.T) {
			parent := NewFeatureCollection()
			child := NewFeatureCollectionWithParent(
				parent,
			)

			f := &testFeature{value: "parent"}
			parent.SetByType(
				reflect.TypeFor[*testFeature](),
				f,
			)

			got := GetLocal[*testFeature](child)
			if got != nil {
				t.Error(
					"GetLocal should not check parent",
				)
			}
		},
	)
}

func TestFeatureCollectionMethods(t *testing.T) {
	t.Run("Has", func(t *testing.T) {
		fc := NewFeatureCollection()
		f := &testFeature{value: "test"}
		fc.SetByType(
			reflect.TypeFor[*testFeature](),
			f,
		)

		if !fc.Has(
			reflect.TypeFor[*testFeature](),
		) {
			t.Error("expected Has to return true")
		}
		if fc.Has(
			reflect.TypeFor[*testFeature2](),
		) {
			t.Error(
				"expected Has to return false for non-existent type",
			)
		}
	})

	t.Run("HasLocal", func(t *testing.T) {
		parent := NewFeatureCollection()
		child := NewFeatureCollectionWithParent(
			parent,
		)

		f := &testFeature{value: "parent"}
		parent.SetByType(
			reflect.TypeFor[*testFeature](),
			f,
		)

		if child.HasLocal(
			reflect.TypeFor[*testFeature](),
		) {
			t.Error(
				"HasLocal should return false for parent features",
			)
		}
		if !parent.HasLocal(
			reflect.TypeFor[*testFeature](),
		) {
			t.Error(
				"HasLocal should return true for local features",
			)
		}
	})

	t.Run("Remove", func(t *testing.T) {
		fc := NewFeatureCollection()
		f := &testFeature{value: "test"}
		fc.SetByType(
			reflect.TypeFor[*testFeature](),
			f,
		)

		removed := fc.Remove(
			reflect.TypeFor[*testFeature](),
		)
		if !removed {
			t.Error(
				"expected Remove to return true",
			)
		}

		got := Get[*testFeature](fc)
		if got != nil {
			t.Error(
				"expected feature to be removed",
			)
		}
	})

	t.Run(
		"Remove non-existent",
		func(t *testing.T) {
			fc := NewFeatureCollection()
			removed := fc.Remove(
				reflect.TypeFor[*testFeature](),
			)
			if removed {
				t.Error(
					"expected Remove to return false for non-existent feature",
				)
			}
		},
	)

	t.Run("SetParent", func(t *testing.T) {
		parent1 := NewFeatureCollection()
		parent2 := NewFeatureCollection()
		child := NewFeatureCollectionWithParent(
			parent1,
		)

		f1 := &testFeature{value: "parent1"}
		f2 := &testFeature{value: "parent2"}
		parent1.SetByType(
			reflect.TypeFor[*testFeature](),
			f1,
		)
		parent2.SetByType(
			reflect.TypeFor[*testFeature](),
			f2,
		)

		// Initially inherits from parent1
		got := Get[*testFeature](child)
		if got.value != "parent1" {
			t.Errorf(
				"expected 'parent1', got '%s'",
				got.value,
			)
		}

		// Change parent to parent2
		child.SetParent(parent2)
		got = Get[*testFeature](child)
		if got.value != "parent2" {
			t.Errorf(
				"expected 'parent2', got '%s'",
				got.value,
			)
		}
	})

	t.Run("Set nil feature", func(_ *testing.T) {
		fc := NewFeatureCollection()
		fc.Set(nil) // Should not panic

		fc.SetByType(
			reflect.TypeFor[*testFeature](),
			nil,
		) // Should not panic
	})
}

func TestThreadSafety(t *testing.T) {
	t.Run("Concurrent reads", func(t *testing.T) {
		fc := NewFeatureCollection()
		f := &testFeature{value: testValueString}
		fc.SetByType(
			reflect.TypeFor[*testFeature](),
			f,
		)

		var wg sync.WaitGroup
		for range 100 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				got := Get[*testFeature](fc)
				if got == nil ||
					got.value != testValueString {
					t.Error(
						"concurrent read failed",
					)
				}
			}()
		}
		wg.Wait()
	})

	t.Run(
		"Concurrent writes",
		func(t *testing.T) {
			fc := NewFeatureCollection()

			var wg sync.WaitGroup
			for i := range 100 {
				wg.Add(1)
				go func(_ int) {
					defer wg.Done()
					f := &testFeature{
						value: "test",
					}
					fc.SetByType(
						reflect.TypeFor[*testFeature](),
						f,
					)
				}(i)
			}
			wg.Wait()

			// Should have one feature at the end
			got := Get[*testFeature](fc)
			if got == nil {
				t.Error(
					"expected feature after concurrent writes",
				)
			}
		},
	)

	t.Run(
		"Concurrent read-write",
		func(_ *testing.T) {
			fc := NewFeatureCollection()
			f := &testFeature{value: "initial"}
			fc.SetByType(
				reflect.TypeFor[*testFeature](),
				f,
			)

			var wg sync.WaitGroup

			// Readers
			for range 50 {
				wg.Add(1)
				go func() {
					defer wg.Done()
					_ = Get[*testFeature](fc)
				}()
			}

			// Writers
			for i := range 50 {
				wg.Add(1)
				go func(_ int) {
					defer wg.Done()
					f := &testFeature{
						value: "updated",
					}
					fc.SetByType(
						reflect.TypeFor[*testFeature](),
						f,
					)
				}(i)
			}

			wg.Wait()
		},
	)
}

func TestGetRequired(t *testing.T) {
	t.Run("Feature exists", func(t *testing.T) {
		fc := NewFeatureCollection()
		f := &testFeature{value: testValueString}
		fc.SetByType(
			reflect.TypeFor[*testFeature](),
			f,
		)

		got := GetRequired[*testFeature](fc)
		if got == nil ||
			got.value != testValueString {
			t.Error("expected to get feature")
		}
	})

	t.Run(
		"Feature missing panics",
		func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Error(
						"expected panic for missing required feature",
					)
				}
			}()

			fc := NewFeatureCollection()
			_ = GetRequired[*testFeature](fc)
		},
	)
}
