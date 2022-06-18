package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//建立模型

type Admin struct {
	gorm.Model
	Name string
	Psw  int64
}

func main() {
	//1.连接数据库
	db, err := gorm.Open("mysql", "root:LKD2020.0921.@tcp/gorm?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接数据库失败!", err.Error())
	}

	//2。将模型关联到mysql
	db.AutoMigrate(&Admin{})
	//3.增 创建数据
	a1 := Admin{Name: "zxw", Psw: 20202524}
	db.Create(&a1)
	//a2 := Admin{Name: "fumin", Psw: 20212703}
	//db.Create(&a2)
	//a0 := Admin{Name: "RobKing", Psw: 20202021}
	//db.Create(&a0)
	//fmt.Println("插入成功！")
	//4.查
	/*
		//第一条数据
		var admin Admin
		db.First(&admin)
		fmt.Printf("第一条数据:%#v\n", admin)
		//最后一条数据
		a1 := new(Admin)
		db.Last(&a1)
		fmt.Printf("最后一条数据:%#v\n", a1)
		//指定名字
		a2 := new(Admin)
		db.Where("Name=?", "zxw").First(&a2)
		fmt.Printf("Name为zxw数据:%#v\n", a2)
		//通过主键
		a3 := new(Admin)
		db.First(&a3, 2)
		fmt.Printf("主键为2的数据:%#v\n", a3)
		//like查询
		a4 := new(Admin)
		db.Where("Name like ?", "%fu%").First(&a4)
		fmt.Printf("Name像fu的数据:%#v\n", a4)
		//struct map 查询
		a5 := new(Admin)
		db.Where(&Admin{Name:"RobKing", Psw: 20202021}).First(&a5)
		fmt.Printf("Name为Robking，Psw为20202021的数据:%#v\n", a5)
		//内联条件
		a6 := new(Admin)
		db.Find(&a6, "Name=?", "zxw")
		fmt.Printf("Name为zxw的数据:%#v\n", a6)
		//firstorinit
		//a13 := new(Admin)
		//db.FirstOrInit(&a13, Admin{Name: "zzz"})
		//db.FirstOrCreate(&a13, Admin{Name: "zzz"})
		//fmt.Printf("Name为zzz的数据:%#v\n", a13)
		////Not条件
		//a7 := new(Admin)
		//db.Not("Name=?", "fumin")
		//fmt.Printf("Name不为fumin的数据:%#v\n", a7)
		////Or条件
		//a8 := new(Admin)
		//db.Find(&a6, "Name=?", "zxw")
		//fmt.Printf("Name为zxw的数据:%#v\n", a8)
		//Order
		a9 := new(Admin)
		db.Order("Psw desc").Order("Name").Find(&a9)
		fmt.Printf("按照密码排序的数据:%#v\n", a9)
		////Limit Offset
		//a10 := new(Admin)
		//db.Find(&a6, "Name=?", "zxw")
		//fmt.Printf("Name为zxw的数据:%#v\n", a10)
		////Joins
		//a11 := new(Admin)
		//db.Find(&a6, "Name=?", "zxw")
		//fmt.Printf("Name为zxw的数据:%#v\n", a11)
		////Scan
		//a12 := new(Admin)
		//db.Find(&a6, "Name=?", "zxw")
		//fmt.Printf("Name为zxw的数据:%#v\n", a12)

	*/
	//5.改
	//a1 := new(Admin)
	//db.Debug().First(&a1)
	//fmt.Printf("更新前第一条的数据:%#v\n", a1)
	//a1.Name = "zxw"
	//a1.Psw = 20202524
	//db.Debug().Save(&a1)
	//db.Debug().Model(&a1).Updates(map[string]interface{}{
	//	"Name":"ZXW",
	//	"Psw":20202524,
	//})
	//db.Debug().First(&a1)
	//让所有的密码加1000
	//db.Model(&Admin{}).Update("Psw", gorm.Expr("Psw+?", 1000))
	//fmt.Printf("更新后第一条的数据:%#v\n", a1)
	//6.删
	//a1 := new(Admin)
	//fmt.Printf("删除前最后一条的数据:%#v\n", a1)
	//db.Debug().Unscoped().Where("Name=?", "zzz").Delete(&a1)
	//fmt.Printf("删除后最后一条的数据:%#v\n", a1)
	db.Unscoped().Where("Name like ?", "%RobK%")
	fmt.Printf("更新后第一条的数据:%#v\n", a1)
}
