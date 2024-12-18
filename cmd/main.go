package main

import (
	"log"
	redis_db "queuing_system/db"
	"sync"
	"time"
)
var unique_name = "me"

func main() {
	rdb:=redis_db.InitRedis()
	// rdb.Rbd.HSet(rdb.Ctx,"Write_lock","booth_1","0")

	var wait sync.WaitGroup
	wait.Add(1)

	go func ()  {
		defer wait.Done()
		for true {
			val,_:=rdb.Rbd.HGet(rdb.Ctx,"Write_lock","booth_1").Result()
			if val == "0"{
				rdb.Rbd.HSet(rdb.Ctx,"Write_lock","booth_1",unique_name)
				log.Println("Lock accquired by ",unique_name)	
				break
			}
			log.Println(unique_name," waiting, lock is acquired by",val)
			time.Sleep(time.Second*1)	
		}
		time.Sleep(20*time.Second)
		rdb.Rbd.HSet(rdb.Ctx,"Write_lock","booth_1","0")
		log.Println("work done by ",unique_name)	
	}()


	wait.Wait()
}