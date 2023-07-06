package albohttp

import (
	"bytes"
	"net/http"
	"time"
)

type FailureResponse struct {
	Success bool   `json:"success" description:"response status" example:"false"`
	Message string `json:"message" description:"response information" example:"failed"`
}

func Failure(msg string) FailureResponse {
	return FailureResponse{
		Success: false,
		Message: msg,
	}
}

type Options struct {
	Endpoint string
	Body     []byte
	Method   string
}

func NewClient(timeout int) *Request {
	return &Request{
		Client: &http.Client{
			Timeout: time.Duration(10) * time.Second,
		},
	}
}

// Request struct
type Request struct {
	Client *http.Client
}

func (r *Request) MakeRequest(o *Options) (*http.Response, error) {
	req, err := http.NewRequest(o.Method, o.Endpoint, bytes.NewBuffer(o.Body))
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Content-Type", "application/json")
	res, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
