package model

import (
	"database/sql/driver"
	"fmt"
)

type PairPoll struct {
	Name string
	Cnt  int
}

type Poll struct {
	ID       int
	PollVars []PairPoll
}

func (p PairPoll) Value() (driver.Value, error) {
	return fmt.Sprintf("('%s', %d)", p.Name, p.Cnt), nil
}


