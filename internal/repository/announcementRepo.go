package repository

import (
	"context"
	"github.com/Sergi-Ch/Marketplace/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AnouncementRepository interface {
	Create(ctx context.Context, announcement *domain.Announcement) *domain.Announcement
	GetAllItems(ctx context.Context) (*domain.Announcement, error)
	Delete(ctx context.Context, id int) error
}

type RepoPg struct {
	db *pgxpool.Pool
}

func NewRepoPg(db *pgxpool.Pool) *RepoPg {
	return &RepoPg{db: db}
}

func (r *RepoPg) Create(ctx context.Context, announcement *domain.Announcement) {

}
