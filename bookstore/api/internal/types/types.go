// Code generated by goctl. DO NOT EDIT.
package types

type AddReq struct {
	Book  string `form:"book"`
	Price int64  `form:"price"`
}

type AddResp struct {
	Ok bool `json:"ok"`
}

type CheckReq struct {
	Book string `form:"book"`
}

type CheckResp struct {
	Found bool  `json:"found"`
	Price int64 `form:"price"`
}
