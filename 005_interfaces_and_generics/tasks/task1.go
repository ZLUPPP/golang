package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

type Notifier interface {
	Notify(msg string) error
}

type Email struct {
	Address string
}

type SMS struct {
	Phone string
}

type Telegram struct {
	Chat  string
	Count int
}

func (email *Email) Notify(msg string) error {
	fmt.Printf(email.Address + ": " + msg + "\n")
	return nil
}

func (sms *SMS) Notify(msg string) error {
	if utf8.RuneCountInString(msg) > 20 {
		return errors.New("Message too long")
	}
	fmt.Printf(sms.Phone + ": " + msg + "\n")
	return nil
}

func (telegram *Telegram) Notify(msg string) error {
	fmt.Printf(telegram.Chat + ": " + msg + "\n")
	telegram.Count++
	return nil
}

func Broadcast(ns []Notifier, msg string) error {
	for _, n := range ns {
		if err := n.Notify(msg); err != nil {
			return err
		}
	}
	return nil
}

func write(w io.Writer, name string, score int) {
	fmt.Fprintf(w, "%s: %d points\n", name, score)
}

func main() {
	email := &Email{Address: "test@example.com"}
	sms := &SMS{Phone: "+77001234567"}
	telegram := &Telegram{Chat: "my_chat"}

	notifiers := []Notifier{email, sms, telegram}

	if err := Broadcast(notifiers, "Hello, World!"); err != nil {
		fmt.Println(err)
	}

	var _ Notifier = &Email{}
	var _ Notifier = &SMS{}
	var _ Notifier = &Telegram{}

	telegram.Notify("Second message")

	fmt.Println("Telegram messages sent:", telegram.Count)

	err := sms.Notify("This message is definitely longer than twenty characters")
	if err != nil {
		fmt.Println("Error:", err)
	}

	write(os.Stdout, "Ilya", 1)

	var sb strings.Builder
	write(&sb, "Didar", 5)
	fmt.Println(sb.String())
}
