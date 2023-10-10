package mysqlaccesscontrol

import (
	entity2 "game-app/internal/entity"
	"game-app/internal/pkg/errormessage"
	"game-app/internal/pkg/richerror"
	"game-app/internal/pkg/slice"
	"game-app/internal/repository/mysql"
	"strings"
)

func (d *DB) GetUserPermissionTitles(userID uint, role entity2.Role) ([]entity2.PermissionTitle, error) {
	const op = "mysql.GetUserPermissionTitles"

	roleACL := make([]entity2.AccessControl, 0)

	rows, err := d.conn.Connection().Query(`select * from access_controls where actor_type = ? and actor_id = ?`, entity2.RoleActorType, role)
	if err != nil {
		return nil, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgSomeThingWentWrong).WithKind(richerror.KindUnexpected)
	}
	defer rows.Close()
	for rows.Next() {
		acl, err := scanAccessControl(rows)
		if err != nil {
			return nil, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgSomeThingWentWrong).WithKind(richerror.KindUnexpected)
		}
		roleACL = append(roleACL, acl)
	}

	if err := rows.Err(); err != nil {
		return nil, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgSomeThingWentWrong).WithKind(richerror.KindUnexpected)
	}

	userACL := make([]entity2.AccessControl, 0)

	userRows, err := d.conn.Connection().Query(`select * from access_controls where actor_type = ? and actor_id = ?`, entity2.UserActorType, userID)
	if err != nil {
		return nil, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgSomeThingWentWrong).WithKind(richerror.KindUnexpected)
	}

	defer userRows.Close()
	for userRows.Next() {
		acl, err := scanAccessControl(userRows)
		if err != nil {
			return nil, richerror.New(op).WithError(err).WithMessage(errormessage.ErrorMsgSomeThingWentWrong).WithKind(richerror.KindUnexpected)
		}
		userACL = append(userACL, acl)
	}

	if err := userRows.Err(); err != nil {
		return nil, richerror.New(op).WithError(err).
			WithMessage(errormessage.ErrorMsgSomeThingWentWrong).WithKind(richerror.KindUnexpected)
	}

	permissionsIDs := make([]uint, 0)
	for _, r := range roleACL {
		if !slice.DoesExist(permissionsIDs, r.PermissionID) {
			permissionsIDs = append(permissionsIDs, r.PermissionID)
		}
	}

	for _, r := range userACL {
		if !slice.DoesExist(permissionsIDs, r.PermissionID) {
			permissionsIDs = append(permissionsIDs, r.PermissionID)
		}
	}

	if len(permissionsIDs) == 0 {
		return nil, nil
	}

	args := make([]any, len(permissionsIDs))

	for i, id := range permissionsIDs {
		args[i] = id
	}

	query := "select * from permissions where id in (?" +
		strings.Repeat(",?", len(permissionsIDs)-1) + ")"

	pRows, err := d.conn.Connection().Query(query, args...)
	if err != nil {
		return nil, richerror.New(op).WithError(err).
			WithMessage(errormessage.ErrorMsgSomeThingWentWrong).WithKind(richerror.KindUnexpected)
	}
	defer pRows.Close()

	permissionTitles := make([]entity2.PermissionTitle, 0)

	for pRows.Next() {
		permission, err := scanPermission(pRows)
		if err != nil {
			return nil, richerror.New(op).WithError(err).
				WithMessage(errormessage.ErrorMsgSomeThingWentWrong).WithKind(richerror.KindUnexpected)
		}

		permissionTitles = append(permissionTitles, permission.Title)
	}

	if err := pRows.Err(); err != nil {
		return nil, richerror.New(op).WithError(err).
			WithMessage(errormessage.ErrorMsgSomeThingWentWrong).WithKind(richerror.KindUnexpected)

	}
	return permissionTitles, nil
}

func scanAccessControl(scanner mysql.Scanner) (entity2.AccessControl, error) {
	var acl entity2.AccessControl
	err := scanner.Scan(&acl.ID, &acl.ActorID, &acl.ActorType, &acl.PermissionID, &acl.CreatedAt)
	return acl, err

}
