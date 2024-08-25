package cpp

import "net/http"

type CppCrawler struct {
	Cookie      map[string]string `json:"cookie"`
	TicketList  []TicketInfo      `json:"ticketList"`
	Personlist  []Person          `json:"personList"`
	ComicList   map[string]string `json:"comicList"`
	Version     string            `json:"version"`
	EventMainId string            `json:"eventMainId"`
	Timesleep   string            `json:"timesleep"`
	Account     string            `json:"account"`
	Password    string            `json:"password"`
	Token       string            `json:"token"`
	BuyTicket   TicketInfo        `json:"buyticket"`
	BuyPerson   Person            `json:"buyperson"`
	client      *http.Client
	orderResult OrderResult
}
type TicketResponse struct {
	TicketTypeList []TicketInfo `json:"ticketTypeList"`
}
type TicketInfo struct {
	ID                int    `json:"id"`
	EventID           int    `json:"eventId"`
	TicketMainID      int    `json:"ticketMainId"`
	TicketName        string `json:"ticketName"`
	TicketAttribute   int    `json:"ticketAttribute"`
	TicketPrice       int    `json:"ticketPrice"`
	PurchaseNum       int    `json:"purchaseNum"`
	RemainderNum      int    `json:"remainderNum"`
	LockNum           int    `json:"lockNum"`
	Session           int    `json:"session"`
	SellStartTime     int64  `json:"sellStartTime"`
	SellEndTime       int64  `json:"sellEndTime"`
	OpenTimer         int64  `json:"openTimer"`
	TicketDescription string `json:"ticketDescription"`
	TicketGPStartTime int64  `json:"ticketGPStartTime"`
	TicketGPEndTime   int64  `json:"ticketGPEndTime"`
	RealnameAuth      bool   `json:"realnameAuth"`
	Square            string `json:"square"`
	CreateTime        any    `json:"createTime"`
	UpdateTime        int64  `json:"updateTime"`
}

type Person struct {
	ID        int    `json:"id"`
	Realname  string `json:"realname"`
	Idcard    string `json:"idcard"`
	Mobile    string `json:"mobile"`
	ValidType int    `json:"validType"`
}
type OrderResult struct {
	Result    Result `json:"result"`
	Message   string `json:"message"`
	IsSuccess bool   `json:"isSuccess"`
}
type Result struct {
	OutTradeNo string `json:"outTradeNo"`
	OrderInfo  string `json:"orderInfo"`
}
