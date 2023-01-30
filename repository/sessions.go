package repository

import (
	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type SessionsRepository struct {
	db db.DB
}

func NewSessionsRepository(db db.DB) SessionsRepository {
	return SessionsRepository{db}
}

func (u *SessionsRepository) ReadSessions() ([]model.Session, error) {
	records, err := u.db.Load("sessions")
	if err != nil {
		return nil, err
	}

	var listSessions []model.Session
	err = json.Unmarshal([]byte(records), &listSessions)
	if err != nil {
		return nil, err
	}

	return listSessions, nil
}

func (u *SessionsRepository) DeleteSessions(tokenTarget string) error {
	listSessions, err := u.ReadSessions()
	if err != nil {
		return err
	}
	list := []model.Session{}
	for _, i := range listSessions{
		if i.Token != tokenTarget {
			list = append(list, i)
		}
	}
	// Select target token and delete from listSessions
	// TODO: answer here

	jsonData, err := json.Marshal(list)
	if err != nil {
		return err
	}

	err = u.db.Save("sessions", jsonData)
	if err != nil {
		return err
	}

	return nil
}

func (u *SessionsRepository) AddSessions(session model.Session) error {
	newData, _ := u.ReadSessions()
	newData = append(newData, session)
	jsonData, err := json.Marshal(newData)
	if err != nil {
		panic(err)
	}
	save := u.db.Save("sessions", jsonData)
	return save // TODO: replace this
}

func (u *SessionsRepository) CheckExpireToken(token string) (model.Session, error) {
	exist, err := u.TokenExist(token)
	if err != nil {
		return model.Session{}, err
	}
	expired := u.TokenExpired(exist)
	if expired {
		return model.Session{}, errors.New("Token is Expired!")
	}
	return exist, nil // TODO: replace this
}

func (u *SessionsRepository) ResetSessions() error {
	err := u.db.Reset("sessions", []byte("[]"))
	if err != nil {
		return err
	}

	return nil
}

func (u *SessionsRepository) TokenExist(req string) (model.Session, error) {
	listSessions, err := u.ReadSessions()
	if err != nil {
		return model.Session{}, err
	}
	for _, element := range listSessions {
		if element.Token == req {
			return element, nil
		}
	}
	return model.Session{}, fmt.Errorf("Token Not Found!")
}

func (u *SessionsRepository) TokenExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}
