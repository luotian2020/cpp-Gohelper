package cpp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/go-toast/toast"
)

func NewCppCrawler() *CppCrawler {
	cppObj := &CppCrawler{
		Timesleep: 500,
		Version:   "3.14.4",
	}
	*cppObj = readjson()
	if cppObj.Account == "" {
		fmt.Println("请输入账户:")
		fmt.Scanln(&cppObj.Account)
	}
	if cppObj.Password == "" {
		fmt.Println("请输入密码:")
		fmt.Scanln(&cppObj.Password)
	}
	if cppObj.Version == "" {
		fmt.Println("请输入当前版本(请填写 3.14.7):")
		fmt.Scanln(&cppObj.Version)
	}
	cppObj.client = &http.Client{}
	writeJson(*cppObj)
	return cppObj
}

/**
* 登录
 */
func (c *CppCrawler) Login() {
	login_url := "https://user.allcpp.cn/api/login/normal"
	params := map[string]string{
		"account":       c.Account,
		"password":      c.Password,
		"deviceId":      "b615637e514d53564a6b6f9da1b94c51",
		"bid":           "cn.comicup.apps.cppub",
		"equipmenttype": "1",
		"deviceversion": "25",
		"devicespec":    "SM-G9810",
		"token":         "",
	}
	formData := url.Values{}
	for k, v := range params {
		formData.Set(k, v)
	}
	resp, err := http.PostForm(login_url, formData)
	if err != nil {
		fmt.Println("login failed,error is ", err)
	}
	defer resp.Body.Close()
	body, errr := io.ReadAll(resp.Body)
	if errr != nil {
		fmt.Println("login body fail ,error is ", errr)
		return
	}
	fmt.Println(string(body))
	cookies := resp.Cookies()
	cookieMap := make(map[string]string)
	for _, cookie := range cookies {
		cookieMap[cookie.Name] = cookie.Value
	}
	fmt.Println("获取到的cookie为", cookieMap)
	c.Cookie = cookieMap
	result := bodyToJson(body)
	c.Token = string(result["token"].(string))
	writeJson(*c)
}

/**
* 获取票信息
 */
func (c *CppCrawler) GetTicketInfo() {
	if c.Token == "" {
		fmt.Println("你还未登录！")
		return
	}
	fmt.Println("请输入活动Id(CP30 ID为 1729):")
	fmt.Scanln(&c.EventMainId)
	header := map[string]string{
		"User-Agent":    "okhttp/3.14.7",
		"Origin":        "https://cp.allcpp.cn",
		"Referer":       "https://cp.allcpp.cn",
		"content-type":  "application/x-www-form-urlencoded",
		"appheader":     "mobile",
		"equipmenttype": "1",
		"deviceversion": "25",
		"devicespec":    "SM-G9810",
		"appversion":    c.Version,
	}
	params := map[string]string{
		"eventMainId":   string(c.EventMainId),
		"ticketMainId":  "0",
		"deviceId":      "b615637e514d53564a6b6f9da1b94c51",
		"bid":           "cn.comicup.apps.cppub",
		"equipmenttype": string(1),
		"deviceversion": string(25),
		"devicespec":    "SM-G9810",
		"token":         c.Token,
	}
	INFO_URL := "https://www.allcpp.cn/allcpp/ticket/getTicketTypeList.do"
	body := c.GetReq(INFO_URL, params, header)
	fmt.Println(string(body))
	var typeResponse TicketResponse
	err := json.Unmarshal(body, &typeResponse)
	if err != nil {
		fmt.Printf("Error parse type:%s", err)
		panic("parse type error!")
	}
	c.TicketList = typeResponse.TicketTypeList
	fmt.Println("展会名称为", typeResponse.TicketMain.Name)
	writeJson(*c)
}

/**
** 选择票
 */
func (c *CppCrawler) ChoseTicket() {
	if len(c.TicketList) == 0 {
		fmt.Println("暂无门票！")
		return
	}
	fmt.Println("票种选择:")
	for index, value := range c.TicketList {
		fmt.Printf("序号为%d 票名：%s,票价：%d \n", index, value.TicketName, value.TicketPrice/100)
	}
	fmt.Println("请输入序号:")
	var a int
	fmt.Scanln(&a)
	c.BuyTicket = c.TicketList[a]
	writeJson(*c)
}

/**
* 获取购买人信息
 */
func (c *CppCrawler) GetPersonInfo() {
	if c.Token == "" {
		fmt.Println("你还未登录！")
		return
	}
	header := map[string]string{
		"User-Agent":    "okhttp/3.14.7",
		"Origin":        "https://cp.allcpp.cn",
		"Referer":       "https://cp.allcpp.cn",
		"content-type":  "application/x-www-form-urlencoded",
		"appheader":     "mobile",
		"equipmenttype": "1",
		"deviceversion": "25",
		"devicespec":    "SM-G9810",
		"appversion":    c.Version,
	}
	params := map[string]string{
		"deviceId":      "b615637e514d53564a6b6f9da1b94c51",
		"bid":           "cn.comicup.apps.cppub",
		"equipmenttype": string(1),
		"deviceversion": string(25),
		"devicespec":    "SM-G9810",
		"token":         c.Token,
	}
	PERSON_URL := "https://www.allcpp.cn/allcpp/user/purchaser/getList.do"
	body := c.GetReq(PERSON_URL, params, header)
	var personlist []Person
	err := json.Unmarshal(body, &personlist)
	if err != nil {
		fmt.Println("error json:", err)
		panic("购买人解析错误！")
	}
	c.Personlist = personlist
	writeJson(*c)
}

/**
* 选择购买人
 */
func (c *CppCrawler) ChosePerson() {
	if len(c.Personlist) == 0 {
		fmt.Println("暂未配置购买人！")
		return
	}
	fmt.Println("购买人选择：")
	for index, value := range c.Personlist {
		fmt.Printf("购买人序号：%d,名称：%s,身份id:%s,电话：%s \n", index, value.Realname, value.Idcard, value.Mobile)
	}
	var b int
	fmt.Println("请输入购票人序号:")
	fmt.Scanln(&b)
	c.BuyPerson = c.Personlist[b]
	writeJson(*c)
}

func (c *CppCrawler) CreateOrder() {
	header := map[string]string{
		"User-Agent":    "okhttp/3.14.7",
		"Origin":        "https://cp.allcpp.cn",
		"Referer":       "https://cp.allcpp.cn",
		"content-type":  "application/x-www-form-urlencoded",
		"appheader":     "mobile",
		"equipmenttype": "1",
		"deviceversion": "25",
		"devicespec":    "SM-G9810",
		"appversion":    c.Version,
	}
	params := map[string]string{
		"count":         "1",
		"purchaserIds":  strconv.Itoa(c.BuyPerson.ID),
		"ticketTypeId":  strconv.Itoa(c.BuyTicket.ID),
		"deviceId":      "b615637e514d53564a6b6f9da1b94c51",
		"bid":           "cn.comicup.apps.cppub",
		"equipmenttype": string(1),
		"deviceversion": string(25),
		"devicespec":    "SM-G9810",
		"token":         c.Token,
	}
	ORDER_URL := "https://www.allcpp.cn/api/ticket/buyticketalipay.do"
	body := c.PostReq(ORDER_URL, params, header)
	var orderResult OrderResult
	err := json.Unmarshal(body, &orderResult)
	if err != nil {
		fmt.Println("error json:", err)
		return 
	}
	c.OrderResult = orderResult
	if orderResult.IsSuccess {
		fmt.Println("抢到票了！")
	} else {
		fmt.Println(orderResult.Message)
	}
}

func bodyToJson(body []byte) map[string]any {
	var result map[string]any
	err := json.Unmarshal(body, &result)
	if err != nil {
		fmt.Printf("body is %s,Error parse json:%s", string(body), err)
		panic("parse json error!")
	}
	return result
}
func readjson() CppCrawler {
	if !fileExists("crawler-cpp.json") {
		return CppCrawler{}
	}
	file, err := os.Open("crawler-cpp.json")
	if err != nil {
		log.Fatalf("无法打开文件：%v", err)
		panic("无法打开文件")
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("读取文件错误：%v", err)
		panic("读取文件错误")
	}
	var c CppCrawler
	err = json.Unmarshal(data, &c)
	if err != nil {
		log.Fatalf("解析 JSON 数据错误: %v", err)
		panic("解析 JSON 数据错误")
	}
	return c
}
func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func writeJson(c CppCrawler) {
	jsonData, err := json.MarshalIndent(c, "", "")
	if err != nil {
		log.Fatalf("编码 JSON 数据错误: %v", err)
		panic("编码 JSON 数据错误")
	}
	err = os.WriteFile("crawler-cpp.json", jsonData, 0644)
	if err != nil {
		log.Fatalf("写入文件错误: %v", err)
		panic("写入文件错误错误")
	}
}
func (c *CppCrawler) GetReq(info_url string, params map[string]string, header map[string]string) []byte {
	rawParams := url.Values{}
	for k, v := range params {
		rawParams.Set(k, v)
	}
	fullUrl := info_url + "?" + rawParams.Encode()
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		log.Fatalf("创建GET请求错误: %v", err)
		panic("创建GET请求错误")
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		log.Fatalf("发送GET请求错误: %v", err)
		panic("发送GET请求错误")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("读取响应错误: %v", err)
		panic("读取响应错误")
	}
	return body
}

func (c *CppCrawler) PostReq(post_url string, data map[string]string, header map[string]string) []byte {
	rawParams := url.Values{}
	for k, v := range data {
		rawParams.Set(k, v)
	}
	formDataStr := rawParams.Encode()
	formDataBytes := []byte(formDataStr)
	formBytesReader := bytes.NewReader(formDataBytes)
	req, err := http.NewRequest("POST", post_url, formBytesReader)
	if err != nil {
		log.Fatalf("创建请求错误: %v", err)
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		log.Fatalf("发送POST请求错误: %v", err)
		panic("发送POST请求错误")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		log.Fatalf("读取响应错误: %v", err)
		panic("读取响应错误")
	}
	return body
}
func (c *CppCrawler) InfoClear() {
	if fileExists("crawler-cpp.json") {
		err := os.Remove("crawler-cpp.json")
		if err != nil {
			fmt.Println("信息重置失败,error为:", err)
		}
	}
}
func (c *CppCrawler) SetTimeSleep() {
	fmt.Println("时间间隔设置(毫秒):")
	fmt.Scanln(&c.Timesleep)
	writeJson(*c)
}
func (c *CppCrawler) GrapTicket() {
	if c.Token == "" {
		fmt.Println("你还未登录！")
		return
	}
	for {
		c.CreateOrder()
		if c.OrderResult.IsSuccess {
			break
		}
		time.Sleep(time.Duration(c.Timesleep) * time.Millisecond)
	}
}
func (c *CppCrawler) CronTicket() {
	if c.Token == "" {
		fmt.Println("你还未登录！")
		return
	}
	if(c.BuyPerson.Realname==""){
		fmt.Println("你还未选择购买人！")
	}
	c.PrintTicketInfo()
	var temp string
	fmt.Println("请输入当天抢票时间(比如 12:00):")
	fmt.Scanln(&temp)
	fmt.Printf("定时时间为:%v,准备抢票中----",temp)
	scheduler := gocron.NewScheduler(time.Local)
	scheduler.Every(1).Day().At(temp).Do(func() {
		c.GrapTicket()
		scheduler.Stop()
	})

	// 启动调度器（阻塞式运行）
	scheduler.StartBlocking()
}
func (c *CppCrawler) SetVersion() {
	fmt.Println("输入版本号:")
	fmt.Scanln(&c.Version)
	writeJson(*c)
}
func (c *CppCrawler) CheckTicketInfo() {
	fmt.Println("请输入活动id:")
	var temp int
	fmt.Scanln(&temp)
	scheduler := gocron.NewScheduler(time.Local)
	scheduler.Every(1).Minute().Do(func() {
		c.NoticeTicketInfo(temp)
	})
	// 启动调度器（阻塞式运行）
	scheduler.StartBlocking()
}

/**
* 获取当前是否有票信息生成
 */
func (c *CppCrawler) NoticeTicketInfo(EventMainId int) {
	header := map[string]string{
		"User-Agent":    "okhttp/3.14.7",
		"Origin":        "https://cp.allcpp.cn",
		"Referer":       "https://cp.allcpp.cn",
		"content-type":  "application/x-www-form-urlencoded",
		"appheader":     "mobile",
		"equipmenttype": "1",
		"deviceversion": "25",
		"devicespec":    "SM-G9810",
		"appversion":    c.Version,
	}
	params := map[string]string{
		"eventMainId":   strconv.Itoa(EventMainId),
		"ticketMainId":  "0",
		"deviceId":      "b615637e514d53564a6b6f9da1b94c51",
		"bid":           "cn.comicup.apps.cppub",
		"equipmenttype": string(1),
		"deviceversion": string(25),
		"devicespec":    "SM-G9810",
		"token":         c.Token,
	}
	INFO_URL := "https://www.allcpp.cn/allcpp/ticket/getTicketTypeList.do"
	body := c.GetReq(INFO_URL, params, header)
	fmt.Println(string(body))
	var typeResponse TicketResponse
	err := json.Unmarshal(body, &typeResponse)
	if err != nil {
		fmt.Printf("Error parse type:%s", err)
		panic("parse type error!")
	}
	c.TicketList = typeResponse.TicketTypeList
	if len(c.TicketList) != 0 {
		//向windows 发送信息
		strvalue := string(body)
		c.SendNotice("Cpp放票信息通知", "放票信息为"+strvalue)
		c.SendMail(strvalue)
	}
}

func (c *CppCrawler) SendNotice(title string, info string) {
	notification := toast.Notification{
		AppID:   "cpp-gohelper",
		Title:   title,
		Message: info,
		Actions: []toast.Action{
			{Type: "protocol", Label: "ok", Arguments: ""},
			{Type: "protocol", Label: "close", Arguments: ""},
		},
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}

func (c *CppCrawler) SendMail(info string) {
	// SMTP 服务器配置
	smtpHost := ""             // SMTP 服务器地址
	smtpPort := ""                     // SMTP 端口
	username := "" // SMTP 用户名
	password := ""        // SMTP 密码

	// 收件人和发件人
	from := "galigali-luotian@qq.com"
	to := []string{"2"}

	// 邮件内容
	message := []byte("From: Give <" + from + ">\r\n" +
		"To: " + to[0] + "\r\n" +
		"Subject: CPP票种信息通知\r\n" +
		"\r\n" + info)

	// 连接到 SMTP 服务器
	auth := smtp.PlainAuth("", username, password, smtpHost)

	// 发送邮件
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Email sent successfully!")
}

func (c *CppCrawler) PrintTicketInfo(){
	fmt.Println("当前信息:")
	fmt.Printf("当前购买票种：%v",c.BuyTicket.TicketName);
	fmt.Println();
	fmt.Printf("当前购买人：%v - %v \n",c.BuyPerson.Realname,string(c.BuyPerson.Idcard))
}
