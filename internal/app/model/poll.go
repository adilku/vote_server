package model

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

type PairPoll struct {
	Name string
	Cnt  int
}

type Poll struct {
	ID       int        `json:"id"`
	PollVars []PairPoll `json:"poll_vars"`
}


func (p PairPoll) Value() (driver.Value, error) {
	return fmt.Sprintf("('%s', %d)", p.Name, p.Cnt), nil
}

func FillNulls(l []string) []PairPoll {
	a := make([]PairPoll, 0)
	for _, i := range l {
		a = append(a, PairPoll{i, 0})
	}
	return a
}

func (p *PairPoll) Scan(src interface{}) error {
	s := string(src.([]byte))
	s = strings.Trim(s, "(")
	s = strings.Trim(s, ")")
	s = strings.Replace(s, "'", "", -1)
	arr := strings.Split(s, ",")
	p.Name = arr[0]
	cnt, err := strconv.Atoi(arr[1])
	if err != nil {
		return err
	}
	p.Cnt = cnt
	return err
}


