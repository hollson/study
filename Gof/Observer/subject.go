// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package observer

type Order struct {
	OrderId   int64   `json:"id"`         // 订单编号
	ProductId string  `json:"product_id"` // 商品编号
	Amount    float64 `json:"amount"`     // 订单金额
	Mark      string  `json:"mark"`       // 备注
}

// 抽象订单主题
type OrderSubject interface {
	Subscribe(Subscriber)   // 开放订阅
	Unsubscribe(Subscriber) // 开放退订
	Publish(*Order)         // 发布订单
}

// 具体订单业务，实现了订单主题接口
type OrderService struct {
	Subscribers []Subscriber
}

func (s *OrderService) Subscribe(f Subscriber) {
	if s == nil {
		return
	}
	for _, c := range s.Subscribers {
		if c == f {
			return
		}
	}
	s.Subscribers = append(s.Subscribers, f)
}

func (s *OrderService) Unsubscribe(f Subscriber) {
	if s == nil {
		return
	}
	for i, c := range s.Subscribers {
		if c == f {
			s.Subscribers = append(s.Subscribers[:i], s.Subscribers[i+1:]...)
		}
	}
}

func (s *OrderService) Publish(o *Order) {
	if s == nil {
		return
	}
	for _, v := range s.Subscribers {
		v.Handler(o)
	}
}
