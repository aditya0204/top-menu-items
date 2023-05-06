package models

type Response struct {
	StatusCode int         `json:"status_code`
	Error      string      `json:"error"`
	Data       interface{} `json:"data"`
}

type TopMenuItems struct{
	TopItemIds []int `json:"top_item_ids`
}
