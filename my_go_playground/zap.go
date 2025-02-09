package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type User struct {
}

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id

	body, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	req, err := http.NewRequest(http.MethodPut, fullURL, bytes.NewBuffer(body))
	if err != nil {
		return User{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-Key", apiKey)

	httpClient := http.Client{}

	res, err := httpClient.Do(req)
	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()

	var result User
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return User{}, err
	}

	return result, nil
}

func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := baseURL + "/" + id

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return User{}, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-Key", apiKey)

	httpClient := http.Client{}

	res, err := httpClient.Do(req)
	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()

	var result User
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return User{}, err
	}

	return result, nil
}

func fetchData(url string) (int, error) {
	res, err := http.Get(url)

	if err != nil {

	}
	defer res.Body.Close()

	return res.StatusCode, nil
}
