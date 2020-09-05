package ras

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

// APIClient - struct client
type APIClient struct {
	client *http.Client
	Env    string
	Key    string
	Secret string
}

// Error - struct error
type Error struct {
	COD      string `json:"COD"`
	Mensagem string `json:"Mensagem"`
}

//New - create a new client
func New(Key, Secret, env string) *APIClient {
	return &APIClient{
		client: &http.Client{Timeout: 60 * time.Second},
		Env:    env,
		Key:    Key,
		Secret: Secret,
	}
}

func (client *APIClient) Request(method, action string, query interface{}, out interface{}) (error, *Error) {
	if client.client == nil {
		client.client = &http.Client{Timeout: 60 * time.Second}
	}
	endpoint := fmt.Sprintf("%s/%s", client.devProd(), action)
	q := url.Values{}
	if query != nil {
		queryStruct := structToMap(query)
		for k, v := range queryStruct {
			if v != "" {
				q.Add(k, fmt.Sprintf("%v", v))
			}

		}
	}
	req, err := http.NewRequest(method, endpoint, bytes.NewBufferString(q.Encode()))
	if err != nil {
		return err, nil
	}

	req.Header.Add("KEY", client.Key)
	req.Header.Add("Secret", client.Secret)
	if method == "POST" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	res, err := client.client.Do(req)
	if err != nil {
		return err, nil
	}
	bodyResponse, err := ioutil.ReadAll(res.Body)
	if res.StatusCode > 201 {
		return errors.New(string(bodyResponse)), nil
	}
	if client.Env == "develop" {
		log.Printf("full url request %s \n", endpoint)
		log.Printf("body response %s \n", string(bodyResponse))
	}
	err = json.Unmarshal(bodyResponse, &out)
	if err != nil {
		return err, nil
	}
	if strings.Contains(string(bodyResponse), "ERR-") {
		var errorAPI *Error
		err = json.Unmarshal(bodyResponse, &errorAPI)
		if err != nil {
			return err, nil
		}

		return nil, errorAPI
	}
	return nil, nil
}

//devProd - check type Env
func (client *APIClient) devProd() string {
	if client.Env == "develop" {
		return "http://ras.business/api"
	}
	return "http://ras.business/api"
}

func structToMap(item interface{}) map[string]interface{} {

	res := map[string]interface{}{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("json")
		field := reflectValue.Field(i).Interface()
		if tag != "" && tag != "-" {
			if v.Field(i).Type.Kind() == reflect.Struct {
				res[tag] = structToMap(field)
			} else {
				res[tag] = field
			}
		}
	}
	return res
}
