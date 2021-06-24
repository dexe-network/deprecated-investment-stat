# Dexe Investing Stat

## Build & Run
### Prerequisites
- go 1.15
- docker

if necessary create block_number file and add there start block number.

Create .env file in root directory and add following values:
```dotenv
DB_DSN=postgres://localhost/dexeinvest2?sslmode=disable
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_DBNAME=dexeinvest2
DB_DIALECT=postgres
DB_MAX_OPEN-CONNS=80
ZAP_LEVEL=2
ETH_NODE=ws://localhost:8545
DEX_PROTOCOL=uniswapV2
DEX_FACTORY_ADDRESS=0x1187D2f98C556a3Cfbb270Be161EbB34EcD2925F
MAX_PARALLEL_BLOCKS=100
DB_DEBUG=true
NETWORK=bsc
```

Also, config options list available from console
```shell
./app --help
Usage of ./app:
      --app-env string              Application environment (default "prod")
      --db-debug                    Debug database
      --db-dialect string           GORM database dialect (default "postgres")
      --db-dsn string               Database data source name
      --db-max-open-conns int       GORM Max database connections (default 80)
      --dex-protocol string         Dex Protocol (default "uniswapV2")
      --dex-router-abi string       Dex Router ABI
      --dex-router-address string   Dex Router address
      --eth-node string             Eth node URL (default "ws://127.0.0.1:8546")
      --max-parallel-blocks int     Max processing blocks in parallel (default 200)
      --migrations-path string      Path to migrations directory (default "migrations.d")
      --zap-level int8              debug -1, info 0, warn 1, error 2, dpanic 3, panic 4, fatal 5 (default -1)
pflag: help requested
```

## Build binary file

Use `make build` to build executable file for linux for other os, you will need to edit build command in a Make file

## Build & Run Locally in a docker-compose 

Use `sudo make run` or `sudo docker-compose up --remove-orphans app`  to build&run project in a docker-compose with postgres and pgadmin (Check docker-compose).

postgres con: postgres://root:root@postgres/dexparser?sslmode=disable

pgadmin: http://localhost:5050/

pgadmin: admin@admin.com

pgadmin: root
