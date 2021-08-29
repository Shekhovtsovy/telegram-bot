package metrics

import "github.com/prometheus/client_golang/prometheus"

// TelegramStat is an interface which provides methods for stat working
type TelegramStat interface {
	IncReceivedMessages()
	IncNewUsers()
	IncTelegramListeningErrors()
	IncProcessingMessageErrors()
	IncProcessingCallbackErrors()
	IncSavingUserErrors()
	IncSavingMessageErrors()
}

type telegramStat struct {
	receivedMessages         prometheus.Counter
	newUsers                 prometheus.Counter
	totalErrors              prometheus.Counter
	telegramListeningErrors  prometheus.Counter
	processingMessageErrors  prometheus.Counter
	processingCallbackErrors prometheus.Counter
	savingUserErrors         prometheus.Counter
	savingMessageErrors      prometheus.Counter
}

// IncReceivedMessages increments counter of received messages
func (t *telegramStat) IncReceivedMessages() {
	t.receivedMessages.Inc()
}

// IncReceivedMessages increments counter of new users
func (t *telegramStat) IncNewUsers() {
	t.newUsers.Inc()
}

// IncTelegramListeningErrors increments counter of telegram listening and total errors
func (t *telegramStat) IncTelegramListeningErrors() {
	t.telegramListeningErrors.Inc()
	t.totalErrors.Inc()
}

// IncProcessingMessageErrors increments counter of processing message and total errors
func (t *telegramStat) IncProcessingMessageErrors() {
	t.processingMessageErrors.Inc()
	t.totalErrors.Inc()
}

// IncProcessingCallbackErrors increments counter of processing callback and total errors
func (t *telegramStat) IncProcessingCallbackErrors() {
	t.processingCallbackErrors.Inc()
	t.totalErrors.Inc()
}

// IncSavingUserErrors increments counter of saving user and total errors
func (t *telegramStat) IncSavingUserErrors() {
	t.savingUserErrors.Inc()
	t.totalErrors.Inc()
}

// IncSavingMessageErrors increments counter of saving message and total errors
func (t *telegramStat) IncSavingMessageErrors() {
	t.savingMessageErrors.Inc()
	t.totalErrors.Inc()
}

// NewTelegramStat returns new Telegram Stat for prometheus
func NewTelegramStat() TelegramStat {
	stat := &telegramStat{
		receivedMessages: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "received_messages_count",
			Help: "Amount of received messages from telegram",
		}),
		newUsers: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "new_users_count",
			Help: "Amount of new user",
		}),
		totalErrors: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "total_errors_count",
			Help: "Amount of all errors",
		}),
		telegramListeningErrors: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "telegram_listening_errors_count",
			Help: "Amount of telegram listening errors",
		}),
		processingMessageErrors: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "processing_message_errors_count",
			Help: "Amount of processing message errors",
		}),
		processingCallbackErrors: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "processing_callback_errors_count",
			Help: "Amount of processing callback errors",
		}),
		savingUserErrors: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "saving_user_errors_count",
			Help: "Amount of saving user errors",
		}),
		savingMessageErrors: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "saving_message_errors_count",
			Help: "Amount of saving message errors",
		}),
	}
	prometheus.MustRegister(
		stat.receivedMessages,
		stat.newUsers,
		stat.totalErrors,
		stat.telegramListeningErrors,
		stat.processingMessageErrors,
		stat.processingCallbackErrors,
		stat.savingUserErrors,
		stat.savingMessageErrors,
	)
	return stat
}
