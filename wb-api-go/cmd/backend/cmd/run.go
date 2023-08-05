/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq" // <------------ here
	"github.com/patrickmn/go-cache"
	"github.com/spf13/cobra"
	"github.com/ssargent/world-builder/wb-api-go/cmd/backend/internal"
	"github.com/ssargent/world-builder/wb-api-go/cmd/backend/internal/config"
)

var runEnvFile string

// runCmd represents the run command.
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		api, err := server()
		if err != nil {
			log.Fatalf("server: %q", err.Error())
		}

		if err := api.ListenAndServe(); err != nil {
			log.Fatalf("ListenAndServe: %q", err.Error())
		}
	},
}

func server() (*internal.API, error) {
	if err := godotenv.Load(runEnvFile); err != nil {
		return nil, fmt.Errorf("godotenv.Load: %w", err)
	}

	var cfg config.Config
	if err := envconfig.Process("worldbuilder", &cfg); err != nil {
		return nil, fmt.Errorf("envconfig.Process: %w", err)
	}

	explainConfig(&cfg)

	db, safeDB, err := database(&cfg)
	if err != nil {
		return nil, fmt.Errorf("database: %w", err)
	}

	fmt.Printf("Connecting to %q\n", safeDB)

	cache := cache.New(cfg.Cache.DefaultExpiration, cfg.Cache.DefaultCleanup)
	return internal.NewAPI(&cfg, db, db, cache), nil
}

//nolint:gochecknoinits // cobra commander has ordered the use of inits.
func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	runCmd.Flags().StringVarP(&runEnvFile, "environment-file", "e", ".env", "contains environment settings")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func database(cfg *config.Config) (*sqlx.DB, string, error) {
	dbURISafe := fmt.Sprintf("postgres://%s:xxxxxxxxxxx@%s/%s?sslmode=disable",
		cfg.Database.Username,
		cfg.Database.Server,
		cfg.Database.Name)
	dbURI := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		cfg.Database.Username,
		strings.TrimSpace(cfg.Database.Password),
		cfg.Database.Server,
		cfg.Database.Name)

	db, err := sqlx.Connect(cfg.Database.Driver, dbURI)

	return db, dbURISafe, err
}

func explainConfig(cfg *config.Config) {
	fmt.Println("-----------------")
	fmt.Println("World Builder Configured Parameters")
	fmt.Println("-----------------")

	fmt.Printf("Config.Port := %q\n", cfg.Port)
	fmt.Printf("Config.Database.Driver := %q\n", cfg.Database.Driver)
	fmt.Printf("Config.Database.Name := %q\n", cfg.Database.Name)
	fmt.Printf("Config.Database.Username := %q\n", cfg.Database.Username)
	fmt.Printf("Config.Database.Server := %q\n", cfg.Database.Server)
	fmt.Printf("Config.Cache.DefaultExpiration := %q\n", cfg.Cache.DefaultExpiration)
	fmt.Printf("Config.Cache.DefaultCleanup := %q\n", cfg.Cache.DefaultCleanup)
	fmt.Println("-----------------")
}
