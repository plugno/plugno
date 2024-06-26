package models

import (
	"database/sql"
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
	LocationName     string  `json:"locationName"`
	Username         string  `json:"username"`
	Email            string  `json:"email"`
	CreatedAt        string  `json:"createdAt"`
}

type ShippingLocation struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type JobObject struct {
	Title            string
	ShortDescription string
	Description      string
	AskingPrice      int
	UserID           int
}

type PlugJobObject struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Place       string `json:"place"`
	UserID      string `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
	IsAccepted  bool   `json:"isAccepted"`
	RequestType string `json:"requestType"`
	Username    string `json:"username,omitempty"`
	Avatar      string `json:"avatar"`
}

type AcceptedJobObject struct {
	JobID       int    `json:"jobId"`
	PlugID      int    `json:"plugId"`
	Status      string `json:"status"`
	UpdatedAt   string `json:"updatedAt"`
	CreatedAt   string `json:"createdAt"`
	UserID      string `json:"userId"`
	Title       string `json:"title"`
	RequestType string `json:"requestType"`
	PlugAvatar  string `json:"plugAvatar"`
}

type NewPlugJobObject struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Place       string `json:"place"`
	UserID      string `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
	RequestType string `json:"requestType"`
}

type PlugAcceptObject struct {
	JobID  int `json:"jobId"`
	PlugID int `json:"plugId"`
}

type ActiveJobObject struct {
	ID            int    `json:"id"`
	JobID         int    `json:"jobId"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
	Status        string `json:"status"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Username      string `json:"username"`
	Avatar        string `json:"avatar"`
	RequestStatus string `json:"requestStatus"`
}

type JobModel struct {
	DB *sql.DB
}

func (model *JobModel) FindOne(id int64) (Job, error) {
	var job Job
	query := `SELECT jobs.title,
       jobs.id,
       jobs.description,
       jobs.short_description as shortDescription,
       jobs.asking_price      as askingPrice,
       jobs.user_id           as userId,
       users.username,
       users.email FROM jobs INNER JOIN users WHERE jobs.id = ? AND users.id = jobs.user_id;`

	err := model.DB.QueryRow(query, id).Scan(&job.Title, &job.ID, &job.Description, &job.ShortDescription, &job.AskingPrice, &job.UserID, &job.Username, &job.Email)
	if err != nil {
		log.Println(err.Error())
		return Job{}, err
	}

	return job, nil
}

func (jm *JobModel) FindAll() ([]Job, error) {
	query := `SELECT title, 
                     id,
                     description,
                     short_description AS shortDescription,
                     asking_price AS askingPrice,
                     user_id as userId,
                     created_at as createdAt,
                     location_name as locationName
              FROM jobs ORDER BY createdAt DESC`

	res, err := jm.DB.Query(query)
	if err != nil {
		log.Printf("Failed to find all jobs. Error: %s", err)
		return nil, err
	}

	defer res.Close()

	jobs := []Job{}
	for res.Next() {
		var job Job
		err := res.Scan(&job.Title, &job.ID, &job.Description, &job.ShortDescription, &job.AskingPrice, &job.UserID, &job.CreatedAt, &job.LocationName)
		if err != nil {
			log.Printf("Failed to scan jobs. Error: %s", err)
			return nil, err
		}
		jobs = append(jobs, job)
	}

	return jobs, nil
}

func (jm *JobModel) Create(jobObject *JobObject) error {
	query := `INSERT INTO jobs (title, short_description, description, asking_price, user_id) VALUES (?, ?, ?, ?, ?)`
	_, err := jm.DB.Exec(query, jobObject.Title, jobObject.ShortDescription, jobObject.Description, jobObject.AskingPrice, jobObject.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (model *JobModel) CreatePlugJob(jobObject *NewPlugJobObject) error {
	query := `INSERT INTO plug_jobs (title, description, place, phone_number, user_id, request_type) VALUES (?, ?, ?, ?, ?, ?)`

	_, err := model.DB.Exec(query, jobObject.Title, jobObject.Description, jobObject.Place, jobObject.PhoneNumber, jobObject.UserID, jobObject.RequestType)
	if err != nil {
		return err
	}

	return nil
}

func (model *JobModel) FindAllPlugJobs(userId string) ([]PlugJobObject, error) {
	query := "SELECT pj.*, u.username, u.avatar FROM plug_jobs pj JOIN users u on pj.user_id = u.id"

	rows, err := model.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	plugJobs := []PlugJobObject{}
	for rows.Next() {
		var plugJob PlugJobObject

		err := rows.Scan(&plugJob.ID, &plugJob.Title, &plugJob.Description, &plugJob.Place, &plugJob.PhoneNumber, &plugJob.UserID, &plugJob.IsAccepted, &plugJob.RequestType, &plugJob.Username, &plugJob.Avatar)
		if err != nil {
			return nil, err
		}

		plugJobs = append(plugJobs, plugJob)
	}

	return plugJobs, nil
}

func (model *JobModel) CreateAcceptedPlugJob(jobAcceptObject *PlugAcceptObject) error {
	query := "INSERT INTO accepted_jobs (job_id, plug_id, status) VALUES (?, ?, ?)"

	_, err := model.DB.Exec(query, jobAcceptObject.JobID, jobAcceptObject.PlugID, "accepted")
	if err != nil {
		return err
	}

	return nil
}

func (model *JobModel) GetActiveJob(jobId int) (*ActiveJobObject, error) {
	var activeJob ActiveJobObject

	query := `select aj.id,
       aj.job_id,
       aj.created_at,
       aj.updated_at,
       aj.status,
       pj.title,
       pj.description,
       aj.request_status,
       u.username,
       u.avatar
    from accepted_jobs aj
         join plug_jobs pj on aj.job_id = ?
         join users u on aj.plug_id = u.id
    where pj.is_accepted = 1`

	err := model.DB.QueryRow(query, jobId).Scan(
		&activeJob.ID,
		&activeJob.JobID,
		&activeJob.CreatedAt,
		&activeJob.UpdatedAt,
		&activeJob.Status,
		&activeJob.Title,
		&activeJob.Description,
		&activeJob.RequestStatus,
		&activeJob.Username,
		&activeJob.Avatar,
	)
	if err != nil {
		return nil, err
	}

	return &activeJob, nil
}

func (model *JobModel) FindAllAcceptedJobs(userId string) ([]AcceptedJobObject, error) {
	query := `select aj.job_id,
       aj.plug_id,
       aj.status,
       aj.updated_at,
       aj.created_at,
       pj.user_id,
       pj.title,
       pj.request_type,
       u.avatar as plugAvatar
        from accepted_jobs aj
             join plug_jobs pj on aj.job_id = pj.id
            join users u on aj.plug_id = u.id
        where pj.user_id = ?
    `

	var acceptedJobObject AcceptedJobObject

	rows, err := model.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}

	acceptedJobs := []AcceptedJobObject{}
	for rows.Next() {
		err := rows.Scan(&acceptedJobObject.JobID, &acceptedJobObject.PlugID, &acceptedJobObject.Status, &acceptedJobObject.UpdatedAt, &acceptedJobObject.CreatedAt, &acceptedJobObject.UserID, &acceptedJobObject.Title, &acceptedJobObject.RequestType, &acceptedJobObject.PlugAvatar)
		if err != nil {
			return nil, err
		}

		acceptedJobs = append(acceptedJobs, acceptedJobObject)
	}

	return acceptedJobs, nil
}
