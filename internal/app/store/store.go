package store

type Store interface {
	 GetPoll() PollRepository
}