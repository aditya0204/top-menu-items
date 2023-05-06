package main

import (
	"top-menu-items/web"
)

func main() {
	// jb,_:=json.MarshalIndent(models.UploadMenuLogRequest{models.MenuLogs{[]models.MenuLogLine{
	// 	{
	// 		EaterId:    100,
	// 		FoodMenuId: 1,
	// 	},
	// 	{
	// 		EaterId:    101,
	// 		FoodMenuId: 1,
	// 	},
	// 	{
	// 		EaterId:    102,
	// 		FoodMenuId: 1,
	// 	},
	// 	{
	// 		EaterId:    102,
	// 		FoodMenuId: 1,
	// 	},
	// 	{
	// 		EaterId:    104,
	// 		FoodMenuId: 1,
	// 	},
	// }}},""," ")
	// fmt.Println(string(jb))
	// return
	web.StartServer()
}
