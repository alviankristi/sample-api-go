package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alviankristi/catalyst-backend-task/pkg/database"
	"github.com/alviankristi/catalyst-backend-task/pkg/server"
	"github.com/spf13/cobra"
)

// run the api server
// command: go run main.go serve
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start http server with configured api",
	Long:  `Starts a http server and serves the configured api`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}

func run() {

	// open db connection
	db := database.Open(Conf)

	// init api server
	server, err := server.NewServer(Conf, db)
	if err != nil {
		log.Panicf("server shutdown: %v", err)
	}

	go server.ListenAndServe()

	log.Printf("Listening on %s...", server.Addr)

	//capture cancel api server
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	sig := <-quit
	log.Printf("Shutting down server... Reason: %v", sig.String())

	ctx, cancel := context.WithCancel(context.Background())

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("server shutdown: %v", err)
	}

	// close all connections
	defer db.Close()
	defer cancel()

	log.Print("API is stopped and close db connection ...")
}
