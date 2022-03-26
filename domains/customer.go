package domains

import (
	"errors"
	"log"
)

// Json eklememizin nedeni ileriki derslerde sunucu olarak dışarı veriyi aktarırken nasıl tanımlanacağını göstermemiz içindir
type Customer struct {
	TcID      string   `json:"tc_id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Location  string   `json:"location"`
	Phone     string   `json:"phone"`
	Orders    []*Order `json:"orders"`
}

var CustomerList map[string]*Customer

func init() {
	log.Println("Customer Paketi Yüklendi...")
	// customerList için boş bir değişken oluşturup bellekte kullanılabilir alan oluşturuyoruz.
	CustomerList = make(map[string]*Customer)
}

// NewCustomer yeni müşteri oluşturur
func NewCustomer(tcID, firstName, lastName, location, phone string) (*Customer, error) {
	// yeni kullanıcı oluşturuyoruz
	customer := &Customer{
		TcID:      tcID,
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		Location:  location,
		Orders:    nil,
	}
	if err := customer.validate(); err != nil {
		return nil, err
	}

	// Tüm Müşterilerin bulunduğu listeye ekliyoruz
	CustomerList[customer.TcID] = customer
	return customer, nil
}

// validate müşteri oluşturma bilgilerini kontrol eder
func (c *Customer) validate() error {
	if c.TcID == ""{
		return errors.New("TC kimlik zorunludur")
	}
	if _, ok := CustomerList[c.TcID]; ok {
		return errors.New("Kullanıcı daha önceden kayıt oluşturmuştur")
	}
	return nil
}

// ChangeLocation müşteri adresini değiştirir
func (c *Customer) ChangeLocation(newAddress string) error {
	if c.Location == newAddress{
		return errors.New("şu anki adresiniz ile yeni girilen adres aynı olamaz")
	}
	c.Location = newAddress
	return nil
}
