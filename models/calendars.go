package models

type Calendar struct {
	ID       int            `json:"id"`
	UserID   int            `json:"userID"`
	Title    string         `json:"title"`
	Password string         `json:"-"`
	Items    []CalendarItem `json:"items"`
}

type CalendarAccess struct {
	ID       int  `json:"id"`
	Unlocked bool `json:"unlocked"`
}

var (
	calendars      = make(map[int]Calendar) // calID -> Calendar
	accessSessions = make(map[int]string)   // calID -> verified
	currentCalID   int
	currentItemID  int
)

type ContentType int

const (
	String ContentType = iota
	Image
	Link
)

var contentType = map[ContentType]string{
	String: "string",
	Image:  "image",
	Link:   "link",
}

type CalendarItem struct {
	ID          int         `json:"id"`
	CalendarID  int         `json:"calendarID"`
	Day         int         `json:"day"`
	Content     string      `json:"content"`
	ContentType ContentType `json:"contentType"`
	IsOpened    bool        `json:"isOpened"`
}
