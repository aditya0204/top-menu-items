package services

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"top-menu-items/models"
)

var (
	ErrDuplicateEaterIdFoodMenuId = errors.New("received duplicate eaterId vs FoodMenuId combination")
)

type MenuService struct {
}

func NewMenuService() *MenuService {
	return &MenuService{}
}

func (ms *MenuService) ValidateLogsLines(menuLogs models.MenuLogs) error {
	eaterIdvsFoodMenuIdMap := make(map[string]int)
	for _, logLine := range menuLogs.Lines {
		key := fmt.Sprintf("%d_%d", logLine.EaterId, logLine.FoodMenuId)
		if _, exist := eaterIdvsFoodMenuIdMap[key]; exist {
			return ErrDuplicateEaterIdFoodMenuId
		}
		eaterIdvsFoodMenuIdMap[key] = logLine.FoodMenuId
	}
	return nil
}
func (ms *MenuService) GetTopItems(menuLogs models.MenuLogs, total int) []int {
	foodItemCount := make(map[int]int)

	for _, line := range menuLogs.Lines {
		foodItemCount[line.FoodMenuId]++
	}

	fmt.Println("foodItems Count=>", foodItemCount)

	type keyValue struct {
		Key   int
		Value int
	}

	var ss []keyValue
	for k, v := range foodItemCount {
		ss = append(ss, keyValue{k, v})
	}

	// Then sorting the slice by value, higher first.
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	var topItems []int

	for i, v := range ss {
		if i >= total {
			fmt.Println("top items ", topItems)
			return topItems
		}
		topItems = append(topItems, v.Key)
	}
	fmt.Println("top items ", topItems)
	return topItems
}

func UploadMenuLogFileService(req models.UploadMenuLogRequest) (*models.TopMenuItems, error) {
	ms := NewMenuService()
	err := ms.ValidateLogsLines(req.MenuLogs)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	topItems := ms.GetTopItems(req.MenuLogs, 3) //currently fetching top 3 items

	return &models.TopMenuItems{TopItemIds: topItems}, nil
}
