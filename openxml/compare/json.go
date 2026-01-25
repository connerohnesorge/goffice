package compare

import (
	"encoding/json"
)

// ToJSON serializes the comparison result to JSON.
func (r *ComparisonResult) ToJSON() ([]byte, error) {
	return json.MarshalIndent(r, "", "  ")
}
