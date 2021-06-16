CREATE TABLE eth_trades (
    "tokenA" character varying(255),
    "tokenB" character varying(255),
    date timestamp with time zone,
    "blockNumber" bigint,
    tx character varying(255),
    protocol character varying(255),
    "priceIn" numeric,
    "priceOut" numeric,
    "amountOut" character varying(255),
    "amountIn" character varying(255),
    wallet character varying(255),
    value character varying(255)
);

-- Indices -------------------------------------------------------

CREATE INDEX "tokenA_tokenB_protocol" ON eth_trades("tokenA" text_ops,"tokenB" text_ops,protocol text_ops);
CREATE INDEX eth_trades_blocknumber_index ON eth_trades("blockNumber" int8_ops);
