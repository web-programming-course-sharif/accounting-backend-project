package otp

type Entries []Entry
type Response struct {
	Return  `json:"return"`
	Entries `json:"entries"`
}
type Return struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
type Entry struct {
	MessageId  int    `json:"messageid"`
	Message    string `json:"message"`
	Status     int    `json:"status"`
	StatusText string `json:"statustext"`
	Sender     string `json:"sender"`
	Receptor   string `json:"receptor"`
	Date       int64  `json:"date"`
	Cost       int    `json:"cost"`
}
