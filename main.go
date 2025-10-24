package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// Save contacts in json file
func saveContacsToFile(contacs []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacs)
	if err != nil {
		return err
	}

	return nil
}

// Load contacts from file json
func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	//Slice by contacts
	var contacts []Contact

	//load exist contacts
	err := loadContactsFromFile(&contacts)

	if err != nil {
		fmt.Println("Error al cargar los contactos: ", err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("=== Gestor de contactos === \n",
			"1. Agregar contacto \n",
			"2. Mostrar todos los contactos \n",
			"3. Salir \n",
			"Elige una opcion: ")

		var option int
		_, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al leer la opcion: ", err)
			return
		}

		//Manejar las opciones del usuario
		switch option {
		case 1:
			// Ingrear y crear contacto
			var c Contact
			fmt.Print("Ingrese nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Ingrese Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Ingrese telefono: ")
			c.Phone, _ = reader.ReadString('\n')

			//Agregar un concto a slice
			contacts = append(contacts, c)

			//Guardar contacto en un archivo json
			if err := saveContacsToFile(contacts); err != nil {
				fmt.Println("Error al guardar el contacto", err)
			}
		case 2:
			fmt.Println("=============================")
			for index, contact := range contacts {
				fmt.Printf("%d. Nombre: %s Email: %s Telefomo: %s\n", index+1, contact.Name, contact.Email, contact.Email)
			}
			fmt.Println("=======================")

		case 3:
			return

		default:
			fmt.Println("Opcion invalida")
		}

	}

}
