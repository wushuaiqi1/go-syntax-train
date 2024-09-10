package tcp

import (
	"fmt"
	"log"
	"net"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestScan(t *testing.T) {
	g := sync.WaitGroup{}
	for i := 1; i < 65535; i++ {
		if i >= 1000 {
			addr := fmt.Sprintf("www.baidu.com:%d", i)
			//addr := fmt.Sprintf("127.0.0.1:%d", i)
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				//log.Println(err)
				return
			}
			err = conn.Close()
			if err != nil {
				return
			}
			log.Println(fmt.Sprintf("连接%d端口成功", i))
		} else {
			g.Add(1)
			go func(x int) {
				//addr := fmt.Sprintf("www.baidu.com:%d", x)
				addr := fmt.Sprintf("127.0.0.1:%d", i)
				conn, err := net.Dial("tcp", addr)
				if err != nil {
					//log.Println(err)
					return
				}
				err = conn.Close()
				if err != nil {
					return
				}
				log.Println(fmt.Sprintf("连接%d端口成功", x))
				g.Add(-1)
			}(i)
		}
	}
	g.Wait()
}

func ScanPorts(domain, proto string, port int, res chan int) {
	addr := fmt.Sprintf(domain+":%d", port)
	conn, err := net.DialTimeout(proto, addr, 10*time.Second)
	if err != nil {
		res <- 0
		return
	}
	conn.Close()
	res <- port
}

func TestScan2(t *testing.T) {
	ip := "127.0.0.1"
	startPort := 1
	endPort := 65535
	var resArray []int
	res := make(chan int, endPort-startPort+1)
	for i := startPort; i <= endPort; i++ {
		go ScanPorts(ip, "tcp", i, res)
	}
	for i := startPort; i <= endPort; i++ {
		port := <-res
		if port != 0 {
			resArray = append(resArray, port)
		}
	}
	log.Println("通:", resArray)
}

func TestMax(t *testing.T) {
	maxGoroutines := 1000000 // 尝试创建100万个协程
	done := make(chan bool, maxGoroutines)

	for i := 0; i < maxGoroutines; i++ {
		go func() {
			done <- true
		}()
	}

	for i := 0; i < maxGoroutines; i++ {
		res := <-done
		log.Println(res)
	}

	fmt.Println("Successfully created", maxGoroutines, "goroutines")
	fmt.Println("Number of Goroutines at the end:", runtime.NumGoroutine())
}
