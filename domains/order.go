package domains

import (
	"errors"
	"github.com/google/uuid"
	"log"
)

const (
	OrderNew       = "ORDER_NEW"
	OrderPickedUp  = "ORDER_PICKED_UP"
	OrderDelivered = "ORDER_DELIVERED"
	OrderCancelled = "ORDER_CANCELLED"
)

var OrderList map[uuid.UUID]*Order

type Order struct {
	ID       uuid.UUID `json:"id"`
	Status   string    `json:"status"`
	Details  string    `json:"details"`
	Sender   *Customer `json:"sender "`
	Receiver *Customer `json:"receiver"`
}

func init() {
	log.Println("Order Paketi Yüklendi...")
	// orderList için boş bir değişken oluşturup bellekte kullanılabilir alan oluşturuyoruz.
	OrderList = make(map[uuid.UUID]*Order)
}

// NewOrder yeni sipariş oluşturulur
func NewOrder(sender, receiver *Customer, details string) (*Order, error) {
	order := &Order{
		ID:       uuid.New(),
		Status:   OrderNew,
		Details:  details,
		Sender:   sender,
		Receiver: receiver,
	}
	if err := order.validate(); err != nil {
		return nil, err
	}
	return order, nil

}

func (o *Order) validate() error {
	if o.Sender == nil || o.Receiver == nil {
		return errors.New("Gönderici ve Alıcı boş gönderilemez")
	}
	if o.Sender == o.Receiver {
		return errors.New("Gönderici ve Alıcı aynı kişi olamaz")
	}
	return nil
}

// Check sipariş durumunu döner
func (o *Order) Check() string {
	return o.Status
}

// PickedUp Yeni gelmiş siparişlerin teslimat adımına geçmesini sağlar
func (o *Order) PickedUp() error {
	if o.Status != OrderNew {
		return errors.New("Sipariş teslimat adımı için uygun değildir")
	}
	o.Status = OrderPickedUp
	return nil
}

// Delivered teslim edilme işlemini eğer sipariş yola çıkmışsa gerçekleştirir
func (o *Order) Delivered() error {
	if o.Status != OrderPickedUp {
		return errors.New("yola çıkmamış kargolar teslim edildi durumuna geçemez.")
	}
	o.Status = OrderDelivered
	return nil
}

// Cancelled iptal işlemini eğer teslim edilmediyse gerçekleştirir
func (o *Order) Cancelled() error {
	if o.Status == OrderDelivered {
		return errors.New("teslim edilmiş kargolar iptal edilemez")
	}
	o.Status = OrderCancelled
	return nil
}
