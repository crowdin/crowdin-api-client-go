package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflowTemplatesListOptionsValues(t *testing.T) {
	tests := []struct {
		name string
		opts *WorkflowTemplatesListOptions
		out  string
	}{
		{
			name: "nil options",
			opts: nil,
		},
		{
			name: "empty options",
			opts: &WorkflowTemplatesListOptions{},
		},
		{
			name: "with group ID = 0",
			opts: &WorkflowTemplatesListOptions{GroupID: toPtr(0)},
			out:  "groupId=0",
		},
		{
			name: "with group ID = 1",
			opts: &WorkflowTemplatesListOptions{GroupID: toPtr(1)},
			out:  "groupId=1",
		},
		{
			name: "with all options",
			opts: &WorkflowTemplatesListOptions{GroupID: toPtr(4), ListOptions: ListOptions{Offset: 1, Limit: 10}},
			out:  "groupId=4&limit=10&offset=1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, ok := tt.opts.Values()
			if len(tt.out) > 0 {
				assert.True(t, ok)
				assert.Equal(t, tt.out, v.Encode())
			} else {
				assert.False(t, ok)
				assert.Empty(t, v)
			}
		})
	}
}
