package mysql

import (
	"database/sql"
	"fmt"
	"game-app/entity"
)

func (d MysqlDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	user := entity.User{}
	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)

	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, fmt.Errorf("can't scan query result: %w", err)
	}
	return false, nil
}
func (d MysqlDB) RegisterUser(user entity.User) (entity.User, error) {
	res, err := d.db.Exec(`insert into users(name, phone_number) values (?, ?)`, user.Name, user.PhoneNumber)
	if err != nil {
		return entity.User{}, fmt.Errorf("can't execute command: %W", err)
	}

	id, _ := res.LastInsertId()
	user.ID = uint(id)

	return user, nil

}
