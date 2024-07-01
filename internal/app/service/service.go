package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Response struct {
	Code int
	Data interface{}
}

type JobsData struct {
	Jobs []struct {
		URL    string `json:"alternate_url"`
		Name   string `json:"name"`
		Salary struct {
			Currency string `json:"currency"`
			From     int    `json:"from"`
			To       int    `json:"to"`
		} `json:"salary"`
	} `json:"items"`
}

type Service struct{}

func New() *Service {
	return &Service{}
}

func generateErrorResponse() Response {
	var respText struct {
		Data string `json:"data"`
	}
	respText.Data = "server error occured"
	return Response{
		Code: 500,
		Data: respText,
	}
}

func (s *Service) GetStatus() Response {
	var respText struct {
		Data string `json:"data"`
	}

	respText.Data = "server currently working"
	return Response{
		Code: 200,
		Data: respText,
	}
}

func (s *Service) GetVacancies() Response {
	u, _ := url.Parse("https://api.hh.ru/vacancies")
	q := u.Query()
	q.Add("text", "Курьер")
	q.Add("currency", "RUR")
	u.RawQuery = q.Encode()
	fmt.Println(u)

	resp, err := http.Get(u.String())
	if err != nil {
		return generateErrorResponse()
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return generateErrorResponse()
	}
	data := JobsData{}
	err = json.Unmarshal(body, &data)

	fmt.Println(data.Jobs[0])
	if err != nil {
		return generateErrorResponse()
	}
	return Response{
		Code: 200,
		Data: data,
	}
}
