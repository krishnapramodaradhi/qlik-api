package model

import "database/sql"

type Utility struct {
	Id                  int     `json:"id"`
	Region              string  `json:"region"`
	Category            string  `json:"category"`
	Role                string  `json:"role"`
	UserId              string  `json:"userId"`
	UserName            string  `json:"userName"`
	ScreenId            string  `json:"screenId"`
	ScreenName          string  `json:"screenName"`
	ActivityDate        string  `json:"activityDate"`
	StartTime           string  `json:"startTime"`
	SessionType         string  `json:"sessionType"`
	Duration            int     `json:"duration"`
	SessionEndTime      string  `json:"sessionEndTime"`
	ParallelScreenCount int     `json:"parallelScreenCount"`
	Overlap             string  `json:"overlap"`
	TimeBreaks          float64 `json:"timeBreaks"`
}

func (u Utility) ScanRows(rows *sql.Rows) ([]Utility, error) {
	utilities := make([]Utility, 0, 20)
	for rows.Next() {
		if err := rows.Scan(&u.Id, &u.Region, &u.Category, &u.Role, &u.UserId, &u.UserName, &u.ScreenId, &u.ScreenName, &u.ActivityDate, &u.StartTime, &u.SessionType, &u.Duration, &u.SessionEndTime, &u.ParallelScreenCount, &u.Overlap, &u.TimeBreaks); err != nil {
			return nil, err
		}
		utilities = append(utilities, u)
	}
	return utilities, nil
}
