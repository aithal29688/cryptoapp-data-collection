CREATE TABLE IF NOT EXISTS cryptoprice (
    id serial not null,
    type text,
    market text,
    fromsymbol text,
    tosymbol text,
    flags text,
    price double precision,
    lastupdate bigint,
    lastvolume double precision,
    lastvolumeto double precision,
    lasttradeid text,
    volumeday double precision,
    volumedayto double precision,
    volume24hour double precision,
    volume24hourto double precision,
    openday double precision,
    highday double precision,
    lowday double precision,
    open24hour double precision,
    high24hour double precision,
    low24hour double precision,
    lastmarket text,
    change24hour double precision,
    changepct24hour double precision,
    changeday double precision,
    changepctday double precision,
    supply double precision,
    mktcap double precision,
    totalvolume24h double precision,
    totalvolume24hto bigint,
    CONSTRAINT cryptoprice_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);

CREATE INDEX IF NOT EXISTS cryptoprice_idx ON cryptoprice (market, fromsymbol, tosymbol, lastupdate);
