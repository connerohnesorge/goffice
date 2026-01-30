package validation

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/connerohnesorge/goffice/wordprocessing/elements"
)

// CommentValidationError represents a validation error for a comment.
type CommentValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Value   string `json:"value,omitempty"`
}

func (e CommentValidationError) Error() string {
	if e.Value != "" {
		return fmt.Sprintf("comment validation error in %s: %s (value: %s)", e.Field, e.Message, e.Value)
	}

	return fmt.Sprintf("comment validation error in %s: %s", e.Field, e.Message)
}

// ValidateComment performs comprehensive validation on a comment.
// Returns nil if valid, or CommentValidationError with details.
func ValidateComment(comment *elements.Comment) *CommentValidationError {
	// Validate ID
	if comment.Id() <= 0 {
		return &CommentValidationError{
			Field:   "id",
			Message: "comment ID must be positive",
			Value:   fmt.Sprintf("%d", comment.Id()),
		}
	}

	// Validate author
	author := comment.Author()
	if strings.TrimSpace(author) == "" {
		return &CommentValidationError{
			Field:   "author",
			Message: "comment author cannot be empty",
		}
	}

	if len(author) > 100 {
		return &CommentValidationError{
			Field:   "author",
			Message: "comment author cannot exceed 100 characters",
			Value:   fmt.Sprintf("%d", len(author)),
		}
	}

	// Validate content has text
	content := comment.ContentWithMentions()
	if strings.TrimSpace(content.Text) == "" {
		return &CommentValidationError{
			Field:   "content",
			Message: "comment content cannot be empty",
		}
	}

	// Validate mention format
	for _, mention := range content.Mentions {
		if !isValidUsername(mention.Username) {
			return &CommentValidationError{
				Field:   "mention",
				Message: "invalid username format in mention",
				Value:   mention.Username,
			}
		}
	}

	// Validate parent reference
	if comment.IsReply() {
		parentId := comment.ParentId()
		if parentId <= 0 {
			return &CommentValidationError{
				Field:   "parentId",
				Message: "reply comment must have valid parent ID",
			}
		}
	}

	return nil
}

// ValidateCommentThread validates a comment thread structure.
func ValidateCommentThread(thread *elements.CommentThread) *CommentValidationError {
	if thread.Root == nil {
		return &CommentValidationError{
			Field:   "root",
			Message: "thread must have a root comment",
		}
	}

	// Validate root comment
	if err := ValidateComment(thread.Root); err != nil {
		return err
	}

	// Validate replies are properly linked
	for _, reply := range thread.Replies {
		if reply.ParentId() != thread.Root.Id() {
			return &CommentValidationError{
				Field:   "reply",
				Message: "reply must reference thread root as parent",
			}
		}

		if err := ValidateComment(reply); err != nil {
			return &CommentValidationError{
				Field:   "reply",
				Message: fmt.Sprintf("reply validation failed: %s", err.Error()),
			}
		}
	}

	return nil
}

// isValidUsername checks if a username meets basic validation criteria.
func isValidUsername(username string) bool {
	if len(username) == 0 || len(username) > 50 {
		return false
	}

	// Allow alphanumeric characters, underscores, and hyphens
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+$`, username)

	return matched
}

// ValidateMentionSyntax checks if @mention syntax is valid in text.
func ValidateMentionSyntax(text string) []CommentValidationError {
	var errors []CommentValidationError
	re := regexp.MustCompile(`@(\w+)`)
	matches := re.FindAllStringSubmatch(text, -1)

	for _, match := range matches {
		if len(match) >= 2 {
			username := match[1]
			if !isValidUsername(username) {
				errors = append(errors, CommentValidationError{
					Field:   "mention",
					Message: "invalid username format in mention",
					Value:   username,
				})
			}
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}
