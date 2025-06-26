package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type Students struct {
	Id         int64     `gorm:"column:id"`
	Name       string    `gorm:"column:name"`
	Age        int64     `gorm:"column:age"`
	Grade      string    `gorm:"column:grade"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

// Account 表示账户信息
type Account struct {
	ID      int64 `gorm:"primaryKey"`
	Balance float64
}

// Transaction 表示转账记录
type Transactions struct {
	ID            int64 `gorm:"primaryKey"`
	FromAccountID int64
	ToAccountID   int64
	Amount        float64
}

func main() {
	//题目一
	db := ConnDB()
	//var students []Students

	// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	//student := Students{
	//	Id:         0,
	//	Name:       "张三",
	//	Age:        20,
	//	Grade:      "三年级",
	//	CreateTime: time.Now(),
	//	UpdateTime: time.Now(),
	//}
	//db.Save(&student)

	//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	//if err := db.Where("age>?", 18).Find(&students).Error; err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	//updateData := make(map[string]interface{})
	//updateData["grade"] = "四年级"
	//updateData["update_time"] = time.Now()
	//if err := db.Table("students").Where("name =?", "张三").Updates(updateData).Error; err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	//if err := db.Where("age < ?", "15").Delete(&students).Error; err != nil {
	//	fmt.Println(err)
	//	return
	//}

	/*编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，
	需要先检查账户 A 的余额是否足够，
	如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
	并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
	*/
	// 开启事务
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // 发生 panic 时回滚事务
			log.Fatalf("失败并回滚 %v", r)
		}
	}()

	// 假设账户 A 和 B 的 ID 分别为 1 和 2
	fromAccountID := int64(1)
	toAccountID := int64(2)
	amount := float64(100.0)

	// 查询账户 A 的余额
	var fromAccount Account
	var toAccount Account
	if err := tx.Table("account").Find(&fromAccount, fromAccountID).Error; err != nil {
		tx.Rollback()
		log.Fatalf("failed to query fromAccountID from account: %v", err)
	}

	if err := tx.Table("account").Find(&toAccount, toAccountID).Error; err != nil {
		tx.Rollback()
		log.Fatalf("failed to query toAccountID from account: %v", err)
	}

	if fromAccount.Balance < amount {
		tx.Rollback()
		log.Fatalf("fromAccount.Balance: %v", fromAccount.Balance)
	}

	// 发送账户扣除100
	if err := tx.Table("account").Where("id", fromAccount.ID).Update("balance", (fromAccount.Balance - amount)).Error; err != nil {
		tx.Rollback()
		log.Fatalf("failed to update from account: %v", err)
	}
	// 接收账户添加100
	if err := tx.Table("account").Where("id", toAccount.ID).Update("balance", (toAccount.Balance + amount)).Error; err != nil {
		tx.Rollback()
		log.Fatalf("failed to update from account: %v", err)
	}

	// 转账记录表中插入转账数据
	//var transaction Transactions
	//transaction.Amount = amount
	//transaction.FromAccountID = fromAccountID
	//transaction.ToAccountID = toAccountID
	//transaction := Transactions{
	//	FromAccountID: fromAccountID,
	//	ToAccountID:   toAccountID,
	//	Amount:        amount,
	//}

	if err := tx.Exec("Insert into transaction (from_account_id,to_account_id,amount) value (?,?,?)", fromAccountID, toAccountID, amount).Error; err != nil {
		//if err := tx.Table("transaction").Create(transaction).Error; err != nil {
		tx.Rollback()
		log.Fatalf("save  transaction info  to transaction: %v", err)
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		log.Fatalf("failed to commit: %v", err)
	}
	log.Println(fromAccount)
}

func ConnDB() *gorm.DB {
	dsn := "root:admin@tcp(127.0.0.1:3306)/gotask?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlDb, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	db, err := mysqlDb.DB()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(2)
	mysqlDb.Debug()
	return mysqlDb
}
