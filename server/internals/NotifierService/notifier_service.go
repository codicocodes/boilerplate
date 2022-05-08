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

func (s *NotifierService) disconnect(id string) {
	defer fmt.Printf("Closing channel for user: %s", id)
	s.mu.Lock()
	defer s.mu.Unlock()
	channel := s.clients[id]
	defer close(channel)
	delete(s.clients, id)
}

func (s *NotifierService) Connect() (chan string, func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	uid, _ := uuid.NewRandom()
	id := uid.String()
	clientChan := make(chan string)
	s.clients[id] = clientChan
	fmt.Printf("Registered new client: %d clients registered\n", len(s.clients))
	return clientChan, func() {
		s.disconnect(id)
	}
}

func (s *NotifierService) BroadcastAll(msg string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	fmt.Printf("Broadcasting \"%s\" to %d clients\n", msg, len(s.clients))
	for _, channel := range s.clients {
		channel <- msg
	}
}
