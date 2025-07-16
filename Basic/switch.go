package main

import "fmt"

type PhoneBook struct {
	Name    string
	Number  string
	Company string
}

var phoneBookList = []PhoneBook{}

func main() {
	for {
		fmt.Println("\n--- Phone Book Menu ---")
		fmt.Println("1. Create Contact")
		fmt.Println("2. Read All Contacts")
		fmt.Println("3. Update Contact")
		fmt.Println("4. Delete Contact")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice (1-5): ")

		var choice int
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1:
			createContact()
		case 2:
			readAllContacts()
		case 3:
			updateContact()
		case 4:
			deleteContact()
		case 5:
			fmt.Println("Exiting Phonebook...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func createContact() {
	var name, number, company string
	fmt.Print("Enter contact name: ")
	fmt.Scanf("%s\n", &name)
	fmt.Print("Enter contact number: ")
	fmt.Scanf("%s\n", &number)
	fmt.Print("Enter contact company: ")
	fmt.Scanf("%s\n", &company)
	contact := PhoneBook{Name: name, Number: number, Company: company}
	phoneBookList = append(phoneBookList, contact)
	fmt.Printf("Contact created: %+v\n", contact)
}

func readAllContacts() {
	for i, contact := range phoneBookList {
		fmt.Printf("%d. %+v\n", i+1, contact)
	}
}

func updateContact() {
	// Implementation for updating a contact
	var name, number, company string
	fmt.Print("Enter contact name to update: ")
	fmt.Scanf("%s\n", &name)
	for i, contact := range phoneBookList {
		if contact.Name == name {

			fmt.Print("Enter contact name: ")
			fmt.Scanf("%s\n", &name)
			phoneBookList[i].Name = name

			fmt.Print("Enter contact number: ")
			fmt.Scanf("%s\n", &number)
			phoneBookList[i].Number = number

			fmt.Print("Enter contact company: ")
			fmt.Scanf("%s\n", &company)
			phoneBookList[i].Company = company
		}
	}
}

func deleteContact() {
	// Implementation for deleting a contact
	var name string
	fmt.Print("Enter contact name to delete: ")
	fmt.Scanf("%s\n", &name)
	for i, contact := range phoneBookList {
		if contact.Name == name {
			phoneBookList = append(phoneBookList[:i], phoneBookList[i+1:]...)
			fmt.Println("Contact deleted.")
			return
		}
	}
	fmt.Println("Contact not found.")
}
