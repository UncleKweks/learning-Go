package main

import "fmt"

type Contact struct {
	ID int
	Name string 
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextID int = 1


func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func AddContact(name string, email string, phone string) {
	contact := Contact{
		ID: nextID,
		Name: name,
		Email: email,
		Phone: phone,
	}
	contactList = append(contactList, contact)
	contactIndexByName[name] = len(contactList) - 1
	nextID++
}

func GetContactByName(name string) *Contact {
	if _, exists := contactIndexByName[name]; exists {
		return &contactList[contactIndexByName[name]]
	}
	return nil
}

func ListContacts() []Contact {
	fmt.Println("Listing all contacts:")
	if len(contactList) == 0 {
		fmt.Println("No contacts found.")
		return contactList
	}
	for _, contact := range contactList {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Phone: %s\n", contact.ID, contact.Name, contact.Email, contact.Phone)
	}
	return contactList
}

func main() {

	AddContact("John Doe", "john@example.com", "555-1234")
	AddContact("Jane Smith", "jane@example.com", "555-5678")
	AddContact("Alice Johnson", "alice@example.com", "555-9012")
	AddContact("Bob Brown", "bob@example.com", "555-0123")
	
	ListContacts()

	bob := GetContactByName("Bob Brown")
	if bob != nil {
		fmt.Printf("Found contact: ID: %d, Name: %s, Email: %s, Phone: %s\n", bob.ID, bob.Name, bob.Email, bob.Phone)
	} else {
		fmt.Println("Contact not found.")
	}

}
