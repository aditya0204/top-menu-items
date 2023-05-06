package models


type MenuLogs struct{
	Lines []MenuLogLine `json:"menu_logs"`
}
type MenuLogLine struct{
	EaterId int `json:"eater_id"`
	FoodMenuId int `json:"foodmenu_id"`
}

type UploadMenuLogRequest struct{
	MenuLogs MenuLogs `json:"menu_logs"`
}