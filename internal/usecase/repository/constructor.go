package repository

import "github.com/v1adhope/crypto-diary/pkg/postgres"

type Repos struct {
	User     *UserRepo
	Position *PositionRepo
}

func New(pg *postgres.Postgres) *Repos {
	return &Repos{
		User:     NewUser(pg),
		Position: NewPosition(pg),
	}
}
