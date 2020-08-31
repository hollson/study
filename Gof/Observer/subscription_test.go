package observer

import (
	"fmt"
	"testing"
)

func TestObserver(t *testing.T) {
	buyer := &Buyer{"阿爽"}
	seller := &Seller{"杜蕾"}
	courier := &Courier{"小顺"}

	var o OrderSubject = new(OrderService)
	o.Subscribe(buyer)
	o.Subscribe(seller)
	o.Subscribe(courier)
	o.Publish(&Order{OrderId: 1001, ProductId: "2020KKMJI01", Amount: 99.99})

	fmt.Println()
	o.Unsubscribe(courier) // 不用快递
	o.Publish(&Order{OrderId: 1002, ProductId: "2020MMMSS02", Amount: 88.88, Mark: "自提"})
}
