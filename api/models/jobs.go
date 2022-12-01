package models

import (
	"database/sql"
	"fmt"
	"log"
)

type Job struct {
	ID               int     `json:"id"`
	Title            string  `json:"title"`
	Description      string  `json:"description"`
	ShortDescription string  `json:"shortDescription"`
	UserID           string  `json:"userId"`
	AskingPrice      int64   `json:"askingPrice"`
	Location         float32 `json:"location"`
	Username         string  `json:"username"`
	Email            string  `json:"email"`
}

type ShippingLocation struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type JobModel struct {
	DB *sql.DB
}

func (jm *JobModel) FindOne(id int64) (Job, error) {
	var job Job
	query := `SELECT jobs.title,
       jobs.id,
       jobs.description,
       jobs.short_description as shortDescription,
       jobs.asking_price      as askingPrice,
       jobs.location,
       jobs.user_id           as userId,
       users.username,
       users.email FROM jobs INNER JOIN users WHERE jobs.id = ? AND users.id = jobs.user_id;`

	err := jm.DB.QueryRow(query, id).Scan(&job.Title, &job.ID, &job.Description, &job.ShortDescription, &job.AskingPrice, &job.Location, &job.UserID, &job.Username, &job.Email)
	if err != nil {
		log.Println(err.Error())
		return Job{}, err
	}

	return job, nil
}

func (jm *JobModel) FindAll() []Job {
	query := `SELECT title, 
                     id,
                     description,
                     short_description AS shortDescription,
                     asking_price AS askingPrice,
                     location,
                     user_id as userId
              FROM jobs`

	res, err := jm.DB.Query(query)
	defer res.Close()

	if err != nil {
		log.Printf("Failed to find all jobs. Error: %s", err)
	}

	jobs := []Job{}
	for res.Next() {
		var job Job
		err := res.Scan(&job.Title, &job.ID, &job.Description, &job.ShortDescription, &job.AskingPrice, &job.Location, &job.UserID)
		if err != nil {
			log.Printf("Failed to scan jobs. Error: %s", err)
		}
		jobs = append(jobs, job)
	}

	fmt.Println("RAW QUERY:\n", jobs)

	return jobs
}