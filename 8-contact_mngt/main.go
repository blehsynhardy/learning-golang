package main

import "fmt"

type Contact struct {
	Id    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact
var contactByIndexName map[string]int
var nextId int = 1

func init() {
	contactList = make([]Contact, 0)
	contactByIndexName = make(map[string]int)
}

func addContact(name, email, phone string) {
	if _, ok := contactByIndexName[name]; ok {
		fmt.Printf("Contact already exists: %v\n", name)
		return
	}

	newContact := Contact{
		Id:    nextId,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	nextId++

	contactList = append(contactList, newContact)
	contactByIndexName[name] = len(contactList) - 1
	fmt.Printf("Contact added: %v\n", name)
}

func findContact(name string) *Contact {
	index, exists := contactByIndexName[name]
	if exists {
		return &contactList[index]
	}
	return nil
}

func listContacts() {
	fmt.Println("===============Lisiting contact list=============")

	if len(contactList) == 0 {
		fmt.Println("No contact list")
		return
	}
	for i, contact := range contactList {
		fmt.Printf("%d. ID: %d, Name %s, Email %s, Phone %s\n", i+1, contact.Id, contact.Name, contact.Email, contact.Phone)
	}

	fmt.Println("==================")
}

func main() {

	addContact("Alice Wonderland", "blessingawodele@gmail.com", "1233232332")
	addContact("Michael Smith", "blessingthrushka@gmail.com", "00289823823")
	addContact("Michael Smith", "blessingthrushka@gmail.com", "00289823823")

	listContacts()

	alice := findContact("Alice Wonderland")

	if alice != nil {
		fmt.Printf("Alice Wonderland: %v\n", alice.Name)
	} else {
		fmt.Println("Alice Wonderland not found")
	}
}
