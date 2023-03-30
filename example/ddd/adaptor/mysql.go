package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

func main() {
	fmt.Println("main----------")
	dbConf := "root:root_password@tcp(127.0.0.1:13306)/ddd?charset=utf8mb4"
	db, err := sqlx.Open("mysql", dbConf)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)
	fmt.Println(db.DB)
	fmt.Println("---------db")

	type userEntity struct {
		ID   int    `db:"id"`
		Name string `db:"name"`
	}
	var users []*userEntity
	q := "SELECT id, name FROM users;"
	err = db.Select(&users, q)
	if err != nil {
		log.Fatal(err)
	}

	for _, u := range users {
		fmt.Println(u.ID)
		fmt.Println(u.Name)
	}

	fmt.Println(users)

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	q = "SELECT id, name FROM users WHERE id = 1 FOR UPDATE;"
	err = db.Select(&users, q)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(users[0].Name)

	newName := "hogaaa"

	_, err = tx.Exec("UPDATE users SET name = ? WHERE id = 1;", newName)

	tx.Rollback()
	//tx.Commit()

	repo := NewTransactionRepository(db)

}

// interface
type TransactionRepository interface {
	Transaction(ctx context.Context, f func(tx *sqlx.Tx) error) error
}

// impl
type transactionRepositoryImpl struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	return &transactionRepositoryImpl{
		db: db,
	}
}

// infra
func (t *transactionRepositoryImpl) Transaction(ctx context.Context, f func(tx *sqlx.Tx) error) error {
	tx, err := t.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	err = f(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

type UserRepository interface {
	//Save(tx *sqlx.Tx) error
	Save(ctx context.Context) error
}

type userRepositoryImpl struct {
	db *sqlx.DB
}

func (u *userRepositoryImpl) Save(ctx context.Context) error {
	//_, err := tx.Exec("UPDATE users SET name = \"hoge\" WHERE id = 1")
	tx, err := u.db.BeginTxx(ctx, nil)
	_, err = tx.Exec("UPDATE users SET name = \"hoge\" WHERE id = 1")
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}

}

var userRepository = NewUserRepository(db)

func ChangeName(ctx context.Context, userID int, newName string) error {
	err := transactionRepository.Transaction(ctx, func(tx *sqlx.Tx) error {
		// repositoryに外からtransactionを与える
		u, err := userRepository.Get(ctx, tx, userID)
		if err != nil {
			return err
		}

		err = u.ChangeName(newName)
		if err != nil {
			return err
		}
		// repositoryに外からtransactionを与える
		return userRepository.Save(ctx, tx, e)
	})
	if err != nil {
		return err
	}
	return nil
}
