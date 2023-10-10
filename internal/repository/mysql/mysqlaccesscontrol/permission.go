package mysqlaccesscontrol

import (
	"game-app/internal/entity"
	"game-app/internal/repository/mysql"
)

func scanPermission(scanner mysql.Scanner) (entity.Permission, error) {
	var p entity.Permission
	err := scanner.Scan(&p.ID, &p.Title, &p.CreatedAt)
	return p, err
}
