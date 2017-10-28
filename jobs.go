package main

import (
	"fmt"
)

// DeliverJob is 快递Job
type DeliverJob struct {
	FromAddr string
	ToAddr   string
	Goods    string
	Fee      string
}

func DefaultDeliverJob() {
	fmt.Println("DefaultDeliverJob")
}

func NewDeliverJob(from string, to string, goods string, fee string) *DeliverJob {
	return &DeliverJob{FromAddr: from, ToAddr: to, Goods: goods, Fee: fee}
}

func (djob DeliverJob) Run() {
	fmt.Println(djob.FromAddr, djob.ToAddr, djob.Goods, djob.Fee)
}
