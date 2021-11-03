// Code generated by goctl. DO NOT EDIT.
package types

type ReqUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ReqUserId struct {
	Id int64 `path:"id"`
}

type ReqUpdateUser struct {
	Id       int64  `path:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RespLogin struct {
	Token string `json:"token"`
}

type CommResp struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}
