BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "trades" (
	"trades_id"	INTEGER NOT NULL,
	"trades_ptf_id"	INTEGER,
	"trades_date"	TEXT,
	"trades_time"	TEXT,
	"trades_order"	TEXT,
	"trades_buy"	REAL,
	"trades_sell"	REAL,
	"trades_quantity"	INTEGER,
	"trades_cost"	REAL,
	"trades_gain"	REAL,
	"trades_gainp"	REAL,
	"trades_cash"	REAL,
	"trades_day"	REAL,
	PRIMARY KEY("trades_id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "quotes" (
	"id"	TEXT,
	"name"	TEXT,
	"date"	TEXT,
	"open"	REAL,
	"high"	REAL,
	"low"	REAL,
	"close"	REAL,
	"close1"	REAL,
	"adjclose"	REAL,
	"volume"	INTEGER
);
CREATE TABLE IF NOT EXISTS "cdays" (
	"cdays_id"	INTEGER NOT NULL,
	"cdays_name"	TEXT,
	"cdays_ptf_id"	INTEGER,
	"cdays_order"	TEXT,
	"cdays_date"	TEXT,
	"cdays_open"	REAL,
	"cdays_close"	REAL,
	"cdays_close1"	REAL,
	"cdays_buy"	REAL,
	"cdays_dvol"	INTEGER,
	"cdays_vevo"	INTEGER,
	"cdays_low"	REAL,
	"cdays_high"	REAL,
	"cdays_time"	TEXT,
	"cdays_percent"	REAL,
	"cdays_rsi"	REAL,
	"cdays_ema"	REAL,
	"cdays_sma"	REAL,
	"cdays_trade"	TEXT,
	"cdays_quantity"	INTEGER,
	"cdays_cost"	REAL,
	"cdays_gain"	REAL,
	"cdays_gainp"	REAL,
	"cdays_volume"	INTEGER,
	PRIMARY KEY("cdays_id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "libelles" (
	"ISIN"	TEXT,
	"nom"	TEXT,
	"ticker"	TEXT
);
CREATE TABLE IF NOT EXISTS "orders" (
	"orders_id"	INTEGER NOT NULL,
	"orders_name"	TEXT,
	"orders_ptf_id"	INTEGER,
	"orders_order"	TEXT,
	"orders_time"	TEXT,
	"orders_quote"	REAL,
	"orders_quantity"	INTEGER,
	"orders_buy"	REAL,
	"orders_sell"	REAL,
	"orders_cost_price"	REAL,
	"orders_cost"	REAL,
	"orders_debit"	REAL,
	"orders_credit"	REAL,
	"orders_gain"	REAL,
	"orders_gainp"	REAL,
	"orders_sell_time"	TEXT,
	"orders_sell_cost"	REAL,
	"orders_sell_gain"	REAL,
	"orders_sell_gainp"	REAL,
	"orders_rem"	TEXT,
	PRIMARY KEY("orders_id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "ptf" (
	"ptf_id"	TEXT NOT NULL,
	"ptf_name"	TEXT,
	"ptf_enabled"	INTEGER,
	"ptf_top"	INTEGER,
	"ptf_gain"	NUMERIC DEFAULT (0),
	"ptf_note"	TEXT,
	"ptf_quote"	REAL,
	"ptf_seuil_achat"	REAL,
	"ptf_seuil_vente"	REAL,
	"ptf_isin"	TEXT,
	"ptf_rem"	TEXT,
	"ptf_trend"	TEXT,
	PRIMARY KEY("ptf_id")
);
CREATE TABLE IF NOT EXISTS "users" (
    user_id integer primary key autoincrement,
    user_username varchar(100),
    user_password varchar(1000),
    user_mail varchar(100)
);
COMMIT;
-- Pour ajouter les nouvelles tables
-- cat resources/schema.sql | sqlite3 /mnt/nas/data/picsou/picsou.sqlite
