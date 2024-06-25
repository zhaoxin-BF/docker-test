package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"sync"
	"time"
)

type Card struct {
	ID        string
	Value     string
	Redeemed  bool
	CreatedAt time.Time
}

type CardStore interface {
	CreateCard(value string) (*Card, error)
	RedeemCard(id string) (*Card, error)
	GetCard(id string) (*Card, error)
}

type MemoryCardStore struct {
	cards map[string]*Card
	mu    sync.Mutex
}

func NewMemoryCardStore() *MemoryCardStore {
	return &MemoryCardStore{
		cards: make(map[string]*Card),
	}
}

func (s *MemoryCardStore) CreateCard(value string) (*Card, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := uuid.New().String()
	card := &Card{
		ID:        id,
		Value:     value,
		CreatedAt: time.Now(),
	}
	s.cards[id] = card
	return card, nil
}

func (s *MemoryCardStore) RedeemCard(id string) (*Card, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	card, ok := s.cards[id]
	if !ok {
		return nil, fmt.Errorf("card not found")
	}
	if card.Redeemed {
		return nil, fmt.Errorf("card already redeemed")
	}
	card.Redeemed = true
	return card, nil
}

func (s *MemoryCardStore) GetCard(id string) (*Card, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	card, ok := s.cards[id]
	if !ok {
		return nil, fmt.Errorf("card not found")
	}
	return card, nil
}

func main() {
	store := NewMemoryCardStore()
	router := gin.Default()

	router.POST("/cards", func(c *gin.Context) {
		value := c.PostForm("value")
		card, err := store.CreateCard(value)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, card)
	})

	router.POST("/cards/:id/redeem", func(c *gin.Context) {
		id := c.Param("id")
		card, err := store.RedeemCard(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, card)
	})

	router.GET("/cards/:id", func(c *gin.Context) {
		id := c.Param("id")
		card, err := store.GetCard(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, card)
	})

	router.Run()
}
