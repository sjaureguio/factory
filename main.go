package main

import (
	"fmt"
)

type PayMethod interface {
	Pay()
}

type PayPal struct{}

type Culqi struct{}

func (p PayPal) Pay() {
	fmt.Println("Estas pagando con paypal")
}

func (c Culqi) Pay() {
	fmt.Println("Estas pagando con culqi")
}

func Factory(method uint8) PayMethod {
	switch method {
	case 1:
		return Culqi{}
	case 2:
		return PayPal{}
	default:
		return nil
	}
}

func main() {
	var method uint8
	fmt.Println("Ingrese el método de pago [1-3]: ")
	_, err := fmt.Scanln(&method)
	if err != nil || method > 2 {
		panic("Debe ingresar un método valido")
	}

	payMethod := Factory(method)
	payMethod.Pay()
}
