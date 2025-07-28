package telegramity

import (
	"sync"

	"github.com/somosbytes/telegramity/internal/configs"
	"github.com/somosbytes/telegramity/internal/telegram/bot"
	"github.com/somosbytes/telegramity/internal/telegramity"
)

var (
	globalClient bot.Client
	globalOnce   sync.Once
	globalErr    error
)

func InitGlobalClient(botToken string, chatID int64, options ...configs.ConfigOption) error {
	globalOnce.Do(func() {
		client, err := telegramity.NewClient(botToken, chatID, options...)
		if err != nil {
			globalErr = err
			return
		}
		globalClient = client
	})
	return globalErr
}

func GetGlobalClient() bot.Client {
	if globalClient == nil {
		panic("telegramity: global client not initialized. Call InitGlobalClient first")
	}
	return globalClient
}

func CloseGlobalClient() error {
	if globalClient != nil {
		return globalClient.Close()
	}
	return nil
}
