package model

import "time"

type Response_afk_toplist struct {
	Username string `json:"username"`
	Userid   struct {
		Int64 int64 `json:"Int64"`
		Valid bool  `json:"Valid"`
	} `json:"userid"`
	Count int `json:"count"`
}

type Request_afk_create struct {
	Username string `json:"username"`
	Userid   int64  `json:"userid"`
}

type Response_afk_create struct {
	CreatedAt struct {
		Time  time.Time `json:"Time"`
		Valid bool      `json:"Valid"`
	} `json:"created_at"`
	Afkid  int `json:"afkid"`
	Userid struct {
		Int64 int64 `json:"Int64"`
		Valid bool  `json:"Valid"`
	} `json:"userid"`
}

type Response_afk_user struct {
	CreatedAt struct {
		Time  time.Time `json:"Time"`
		Valid bool      `json:"Valid"`
	} `json:"created_at"`
	Afkid  int `json:"afkid"`
	Userid struct {
		Int64 int64 `json:"Int64"`
		Valid bool  `json:"Valid"`
	} `json:"userid"`
}
