// File: main.go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// LambdaHandler is the function that AWS Lambda will invoke.
// this is being invoked every time Telegram sends a request to this Lambda function.
func LambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	
	// 1. Create the bot instance inside the handler.
	//    The options now include your actual handler logic.
	opts := []bot.Option{
		bot.WithDefaultHandler(myBotHandler),
		// The secret token is validated by API Gateway or in the handler, not by the bot library directly in this mode.
	}

	b, err := bot.New(os.Getenv("TELEGRAM_BOT_TOKEN"), opts...)
	if err != nil {
		log.Printf("failed to create bot: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	// 2. Parse the update from the incoming request body.
	var update models.Update
	err = json.Unmarshal([]byte(request.Body), &update)
	if err != nil {
		log.Printf("failed to unmarshal update: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, nil
	}

	// 3. Manually process the update. This is the key difference.
	//    This call is synchronous and does not block forever.
	b.ProcessUpdate(ctx, &update)

	// 4. Return a 200 OK response to API Gateway to acknowledge receipt.
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: "OK"}, nil
}

// This is your actual bot logic, renamed to avoid conflict.
func myBotHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Hi there from AWS Lambda!",
	})
}

// The main function now just starts the Lambda handler.
func main() {
	lambda.Start(LambdaHandler)
}

// explanation of webhooks: 
// instead of API pattern, which is us call the API, 
// webhooks call to us whenever some kind of event worth listening to happens.
// eg. in this case, when the client messages the shaw telegram bot.

// usually, webhooks are HTTP servers which hook (call) to a provided address (which in this case, is our FaaS.)
// these webhooks make the HTTP requests to our FaaS, and usually it includes some kind of auth so we dont just 
// accept HTTP requests from anyone (as that could be a hacker). 	




