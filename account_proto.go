package duobb_proto

type CreateDuobbAccountReq struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	PicUrl   string `json:"picUrl"`
}

type UpdateDuobbAccountPasswordReq struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type UpdateDuobbAccountPhoneReq struct {
	User  string `json:"user"`
	Phone string `json:"phone"`
}

type UpdateDuobbAccountPicUrlReq struct {
	User   string `json:"user"`
	PicUrl string `json:"picUrl"`
}

type GetDuobbAccountReq struct {
	User string `json:"user"`
}

type GetDuobbAccountFromPhoneReq struct {
	Phone string `json:"phone"`
}

type DuobbLogin struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type DuobbLogout struct {
	User string `json:"user"`
}

type DuobbHeartbeat struct {
	User string `json:"user"`
}

type DuobbAccountUploadDataReq struct {
	User              string  `json:"user"`
	AlimamaName       string  `json:"alimamaName"`
	Day               string  `json:"day"`
	TodaySendItemsNum int64   `json:"todaySendItemsNum"`
	TodayBuyItemsNum  int64   `json:"todayBuyItemsNum"`
	TodayCommission   float32 `json:"todayCommission"`
}

type GetAllDuobbDataReq struct {
	User string `json:"user"`
}

type DuobbAccountUploadAlimamaCookieReq struct {
	User        string `json:"user"`
	AlimamaName string `json:"ali"`
	Cookie      string `json:"c"`
}

type DuobbAccountGetAlimamaCookieReq struct {
	User        string `json:"user"`
	AlimamaName string `json:"ali"`
}
