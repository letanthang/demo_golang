/*
**
detect race condition by: $ go run -race main.go
race condition vs data race ??????

Data Race - định nghĩa :
+ Từ 2 thread/process trở lên cùng truy cập vào vùng nhớ chung (shared resource).
+ Ít nhất 1 thread/process thay đổi giá trị của vùng nhớ chung đó.

Race Condition: 
	Vấn đề sai sót về mặt thời gian hoặc thứ tự thực thi của các thread trong chương trình khiến cho
kết quả cuối cùng không đúng như mong muốn.

Race condition khó sửa hơn race data và không detect duoc bang $ go run -race
*/
package main

import (
	"app/race_condition/global_var"
	"app/race_condition/primitive_var"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// ConcurrentWriteMap()
	// ParallelWriteFixByCreate([]byte("HE HE"))
	// ProtectedGlobalVar()

	SafeWatchdog()
}

func SafeWatchdog() {
	wd := primitive_var.SafeWatchdog{}
	wd.Start()

	isRun := true
	time.AfterFunc(time.Second*5, func() { 	// this 's also a race data
		isRun = false						// commit these lines to fix !!!!
	})										// ....

	for isRun {
		wd.KeepAlive()
		time.Sleep(time.Millisecond * 100)
	}

	forever := make(chan int, 1)
	<-forever
}

func Watchdog() {
	wd := primitive_var.Watchdog{}
	wd.Start()

	isRun := true
	time.AfterFunc(time.Second*5, func() {
		isRun = false
	})

	for isRun {
		wd.KeepAlive()
		time.Sleep(time.Millisecond * 100)
	}

	forever := make(chan int, 1)
	<-forever
}

func UnprotectedGlobalVar() {
	addr1, err := net.ResolveTCPAddr("tcp", "1.1.1.1:27017")
	if err != nil {
		panic(err)
	}

	addr2, err := net.ResolveTCPAddr("tcp", "1.2.3.4:3600")
	if err != nil {
		panic(err)
	}

	go global_var.RegisterService("mongo", addr1)
	global_var.RegisterService("mysql", addr2)
}

func ProtectedGlobalVar() {
	addr1, err := net.ResolveTCPAddr("tcp", "1.1.1.1:27017")
	if err != nil {
		panic(err)
	}

	addr2, err := net.ResolveTCPAddr("tcp", "1.2.3.4:3600")
	if err != nil {
		panic(err)
	}

	go global_var.ProtectRegisterService("mongo", addr1)
	global_var.ProtectRegisterService("mysql", addr2)
}

func ParallelWriteFixByCreate(data []byte) chan error {
	res := make(chan error, 2)
	f1, err := os.Create("file1")
	if err != nil {
		res <- err
	} else {
		go func() {
			// This err is shared with the main goroutine,
			// so the write races with the write below.
			_, err := f1.Write(data)
			res <- err
			f1.Close()
		}()
	}
	f2, err := os.Create("file2") // The second conflicting write to err.
	if err != nil {
		res <- err
	} else {
		go func() {
			_, err := f2.Write(data)
			res <- err
			f2.Close()
		}()
	}
	return res
}

func ParallelWrite(data []byte) chan error {
	res := make(chan error, 2)
	f1, err := os.Create("file1")
	if err != nil {
		res <- err
	} else {
		go func() {
			// This err is shared with the main goroutine,
			// so the write races with the write below.
			_, err = f1.Write(data)
			res <- err
			f1.Close()
		}()
	}
	f2, err := os.Create("file2") // The second conflicting write to err.
	if err != nil {
		res <- err
	} else {
		go func() {
			_, err = f2.Write(data)
			res <- err
			f2.Close()
		}()
	}
	return res
}

func ConcurrentWriteMap() {
	done := make(chan bool)
	m := make(map[string]string)
	m["name"] = "world"
	go func() {
		m["name"] = "data race"
		done <- true
	}()
	fmt.Println("Hello,", m["name"])
	<-done
	fmt.Println("Hello,", m["name"])
}
