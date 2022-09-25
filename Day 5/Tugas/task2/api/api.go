package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Description struct {
	Description string `json:"description"`
}

type Job struct {
	Title      string `json:"title"`
	WorkFrom   string `json:"work_from"`
	Department string `json:"department"`
}

type Aggregation struct {
	Jobs []Job
	Description
}

func fetch(url string, data interface{}) error {
	var err error
	var response *http.Response
	var buffer []byte

	response, err = http.Get(url)
	buffer, err = io.ReadAll(response.Body)

	err = json.Unmarshal(buffer, data)

	return err
}

type Cache struct {
	Aggregation
	IsValid bool
}

func (c *Cache) Aggregate() (Aggregation, error) {
	defer caculateTime(time.Now())

	c.IsValid = !c.IsValid

	var err error
	if c.IsValid {
		descChannel := make(chan Description)
		descErrChannel := make(chan error)

		go func(descChan chan Description, descErrChan chan error) {
			var data Description
			err := fetch("https://workspace-rho.vercel.app/api/description", &data)

			descChan <- data
			descErrChan <- err
		}(descChannel, descErrChannel)

		jobsChannel := make(chan []Job)
		jobsErrChannel := make(chan error)
		go func(jobsChan chan []Job, jobsErrChan chan error) {
			var jobs []Job
			err := fetch("https://workspace-rho.vercel.app/api/jobs", &jobs)

			jobsChan <- jobs
			jobsErrChan <- err
		}(jobsChannel, jobsErrChannel)

		desc := <-descChannel
		descErr := <-descErrChannel
		jobs := <-jobsChannel
		jobsErr := <-jobsErrChannel

		if descErr != nil {
			err = descErr
		}
		if jobsErr != nil {
			err = jobsErr
		}

		c.Aggregation = Aggregation{
			Description: desc,
			Jobs:        jobs,
		}
	}

	return c.Aggregation, err
}

func caculateTime(start time.Time) {
	fmt.Println("Fetch Data Succeed At :", start)
	fmt.Println("Fetch Time :", time.Since(start))
	fmt.Println("==================================================================================")
}
