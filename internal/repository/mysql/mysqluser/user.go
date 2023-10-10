package mysqluser

import (
	"context"
	"database/sql"
	"fmt"
	entity2 "game-app/internal/entity"
	"game-app/internal/pkg/errormessage"
	"game-app/internal/pkg/richerror"
	"game-app/internal/repository/mysql"
)

func (d *DB) IsPhoneNumberUnique(ctx context.Context, phoneNumber string) (bool, error) {
	row := d.conn.Connection().QueryRow(`select * from users where phone_number = ?`, phoneNumber)

	_, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, richerror.New("mysql.IsPhoneNumberUnique").WithError(err).WithMessage(errormessage.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return false, nil
}
func (d *DB) RegisterUser(ctx context.Context, user entity2.User) (entity2.User, error) {
	res, err := d.conn.Connection().Exec(`insert into users(name, phone_number, password, role) values (?, ? , ?, ?)`, user.Name, user.PhoneNumber, user.Password, user.Role.String())
	if err != nil {
		return entity2.User{}, fmt.Errorf("can't execute command: %W", err)
	}

	id, _ := res.LastInsertId()
	user.ID = uint(id)

	return user, nil

}

func (d *DB) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (entity2.User, error) {
	const op = "mysql.GetUserByPhoneNumber"
	row := d.conn.Connection().QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity2.User{}, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgNotFound).WithKind(richerror.KindNotFound)

		}

		//TODO log unexpected error for better observability

		return entity2.User{}, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return user, nil
}

func (d *DB) GetUserByID(ctx context.Context, UserID uint) (entity2.User, error) {
	const op = "mysql.GetUserByID"
	row := d.conn.Connection().QueryRow(`select * from users where id = ?`, UserID)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity2.User{}, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgNotFound).WithKind(richerror.KindNotFound)

		}
		return entity2.User{}, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return user, nil
}

func scanUser(scanner mysql.Scanner) (entity2.User, error) {
	var user entity2.User

	var roleStr string

	err := scanner.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.CreatedAt, &user.Password, &roleStr)

	user.Role = entity2.MapToRoleEntity(roleStr)
	return user, err

}
