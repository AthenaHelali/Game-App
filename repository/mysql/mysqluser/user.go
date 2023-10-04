package mysqluser

import (
	"database/sql"
	"fmt"
	"game-app/entity"
	"game-app/pkg/errormessage"
	"game-app/pkg/richerror"
	"game-app/repository/mysql"
)

func (d *DB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
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
func (d *DB) RegisterUser(user entity.User) (entity.User, error) {
	res, err := d.conn.Connection().Exec(`insert into users(name, phone_number, password, role) values (?, ? , ?, ?)`, user.Name, user.PhoneNumber, user.Password, user.Role.String())
	if err != nil {
		return entity.User{}, fmt.Errorf("can't execute command: %W", err)
	}

	id, _ := res.LastInsertId()
	user.ID = uint(id)

	return user, nil

}

func (d *DB) GetUserByPhoneNumber(phoneNumber string) (entity.User, error) {
	const op = "mysql.GetUserByPhoneNumber"
	row := d.conn.Connection().QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgNotFound).WithKind(richerror.KindNotFound)

		}

		//TODO log unexpected error for better observability

		return entity.User{}, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return user, nil
}

func (d *DB) GetUserByID(UserID uint) (entity.User, error) {
	const op = "mysql.GetUserByID"
	row := d.conn.Connection().QueryRow(`select * from users where id = ?`, UserID)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgNotFound).WithKind(richerror.KindNotFound)

		}
		return entity.User{}, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return user, nil
}

func scanUser(scanner mysql.Scanner) (entity.User, error) {
	var user entity.User

	var roleStr string

	err := scanner.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.CreatedAt, &user.Password, &roleStr)

	user.Role = entity.MapToRoleEntity(roleStr)
	return user, err

}
