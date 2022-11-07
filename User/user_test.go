package User

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	fmt.Println("开始测试:")
	t.Run("测试查询一条记录:", testUser_SelectUserByEmail)
}

func testUser_AddUser(t *testing.T) {
	fmt.Println("添加用户:")
	user := User{
		Username: "liweijun",
		Password: "liweijun",
		Email:    "123@gmail.com",
	}
	err := user.AddUser()
	if err != nil {
		fmt.Println("添加用户失败:", err)
		return
	}
}
func testUser_SelectUserByEmail(t *testing.T) {
	fmt.Println("测试查询一条记录")
	user := User{
		Username: "",
		Password: "",
		Email:    "admin@admin.com",
	}
	err := user.SelectUserByEmail()
	if err != nil {
		return
	}
}
