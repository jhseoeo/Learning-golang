package main

import (
	"errors"
	"fmt"
)

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
	Err     error
}

func (se StatusErr) Error() string {
	return se.Message
}

func (se StatusErr) Unwrap() error {
	return se.Err
}

func login(uid, pwd string) error {
	// perform authentification
	return errors.New("hi")
}

func getData(file string) ([]byte, error) {
	// data reading
	return []byte(""), errors.New("there")
}

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("Invalid Credentials for user %s", uid),
			Err:     err,
		}
	}

	data, err := getData(file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("file %s is not found", file),
			Err:     err,
		}
	}

	return data, nil
}
