package mysql

import (
	"fmt"
	"rates-listener/internal/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type TicksRepository struct {
	db *sqlx.DB //указатель на экземпляр ДБ нужен, чтобы при передаче структуры в другие функции они могли изменять состояние именно этого экземпляра, а не его копии.
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewTickRepository(config Config) (*TicksRepository, error) { //звездочка у возвращаемого значения говорит о том, что из функции вернется указатель на адрес в памяти
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Username, config.Password, config.Host, config.Port, config.DBName))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &TicksRepository{db: db}, nil //амперсанд говорит о том, что тут возвращается указатель на созданный экземпляр. Указатель будет использоваться для внесения изменений в объект.
}

type TickDataBaseDTO struct {
	timestamp int64
	symbol    string
	best_bid  float64
	best_ask  float64
}

func (r *TicksRepository) CreateBatch(ticks []service.Tick) error {
	ticksDB := mapTickSliceToTicksDTOSlice(ticks)

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO ticks (timestamp, symbol, best_bid, best_ask) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, tickDB := range ticksDB {
		_, err := stmt.Exec(tickDB.timestamp, tickDB.symbol, tickDB.best_bid, tickDB.best_ask)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func mapTickSliceToTicksDTOSlice(ticks []service.Tick) []TickDataBaseDTO {
	var ticksDB []TickDataBaseDTO
	for _, tick := range ticks {
		ticksDB = append(ticksDB, mapTickToTicksDTO(tick))
	}
	return ticksDB
}

func mapTickToTicksDTO(tick service.Tick) TickDataBaseDTO {
	var tickDB TickDataBaseDTO
	tickDB.timestamp = tick.Timestamp
	tickDB.symbol = tick.Symbol
	tickDB.best_bid = tick.Best_bid
	tickDB.best_ask = tick.Best_ask
	return tickDB
}
