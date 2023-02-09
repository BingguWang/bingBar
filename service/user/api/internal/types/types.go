// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	Id       int64  `json:"id"`
	Mobile   string `json:"mobile"`
	Nickname string `json:"nickname"`
	Sex      int64  `json:"sex"`
	Avatar   string `json:"avatar"`
	Info     string `json:"info"`
}

type RegisterReq struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type RegisterResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type LoginReq struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type FollowReq struct {
	Uid int64 `json:"uid"`
}

type FollowResp struct {
	RetMsg string `json:"retMsg"`
}

type UnFollowReq struct {
	Uid int64 `json:"uid"`
}

type UnFollowResp struct {
	RetMsg string `json:"retMsg"`
}

type GetFollowListReq struct {
	PageNo   int64 `json:"pageNo"`
	PageSize int64 `json:"pageSize"`
}

type GetFollowListResp struct {
	RetMsg   string `json:"retMsg"`
	UserList []User `json:"userList"`
}

type GetFansListReq struct {
	PageNo   int64 `json:"pageNo"`
	PageSize int64 `json:"pageSize"`
}

type GetFansListResp struct {
	RetMsg   string `json:"retMsg"`
	UserList []User `json:"userList"`
}

type UserInfoReq struct {
}

type UserInfoResp struct {
	UserInfo User `json:"userInfo"`
}
