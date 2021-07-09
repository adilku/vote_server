package model

import "testing"

func TestPoll(t *testing.T) *Poll {
	t.Helper()
	return &Poll{
		PollVars: []PairPoll{
			{Name: "1", Cnt: 0},
		},
	}
}
