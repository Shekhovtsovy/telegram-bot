package metrics

import "github.com/prometheus/client_golang/prometheus"

type telegramStat struct {
	receivedMessages prometheus.Counter
}

// TelegramStat is an interface which provides methods for stat working
type TelegramStat interface {
	IncReceivedMessages()
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
