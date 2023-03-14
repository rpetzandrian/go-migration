migration:
	@read -p "migration name: " mname; \
	touch ./db/migrations/$(shell date '+%Y%m%d%H%M%S')_$$mname.up.sql \
	touch ./db/migrations/$(shell date '+%Y%m%d%H%M%S')_$$mname.down.sql

migrate:
	@go run ./db/main.go $(db) $(step)
