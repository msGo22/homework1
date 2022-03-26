package main

import (
	"fmt"
	"github.com/msGo22/homework1/domains"
	"log"
)

func main()  {
	alameddin, err := domains.NewCustomer("1234567", "Alameddin", "Çelik", "istanbul üsküdar", "05417907817")
	if err != nil {
		log.Fatal(err)
	}
	necati, err := domains.NewCustomer("23456", "Necati", "Şaşmaz", "istanbul kadıköy", "05110023200")
	if err != nil {
		log.Fatal(err)
	}
	ilkSiparis, err := domains.NewOrder(alameddin, necati, "1 kg leblebi")
	if err != nil {
		log.Fatal(err)
	}

	// Kargo işlemleri
	fmt.Println(ilkSiparis.Check())
	if err := ilkSiparis.PickedUp(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(ilkSiparis.Check())
	if err := ilkSiparis.Delivered(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(ilkSiparis.Check())

}
