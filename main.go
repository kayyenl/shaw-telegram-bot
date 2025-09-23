package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
)

// Send any text message to the bot after the bot has been started

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// configure the new bot by placing With functions in the opts object,
	// part of the functional options pattern, making it easier to init structs when too many params.
	opts := []bot.Option{
		bot.WithDefaultHandler(defaultHandler),
		bot.WithWebhookSecretToken(os.Getenv("TELEGRAM_BOT_TOKEN")),
	}
    
    // get telegram token from terminal and handle if token cannot be gotten
    token := os.Getenv("TELEGRAM_BOT_TOKEN")
	b, err := bot.New(token, opts...)
	if err != nil {
		panic(err)
	}

	// takes the updates received by the webserver (the http.ListenAndServe), and dispatches them
	// to the respective handler functions.
	// StartWebhook has to run in parallel (in a goroutine), because if not the listenAndServe cannot execute, and will be forever blocked.
	go b.StartWebhook(ctx)

	// blocks forever, listening for incoming HTTP requests coming from Telegram.
	// known as a receiver.
	http.ListenAndServe(":2000", )
}

// explanation of telegram bot


// FOR THE NORMAL CASE

// Telegram API is hosted online by Telegram, separate from this code (which runs locally or any method defined by me, different from telegram)
// When i say "xxx" as a Telegram user to my bot, Telegram backend sees the message and sends it to my machine which should be actively hosting this code. By hosting this code, it is sending a long polling HTTP GET to the Telegram API (hosted by Telegram).
// when machine running this code receives HTTP response from long polling Telegram API,
// the internal code of this machine deserializes it into models.Update struct. 
// Then our b.SendMessage functionality takes these info from update, and makes use of it (find chatID who sent it here, in order to send it back, etc.)


// FOR THE WEBHOOK CASE

// Instead of calling Telegram API via long polling to find out if messages are being sent from the telegram client,
// This backend becomes the server instead. 
// We register this server's public URL with Telegram instead, and whenever a user sends a message to this bot,
// Telegram wil send a HTTP POST directly to the registered URL. This server receives it and our handlers handle it.