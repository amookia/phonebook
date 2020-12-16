package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

//JUST FOR TEST :D 2020/12/16 20:30
const base_url = "http://localhost:8000/api/v1/"
var respjson map[string]string
var respinter map[string]interface{}
var token string
var id float64


func TestAuth(t *testing.T){
	url := base_url + "register"
	payload := map[string]interface{}{
		"email":"testing@test.com",
		"password":"justfortesting",
	}
	jsonValue, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp,_ := client.Do(req)
	if resp.StatusCode != 200 {
		loginurl := base_url + "login"
		req, _ = http.NewRequest("POST", loginurl, bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, _ := client.Do(req)
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			json.NewDecoder(resp.Body).Decode(&respjson)
			token = respjson["token"]
		}
	}else{
		json.NewDecoder(resp.Body).Decode(&respjson)
		token = respjson["token"]
	}
	fmt.Println(respjson)
	assert.NotNil(t,token)
	assert.Equal(t,200,resp.StatusCode)
}

func TestAddContact(t *testing.T){
	url := base_url + "contact/add"
	payload := map[string]interface{}{
		"name":"Test GOLANG",
		"phone_numbers":[]string{"09031515222","09125632141"},
		"description":"test test test test test",
	}
	jsonValue, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization",token)
	client := &http.Client{}
	resp,_ := client.Do(req)
	json.NewDecoder(resp.Body).Decode(&respinter)
	id64 := respinter["id"]
	id = id64.(float64)
	assert.NotZero(t,id)
	assert.Equal(t,200,resp.StatusCode)
}

func TestUpdateContact(t *testing.T){
	url := fmt.Sprintf("%s/contact/%.0f",base_url,id)
	payload := map[string]interface{}{
		"name":"Update GOLANG",
		"phone_numbers":[]string{"09858585858","0969696969696"},
		"description":"update update update update",
	}
	jsonValue, _ := json.Marshal(payload)
	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization",token)
	client := &http.Client{}
	resp,_ := client.Do(req)
	json.NewDecoder(resp.Body).Decode(&respinter)
	assert.Equal(t,200,resp.StatusCode)
}

