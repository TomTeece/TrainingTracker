package models

type Climbing struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"userId"`
	Date     string `json:"date"`
	Activity string `json:"activity"`
}

// AddClimbing adds climbing
func AddClimbing(climbing Climbing) (int64, error) {
	res, err := db.Exec("INSERT INTO Climbing (userId, date, activity) VALUES ($1, $2, $3)", climbing.UserID, climbing.Date, climbing.Activity)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func ClimbingByUserID(climbing Climbing) ([]*Climbing, error) {
	rows, err := db.Query("SELECT * FROM Climbing WHERE userId = ($1)", climbing.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usrClimbing := make([]*Climbing, 0)
	for rows.Next() {
		climb := new(Climbing)
		err := rows.Scan(&climb.ID, &climb.UserID, &climb.Date, &climb.Activity)
		if err != nil {
			return nil, err
		}
		usrClimbing = append(usrClimbing, climb)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return usrClimbing, nil
}
