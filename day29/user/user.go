package user

import "fmt"

type User struct {
	Name  string
	Pwd   string
	Count int
}

func NewUser() *User {
	return &User{
		"",
		"",
		3,
	}
}

func (this *User) Check() {
	for {
		fmt.Println("请输入用户名：")
		fmt.Scanln(&this.Name)
		fmt.Println("请输入密码：")
		fmt.Scanln(&this.Pwd)
		if this.Name != "ljw" || this.Pwd != "123" {
			fmt.Println("用户名或密码不正确！请重新输入！")
			this.Count--
			fmt.Printf("你还有%v次重试机会\n", this.Count)
			if this.Count == 0 {
				break
			}

		} else {
			fmt.Println("恭喜你！login sucess!")
			break
		}

	}
}
