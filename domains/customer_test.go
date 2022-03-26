package domains_test

import (
	"github.com/msGo22/homework1/domains"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCustomer(t *testing.T) {
	type input struct {
		tcNo      string
		firstName string
		lastName  string
		location  string
		phone     string
	}
	type testCase struct {
		title string
		input
		result    *domains.Customer
		resultErr string
	}
	// tüm ihtimalleri test olarak yazdık
	// (Uyarı: Normal şartlarda duplicated Customer için farklı bağımsız bir test yazmamız gerekirdi)
	// (Uyarı açıklaması: Test case'ler birbirlerinden bağımsız veya karışık çalıştırılabilir bu durumda hata alınır)
	testCases := []testCase{
		{"Success New Customer", input{"123", "abc", "abc", "abc", "0123"}, &domains.Customer{TcID: "123", FirstName: "abc", LastName: "abc", Location: "abc", Phone: "0123"}, ""},
		{"Duplicated Customer", input{"123", "efg", "def", "bcd", "0124"}, nil, "Kullanıcı daha önceden kayıt oluşturmuştur"},
		{"Empty TC No Customer", input{"", "abc", "abc", "abc", "0123"}, nil, "TC kimlik zorunludur"},
	}
	for _, v := range testCases {
		// testCases adlı dizideki tüm testleri aşağıda denemek için birer t.Run ile eşzamanlı çalıştırıcılar ekledik
		t.Run(v.title, func(t *testing.T) {
			// testCase deki bilgileri kullanarak yeni bir kullanıcı oluşturduk
			customer, err := domains.NewCustomer(v.tcNo, v.firstName, v.lastName, v.location, v.phone)
			// çıktıdaki müşteri bilgisi bizim bilgiler ile aynı mı diye kontrol edilyoruz
			assert.Equal(t, v.result, customer)
			if err != nil {
				// eğer hatalı çıktı bekliyorsak o zaman hata mesajını error.Error() ile string olarak kontrol ediyoruz
				assert.Equal(t, v.resultErr, err.Error())
			}
		})
	}

}
