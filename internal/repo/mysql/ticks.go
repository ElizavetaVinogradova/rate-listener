package mysql

import (
	"fmt"
	"rates-listener/internal/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type TicksRepository struct {
	db *sqlx.DB
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewTickRepository(config Config) (*TicksRepository, error) {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Username, config.Password, config.Host, config.Port, config.DBName))
	if err != nil {
		return nil, fmt.Errorf("open sql connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return &TicksRepository{db: db}, nil
}

func (r *TicksRepository) Close(){
	r.db.Close()
}

type TickDataBaseDTO struct {
	timestamp int64
	symbol    string
	bestBid   float64
	bestAsk   float64
}

func (r *TicksRepository) CreateBatch(ticks []service.Tick) error {
	ticksDB := mapTickSliceToTicksDTOSlice(ticks)

	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO ticks (timestamp, symbol, best_bid, best_ask) VALUES (?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("prepare statement: %w", err)
	}
	defer stmt.Close()

	for _, tickDB := range ticksDB {
		_, err := stmt.Exec(tickDB.timestamp, tickDB.symbol, tickDB.bestBid, tickDB.bestAsk)
		if err != nil {
			return fmt.Errorf("execute statement: %w", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func mapTickSliceToTicksDTOSlice(ticks []service.Tick) []TickDataBaseDTO {
	ticksDB := make([]TickDataBaseDTO, 0, len(ticks))
	for _, tick := range ticks {
		ticksDB = append(ticksDB, mapTickToTicksDTO(tick))
	}
	return ticksDB
}

func mapTickToTicksDTO(tick service.Tick) TickDataBaseDTO {
	var tickDB TickDataBaseDTO
	tickDB.timestamp = tick.Timestamp
	tickDB.symbol = tick.Symbol
	tickDB.bestBid = tick.BestBid
	tickDB.bestAsk = tick.BestAsk
	return tickDB
}
