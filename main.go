package main

import (
	"fmt"
	"github.com/bolt-master/bolt-master"
)

func main() {
	//mode: 模式
	db, err := bolt.Open("datas.db",0600,nil)
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	fmt.Println(db)
	//1、设置数据
	db.Update(func(tx *bolt.Tx) error {
		//Tx: transaction: 交易
		tong1, err := tx.CreateBucket([]byte("male"))
		if err != nil {
			panic(err.Error())
		}
		//存储第一个男生信息
		tong1.Put([]byte("516541565"),[]byte("于洪伟"))
		tong1.Put([]byte("516194985"),[]byte("隔壁老王"))

		//tong2
		tong2, err := tx.CreateBucket([]byte("feamle"))
		if err != nil {
			panic(err.Error())
		}
		tong2.Put([]byte("52511985"),[]byte("翠花"))
		return nil
	})

	//2、读取数据
	db.View(func(tx *bolt.Tx) error {
		//读取数据
		bucket := tx.Bucket([]byte("male"))

		if bucket != nil {
			panic("读取数据失败，请重试")
		}
		dataBytes := bucket.Get([]byte("516541565"))
		fmt.Println("读取到的数据：",dataBytes)

		return nil
	})

}


