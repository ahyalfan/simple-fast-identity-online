package repository

import (
	"context"
	"database/sql"
	"golang_biomtrik_login_fido/domain"

	"github.com/doug-martin/goqu/v9"
)

type userRepositoryImpl struct {
	db *goqu.Database //disini kita coba untuk menggunkan goqu, yaitu query builder
}

func NewUser(con *sql.DB) domain.UserRepository {
	return &userRepositoryImpl{
		// melakukan connection dengan goqu, agar bisa bikin query builder
		db: goqu.New("default", con),
	}
}

func (u *userRepositoryImpl) Save(ctx context.Context, user *domain.User) error {
	executor := u.db.Insert("users").Rows(user).Executor()
	_, err := executor.ExecContext(ctx) // karena ini insert maka kita tidak butuh respons dari insert
	return err
}

func (u *userRepositoryImpl) FindByDeviceId(ctx context.Context, id string) (user domain.User, err error) {
	println("asdas")
	print(id)
	dataset := u.db.From("users").Where(goqu.C("device_id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &user) // kita kembalikan pointer
	return
	// kenapa return kita kosong, ini disebakan kita langusng inisialisai di atas yg mana itu akan dijadika variabel return value
	// yg aman kita mengisinya dengan user pointer dan err
}
