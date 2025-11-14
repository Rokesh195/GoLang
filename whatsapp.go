package main

import "fmt"

type mess struct {
	Media []string
}

type whatsappstruct struct {
	//User-chat-message-encryption-server-receiptient-device-media
	layout map[string]map[string]map[string]map[string]map[string]map[string]mess
}

func whatsapp() {
	whatsapp := make(map[string]whatsappstruct)
	whatsapp["User"] = whatsappstruct{layout: make(map[string]map[string]map[string]map[string]map[string]map[string]mess)}
	whatsapp["User"].layout["chat"] = make(map[string]map[string]map[string]map[string]map[string]mess)
	whatsapp["User"].layout["chat"]["Message"] = make(map[string]map[string]map[string]map[string]mess)
	whatsapp["User"].layout["chat"]["Message"]["Encryption"] = make(map[string]map[string]map[string]mess)
	whatsapp["User"].layout["chat"]["Message"]["Encryption"]["Server"] = make(map[string]map[string]mess)
	whatsapp["User"].layout["chat"]["Message"]["Encryption"]["Server"]["Recipient"] = make(map[string]mess)
	whatsapp["User"].layout["chat"]["Message"]["Encryption"]["Server"]["Recipient"]["Device"] = mess{
		Media: []string{"image", "text", "audio"},
	}

	fmt.Println((whatsapp))
}

func main() {
	whatsapp()
}
