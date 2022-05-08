package notifierservice

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type NotifierService struct {
	mu      sync.Mutex
	clients map[string]chan string
}

func New() *NotifierService {
	clients := map[string]chan string{}
	return &NotifierService{
		clients: clients,
	}
}

func (s *NotifierService) Disconnect(id string) {
	fmt.Printf("Disconnecting client\n")
	s.mu.Lock()
	defer s.mu.Unlock()
	channel := s.clients[id]
	defer close(channel)
	delete(s.clients, id)
}

func (s *NotifierService) Connect() (chan string, string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	id, _ := uuid.NewRandom()
	clientChan := make(chan string)
	s.clients[id.String()] = clientChan
	fmt.Printf("Registered new client: %d clients registered\n", len(s.clients))
	return clientChan, id.String()
}

func (s *NotifierService) BroadcastAll(msg string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	fmt.Printf("Broadcasting \"%s\" to %d clients\n", msg, len(s.clients))
	for _, channel := range s.clients {
		channel <- msg
	}
}
