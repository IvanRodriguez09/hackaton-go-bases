package file

import (
	"strconv"
	"strings"

	"github.com/IvanRodriguez09/hackaton-go-bases/internal/service"
)

type File struct {
	Path    string
	Content []string
}

func (f *File) Read() ([]service.Ticket, error) {
	var tickets []service.Ticket
	for _, line := range f.Content {
		if len(line) != 0 {
			data := strings.Split(line, ",")
			id, errid := strconv.ParseInt(data[0], 10, 0)
			price, errprice := strconv.ParseInt(data[5], 10, 0)
			if errprice != nil || errid != nil {
				panic("Error parsing id or price")
			}
			ticket := service.Ticket{
				Id:          int(id),
				Names:       data[1],
				Email:       data[2],
				Destination: data[3],
				Date:        data[4],
				Price:       int(price),
			}
			tickets = append(tickets, ticket)
		}
	}
	return tickets, nil
}

func (f *File) Write(service.Ticket) error {
	return nil
}
