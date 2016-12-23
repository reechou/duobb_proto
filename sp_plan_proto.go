package duobb_proto

type CreateSpPlanReq struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Remark   string `json:"remark"`
}

type SyncSpPlanSourceReq struct {
	User             string `json:"user"`
	PlanId           int64  `json:"planId"`
	SourceFromId     int64  `json:"sourceFromId"`
	SourceIdPassword string `json:"sourceIdPassword"`
}

type GetSpPlanListFromUserReq struct {
	User   string `json:"user"`
	Offset int64  `json:"offset"`
	Num    int64  `json:"num"`
}

type GetSpPlanListPublicReq struct {
	Offset int64 `json:"offset"`
	Num    int64 `json:"num"`
}

type GetSpPlanInfoFromUserReq struct {
	User   string `json:"user"`
	PlanId int64  `json:"planId"`
}

type GetSpPlanFromPasswordReq struct {
	PlanId   int64  `json:"planId"`
	Password string `json:"password"`
}

type UpdateSpPlanItemsReq struct {
	User          string  `json:"user"`
	PlanId        int64   `json:"planId"`
	ItemsNum      int64   `json:"itemsNum"`
	ItemsAvgPrice float32 `json:"itemsAvgPrice"`
	AvgCommission float32 `json:"avgCommission"`
	ItemsList     string  `json:"itemsList"`
}

type UpdateSpPlanPasswordReq struct {
	User     string `json:"user"`
	PlanId   int64  `json:"planId"`
	Password string `json:"password"`
}

type UpdateSpPlanRemarkReq struct {
	User   string `json:"user"`
	PlanId int64  `json:"planId"`
	Remark string `json:"remark"`
}
