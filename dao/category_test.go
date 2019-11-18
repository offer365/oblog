package dao

import (
	"fmt"
	"testing"
)

func TestGetCategoryList(t *testing.T) {
	err := Init("117.107.146.194:6603", "oblog", "qwe123!@#QWE", "oblog", "utf8mb4", true)
	if err != nil {
		t.Fatal(err)
	}
	cs, err := GetCategoryList([]int64{1, 2, 3})
	fmt.Println(cs[0].Name)
}
