package cpp

import "net/http"

type CppCrawler struct {
	Cookie      map[string]string `json:"cookie"`
	TicketList  []TicketInfo      `json:"ticketList"`
	Personlist  []Person          `json:"personList"`
	ComicList   map[string]string `json:"comicList"`
	Version     string            `json:"version"`
	EventMainId string            `json:"eventMainId"`
	Timesleep   int               `json:"timesleep"`
	Account     string            `json:"account"`
	Password    string            `json:"password"`
	Token       string            `json:"token"`
	BuyTicket   TicketInfo        `json:"buyticket"`
	BuyPerson   Person            `json:"buyperson"`
	client      *http.Client
	OrderResult OrderResult
}
type TicketMain struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	EventName        string `json:"eventName"`
	Description      string `json:"description"`
	EventDescription string `json:"eventDescription"`
	CoverPicID       int    `json:"coverPicId"`
	CoverPicURL      string `json:"coverPicUrl"`
	PicID            int    `json:"picId"`
	Priority         int    `json:"priority"`
	Enabled          int    `json:"enabled"`
	EventMainID      int    `json:"eventMainId"`
	Type             int    `json:"type"`
	CreateTime       int64  `json:"createTime"`
	UpdateTime       int64  `json:"updateTime"`
	ConfirmableVO    any    `json:"confirmableVO"`
}
type TicketResponse struct {
	TicketMain     TicketMain   `json:"ticketMain"`
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
