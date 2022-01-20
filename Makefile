
run-postgres-migrations: check-cdb-env
	migrate -database ${POSTGRESQL_URL} -path /Users/Belal/GoFinder/linkgraph/store/postgresql/migrations up

drop-postgres-migrations: check-cdb-env
	migrate -database ${POSTGRESQL_URL} -path /Users/Belal/GoFinder/linkgraph/store/postgresql/migrations down

define PostgresSQL_missing_error

POSTGRESQL_URL is undefined. To run the migrations this POSTGRESQL_URL
must point to a POSTGRESQL db instance. For example.

endef

export POSTGRESQL_URL='postgres://postgres:123@localhost:5432/goFinder?sslmode=disable'

export PostgresSQL_missing_error

check-cdb-env:
ifndef POSTGRESQL_URL
	$(error ${PostgresSQL_missing_error})
endif