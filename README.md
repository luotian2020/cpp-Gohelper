# CPP 抢票助手
仅供交流学习之用，禁止用于任何商业用途

window 系统的用户下载exe 文件食用

请按照顺序填写相关信息

2->3->4

登录->票种选择->购票人选择

直接抢票可以输入5

定时抢票仅可以定时当天，输入时间为 HH:mm,比如 12:30

后续持续优化中
待优化：
 - 未登录提示
 - 抢票前信息检查
 - 输出信息显示时间戳
 - 抢到票后回到首页 ？
 - 异常信息处理


## 附录：
注：本文档仅供参考，接口抓包时间为2023年8月20日，截止到2024年8月27日 仍然可用
### 参考接口：
#### cpp 

手机端抓包

抓包工具：

Charles

cpp版本  3.11.6（中间3.11.5升级到3.11.6了）

请求头可以带上

```json
"Origin": "https://cp.allcpp.cn",
"Referer": "https://cp.allcpp.cn",
```



#### token获取

```
https://user.allcpp.cn/api/login/normal
```

方法：post

header

| content-type    | application/x-www-form-urlencoded                            |
| --------------- | ------------------------------------------------------------ |
| accept-encoding | gzip                                                         |
| cookie          | token=; JALKSJFJASKDFJKALSJDFLJSF=141371191544379885c59f0493c9bc7e8ee1a49f499139.227.63.135_194105186121 |
| user-agent      | okhttp/3.14.7                                                |
| content-length  | 183                                                          |
| appheader       | mobile                                                       |
| equipmenttype   | 1                                                            |
| deviceversion   | 25                                                           |
| devicespec      | SM-G9810                                                     |
| appversion      | 3.11.6                                                       |

传输内容：

| account       | 18521310444                      |
| ------------- | -------------------------------- |
| password      | 123124234                        |
| deviceId      | b615637e514d53564a6b6f9da1b94c51 |
| bid           | cn.comicup.apps.cppub            |
| appVersion    | 3.11.6                           |
| deviceSpec    | SM-G9810                         |
| token         |                                  |
| equipmentType | 1                                |
| deviceVersion | 25                               |

返回内容：

```json
{
"token": "SSmHIrXa4M+P/JqJ8a2VUkjmZmcg04bm+Hyu+XgOet5WTXpRYuauh2HRThDWAMjUchcSkulZhMFQVzepHtSJXx+mqlD1z4+PBWs4DSUUOjZIcpRrsoxnqFAN8hvnAnPCV8R9IMcob4fwFdi8pOoHbnOVjfQmPnYhc38q6e3SQfg=",
"readyToken": "",
"needSetPassword": false,
"needUpdateUserInfo": true,
"thirdAccountBindToken": "",
"phoneAccountBindToken": "",
"isThirdAccountBindSuccess": false,
"isPhoneAccountBindSuccess": false
}
```



#### 漫展具体信息

获取某个漫展的具体信息

这里有两个链接

```
https://www.allcpp.cn/api/event/geteventdetail2.do
```

传输方式：post

传输内容：

| eventid       | 1541                                                         |
| ------------- | ------------------------------------------------------------ |
| token         | Kr8uDRiWM/GH4pq6FTSaatd7L3jk26SsXiPJmGtcbN221hTQr7gHDMV1f6Q2Zm4ya/qv6k/pWCz2XSHaN26AH0c1rVsxaGZxdpf+zrtZjOSxA6GK3oU2Sv3g2jrnVi8Ca5zZicYfrXdcDrzC80MoHGEqrYtt3BF8B4NnQLPV5gY= |
| deviceId      | b615637e514d53564a6b6f9da1b94c51                             |
| bid           | cn.comicup.apps.cppub                                        |
| appVersion    | 3.11.6                                                       |
| deviceSpec    | SM-G9810                                                     |
| token         | Kr8uDRiWM/GH4pq6FTSaatd7L3jk26SsXiPJmGtcbN221hTQr7gHDMV1f6Q2Zm4ya/qv6k/pWCz2XSHaN26AH0c1rVsxaGZxdpf+zrtZjOSxA6GK3oU2Sv3g2jrnVi8Ca5zZicYfrXdcDrzC80MoHGEqrYtt3BF8B4NnQLPV5gY= |
| equipmentType | 1                                                            |
| deviceVersion | 25                                                           |

cookie格式化:

| JALKSJFJASKDFJKALSJDFLJSF | 134157253982e510d5135ea40e1ac10c84ccb201ab4139.227.63.135_138243844 |
| ------------------------- | ------------------------------------------------------------ |
| JSESSIONID                | 200D2F8F81A8CA71956B2E13A60B28D9                             |
| token                     | Kr8uDRiWM/GH4pq6FTSaatd7L3jk26SsXiPJmGtcbN221hTQr7gHDMV1f6Q2Zm4ya/qv6k/pWCz2XSHaN26AH0c1rVsxaGZxdpf+zrtZjOSxA6GK3oU2Sv3g2jrnVi8Ca5zZicYfrXdcDrzC80MoHGEqrYtt3BF8B4NnQLPV5gY= |
| cdn_sec_tc                | 8be3e64616908159512216787e940a5b4c2a9442665a606fa4def28c03   |
| acw_tc                    | 8be3e64616908159512216787e940a5b4c2a9442665a606fa4def28c03   |

header:

| :method         | POST                                                         |
| --------------- | ------------------------------------------------------------ |
| :path           | /api/event/geteventdetail2.do                                |
| :authority      | www.allcpp.cn                                                |
| :scheme         | https                                                        |
| content-type    | application/x-www-form-urlencoded                            |
| accept-encoding | gzip                                                         |
| cookie          | JALKSJFJASKDFJKALSJDFLJSF=134157253982e510d5135ea40e1ac10c84ccb201ab4139.227.63.135_138243844; JSESSIONID=200D2F8F81A8CA71956B2E13A60B28D9; token=Kr8uDRiWM/GH4pq6FTSaatd7L3jk26SsXiPJmGtcbN221hTQr7gHDMV1f6Q2Zm4ya/qv6k/pWCz2XSHaN26AH0c1rVsxaGZxdpf+zrtZjOSxA6GK3oU2Sv3g2jrnVi8Ca5zZicYfrXdcDrzC80MoHGEqrYtt3BF8B4NnQLPV5gY=; cdn_sec_tc=8be3e64616908159512216787e940a5b4c2a9442665a606fa4def28c03; acw_tc=8be3e64616908159512216787e940a5b4c2a9442665a606fa4def28c03 |
| user-agent      | okhttp/3.14.7                                                |
| content-length  | 529                                                          |
| appheader       | mobile                                                       |
| equipmenttype   | 1                                                            |
| deviceversion   | 25                                                           |
| devicespec      | SM-G9810                                                     |
| appversion      | 3.11.6                                                       |

返回内容为

```json
{
	"result": {
		"templateType": 0,
		"eventId": 1541,
		"positionEnabledCode": "0",
		"enterAddress": "ShowHit LiveHouse",
		"tbSellLink": "",
		"description": "",
		"rule": "",
		"isWannaGo": false,
		"positionEnabled": "已结束",
		"isJoin": 1,
		"isExclusive": 1,
		"isLogin": true,
		"eventMainId": 1410,
		"name": "深圳Vtuber同人Only",
		"positionPicUrl": "",
		"eventMainEnabled": 0,
		"WannaGoEnabled": -1,
		"wannaGoCount": 190,
		"enterTime": "2023-08-05",
		"logoPicUrl": "https://imagecdn3.allcpp.cn/upload/2023/6/5e6febde-3940-409d-b66d-5d29175068ab.jpeg",
		"approveClosed": 1
	},
	"message": "",
	"isSuccess": true
}
```

第二个链接：

```
https://www.allcpp.cn/api/event/getEventMainDetail.do
```

传输方式：post

传输内容：

| eventmainid   | 1410                                                         |
| ------------- | ------------------------------------------------------------ |
| deviceId      | b615637e514d53564a6b6f9da1b94c51                             |
| bid           | cn.comicup.apps.cppub                                        |
| appVersion    | 3.11.6                                                       |
| deviceSpec    | SM-G9810                                                     |
| token         | Kr8uDRiWM/GH4pq6FTSaatd7L3jk26SsXiPJmGtcbN221hTQr7gHDMV1f6Q2Zm4ya/qv6k/pWCz2XSHaN26AH0c1rVsxaGZxdpf+zrtZjOSxA6GK3oU2Sv3g2jrnVi8Ca5zZicYfrXdcDrzC80MoHGEqrYtt3BF8B4NnQLPV5gY= |
| equipmentType | 1                                                            |
| deviceVersion | 25                                                           |

cookie格式化：

| JALKSJFJASKDFJKALSJDFLJSF | 134157253982e510d5135ea40e1ac10c84ccb201ab4139.227.63.135_138243844 |
| ------------------------- | ------------------------------------------------------------ |
| JSESSIONID                | 200D2F8F81A8CA71956B2E13A60B28D9                             |
| token                     | Kr8uDRiWM/GH4pq6FTSaatd7L3jk26SsXiPJmGtcbN221hTQr7gHDMV1f6Q2Zm4ya/qv6k/pWCz2XSHaN26AH0c1rVsxaGZxdpf+zrtZjOSxA6GK3oU2Sv3g2jrnVi8Ca5zZicYfrXdcDrzC80MoHGEqrYtt3BF8B4NnQLPV5gY= |
| cdn_sec_tc                | 8be3e64616908159512216787e940a5b4c2a9442665a606fa4def28c03   |
| acw_tc                    | 8be3e64616908159512216787e940a5b4c2a9442665a606fa4def28c03   |

header:

| :method         | POST                                                         |
| --------------- | ------------------------------------------------------------ |
| :path           | /api/event/getEventMainDetail.do                             |
| :authority      | www.allcpp.cn                                                |
| :scheme         | https                                                        |
| content-type    | application/x-www-form-urlencoded                            |
| accept-encoding | gzip                                                         |
| cookie          | JALKSJFJASKDFJKALSJDFLJSF=134157253982e510d5135ea40e1ac10c84ccb201ab4139.227.63.135_138243844; JSESSIONID=200D2F8F81A8CA71956B2E13A60B28D9; token=Kr8uDRiWM/GH4pq6FTSaatd7L3jk26SsXiPJmGtcbN221hTQr7gHDMV1f6Q2Zm4ya/qv6k/pWCz2XSHaN26AH0c1rVsxaGZxdpf+zrtZjOSxA6GK3oU2Sv3g2jrnVi8Ca5zZicYfrXdcDrzC80MoHGEqrYtt3BF8B4NnQLPV5gY=; cdn_sec_tc=8be3e64616908159512216787e940a5b4c2a9442665a606fa4def28c03; acw_tc=8be3e64616908159512216787e940a5b4c2a9442665a606fa4def28c03 |
| user-agent      | okhttp/3.14.7                                                |
| content-length  | 344                                                          |
| appheader       | mobile                                                       |
| equipmenttype   | 1                                                            |
| deviceversion   | 25                                                           |
| devicespec      | SM-G9810                                                     |
| appversion      | 3.11.6                                                       |

返回内容：

```json
{
	"result": {
		"eventMainId": 1410,
		"eventName": "深圳Vtuber同人Only",
		"picUrl": "https://imagecdn3.allcpp.cn/upload/2023/6/5e6febde-3940-409d-b66d-5d29175068ab.jpeg",
		"positionPicUrl": [],
		"typeId": 3,
		"type": "ONLY",
		"enterTime": 1691164800000,
		"endTime": 1691164800000,
		"enterAddress": "ShowHit LiveHouse",
		"description": "",
		"ticketDescription": "<p>全年龄段需购票进入</p>",
		"isExclusive": 1,
		"templateType": 0,
		"wannaGoCount": 190,
		"worksCount": 2,
		"eventList": [{
			"eventId": 1541,
			"eventMainId": 0,
			"name": "深圳Vtuber同人Only",
			"event_enabled": 0
		}],
		"haveTicketCode": 1,
		"userId": 82755,
		"userName": "",
		"enabled": 0
	},
	"message": "",
	"isSuccess": true
}
```

#### 用户信息获取（貌似没什么用）

获取用户信息

```
https://www.allcpp.cn/api/user/getUserBasic.do
```

传输方式：Post

传输内容：

| userid        | 2095394                                                      |
| ------------- | ------------------------------------------------------------ |
| deviceId      | b615637e514d53564a6b6f9da1b94c51                             |
| bid           | cn.comicup.apps.cppub                                        |
| appVersion    | 3.11.6                                                       |
| deviceSpec    | SM-G9810                                                     |
| token         | Kr8uDRiWM/GH4pq6FTSaatd7L3jk26SsXiPJmGtcbN221hTQr7gHDMV1f6Q2Zm4ya/qv6k/pWCz2XSHaN26AH0c1rVsxaGZxdpf+zrtZjOSxA6GK3oU2Sv3g2jrnVi8Ca5zZicYfrXdcDrzC80MoHGEqrYtt3BF8B4NnQLPV5gY= |
| equipmentType | 1                                                            |
| deviceVersion | 25                                                           |

header:

| :method         | POST                                                         |
| --------------- | ------------------------------------------------------------ |
| :path           | /api/user/getUserBasic.do                                    |
| :authority      | www.allcpp.cn                                                |
| :scheme         | https                                                        |
| content-type    | application/x-www-form-urlencoded                            |
| accept-encoding | gzip                                                         |
| cookie          | JALKSJFJASKDFJKALSJDFLJSF=134157253982e510d5135ea40e1ac10c84ccb201ab4139.227.63.135_138243844; JSESSIONID=10E3143788F652CB10B088D32E6B8F95; token=Kr8uDRiWM/GH4pq6FTSaatd7L3jk26SsXiPJmGtcbN221hTQr7gHDMV1f6Q2Zm4ya/qv6k/pWCz2XSHaN26AH0c1rVsxaGZxdpf+zrtZjOSxA6GK3oU2Sv3g2jrnVi8Ca5zZicYfrXdcDrzC80MoHGEqrYtt3BF8B4NnQLPV5gY=; cdn_sec_tc=8be3e64916909016309925774e2ec70d7d9f93a63ada429284febfc5c1; acw_tc=8be3e64916909016309925774e2ec70d7d9f93a63ada429284febfc5c1 |
| user-agent      | okhttp/3.14.7                                                |
| content-length  | 342                                                          |
| appheader       | mobile                                                       |
| equipmenttype   | 1                                                            |
| deviceversion   | 25                                                           |
| devicespec      | SM-G9810                                                     |
| appversion      | 3.11.6                                                       |

返回内容：

```
{
	"result": {
		"circle2DUrl": "1098658",
		"facePicUrl": "https://imagecdn3.allcpp.cn/face/2023/4/7056acc0-a054-4e56-b32f-e8288bb866ed.png",
		"sex": 2,
		"name": "洛天_official",
		"followMyCircleCount": 1,
		"circlePicUrl": null,
		"userId": 2095394,
		"circleName": "洛天_official的社团",
		"circleMemberCount": 1,
		"user2DUrl": "2095394",
		"collectUserCount": 0,
		"followMeCount": 0
	},
	"message": "",
	"isSuccess": true
}
```

#### 漫展票价信息

获取票价

```
https://www.allcpp.cn/allcpp/ticket/getTicketTypeList.do
```

传输方法：get

传输内容：

| eventMainId   | 1410                                                         |
| ------------- | ------------------------------------------------------------ |
| ticketMainId  | 0                                                            |
| deviceId      | b615637e514d53564a6b6f9da1b94c51                             |
| bid           | cn.comicup.apps.cppub                                        |
| appVersion    | 3.11.6                                                       |
| deviceSpec    | SM-G9810                                                     |
| token         | Kr8uDRiWM/GH4pq6FTSaatd7L3jk26SsXiPJmGtcbN221hTQr7gHDMV1f6Q2Zm4ya/qv6k/pWCz2XSHaN26AH0c1rVsxaGZxdpf+zrtZjOSxA6GK3oU2Sv3g2jrnVi8Ca5zZicYfrXdcDrzC80MoHGEqrYtt3BF8B4NnQLPV5gY= |
| equipmentType | 1                                                            |
| deviceVersion | 25                                                           |

header:

| :method         | GET                                                          |
| --------------- | ------------------------------------------------------------ |
| :path           | /allcpp/ticket/getTicketTypeList.do?eventMainId=1410&ticketMainId=0&deviceId=b615637e514d53564a6b6f9da1b94c51&bid=cn.comicup.apps.cppub&appVersion=3.11.5&deviceSpec=SM-G9810&token=Kr8uDRiWM%2FGH4pq6FTSaatd7L3jk26SsXiPJmGtcbN221hTQr7gHDMV1f6Q2Zm4ya%2Fqv6k%2FpWCz2XSHaN26AH0c1rVsxaGZxdpf%2BzrtZjOSxA6GK3oU2Sv3g2jrnVi8Ca5zZicYfrXdcDrzC80MoHGEqrYtt3BF8B4NnQLPV5gY%3D&equipmentType=1&deviceVersion=25 |
| :authority      | www.allcpp.cn                                                |
| :scheme         | https                                                        |
| accept-encoding | gzip                                                         |
| cookie          | JALKSJFJASKDFJKALSJDFLJSF=134157253982e510d5135ea40e1ac10c84ccb201ab4139.227.63.135_138243844; JSESSIONID=10E3143788F652CB10B088D32E6B8F95; token=Kr8uDRiWM/GH4pq6FTSaatd7L3jk26SsXiPJmGtcbN221hTQr7gHDMV1f6Q2Zm4ya/qv6k/pWCz2XSHaN26AH0c1rVsxaGZxdpf+zrtZjOSxA6GK3oU2Sv3g2jrnVi8Ca5zZicYfrXdcDrzC80MoHGEqrYtt3BF8B4NnQLPV5gY=; cdn_sec_tc=8be3e64916909016309925774e2ec70d7d9f93a63ada429284febfc5c1; acw_tc=8be3e64916909016309925774e2ec70d7d9f93a63ada429284febfc5c1 |
| user-agent      | okhttp/3.14.7                                                |

返回内容：

```
{
	"ticketMain": {
		"id": 587,
		"name": "深圳Vtuber同人Only",
		"eventName": "深圳Vtuber同人Only",
		"description": "<p>全年龄段需购票进入</p>",
		"eventDescription": "<p>普票</p>",
		"coverPicId": 0,
		"coverPicUrl": "",
		"picId": 0,
		"priority": 10,
		"enabled": 1,
		"eventMainId": 1410,
		"type": 0,
		"createTime": 1686112434581,
		"updateTime": 1686112434581,
		"confirmableVO": null
	},
	"ticketTypeList": [{
		"id": 1968,
		"eventId": 1541,
		"ticketMainId": 587,
		"ticketName": "普票",
		"ticketAttribute": 1,
		"ticketPrice": 5800,
		"purchaseNum": 5,
		"remainderNum": 88,
		"lockNum": 0,
		"session": 0,
		"sellStartTime": 1686117600000,
		"sellEndTime": 1691136000000,
		"openTimer": -4784413351,
		"ticketDescription": "",
		"ticketGPStartTime": 1691107200000,
		"ticketGPEndTime": 1691136000000,
		"realnameAuth": false,
		"square": "",
		"createTime": null,
		"updateTime": 1690891560459
	}]
}
```

#### 创建订单

买票（创建订单）

```
https://www.allcpp.cn/api/ticket/buyticketalipay.do
```

传输方式：post

传输内容:

| count         | 1                                                            |
| ------------- | ------------------------------------------------------------ |
| purchaserIds  | 123456（这里指购票人的id)                                    |
| ticketTypeId  | 1968                                                         |
| deviceId      | b615637e514d53564a6b6f9da1b94c51                             |
| bid           | cn.comicup.apps.cppub                                        |
| appVersion    | 3.11.6                                                       |
| deviceSpec    | SM-G9810                                                     |
| token         | Kr8uDRiWM/GH4pq6FTSaatd7L3jk26SsXiPJmGtcbN221hTQr7gHDMV1f6Q2Zm4ya/qv6k/pWCz2XSHaN26AH0c1rVsxaGZxdpf+zrtZjOSxA6GK3oU2Sv3g2jrnVi8Ca5zZicYfrXdcDrzC80MoHGEqrYtt3BF8B4NnQLPV5gY= |
| equipmentType | 1                                                            |
| deviceVersion | 25                                                           |

header:

```
:method	POST
:path	/api/ticket/buyticketalipay.do
:authority	www.allcpp.cn
:scheme	https
content-type	application/x-www-form-urlencoded
accept-encoding	gzip
cookie	JALKSJFJASKDFJKALSJDFLJSF=134157253982e510d5135ea40e1ac10c84ccb201ab4139.227.63.135_138243844; JSESSIONID=10E3143788F652CB10B088D32E6B8F95; token=Kr8uDRiWM/GH4pq6FTSaatd7L3jk26SsXiPJmGtcbN221hTQr7gHDMV1f6Q2Zm4ya/qv6k/pWCz2XSHaN26AH0c1rVsxaGZxdpf+zrtZjOSxA6GK3oU2Sv3g2jrnVi8Ca5zZicYfrXdcDrzC80MoHGEqrYtt3BF8B4NnQLPV5gY=; cdn_sec_tc=8be3e64916909016309925774e2ec70d7d9f93a63ada429284febfc5c1; acw_tc=8be3e64916909016309925774e2ec70d7d9f93a63ada429284febfc5c1
user-agent	okhttp/3.14.7
content-length	367
appheader	mobile
equipmenttype	1
deviceversion	25
devicespec	SM-G9810
appversion	3.11.5
```

返回内容：

```
{
	"result": {
		"outTradeNo": "10162071",
		"orderInfo": "alipay_root_cert_sn=687b59193f3f462dd5336e5abf83c5d8_02941eef3187dddf3d3b83462e1dfcf6&alipay_sdk=alipay-sdk-java-dynamicVersionNo&app_cert_sn=827937bd264465c567fe59847d422de7&app_id=2017050507134669&biz_content=%7B%22out_trade_no%22%3A%2210162071%22%2C%22total_amount%22%3A%2258.0%22%2C%22subject%22%3A%22CPP%E7%94%B5%E5%AD%90%E7%A5%A8%22%2C%22timeout_express%22%3A%2230m%22%2C%22product_code%22%3A%22QUICK_MSECURITY_PAY%22%7D&charset=utf-8&format=json&method=alipay.trade.app.pay&notify_url=https%3A%2F%2Fwww.allcpp.cn%3A443%2Fapi%2Fticket%2FAliPayCallback.do&return_url=https%3A%2F%2Fwww.allcpp.cn%3A443%2Fallcpp%2Fticket%2Falipay_return.do&sign=ttKC9tnjpPYm5yp7X5YXk0bFNWr%2Fu47ktbOFow1Pt1F5AuwQDiM9gfLN2Xw8QNDaC2dQWTkpq1JedNkQrVFocAeDfxSr0LZFYf0TRHTxabCMA8omZLQZQQ2Kh9uBefvMzYbXDJ4u9CPnFWB2S%2FVZoeesSB%2FHrP1p2%2FFLaTZfoWswYGxO4UVxinusHceuZuLcNTtOt%2B2NSfyVZde6T59%2BiIt9WvO7KJpMCESeH%2BW0WGCpCWZRa13sO2yWhMkf7bsqqjQrz8X%2Ff0IXqYw7u%2FSGNYvUmUwPhoi%2FbrgnO9HevPoQf8Hslt6SA79EDBORv6gJXfh0Abk19N8qEqSVo7PmeQ%3D%3D&sign_type=RSA2&timestamp=2023-08-01+23%3A15%3A19&version=1.0"
	},
	"message": "",
	"isSuccess": true
}
```

#### 漫展搜索

根据关键词搜索漫展信息

url

```
https://www.allcpp.cn/api/event/getEventMainList.do
```

方法：post

header:

| :method         | POST                                                         |
| --------------- | ------------------------------------------------------------ |
| :path           | /api/event/getEventMainList.do                               |
| :authority      | www.allcpp.cn                                                |
| :scheme         | https                                                        |
| content-type    | application/x-www-form-urlencoded                            |
| accept-encoding | gzip                                                         |
| cookie          | JSESSIONID=9E18CB05D2DD6B7C83D51EB726D40E84; token=FX10FqMz83mKKKFdcdz75aJ/S6Geucge5K8leqNXJVnIwp4nLMNAmo/vkKHFXVBpHQraYj/1G4JgAXwoCDvABGZoO7MoZCPEJXwkMifS/QX7ls4pov5t46N8gwXLjoyagxOJl+w9d/EyKIXuHzD7QlPtvkLrBcgwy4PJareUyw8=; cdn_sec_tc=8be3e64716918448359684771e861788f98485299d1b2bbcc736d80489; acw_tc=8be3e64716918448359684771e861788f98485299d1b2bbcc736d80489; JALKSJFJASKDFJKALSJDFLJSF=781819417288b6f58984d84affab4c32989a45bd67139.227.63.135_6922210595 |
| user-agent      | okhttp/3.14.7                                                |
| content-length  | 409                                                          |
| appheader       | mobile                                                       |
| equipmenttype   | 1                                                            |
| deviceversion   | 25                                                           |
| devicespec      | SM-G9810                                                     |
| appversion      | 3.11.6                                                       |

传输内容：

| cityid        | 0                                                            |
| ------------- | ------------------------------------------------------------ |
| isnew         | 1                                                            |
| orderbyid     | 1                                                            |
| searchstring  | CP                                                           |
| typeid        | 0                                                            |
| pageSize      | 20                                                           |
| pageindex     | 1                                                            |
| deviceId      | b615637e514d53564a6b6f9da1b94c51                             |
| bid           | cn.comicup.apps.cppub                                        |
| appVersion    | 3.11.6                                                       |
| deviceSpec    | SM-G9810                                                     |
| token         | FX10FqMz83mKKKFdcdz75aJ/S6Geucge5K8leqNXJVnIwp4nLMNAmo/vkKHFXVBpHQraYj/1G4JgAXwoCDvABGZoO7MoZCPEJXwkMifS/QX7ls4pov5t46N8gwXLjoyagxOJl+w9d/EyKIXuHzD7QlPtvkLrBcgwy4PJareUyw8= |
| equipmentType | 1                                                            |
| deviceVersion | 25                                                           |

返回内容：

```json
{
	"result": {
		"pageCount": 1,
		"total": 3,
		"list": [{
			"eventId": 1490,
			"eventMainId": 1360,
			"eventName": "COMICUPGZ05",
			"picUrl": "https://imagecdn3.allcpp.cn/upload/2023/5/0bd81190-ae04-4f3e-9216-76de399ef623.jpg",
			"typeId": 2,
			"type": "综合同人展",
			"enterTime": 1692979200000,
			"endTime": 1693065600000,
			"enterAddress": "保利世贸博览馆",
			"doujinshiCount": 69420,
			"circleCount": 5545,
			"wannaGoCount": 67696,
			"ticketEnabled": "电子票未开票",
			"ticketEnabledCode": 2,
			"positionIsAnnounced": "未公布",
			"positionIsAnnouncedCode": 2,
			"positionEnabled": "申摊已结束",
			"positionEnabledCode": 4,
			"isExclusive": 1,
			"priority": 999,
			"enabled": 0,
			"tbSellLink": "",
			"canAddDoujinshi": 0,
			"canApplyPosition": 1,
			"isApplyPositionStart": 1,
			"approveClosed": 1
		}, {
			"eventId": 1334,
			"eventMainId": 1216,
			"eventName": "五悠CP向主题同好会-广州",
			"picUrl": "https://imagecdn3.allcpp.cn/upload/2023/7/cb0c6ecd-4857-452f-b13e-ff1a627cd664.jpg",
			"typeId": 3,
			"type": "ONLY",
			"enterTime": 1696262400000,
			"endTime": 1696262400000,
			"enterAddress": "广州市天河区黄埔大道中322号粤大金融城国际酒店（地铁5号线科韵路站A口）",
			"doujinshiCount": 143,
			"circleCount": 23,
			"wannaGoCount": 883,
			"ticketEnabled": "电子票开票中",
			"ticketEnabledCode": 1,
			"positionIsAnnounced": "未公布",
			"positionIsAnnouncedCode": 2,
			"positionEnabled": "申摊开放中",
			"positionEnabledCode": 2,
			"isExclusive": 1,
			"priority": 0,
			"enabled": 0,
			"tbSellLink": "",
			"canAddDoujinshi": 0,
			"canApplyPosition": 1,
			"isApplyPositionStart": 1,
			"approveClosed": 0
		}, {
			"eventId": 1673,
			"eventMainId": 1530,
			"eventName": "五悠CP向主题同好会-上海",
			"picUrl": "https://imagecdn3.allcpp.cn/upload/2023/8/71aeeca1-31b8-4111-92db-f07364bd6e34.jpg",
			"typeId": 3,
			"type": "ONLY",
			"enterTime": 1702051200000,
			"endTime": 1702051200000,
			"enterAddress": "待定",
			"doujinshiCount": 9,
			"circleCount": 2,
			"wannaGoCount": 112,
			"ticketEnabled": "电子票未开票",
			"ticketEnabledCode": 2,
			"positionIsAnnounced": "未公布",
			"positionIsAnnouncedCode": 2,
			"positionEnabled": "申摊开放中",
			"positionEnabledCode": 2,
			"isExclusive": 1,
			"priority": 0,
			"enabled": 0,
			"tbSellLink": "",
			"canAddDoujinshi": 0,
			"canApplyPosition": 1,
			"isApplyPositionStart": 1,
			"approveClosed": 0
		}],
		"currentPage": 3
	},
	"message": "",
	"isSuccess": true
}
```

#### 获取购票人信息

url

```
https://www.allcpp.cn/allcpp/user/purchaser/getList.do
```

方法：get

header:

| :method         | GET                                                          |
| --------------- | ------------------------------------------------------------ |
| :path           | /allcpp/user/purchaser/getList.do?deviceId=b615637e514d53564a6b6f9da1b94c51&bid=cn.comicup.apps.cppub&appVersion=3.11.6&deviceSpec=SM-G9810&token=FX10FqMz83mKKKFdcdz75aJ%2FS6Geucge5K8leqNXJVnIwp4nLMNAmo%2FvkKHFXVBpHQraYj%2F1G4JgAXwoCDvABGZoO7MoZCPEJXwkMifS%2FQX7ls4pov5t46N8gwXLjoyagxOJl%2Bw9d%2FEyKIXuHzD7QlPtvkLrBcgwy4PJareUyw8%3D&equipmentType=1&deviceVersion=25 |
| :authority      | www.allcpp.cn                                                |
| :scheme         | https                                                        |
| accept-encoding | gzip                                                         |
| cookie          | JSESSIONID=9E18CB05D2DD6B7C83D51EB726D40E84; token=FX10FqMz83mKKKFdcdz75aJ/S6Geucge5K8leqNXJVnIwp4nLMNAmo/vkKHFXVBpHQraYj/1G4JgAXwoCDvABGZoO7MoZCPEJXwkMifS/QX7ls4pov5t46N8gwXLjoyagxOJl+w9d/EyKIXuHzD7QlPtvkLrBcgwy4PJareUyw8=; cdn_sec_tc=8be3e64716918448359684771e861788f98485299d1b2bbcc736d80489; acw_tc=8be3e64716918448359684771e861788f98485299d1b2bbcc736d80489; JALKSJFJASKDFJKALSJDFLJSF=781819417288b6f58984d84affab4c32989a45bd67139.227.63.135_6922210595 |
| user-agent      | okhttp/3.14.7                                                |

传输内容:

| deviceId      | b615637e514d53564a6b6f9da1b94c51                             |
| ------------- | ------------------------------------------------------------ |
| bid           | cn.comicup.apps.cppub                                        |
| appVersion    | 3.11.6                                                       |
| deviceSpec    | SM-G9810                                                     |
| token         | FX10FqMz83mKKKFdcdz75aJ/S6Geucge5K8leqNXJVnIwp4nLMNAmo/vkKHFXVBpHQraYj/1G4JgAXwoCDvABGZoO7MoZCPEJXwkMifS/QX7ls4pov5t46N8gwXLjoyagxOJl+w9d/EyKIXuHzD7QlPtvkLrBcgwy4PJareUyw8= |
| equipmentType | 1                                                            |
| deviceVersion | 25                                                           |

返回内容：

```json
[{
	"id": 1338422,
	"realname": "XXX",
	"idcard": "412721190001570694",
	"mobile": "18522340756",
	"validType": 0 //身份证类型
}]
```
