-- accounts가 FK로 엮여있기 때문에 entries, transfers 다 지우고, accounts를 지워야 한다.
DROP TABLE IF EXISTS entries;
DROP TABLE IF EXISTS transfers;
DROP TABLE IF EXISTS accounts;