package models

type Training struct {
	ID       int64  `json:"id"`
	UserID   int64  `json:"userId"`
	Date     string `json:"date"`
	Activity string `json:"activity"`
}

// AddTraining adds training
func AddTraining(training Training) (int64, error) {
	res, err := db.Exec("INSERT INTO Training (userId, date, activity) VALUES ($1, $2, $3)", training.UserID, training.Date, training.Activity)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func TrainingByUserID(training Training) ([]*Training, error) {
	rows, err := db.Query("SELECT * FROM Training WHERE userId = ($1)", training.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usrTraining := make([]*Training, 0)
	for rows.Next() {
		train := new(Training)
		err := rows.Scan(&train.ID, &train.UserID, &train.Date, &train.Activity)
		if err != nil {
			return nil, err
		}
		usrTraining = append(usrTraining, train)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return usrTraining, nil
}
