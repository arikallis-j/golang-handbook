package main

import (
	"fmt"
	// импортируем пакет из поддиректории internal/contacts
	"contacts"
)

func test_scope() {
	contacts.SetSupport("Служба поддержки")
	fmt.Println(contacts.GetContact())
	fmt.Println("Email:", contacts.Email)
}
