package main

import (
	"fmt"
	"github.com/ajvb/kala/client"
	"github.com/ajvb/kala/job"
	"net/http"
	"bytes"
	"io/ioutil"
	"time"
	"strings"
)

const kalaUrl = "http://127.0.0.1:8000"
//const kalaUrl = "http://kala.d6mqxyc2pm.eu-central-1.elasticbeanstalk.com/"
const tseUrl = "http://127.0.0.1:8080/"
const timesToRepeat = "R3"
const noRepeat = "R1"
const intervalBetweenRun = "PT10S"

func main() {
	id := createJobviaGo()
	fmt.Println(id)
	//createJobviaHttp()
}

func createJobviaGo() string {
	c := client.New(kalaUrl)
	// create a new job
	remote := job.RemoteProperties{}
	remote.Body = ""
	remote.ExpectedResponseCodes = []int{200, 201}
	m := make(map[string][]string)
	m["Authorization"] = []string{"Token 123"}
	remote.Headers = m
	remote.Method = "GET"
	remote.Timeout = 200
	remote.Url = tseUrl + "api/test222"
	t := time.Now().UTC().Add(time.Duration(2*time.Minute)).Format(time.RFC3339)
	scheduler := strings.Join([]string{noRepeat, t,intervalBetweenRun}, "/")
	body := &job.Job{
		//Schedule: "R0/2017-08-28T19:25:16.828696-07:00/PT10S",
		Schedule: scheduler,
		Name:     "test_job_01",
		//Command:  "bash -c 'date'",
		RemoteProperties: remote,
		JobType: 1,
		Retries: 3,
		Epsilon: "PT20S",
	}
	
	id, err := c.CreateJob(body)
	if err != nil {
		panic(err)
	}
	return id
}

func createJobviaHttp() {
	body := []byte(`
	{
		"epsilon": "PT5S", 
		"command": "bash -c 'date", 
		"name": "test_job_2", 
		"schedule": "R2/2017-09-08T19:25:16.828696-07:00/PT10S"
	}`)
	url := kalaUrl + "/api/v1/job/"
	fmt.Println(url)
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	fmt.Println("response Status:", res.Status)
	fmt.Println("response Headers:", res.Header)
	fmt.Println("response body:", string(out))
}
