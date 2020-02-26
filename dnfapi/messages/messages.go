package messages
//这是一个用于秒杀的模块,暂时没有什么作用

import (
	"fmt"
	"strconv"
	"sync/atomic"
	"time"
)

type iItem struct{
	ch chan interface{}
	item int64
}
var items iItem
var timer *time.Timer
func init()  {
	items=iItem{
		ch:make(chan interface{}),
		item:10,
	}
	serviceStart();
}

func resetStart(){
	 fmt.Println("刷新模块启动")
	timer=time.NewTimer(time.Second*10)
	for{
		_=<-timer.C
		if(items.item<1){
			items.item=10;
			fmt.Println("刷新了物品数量,剩余数量:"+strconv.FormatInt(items.item,10))
		}
		
	}
}

func serviceStart(){
	defer fmt.Println("消息队列启动")
	for i:=0;i<10;i++{
		go saleitem()
	}
}
func saleitem(){
	for{
		w:=<-items.ch;
		if(timer.Reset(time.Second*10)){

		}else{
			timer=time.NewTimer(time.Second*10)
		}
		if(items.item>0){
			atomic.AddInt64(&items.item,-1)
			w.(chan int64)<-atomic.LoadInt64(&items.item)
			fmt.Println("剩余数量:"+strconv.FormatInt(items.item,10))
		}else{
			items.item=-1
			w.(chan int64)<-items.item
		}
		
	}
}
//Push a
func Push(msg interface{})  {
	items.ch<- msg
}