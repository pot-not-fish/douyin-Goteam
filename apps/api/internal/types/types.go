// Code generated by goctl. DO NOT EDIT.
package types

type RegisterReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type RegisterResp struct {
	Status_code int8   `json:"status_code"`
	Status_msg  string `json:"status_msg"`
	User_id     int64  `json:"user_id"`
	Token       string `json:"token"`
}

type LoginReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type LoginResp struct {
	Status_code int8   `json:"status_code"`
	Status_msg  string `json:"status_msg"`
	User_id     int64  `json:"user_id"`
	Token       string `json:"token"`
}

type UserinfoReq struct {
	User_id int64  `form:"user_id"`
	Token   string `form:"token"`
}

type UserinfoResp struct {
	Status_code int8   `json:"status_code"`
	Status_msg  string `json:"status_msg"`
	User        *User  `json:"user"`
}

type User struct {
	Id               int64  `json:"id"`
	Name             string `json:"name"`
	Follow_count     int64  `json:"follow_count"`
	Follower_count   int64  `json:"follower_count"`
	Is_follow        bool   `json:"is_follow"`
	Avatar           string `json:"avatar"`
	Background_image string `json:"background_image"`
	Signature        string `json:"signature"`
	Total_favorited  string `json:"total_favorited"`
	Work_count       int64  `json:"work_count"`
	Favorite_count   int64  `json:"favorite_count"`
}