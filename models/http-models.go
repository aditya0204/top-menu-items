package models

type Response struct {
	StatusCode int         `json:"status_code`
	Error      string      `json:"error"`
	Data       interface{} `json:"data"`
}

type UploadMenuLogRequest struct{
	MenuLogs MenuLogs `json:"menu_logs"`
}

type TopMenuItems struct{
	TopItemIds []int `json:"top_food_item_ids"`
}
