// Code generated by goctl. DO NOT EDIT.
package types

type CreateRequest struct {
	Pid int64 `json:"pid"`
	Num int64 `json:"num"`
}

type CreateResponse struct {
	Id int64 `json:"id"`
}

type UpdateRequest struct {
	Id     int64 `json:"id"`
	Pid    int64 `json:"pid,optional"`
	Amount int64 `json:"amount,optional"`
	Status int64 `json:"status,optional"`
}

type UpdateResponse struct {
}

type RemoveRequest struct {
	Id int64 `json:"id"`
}

type RemoveResponse struct {
}

type DetailRequest struct {
	Id int64 `json:"id"`
}

type DetailResponse struct {
	Id     int64 `json:"id"`
	Uid    int64 `json:"uid"`
	Pid    int64 `json:"pid"`
	Num    int64 `json:"num"`
	Amount int64 `json:"amount"`
	Status int64 `json:"status"`
}

type ListRequest struct {
}

type ListResponse struct {
	Id     int64 `json:"id"`
	Uid    int64 `json:"uid"`
	Pid    int64 `json:"pid"`
	Num    int64 `json:"num"`
	Amount int64 `json:"amount"`
	Status int64 `json:"status"`
}
