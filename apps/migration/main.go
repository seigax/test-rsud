package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/viper"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/config"
)

func init() {
	godotenv.Load()
	config.ViperConfig()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide action!")
		return
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("POSTGRES_HOST"), viper.GetString("POSTGRES_PORT"),
		viper.GetString("POSTGRES_USERNAME"), viper.GetString("POSTGRES_PASSWORD"),
		viper.GetString("POSTGRES_DATABASE"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	migrations := &migrate.FileMigrationSource{
		Dir: "migration",
	}

	switch os.Args[1] {
	case "up":
		m, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Applied %d migrations!\n", m)
	case "down":
		m, err := migrate.ExecMax(db, "postgres", migrations, migrate.Down, 1)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Applied %d migrations!\n", m)
	case "status":
		m, err := migrate.GetMigrationRecords(db, "postgres")
		if err != nil {
			panic(err)
		}
		for _, record := range m {
			fmt.Printf("%s - %s\n", record.Id, record.AppliedAt)
		}
	case "now":
		loc, _ := time.LoadLocation("Asia/Jakarta")
		time.Local = loc
		fmt.Println(time.Now().Format("20060102150405"))
	case "create":
		fileName := os.Args[2]
		loc, _ := time.LoadLocation("Asia/Jakarta")
		time.Local = loc
		nowStr := time.Now().Format("20060102150405")
		os.Create("./migration/" + nowStr + "-" + fileName + ".sql")
	default:
		fmt.Println("command not found")
	}

}
