mainmigrateup:
	migrate -path migration/postgres -database postgresql://root:password@localhost:5432/site_pinger?sslmode=disable -verbose up
	
mainmigratedown:
	migrate -path migration/postgres -database postgresql://root:password@localhost:5432/site_pinger?sslmode=disable -verbose down

statsmigrateup:
	migrate -path migration/clickhouse -database clickhouse://localhost:9000?username=default -verbose up

statsmigratedown:
	migrate -path migration/clickhouse -database clickhouse://localhost:9000?username=default -verbose down

.PHONY: mainmigrateup mainmigratedown statsmigrateup statsmigratedown