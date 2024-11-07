package database

import "time"

var UserStories []UserStory
var Users []User

type User struct {
	Id         string
	FName      string
	LName      string
	Email      string
	Password   string
	Role       string
	CustomerId string
	CreatedAt  time.Time
	ModifiedAt time.Time
}
type Session struct {
	SessionId string
	UserId    string
	EmailId   string
	IpAddress string
	UserAgent string
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiresAt time.Time
}
type UserStory struct {
	Number     int
	Title      string
	Due        time.Time
	Assignee   string
	AssigneeID string
	CreatedAt  time.Time
	ModifiedAt time.Time
	WorkLogs   []WorkLog
}

type WorkLog struct {
	UserStoryID  int
	Number       int
	Title        string
	Assignee     string
	AssigneeID   string
	Status       string
	TotalTime    time.Time
	ExpectedTime time.Time
	ActualTime   time.Time
	CreatedAt    time.Time
	ModifiedAt   time.Time
	Commands     string
	SubWorkLogs  []SubWorkLog
}

type SubWorkLog struct {
	Number       int
	Title        string
	Assignee     string
	AssigneeID   string
	TaskType     string
	Status       string
	TotalTime    time.Time
	ExpectedTime time.Time
	ActualTime   time.Time
	CreatedAt    time.Time
	ModifiedAt   time.Time
	Commands     string
}
