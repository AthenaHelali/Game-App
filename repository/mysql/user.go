package mysql

import (
	"database/sql"
	"fmt"
	"game-app/entity"
	"game-app/pkg/errormessage"
	"game-app/pkg/richerror"
)

func (d *MysqlDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)

	_, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, fmt.Errorf("can't scan query result: %w", err)
	}
	return false, nil
}
func (d *MysqlDB) RegisterUser(user entity.User) (entity.User, error) {
	res, err := d.db.Exec(`insert into users(name, phone_number, password) values (?, ? , ?)`, user.Name, user.PhoneNumber, user.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("can't execute command: %W", err)
	}

	id, _ := res.LastInsertId()
	user.ID = uint(id)

	return user, nil

}

func (d *MysqlDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error) {
	const op = "mysql.GetUserByPhoneNumber"
	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, false, nil
		}

		return entity.User{}, false, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return user, true, nil
}
func (d *MysqlDB) GetUserByID(UserID uint) (entity.User, error) {
	const op = "mysql.GetUserByID"
	row := d.db.QueryRow(`select * from users where id = ?`, UserID)
	user, err := scanUser(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgNotFound).WithKind(richerror.KindNotFound)

		}
		return entity.User{}, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}
	return user, nil
}

func scanUser(row *sql.Row) (entity.User, error) {
	var user entity.User
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.CreatedAt, &user.Password)
	return user, err

}
