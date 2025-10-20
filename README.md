# CARA MENJALANKAN

pertama buat database di postgreseql versi 14. dengan nama tabel test_rsud. apabila credential nya berbeda maka edit file config.yaml dibagian 

POSTGRES_USERNAME: "postgres"
POSTGRES_PASSWORD: "postgres"
POSTGRES_HOST: "127.0.0.1"
POSTGRES_PORT: "5432"
POSTGRES_DATABASE: "test_rsud"

1. go run apps/migration/main.go up
2. go run apps/api/main.go
3. kemudian dapat menjalankan pada browser atau pada postman http://localhost:8081/master/province
