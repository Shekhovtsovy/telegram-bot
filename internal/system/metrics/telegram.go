package metrics

import "github.com/prometheus/client_golang/prometheus"

// TelegramStat is an interface which provides methods for stat working
type TelegramStat interface {
	IncReceivedMessages()
}

type telegramStat struct {
	receivedMessages prometheus.Counter
}

// Increment counter of received messages
func (t *telegramStat) IncReceivedMessages() {
	t.receivedMessages.Inc()
}

// NewTelegramStat return a new Telegram Stat
func NewTelegramStat() TelegramStat {
	stat := &telegramStat{
		receivedMessages: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "received_messages_count",
			Help: "Amount of received messages from telegram",
		}),
	}
	prometheus.MustRegister(
		stat.receivedMessages,
	)
	return stat
}
