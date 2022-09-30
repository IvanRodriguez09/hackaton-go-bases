package service

import (
	"fmt"
)

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	for _, ticket := range b.Tickets {
		if ticket.Id == t.Id {
			panic("Ticket already exists")
		}
	}
	b.Tickets = append(b.Tickets, t)
	fmt.Println("Ticket created successfully")
	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	var exists bool = false
	var ticket Ticket
	for _, t := range b.Tickets {
		if t.Id == id {
			exists = true
			ticket = t
			fmt.Println("Ticket read successfully")
		}
	}
	if !exists {
		panic("ticket not found")
	}
	return ticket, nil
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	var exists bool = false
	for _, ticket := range b.Tickets {
		if ticket.Id == id {
			exists = true
			ticket = t
			fmt.Println("Ticket updated successfully")
		}
	}
	if !exists {
		panic("ticket not found")
	}
	return t, nil
}

func (b *bookings) Delete(id int) (int, error) {
	var exists bool = false
	for _, ticket := range b.Tickets {
		if ticket.Id == id {
			b.Tickets = append(b.Tickets[:id-1], b.Tickets[id:]...)
			fmt.Println("Ticket removed successfully")
			exists = true
			return 0, nil
		}
	}
	if !exists {
		panic("ticket not found")
	}
	return 0, nil
}
