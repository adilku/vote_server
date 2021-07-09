package store

import "github.com/adilku/vote_server/internal/app/model"

type PollRepository interface {
	Create(poll *model.Poll) error
	FindById(id int) (*model.Poll, error)
	//ChangeById(id int, poll *model.Poll) error
}