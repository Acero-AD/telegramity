#!/bin/bash

# Setup script for Telegramity SDK testing
echo "🤖 Setting up environment for Telegramity SDK testing..."

# Check if .env file exists
if [ ! -f .env ]; then
    echo "📝 Creating .env file from template..."
    cp .env.example .env
    echo "✅ .env file created!"
    echo ""
    echo "📋 Please edit .env file with your actual values:"
    echo "   1. Replace 'your_bot_token_here' with your bot token from @BotFather"
    echo "   2. Replace '123456789' with your actual chat ID"
    echo ""
    echo "🔗 To get your chat ID:"
    echo "   1. Start a chat with your bot"
    echo "   2. Send a message to your bot"
    echo "   3. Visit: https://api.telegram.org/bot<YOUR_BOT_TOKEN>/getUpdates"
    echo "   4. Look for 'chat':{'id':123456789} in the response"
    echo ""
else
    echo "✅ .env file already exists!"
fi

echo ""
echo "🚀 To test the bot client, run:"
echo "   go run cmd/example/main.go"
echo ""
echo "🧪 To run tests, run:"
echo "   go test ./tests/unit/ -v" 