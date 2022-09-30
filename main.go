package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/IvanRodriguez09/hackaton-go-bases/internal/file"
	"github.com/IvanRodriguez09/hackaton-go-bases/internal/service"
)

const filepath string = "./tickets.csv"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Function to get tickets from csv file
func GetFileTickets() []service.Ticket {
	defer func() {
		err := recover()
		if err != nil {
			println("error getting tickets from csv file")
		}
	}()
	// Declare file struct
	filesvc := file.File{
		Path: filepath,
	}
	// Reading the file svc
	fileData, err := os.ReadFile(filesvc.Path)
	check(err)
	filesvc.Content = strings.Split(string(fileData), "\n")
	// Get ticket slice
	tickets, err := filesvc.Read()
	check(err)
	return tickets
}

func TestOperations(booking service.Bookings) {
	defer func() {
		if err := recover(); err != nil {
			println("Something went wrong with bookings")
		}
		fmt.Println("checking operations successfully")
		fmt.Printf("%v\n", booking)
	}()
	ticket := service.Ticket{
		Id:          1001,
		Names:       "Ivan Rodriguez",
		Email:       "ivan@email.com",
		Destination: "Bahamas",
		Date:        "6:00",
		Price:       250,
	}
	_, err := booking.Create(ticket)
	check(err)
	ticket, err = booking.Read(1001)
	check(err)
	ticket.Price = 300
	check(err)
	ticket, err = booking.Update(1001, ticket)
	check(err)
	_, err = booking.Delete(1002)
	check(err)
}

func main() {
	// Funcion para obtener tickets del archivo csv
	tickets := GetFileTickets()
	booking := service.NewBookings(tickets)
	TestOperations(booking)
}
