package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTodo_IsOverDue(t *testing.T) {
	type fields struct {
		complete bool
		deadline time.Time
	}

	tests := []struct {
		name   string
		fields fields
		expect bool
	}{
		{
			name: "not overdue",
			fields: fields{
				complete: false,
				deadline: time.Now().Add(1 * time.Hour),
			},
			expect: false,
		},
		{
			name: "overdue",
			fields: fields{
				complete: false,
				deadline: time.Now().Add(-1 * time.Hour),
			},
			expect: true,
		},
		{
			name: "completed",
			fields: fields{
				complete: true,
				deadline: time.Now().Add(1 * time.Hour),
			},
			expect: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := &Todo{
				complete: test.fields.complete,
				deadline: test.fields.deadline,
			}
			// if result = e.IsOverdue (false if e complete, true if time.Now after deadline) , result != test.expect
			if result := e.IsOverdue(); result != test.expect {
				t.Errorf("IsOverdue() = %v, want %v", result, test.expect)
			}
		})
	}
}

func TestTodo_ExtendDeadline(t *testing.T) {
	now := time.Now()
	type fields struct {
		complete bool
		deadline time.Time
	}

	type args struct {
		addition time.Duration
	}

	tests := []struct {
		name             string
		fields           fields
		args             args
		expectedDeadline time.Time
	}{
		{
			name: "extended",
			fields: fields{
				complete: false,
				deadline: now,
			},
			args: args{
				addition: 1 * time.Hour,
			},
			expectedDeadline: now.Add(1 * time.Hour),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := &Todo{
				complete: test.fields.complete,
				deadline: test.fields.deadline,
			}
			e.ExtendDealine(test.args.addition)
			assert.Equal(t, test.expectedDeadline, e.deadline)
		})
	}
}
