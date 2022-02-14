package wander

import (
	"io"
)

type Discord struct {
	name string
}

func NewDiscord() Discord {
	discord := Discord{}
	return discord
}

func Name() string {
	return "discord"
}

func UserName() string {
	return ""
}

func UserId() string {
	return ""
}

func SendMessage(channel, message string) error {
	return nil
}

func SendAction(channel, message string) error {
	return nil
}

func DeleteMessage(channel, messageId string) error {
	return nil
}

func SendFile(channel, name string, r io.Reader) error {
	return nil
}

func BanUser(channel, userId string, duration int) error {
	return nil
}

func UnbanUser(channel, userID string) error {
	return nil
}

func Join(join string) error {
	return nil
}

func Typing(channel string) error {
	return nil
}

func PrivateMessage(userID, messageID string) error {
	return nil
}

func SupportsPrivateMessages() bool {
	return false
}

func SupportsMultiline() bool {
	return false
}

func CommandPrefix() string {
	return ""
}

func ChannelCount() int {
	return 0
}

func Log(Medium, LogLevel, string) {
}
