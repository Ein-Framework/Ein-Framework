package cli

type ContextKey string

const (
	configKey ContextKey = "config"
	loggerKey ContextKey = "logger"
)
