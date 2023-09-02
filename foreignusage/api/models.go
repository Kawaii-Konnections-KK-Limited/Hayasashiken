package api

type link struct {
	ID   int    `json:"id"`
	Link string `json:"link"`
}
type requestLinks struct {
	Links               []link `json:"links"`
	Timeout             int32  `json:"timeout"`
	UpperBoundPingLimit int32  `json:"upperbound"`
	TestUrl             string `json:"testurl"`
}
type responseLink struct {
	ID   int    `json:"id"`
	Ping int32  `json:"ping"`
	Link string `json:"link"`
}
type responseLinks struct {
	Links []responseLink `json:"links"`
}
type responseError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
