package mysql

import (
	"database/sql"
	"fmt"
	"game-app/entity"
)

func (d *MysqlDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	user := entity.User{}
	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)

	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password)
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
	var createdAt []uint8
	user := entity.User{}
	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)

	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, false, nil
		}

		return entity.User{}, false, fmt.Errorf("can't scan query result: %w", err)
	}
	return user, true, nil
}
