package main

import (
	"testing"
	"time"
)

/* rpc
公平发放模式下,不会丢失信息
*/
func TestRpc(t *testing.T) {
	go rpcClient(12)
	go rpcClient(11)
	time.Sleep(time.Second)
	go rpcServer()
	time.Sleep(time.Second)
}

/* 公平发放模式

 */
func Test1(t *testing.T) {
	newTask()

	go worker(time.Millisecond*100, 1, false)
	go worker(time.Millisecond*200, 2, false)
	go worker(time.Millisecond*300, 3, false)

	time.Sleep(time.Second * 20)

}

// 轮询模式
func TestPolling(t *testing.T) {
	send()

	go recv(time.Millisecond*100, 1)
	go recv(time.Millisecond*100, 2)
	go recv(time.Millisecond*100, 3)

	time.Sleep(time.Second * 5)

}

// 订阅模式

// topic模式
