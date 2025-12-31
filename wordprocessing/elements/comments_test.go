package elements

import (
	"testing"
	"time"
)

func TestCommentDone(t *testing.T) {
	t.Run("DefaultNotDone", func(t *testing.T) {
		// Create a comment without setting Done
		comment := NewComment(
			1,
			"John Doe",
			"This is a comment",
		)

		// Should return false by default
		if comment.Done() {
			t.Error(
				"Expected Done() to return false by default",
			)
		}
	})

	t.Run("SetDoneTrue", func(t *testing.T) {
		// Create a comments root and add a comment
		comments := NewComments()
		comment := comments.AddComment(
			"Jane Smith",
			"Another comment",
		)

		// Set Done to true
		comment.SetDone(true)

		// Should return true
		if !comment.Done() {
			t.Error(
				"Expected Done() to return true after SetDone(true)",
			)
		}

		// Verify the attribute is set correctly
		attr, found := comment.GetAttribute(
			"done",
			NamespaceW15,
		)
		if !found {
			t.Error(
				"Expected w15:done attribute to be set",
			)
		}
		if attr.Value() != "1" {
			t.Errorf(
				"Expected w15:done='1', got '%s'",
				attr.Value(),
			)
		}
	})

	t.Run("SetDoneFalse", func(t *testing.T) {
		// Create a comments root and add a comment
		comments := NewComments()
		comment := comments.AddComment(
			"Bob Johnson",
			"Yet another comment",
		)

		// Set Done to true first
		comment.SetDone(true)
		if !comment.Done() {
			t.Error(
				"Expected Done() to return true after SetDone(true)",
			)
		}

		// Now set it to false
		comment.SetDone(false)

		// Should return false
		if comment.Done() {
			t.Error(
				"Expected Done() to return false after SetDone(false)",
			)
		}

		// Verify the attribute is removed
		_, found := comment.GetAttribute(
			"done",
			NamespaceW15,
		)
		if found {
			t.Error(
				"Expected w15:done attribute to be removed when SetDone(false)",
			)
		}
	})

	t.Run("RoundtripDone", func(t *testing.T) {
		// Create a comments root and add a comment
		comments := NewComments()
		comment := comments.AddComment(
			"Alice Cooper",
			"Roundtrip test",
		)

		// Set Done to true
		comment.SetDone(true)

		// Verify we can read it back
		if !comment.Done() {
			t.Error(
				"Expected Done() to return true after SetDone(true)",
			)
		}

		// Set to false
		comment.SetDone(false)

		// Verify we can read it back as false
		if comment.Done() {
			t.Error(
				"Expected Done() to return false after SetDone(false)",
			)
		}

		// Set to true again
		comment.SetDone(true)

		// Verify we can read it back as true again
		if !comment.Done() {
			t.Error(
				"Expected Done() to return true after second SetDone(true)",
			)
		}
	})
}

func TestCommentsW15Namespace(t *testing.T) {
	t.Run(
		"NamespaceAddedWhenDoneSet",
		func(t *testing.T) {
			// Create a comments root and add a comment
			comments := NewComments()
			comment := comments.AddComment(
				"Test Author",
				"Test comment",
			)

			// Set Done to true
			comment.SetDone(true)

			// Now w15 namespace should be declared
			attr, found := comments.GetAttribute(
				"w15",
				"http://www.w3.org/2000/xmlns/",
			)
			if !found {
				t.Error(
					"Expected xmlns:w15 attribute to be declared after SetDone(true)",
				)
			} else if attr.Value() != NamespaceW15 {
				t.Errorf("Expected xmlns:w15='%s', got '%s'", NamespaceW15, attr.Value())
			}
		},
	)

	t.Run(
		"AddW15NamespaceDirectly",
		func(t *testing.T) {
			// Create a comments root
			comments := NewComments()

			// Add w15 namespace directly
			comments.AddW15Namespace()

			// Verify it's declared
			attr, found := comments.GetAttribute(
				"w15",
				"http://www.w3.org/2000/xmlns/",
			)
			if !found {
				t.Error(
					"Expected xmlns:w15 attribute to be declared after AddW15Namespace()",
				)
			} else if attr.Value() != NamespaceW15 {
				t.Errorf("Expected xmlns:w15='%s', got '%s'", NamespaceW15, attr.Value())
			}

			// Calling it again should be idempotent
			comments.AddW15Namespace()
			attr, found = comments.GetAttribute(
				"w15",
				"http://www.w3.org/2000/xmlns/",
			)
			if !found {
				t.Error(
					"Expected xmlns:w15 attribute to still be declared after second AddW15Namespace()",
				)
			} else if attr.Value() != NamespaceW15 {
				t.Errorf("Expected xmlns:w15='%s', got '%s'", NamespaceW15, attr.Value())
			}
		},
	)
}

func TestCommentsByAuthor(t *testing.T) {
	t.Run(
		"ReturnsMatchingComments",
		func(t *testing.T) {
			// Create comments with different authors
			comments := NewComments()
			c1 := comments.AddComment(
				"Alice",
				"First comment",
			)
			c2 := comments.AddComment(
				"Bob",
				"Second comment",
			)
			c3 := comments.AddComment(
				"Alice",
				"Third comment",
			)
			c4 := comments.AddComment(
				"Charlie",
				"Fourth comment",
			)

			// Get all comments by Alice
			aliceComments := comments.ByAuthor(
				"Alice",
			)

			// Should have exactly 2 comments
			if len(aliceComments) != 2 {
				t.Errorf(
					"Expected 2 comments by Alice, got %d",
					len(aliceComments),
				)
			}

			// Verify the correct comments are returned
			foundC1, foundC3 := false, false
			for _, comment := range aliceComments {
				if comment.Id() == c1.Id() {
					foundC1 = true
				}
				if comment.Id() == c3.Id() {
					foundC3 = true
				}
			}
			if !foundC1 || !foundC3 {
				t.Error(
					"Expected to find both of Alice's comments",
				)
			}

			// Verify Bob's comments
			bobComments := comments.ByAuthor(
				"Bob",
			)
			if len(bobComments) != 1 {
				t.Errorf(
					"Expected 1 comment by Bob, got %d",
					len(bobComments),
				)
			}
			if len(bobComments) > 0 &&
				bobComments[0].Id() != c2.Id() {
				t.Error(
					"Expected to find Bob's comment",
				)
			}

			// Verify Charlie's comments
			charlieComments := comments.ByAuthor(
				"Charlie",
			)
			if len(charlieComments) != 1 {
				t.Errorf(
					"Expected 1 comment by Charlie, got %d",
					len(charlieComments),
				)
			}
			if len(charlieComments) > 0 &&
				charlieComments[0].Id() != c4.Id() {
				t.Error(
					"Expected to find Charlie's comment",
				)
			}
		},
	)

	t.Run(
		"ReturnsEmptySliceWhenNoMatches",
		func(t *testing.T) {
			// Create comments with different authors
			comments := NewComments()
			comments.AddComment(
				"Alice",
				"First comment",
			)
			comments.AddComment(
				"Bob",
				"Second comment",
			)

			// Search for non-existent author
			result := comments.ByAuthor("Dave")

			// Should return empty slice (not nil)
			if result == nil {
				t.Error(
					"Expected non-nil slice when no matches",
				)
			}
			if len(result) != 0 {
				t.Errorf(
					"Expected empty slice, got %d comments",
					len(result),
				)
			}
		},
	)

	t.Run(
		"CaseSensitiveMatch",
		func(t *testing.T) {
			// Create comments with different case variations
			comments := NewComments()
			comments.AddComment(
				"Alice",
				"First comment",
			)
			comments.AddComment(
				"alice",
				"Second comment",
			)
			comments.AddComment(
				"ALICE",
				"Third comment",
			)

			// Search for exact case
			result := comments.ByAuthor("Alice")
			if len(result) != 1 {
				t.Errorf(
					"Expected 1 comment for 'Alice', got %d",
					len(result),
				)
			}

			result = comments.ByAuthor("alice")
			if len(result) != 1 {
				t.Errorf(
					"Expected 1 comment for 'alice', got %d",
					len(result),
				)
			}

			result = comments.ByAuthor("ALICE")
			if len(result) != 1 {
				t.Errorf(
					"Expected 1 comment for 'ALICE', got %d",
					len(result),
				)
			}
		},
	)

	t.Run("EmptyAuthorName", func(t *testing.T) {
		// Create comments with empty and non-empty authors
		comments := NewComments()
		c1 := comments.AddComment(
			"",
			"Anonymous comment",
		)
		comments.AddComment(
			"Alice",
			"Named comment",
		)

		// Search for empty author
		result := comments.ByAuthor("")
		if len(result) != 1 {
			t.Errorf(
				"Expected 1 comment with empty author, got %d",
				len(result),
			)
		}
		if len(result) > 0 &&
			result[0].Id() != c1.Id() {
			t.Error(
				"Expected to find the anonymous comment",
			)
		}
	})
}

func TestCommentsByDateRange(t *testing.T) {
	t.Run(
		"ReturnsCommentsWithinRange",
		func(t *testing.T) {
			// Create comments with specific dates
			comments := NewComments()

			date1 := time.Date(
				2024,
				1,
				1,
				12,
				0,
				0,
				0,
				time.UTC,
			)
			date2 := time.Date(
				2024,
				1,
				5,
				12,
				0,
				0,
				0,
				time.UTC,
			)
			date3 := time.Date(
				2024,
				1,
				10,
				12,
				0,
				0,
				0,
				time.UTC,
			)
			date4 := time.Date(
				2024,
				1,
				15,
				12,
				0,
				0,
				0,
				time.UTC,
			)

			c1 := comments.AddComment(
				"Alice",
				"First",
			)
			c1.SetDate(date1)

			c2 := comments.AddComment(
				"Bob",
				"Second",
			)
			c2.SetDate(date2)

			c3 := comments.AddComment(
				"Charlie",
				"Third",
			)
			c3.SetDate(date3)

			c4 := comments.AddComment(
				"Dave",
				"Fourth",
			)
			c4.SetDate(date4)

			// Search for comments between Jan 5 and Jan 10 (inclusive)
			from := time.Date(
				2024,
				1,
				5,
				0,
				0,
				0,
				0,
				time.UTC,
			)
			to := time.Date(
				2024,
				1,
				10,
				23,
				59,
				59,
				0,
				time.UTC,
			)
			result := comments.ByDateRange(
				from,
				to,
			)

			// Should return c2 and c3
			if len(result) != 2 {
				t.Errorf(
					"Expected 2 comments in range, got %d",
					len(result),
				)
			}

			foundC2, foundC3 := false, false
			for _, comment := range result {
				if comment.Id() == c2.Id() {
					foundC2 = true
				}
				if comment.Id() == c3.Id() {
					foundC3 = true
				}
			}
			if !foundC2 || !foundC3 {
				t.Error(
					"Expected to find c2 and c3 in the date range",
				)
			}
		},
	)

	t.Run(
		"ReturnsEmptySliceWhenNoMatches",
		func(t *testing.T) {
			// Create comments with specific dates
			comments := NewComments()

			c1 := comments.AddComment(
				"Alice",
				"First",
			)
			c1.SetDate(
				time.Date(
					2024,
					1,
					1,
					12,
					0,
					0,
					0,
					time.UTC,
				),
			)

			c2 := comments.AddComment(
				"Bob",
				"Second",
			)
			c2.SetDate(
				time.Date(
					2024,
					1,
					20,
					12,
					0,
					0,
					0,
					time.UTC,
				),
			)

			// Search for range with no comments
			from := time.Date(
				2024,
				1,
				5,
				0,
				0,
				0,
				0,
				time.UTC,
			)
			to := time.Date(
				2024,
				1,
				10,
				0,
				0,
				0,
				0,
				time.UTC,
			)
			result := comments.ByDateRange(
				from,
				to,
			)

			// Should return empty slice (not nil)
			if result == nil {
				t.Error(
					"Expected non-nil slice when no matches",
				)
			}
			if len(result) != 0 {
				t.Errorf(
					"Expected empty slice, got %d comments",
					len(result),
				)
			}
		},
	)

	t.Run(
		"NoLowerBoundWhenFromIsZero",
		func(t *testing.T) {
			// Create comments with specific dates
			comments := NewComments()

			c1 := comments.AddComment(
				"Alice",
				"First",
			)
			c1.SetDate(
				time.Date(
					2024,
					1,
					1,
					12,
					0,
					0,
					0,
					time.UTC,
				),
			)

			c2 := comments.AddComment(
				"Bob",
				"Second",
			)
			c2.SetDate(
				time.Date(
					2024,
					1,
					10,
					12,
					0,
					0,
					0,
					time.UTC,
				),
			)

			c3 := comments.AddComment(
				"Charlie",
				"Third",
			)
			c3.SetDate(
				time.Date(
					2024,
					1,
					20,
					12,
					0,
					0,
					0,
					time.UTC,
				),
			)

			// Search with zero from date and upper bound
			to := time.Date(
				2024,
				1,
				10,
				23,
				59,
				59,
				0,
				time.UTC,
			)
			result := comments.ByDateRange(
				time.Time{},
				to,
			)

			// Should return c1 and c2 (no lower bound)
			if len(result) != 2 {
				t.Errorf(
					"Expected 2 comments (no lower bound), got %d",
					len(result),
				)
			}
		},
	)

	t.Run(
		"NoUpperBoundWhenToIsZero",
		func(t *testing.T) {
			// Create comments with specific dates
			comments := NewComments()

			c1 := comments.AddComment(
				"Alice",
				"First",
			)
			c1.SetDate(
				time.Date(
					2024,
					1,
					1,
					12,
					0,
					0,
					0,
					time.UTC,
				),
			)

			c2 := comments.AddComment(
				"Bob",
				"Second",
			)
			c2.SetDate(
				time.Date(
					2024,
					1,
					10,
					12,
					0,
					0,
					0,
					time.UTC,
				),
			)

			c3 := comments.AddComment(
				"Charlie",
				"Third",
			)
			c3.SetDate(
				time.Date(
					2024,
					1,
					20,
					12,
					0,
					0,
					0,
					time.UTC,
				),
			)

			// Search with lower bound and zero to date
			from := time.Date(
				2024,
				1,
				10,
				0,
				0,
				0,
				0,
				time.UTC,
			)
			result := comments.ByDateRange(
				from,
				time.Time{},
			)

			// Should return c2 and c3 (no upper bound)
			if len(result) != 2 {
				t.Errorf(
					"Expected 2 comments (no upper bound), got %d",
					len(result),
				)
			}
		},
	)

	t.Run(
		"BothBoundsZeroReturnsAll",
		func(t *testing.T) {
			// Create comments with specific dates
			comments := NewComments()

			c1 := comments.AddComment(
				"Alice",
				"First",
			)
			c1.SetDate(
				time.Date(
					2024,
					1,
					1,
					12,
					0,
					0,
					0,
					time.UTC,
				),
			)

			c2 := comments.AddComment(
				"Bob",
				"Second",
			)
			c2.SetDate(
				time.Date(
					2024,
					1,
					10,
					12,
					0,
					0,
					0,
					time.UTC,
				),
			)

			c3 := comments.AddComment(
				"Charlie",
				"Third",
			)
			c3.SetDate(
				time.Date(
					2024,
					1,
					20,
					12,
					0,
					0,
					0,
					time.UTC,
				),
			)

			// Search with both dates zero
			result := comments.ByDateRange(
				time.Time{},
				time.Time{},
			)

			// Should return all comments
			if len(result) != 3 {
				t.Errorf(
					"Expected 3 comments (no bounds), got %d",
					len(result),
				)
			}
		},
	)

	t.Run(
		"SkipsCommentsWithZeroDates",
		func(t *testing.T) {
			// Create comments, some with dates, some without
			comments := NewComments()

			c1 := comments.AddComment(
				"Alice",
				"First",
			)
			c1.SetDate(
				time.Date(
					2024,
					1,
					1,
					12,
					0,
					0,
					0,
					time.UTC,
				),
			)

			// c2 has no date set (zero time)
			comments.AddComment("Bob", "Second")

			c3 := comments.AddComment(
				"Charlie",
				"Third",
			)
			c3.SetDate(
				time.Date(
					2024,
					1,
					20,
					12,
					0,
					0,
					0,
					time.UTC,
				),
			)

			// Search with wide range
			from := time.Date(
				2024,
				1,
				1,
				0,
				0,
				0,
				0,
				time.UTC,
			)
			to := time.Date(
				2024,
				1,
				31,
				0,
				0,
				0,
				0,
				time.UTC,
			)
			result := comments.ByDateRange(
				from,
				to,
			)

			// Should return only c1 and c3 (c2 has zero date)
			if len(result) != 2 {
				t.Errorf(
					"Expected 2 comments (skipping zero dates), got %d",
					len(result),
				)
			}
		},
	)

	t.Run("InclusiveBounds", func(t *testing.T) {
		// Create comments with dates exactly on the boundaries
		comments := NewComments()

		exactDate := time.Date(
			2024,
			1,
			10,
			12,
			0,
			0,
			0,
			time.UTC,
		)

		c1 := comments.AddComment(
			"Alice",
			"On boundary",
		)
		c1.SetDate(exactDate)

		// Search with from == to == exactDate
		result := comments.ByDateRange(
			exactDate,
			exactDate,
		)

		// Should return c1 (inclusive bounds)
		if len(result) != 1 {
			t.Errorf(
				"Expected 1 comment (inclusive bounds), got %d",
				len(result),
			)
		}
		if len(result) > 0 &&
			result[0].Id() != c1.Id() {
			t.Error(
				"Expected to find the comment on the exact boundary",
			)
		}
	})
}

func TestCommentsCount(t *testing.T) {
	t.Run(
		"ReturnsCorrectCount",
		func(t *testing.T) {
			// Create comments
			comments := NewComments()

			// Initially should be 0
			if count := comments.Count(); count != 0 {
				t.Errorf(
					"Expected count 0 for empty comments, got %d",
					count,
				)
			}

			// Add one comment
			comments.AddComment("Alice", "First")
			if count := comments.Count(); count != 1 {
				t.Errorf(
					"Expected count 1, got %d",
					count,
				)
			}

			// Add more comments
			comments.AddComment("Bob", "Second")
			comments.AddComment(
				"Charlie",
				"Third",
			)
			if count := comments.Count(); count != 3 {
				t.Errorf(
					"Expected count 3, got %d",
					count,
				)
			}

			// Remove a comment
			comments.RemoveComment(
				2,
			) // Remove Bob's comment
			if count := comments.Count(); count != 2 {
				t.Errorf(
					"Expected count 2 after removal, got %d",
					count,
				)
			}
		},
	)

	t.Run(
		"CountWithNoComments",
		func(t *testing.T) {
			comments := NewComments()
			if count := comments.Count(); count != 0 {
				t.Errorf(
					"Expected count 0 for new Comments, got %d",
					count,
				)
			}
		},
	)
}
