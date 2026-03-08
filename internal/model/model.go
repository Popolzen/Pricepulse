package model

import "time"

type Item struct {
	ID           int64
	URL          string
	ShopType     string
	Name         string
	CurrentPrice int
	TargetPrice  int
	CreatedAt    time.Time
}

type PriceHistory struct {
	ID        int64
	ItemID    int64
	Price     int
	CheckedAt time.Time
}
