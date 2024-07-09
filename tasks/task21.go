package main

import "fmt"

type Charger interface { // Target
	Charging()
}

type Phone struct{}

func (ph *Phone) Charging() {
	fmt.Println("Charging Mobile")
}

type Notebook struct{} // adaptee

func (nb *Notebook) ChargeNotebook() {
	fmt.Println("Charging Notebook")
}

type ChargerAdapter struct { // adapter
	Notebook *Notebook
}

func (ca *ChargerAdapter) Charging() {
	ca.Notebook.ChargeNotebook()
}

type Client struct{}

func (cl *Client) ChargeMyMobile(charger Charger) {
	charger.Charging()
}

func main() {
	phone := &Phone{}
	notebook := &Notebook{}
	client := &Client{}

	chargerAdapter := &ChargerAdapter{
		Notebook: notebook,
	}

	client.ChargeMyMobile(phone)
	client.ChargeMyMobile(chargerAdapter)
}
