package wander

import (
	"io"
)

type LogLevel string

const (
	LogLevelDebug LogLevel = "[D]"
	LogLevelInfo LogLevel = "[I]"
	LogLevelWarn LogLevel = "[W]"
	LogLevelError LogLevel = "[E]"
	LogLevelFatal LogLevel = "[F]"
)

// CRUD, but with messages.
type MessageState string

const (
	MessageTypeCreate MessageState = "create"
	MessageTypeUpdate MessageState = "update"
	MessageTypeDelete MessageState = "delete"
)

type MediumName string

const (
	MediumNameDiscord MediumName = "discord"
	MediumNameIrc MediumName = "irc"
	MediumNameMatrix MediumName = "matrix"
)

type Message interface {
	Channel() string
	UserName() string
	UserId() string
	UserAvatar() string
	Message() string
	RawMessage() string
	MessageId() string
	// Test if message is a certain command.
	MatchesCommand(prefix, command string)
	// Test if there is a prefix. Return first word, followed by rest, followed by match.
	ParseCommand(prefix string) (command string, rest []string, match bool)
}

// Similar to Bruxism. Example applications: Discord, IRC, Matrix
type Medium interface {
	Name() string
	UserName() string
	UserId() string
	SendMessage(channel, message string) error
	SendAction(channel, message string) error
	DeleteMessage(channel, messageId string) error
	SendFile(channel, name string, r io.Reader) error
	BanUser(channel, userId string, duration int) error
	UnbanUser(channel, userID string) error
	// Join can be a channel, a server, etc.
	Join(join string) error
	// Are we currently typing in a channel?
	Typing(channel string) error
	// Send a private message.
	PrivateMessage(userID, messageID string) error
	// Do we support private messages?
	SupportsPrivateMessages() bool
	SupportsMultiline() bool
	CommandPrefix() string
	ChannelCount() int
	Log(Medium, LogLevel, string)
}

// Basic function signatures to implement.
type LoadFunc func(*Bot, Medium, []byte) error
type SaveFunc func() ([]byte, error)
type HelpFunc func(*Bot, Medium, Message, bool) []string
type MessageFunc func(*Bot, Medium, Message)
type StatsFunc func(*Bot, Medium, Message) []string

type Module interface {
	Name() string
	// Saving and loading maintains a simple state.
	Load(*Bot, Medium, []byte) error
	Save() ([]byte, error)
	// Message help.
	Help(*Bot, Medium, Message, bool) []string
	// Process message.
	Message(*Bot, Medium, Message)
	Stats(*Bot, Medium, Message) []string
}
