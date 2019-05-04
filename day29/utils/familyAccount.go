package utils

import "fmt"

type FamilyAccount struct {
	key     string
	balance float64
	money   float64
	note    string
	flag    bool
	details string
	choice  string
	loop    bool
}

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		balance: 10000,
		money:   0.0,
		note:    "",
		details: "账户\t账户金额\t收支金额\t\t说 明",
		flag:    false,
		choice:  "",
		loop:true,
	}
}

func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n------------------家庭收支记账软件--------------------")
		fmt.Println("                  1. 收支明细")
		fmt.Println("                  2. 登记收入")
		fmt.Println("                  3. 登记支出")
		fmt.Println("                  4. 转账")
		fmt.Println("                  5. 退出")
		fmt.Println("请选择(1-4):")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showinfo()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.tranfer()
		case "5":
			this.exit()
		default:
			fmt.Println("请输入正确的选项")
		}
		if !this.loop {
			break
		}
	}
}

func (this *FamilyAccount) showinfo() {
	fmt.Println("------------------当前收支明细记录--------------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有任何收支，请来一笔把！")
	}
}

func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) pay() {
	fmt.Println("本次支出金额")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("本次支出余额不足")
	}
	this.balance -= this.money
	fmt.Println("本次支出说明")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) tranfer() {
	fmt.Println("本次转账金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("所剩余额不足")
	}
	this.balance -= this.money
	fmt.Println("本次转账说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n转账\t%v\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}
func (this *FamilyAccount) exit() {
	fmt.Println("你确定要退出？(y/n)")

	for {
		fmt.Scanln(&this.choice)
		if this.choice == "y" || this.choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入！(y/n)")

	}
	if this.choice == "y" {
		this.loop = false
	}

}
