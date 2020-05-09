package models

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
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

func AddUser(firstname string, lastname string) (int64, error) {
	res, err := db.Exec("INSERT INTO Users (firstname, lastname) VALUES ($1, $2)", firstname, lastname)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
