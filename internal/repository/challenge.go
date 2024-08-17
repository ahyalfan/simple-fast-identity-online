package repository

import (
	"context"
	"database/sql"
	"golang_biomtrik_login_fido/domain"

	"github.com/doug-martin/goqu/v9"
)

type challageRepositoryImpl struct {
	db *goqu.Database
}

func NewChallenge(con *sql.DB) domain.ChallegeRepository {
	return &challageRepositoryImpl{
		db: goqu.New("default", con),
	}
}

func (c *challageRepositoryImpl) Save(ctx context.Context, challenge *domain.Challenge) error {
	executor := c.db.Insert("challenge").Rows(challenge).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (c *challageRepositoryImpl) Update(ctx context.Context, challenge *domain.Challenge) error {
	executor := c.db.Update("challenge").Set(challenge).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (c *challageRepositoryImpl) FindById(ctx context.Context, id string) (challenge domain.Challenge, err error) {
	dataset := c.db.From("challenge").Where(goqu.C("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &challenge)
	return
	// kenapa return kita kosong, ini disebakan kita langusng inisialisai di atas yg mana itu akan dijadika variabel return value
	// yg aman kita mengisinya dengan challeng pointer dan err
}
