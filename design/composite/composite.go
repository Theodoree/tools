package composite

import "sync"

// Composite Design Pattern

type SellerStats interface {
	SellerStats() uint64
	Name() string
}

type Manager interface {
	SellerStats
	AddSeller(seller Seller)
	DeleteSeller(seller Seller)
}

type manager struct {
	seller       map[Seller]struct{}
	sellerAmount uint64
	name         string
	sync.Mutex
}

func (m *manager) SellerStats() uint64 {
	return m.sellerAmount
}

func (m *manager) AddSeller(seller Seller) {
	m.Lock()
	if _, ok := m.seller[seller]; !ok {
		m.seller[seller] = struct{}{}
		m.sellerAmount += seller.SellerStats()
	}
	m.Unlock()
}

func (m *manager) DeleteSeller(seller Seller) {
	if m.sellerAmount < seller.SellerStats() {
		panic("uint64 溢出")
	}
	m.Lock()
	m.sellerAmount -= seller.SellerStats()
	delete(m.seller, seller)
	m.Unlock()
}

func (m *manager) Name() string {
	return m.name
}

func NewManager(name string) Manager {
	return &manager{
		name:   name,
		seller: make(map[Seller]struct{}),
		Mutex:  sync.Mutex{},
	}

}

type Seller interface {
	SellerStats
}

type seller struct {
	salesAmount uint64
	name        string
}

func (s *seller) SellerStats() uint64 {
	return s.salesAmount
}

func NewSeller(name string, salesAmount uint64) Seller {
	return &seller{
		salesAmount: salesAmount,
		name:        name,
	}
}

func (s *seller) Name() string {
	return s.name
}
