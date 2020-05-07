package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func AllUsers() ([]*User, error) {
	rows, err := db.Query("SELECT id, firstname, lastname FROM Users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usrs := make([]*User, 0)
	for rows.Next() {
		usr := new(User)
		err := rows.Scan(&usr.ID, &usr.FirstName, &usr.LastName)
		if err != nil {
			return nil, err
		}
		usrs = append(usrs, usr)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return usrs, nil
}
