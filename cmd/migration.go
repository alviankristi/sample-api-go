package cmd

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"

	db "github.com/alviankristi/catalyst-backend-task/pkg/database"
)

var (
	migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Migrate database",
		Long:  "Migrate database",
		Run: func(cmd *cobra.Command, args []string) {
			err := migrateDb()
			if err != nil {
				log.Printf("error migrate : %v", err)
			}
		},
	}
)

func openConnection() (*migrate.Migrate, error) {
	db := db.Open(Conf)
	defer db.Close()

	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://sql_script",
		"mysql", driver)

	if err != nil {
		return nil, err
	}
	return m, nil
}

func migrateDb() error {
	m, err := openConnection()
	if err != nil {
		log.Fatalf("failed open connection: %v", err)
		return err
	}
	err = m.Up()
	if err != nil {
		log.Fatalf("failed migrate: %v", err)
		return err
	}
	err = m.Down()
	if err != nil {
		log.Fatalf("failed migrate: %v", err)
		return err
	}
	fmt.Println("migrate sql to database succeed")
	return nil
}

func init() {
	RootCmd.AddCommand(migrateCmd)
}
