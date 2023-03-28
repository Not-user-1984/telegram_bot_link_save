package main

import (
	"flag"
	"log"
	"telegram_bot_link/clients/telegram"
)
const (
	tgBotHost = "api.telegram.org"
)
func main() {

	tgClient = telegram.New(tgBotHost, mustToken())

}

func mustToken() string {
    // Remove argument names.
    token := flag.String(
        "token-bot-token",
        "",
        "token for access to telegram bot",
    )
    
    flag.Parse()
    if *token == "" {
        // Change error message to provide more information.
        log.Fatal("Error: token not provided")
    }

    return *token
}