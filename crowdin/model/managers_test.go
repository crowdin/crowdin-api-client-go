package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManagerListOptionsValues(t *testing.T) {
	tests := []struct {
		name string
		opts *ManagerListOptions
		out  string
	}{
		{
			name: "nil options",
			opts: nil,
		},
		{
			name: "empty options",
			opts: &ManagerListOptions{},
		},
		{
			name: "with parent ID",
			opts: &ManagerListOptions{TeamIds: 123},
			out:  "parentId=123",
		},
		{
			name: "with list options",
			opts: &ManagerListOptions{ListOptions: ListOptions{Limit: 10, Offset: 5}},
			out:  "limit=10&offset=5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, ok := tt.opts.Values()
			if len(tt.out) > 0 {
				assert.True(t, ok)
				assert.Equal(t, tt.out, val.Encode())
			} else {
				assert.False(t, ok)
				assert.Empty(t, val)
			}
		})
	}
}
