CREATE TABLE IF NOT EXISTS ticks.ticks (
	Id BIGINT PRIMARY KEY auto_increment NOT NULL,
	`timestamp` BIGINT NULL COMMENT 'Unixtime in milliseconds',
	symbol VARCHAR(100) NULL COMMENT 'name of insrtruments',
	best_bid DOUBLE NULL COMMENT 'the best sell offer',
	best_ask DOUBLE NULL COMMENT 'the best bye offer'
)
