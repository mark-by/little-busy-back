create table Records (
    ClientID Nullable(UInt64),
    EventID Nullable(UInt64),
    Type String,
    Value Float32,
    DateTime DateTime,
    Description Nullable(String)

) ENGINE = MergeTree
ORDER BY (DateTime);