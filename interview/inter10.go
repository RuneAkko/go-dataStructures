package main

/**
GO里面MAP:
实现: key不存在 get操作等待 直到key存在或者超时，保证并发安全，且需要实现以下接口：
type sp interface {
    Out(key string, val interface{})  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
    Rd(key string, timeout time.Duration) interface{}  //读取一个key，如果key不存在阻塞，等待key存在或者超时
}
????
*/

func inter10() {

}
