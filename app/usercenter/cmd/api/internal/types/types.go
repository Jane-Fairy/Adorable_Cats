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

type WXMiniAuthReq struct {
	Code          string `json:"code"`
	IV            string `json:"iv"`
	EncryptedData string `json:"encryptedData"`
}

type WXMiniAuthResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type UserInfoReq struct {
}

type UserInfoResp struct {
	UserInfo User `json:"userInfo"`
}

type PetInfo struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Breed        string `json:"breed"`
	DateOfBirth  int64  `json:"date_of_birth,optional"`
	Gender       string `json:"gender"`
	OwnerName    string `json:"owner_name"`
	OwnerContact string `json:"owner_contact"`
	Note         string `json:"note"`
	CreateTime   int64  `json:"create_string,optional"`
	UpdateTime   int64  `json:"update_string,optional"`
	DeleteTime   int64  `json:"delete_string,optional"`
	DelState     int64  `json:"del_state"`
}

type PetReq struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Breed        string `json:"breed"`
	DateOfBirth  int64  `json:"date_of_birth"`
	Gender       string `json:"gender"`
	OwnerName    string `json:"owner_name"`
	OwnerContact string `json:"owner_contact"`
	Note         string `json:"note"`
}

type PetResp struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Breed        string `json:"breed"`
	DateOfBirth  int64  `json:"date_of_birth"`
	Gender       string `json:"gender"`
	OwnerName    string `json:"owner_name"`
	OwnerContact string `json:"owner_contact"`
	Note         string `json:"note"`
	CreateTime   int64  `json:"create_string,optional"`
	UpdateTime   int64  `json:"update_string,optional"`
	DeleteTime   int64  `json:"delete_string,optional"`
	DelState     int64  `json:"del_state"`
	Version      int64  `json:"version"`
}
