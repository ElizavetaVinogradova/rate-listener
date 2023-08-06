package mysql

type Tick struct {
	timestamp int64   //(BIGINT) - Unixtime в миллисекундах,
	symbol    string  //(VARCHAR) - название инструментов,
	best_bid  float64 //(DOUBLE) - цена предложения продажи,
	best_ask  float64 //(DOUBLE) - цена предложения покупки.
}
