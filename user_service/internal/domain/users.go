package domain

import (
	"time"
)

type userImpl struct {
	UUID        string
	Name        string
	Description string
	BornDay     time.Time
	City        string
	Links       map[string]string
}

type User interface {
	SetUUID(uuid string)
	SetName(name string)
	SetDescription(dis string)
	SetBornDay(day time.Time)
	SetCity(city string)
	SetLinks(links map[string]string)
	AddLink(alias, link string)
	GetUUID() string
	GetName() string
	GetDescription() string
	GetBornDay() time.Time
	GetCity() string
	GetLinks() map[string]string
	GetLink(alias string) string
}

func NewUser() User {
	return &userImpl{}
}

func (u *userImpl) SetUUID(uuid string) {
	u.UUID = uuid
}

func (u *userImpl) SetName(name string) {
	u.Name = name
}

func (u *userImpl) SetDescription(description string) {
	u.Description = description
}

func (u *userImpl) SetBornDay(day time.Time) {
	u.BornDay = day
}

func (u *userImpl) SetCity(city string) {
	u.City = city
}

func (u *userImpl) SetLinks(links map[string]string) {
	u.Links = links
}

func (u *userImpl) AddLink(alias, link string) {
	u.Links[alias] = link
}

func (u *userImpl) GetUUID() string {
	return u.UUID
}

func (u *userImpl) GetName() string {
	return u.Name
}

func (u *userImpl) GetDescription() string {
	return u.Description
}

func (u *userImpl) GetBornDay() time.Time {
	return u.BornDay
}

func (u *userImpl) GetCity() string {
	return u.City
}

func (u *userImpl) GetLinks() map[string]string {
	return u.Links
}

func (u *userImpl) GetLink(alias string) string {
	return u.Links[alias]
}
