package duobb_proto

type QGMemberInviteReq struct {
	UserName  string `json:"userName"`
	GroupName string `json:"groupName"`
	Member    string `json:"member"`
	Inviter   string `json:"inviter"`
}

type QGMemberExitReq struct {
	UserName  string `json:"userName"`
	GroupName string `json:"groupName"`
	Member    string `json:"member"`
}

type QGMemberSign struct {
	UserName  string `json:"userName"`
	GroupName string `json:"groupName"`
	Member    string `json:"member"`
}

type QGMemberCommitOrder struct {
	UserName   string  `json:"userName"`
	GroupName  string  `json:"groupName"`
	Member     string  `json:"member"`
	OrderId    string  `json:"orderId"`
	Commission float32 `json:"commission"`
}

type QGMemberConfirmOrder struct {
	UserName  string `json:"userName"`
	GroupName string `json:"groupName"`
	Member    string `json:"member"`
	OrderId   string `json:"orderId"`
}

type QGScoreSettingReq struct {
	UserName          string `json:"userName"`
	SignScore         int64  `json:"signScore"`
	InviteScore       int64  `json:"inviteScore"`
	OrderOneYuanScore int64  `json:"orderOneYuanScore"`
	OrderScoreScale   int64  `json:"orderScoreScale"`
}

type QGAddBlackListReq struct {
	UserName  string `json:"userName"`
	GroupName string `json:"groupName"`
	Member    string `json:"member"`
}

type QGDelBlackListReq struct {
	UserName  string `json:"userName"`
	GroupName string `json:"groupName"`
	Member    string `json:"member"`
}

type QGGetBlackListReq struct {
	UserName  string `json:"userName"`
	GroupName string `json:"groupName"`
	Offset    int64  `json:"offset"`
	Num       int64  `json:"num"`
}

type QGAddWhiteListReq struct {
	UserName  string `json:"userName"`
	GroupName string `json:"groupName"`
	Member    string `json:"member"`
	Position  int64  `json:"position"` // 0: 群管 1:发单员
}

type QGGetWhiteListReq struct {
	UserName  string `json:"userName"`
	GroupName string `json:"groupName"`
}

type QGDelWhiteListReq struct {
	UserName  string `json:"userName"`
	GroupName string `json:"groupName"`
	Member    string `json:"member"`
}
