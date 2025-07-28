package telegramity

import (
	"sync"
)

var (
	globalClient Client
	globalOnce   sync.Once
	globalErr    error
)

// InitGlobalClient initializes the global client singleton
// Call this once at the start of your application
func InitGlobalClient(botToken string, chatID int64, options ...ConfigOption) error {
	globalOnce.Do(func() {
		client, err := NewClient(botToken, chatID, options...)
		if err != nil {
			globalErr = err
			return
		}
		globalClient = client
	})
	return globalErr
}

// GetGlobalClient returns the global client instance
// Panics if InitGlobalClient hasn't been called
func GetGlobalClient() Client {
	if globalClient == nil {
		panic("telegramity: global client not initialized. Call InitGlobalClient first")
	}
	return globalClient
}

// CloseGlobalClient closes the global client
func CloseGlobalClient() error {
	if globalClient != nil {
		return globalClient.Close()
	}
	return nil
}
