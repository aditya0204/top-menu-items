package services_test

import (
	"fmt"
	"testing"
	"top-menu-items/models"
	"top-menu-items/services"
)

type ValidateLogsLinesTestCase struct {
	MenuLogs models.MenuLogs
	Expected error
}

var testCasesForValidateLogLines = []ValidateLogsLinesTestCase{
	// TestCase 1: this time validate is pass when error occurs
	//because we have 2 customer combination repeated for eaterId:102
	ValidateLogsLinesTestCase{
		MenuLogs: models.MenuLogs{[]models.MenuLogLine{
			{
				EaterId:    100,
				FoodMenuId: 1,
			},
			{
				EaterId:    101,
				FoodMenuId: 1,
			},
			{
				EaterId:    102,
				FoodMenuId: 1,
			},
			{
				EaterId:    102,
				FoodMenuId: 1,
			},
			{
				EaterId:    104,
				FoodMenuId: 1,
			},
		}},
		Expected: services.ErrDuplicateEaterIdFoodMenuId,
	},
	// TestCase 2: this time validate is pass when error occurs
	//because we have 2 customer combination repeated for eaterId:102
	ValidateLogsLinesTestCase{
		MenuLogs: models.MenuLogs{[]models.MenuLogLine{
			{
				EaterId:    100,
				FoodMenuId: 1,
			},
			{
				EaterId:    101,
				FoodMenuId: 1,
			},
			{
				EaterId:    102,
				FoodMenuId: 1,
			},
			{
				EaterId:    103,
				FoodMenuId: 1,
			},
			{
				EaterId:    104,
				FoodMenuId: 1,
			},
		}},
		Expected: nil, //no error
	},
}

func TestValidateLogLines(t *testing.T) {
	ms := services.NewMenuService()
	for _, testCase := range testCasesForValidateLogLines {
		output := ms.ValidateLogsLines(testCase.MenuLogs)
		if output != testCase.Expected {
			t.Errorf("Output=> %v not equal to expected=> %v", output, testCase.Expected)
		}
	}
}

type GetTopItemsTestCase struct {
	MenuLogs models.MenuLogs
	Expected map[int]bool
}

var testCasesForGetTopItems = []GetTopItemsTestCase{
	// TestCase 1: this time validate is pass when error occurs
	//because we have 2 customer combination repeated for eaterId:102
	GetTopItemsTestCase{
		MenuLogs: models.MenuLogs{[]models.MenuLogLine{
			{
				EaterId:    100,
				FoodMenuId: 1,
			},
			{
				EaterId:    101,
				FoodMenuId: 1,
			},
			{
				EaterId:    102,
				FoodMenuId: 1,
			},
			{
				EaterId:    102,
				FoodMenuId: 1,
			},
			{
				EaterId:    104,
				FoodMenuId: 1,
			},
		}},
		Expected: map[int]bool{
			1: true,
		},
	},
	// TestCase 2: this time validate is pass when error occurs
	//because we have 2 customer combination repeated for eaterId:102
	GetTopItemsTestCase{
		MenuLogs: models.MenuLogs{[]models.MenuLogLine{
			{
				EaterId:    100,
				FoodMenuId: 1,
			},
			{
				EaterId:    101,
				FoodMenuId: 2,
			},
			{
				EaterId:    102,
				FoodMenuId: 3,
			},
			{
				EaterId:    103,
				FoodMenuId: 4,
			},
			{
				EaterId:    104,
				FoodMenuId: 5,
			},
		}},
		Expected: map[int]bool{ //any 3 items can be top due same number of times ordered
			1: true,
			2: true,
			3: true,
			4:true,
			5:true,
		},
	},
}

func TestGetTopItems(t *testing.T) {
	ms := services.NewMenuService()
	for i, testCase := range testCasesForGetTopItems {
		fmt.Println("\n\nRunning test case number: ",i)
		output := ms.GetTopItems(testCase.MenuLogs, 3) //testing for top 3 items
		for _, o := range output {
			if _, exist := testCase.Expected[o]; !exist {
				t.Errorf("Output=> %v not equal to expected=> %v", o, testCase.Expected)
			}
		}

	}
}
