package observer

import "fmt"

// 抽象观察者
type Subscriber interface {
	Handler(*Order)
}

// 具体观察者(卖家)
type Seller struct {
	Name string
}

func (s *Seller) Handler(o *Order) {
	fmt.Printf("【卖家】%s - 接到订单(订单编号:%d),正在出库中.\n", s.Name, o.OrderId)
}

// 具体观察者(快递员)
type Courier struct {
	Name string
}

func (s *Courier) Handler(o *Order) {
	fmt.Printf("【快递】%s - 收到快件(订单编号:%d),加紧配送中.\n", s.Name, o.OrderId)
}

// 具体观察者(买家)
type Buyer struct {
	Name string
}

func (s *Buyer) Handler(o *Order) {
	fmt.Printf("【买家】%s - 下单成功(订单编号:%d),耐心等待哦.\n", s.Name, o.OrderId)
}
