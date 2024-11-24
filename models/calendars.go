package models

type Calendar struct {
	ID     int            `json:"id"`
	UserID int            `json:"userID"`
	Title  string         `json:"title"`
	Items  []CalendarItem `json:"items"`
}

type ContentType int

const (
	String ContentType = iota
	Image
	Link
)

type CalendarItem struct {
	ID          int         `json:"id"`
	CalendarID  int         `json:"calendarID"`
	Day         int         `json:"day"`
	Content     string      `json:"content"`
	ContentType ContentType `json:"contentType"`
	IsOpened    bool        `json:"isOpened"`
}
