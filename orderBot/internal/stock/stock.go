package model

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Stock struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	Currency string    `json:"currency"`
	Time     time.Time `json:"time"`
}

func NewStock(name string, price float64, currency string) (Stock, error) {
	s := Stock{}

	if err := s.SetName(name); err != nil {
		return nil, err
	}

	if err := s.SetCurrency(currency); err != nil {
		return nil, err
	}

	if err := s.SetPrice(price); err != nil {
		return nil, err
	}

	if err := s.SetTime(); err != nil {
		return nil, err
	}

	s.Id = uuid.New()
	return &s, nil
}

func (s *Stock) Update(name string, price float64, currency string) {
	if currency == "" {
		currency = "usd"
	}

	s.Name = name
	s.Price = price
	s.Currency = currency
}

func (s *Stock) SetCurrency(currency string) error {
	if len(currency) == 0 || len(currency) > 5 {
		return fmt.Errorf("bad currency <%v>", currency)
	}

	if currency == "" {
		currency = "usd"
	}

	s.Currency = currency
	return nil
}

func (s *Stock) SetName(name string) error {
	if len(name) == 0 || len(name) > 10 {
		return fmt.Errorf("bad name <%v>", name)
	}

	s.Name = name
	return nil
}

func (s *Stock) SetPrice(price float64) error {
	if price < 0 {
		return fmt.Errorf("price < 0 <%v>", price)
	}
	s.Price = price
	return nil
}

func (s *Stock) SetTime() error {
	s.Time = time.Now()
	return nil
}

func (s *Stock) String() string {
	return fmt.Sprintf(
		"%d: %s - %.3f %s; Time: %v-%v-%v",
		s.Id,
		s.Name,
		s.Price,
		s.Currency,
		s.Time.Year(),
		s.Time.Month(),
		s.Time.Day(),
	)
}

func (s *Stock) Marshal() ([]byte, error) {
	return json.Marshal(s)
}
