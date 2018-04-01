// https://github.com/silentbay/webcambot
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/brian-armstrong/gpio"

	"github.com/Syfaro/telegram-bot-api"
)

var (
	BotToken string
	BotUser  string
	GpioPin  string
	Watchdog bool = false
)

type BotConfig struct {
	BotToken string
	BotUser  string
	GpioPin  string
}

func init() {
	file, _ := os.Open("botconfig.json")
	decoder := json.NewDecoder(file)
	config := BotConfig{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	BotToken = config.BotToken
	BotUser = config.BotUser
	GpioPin = config.GpioPin
	fmt.Println("Initialization complete", BotToken, BotUser, GpioPin)

}

func watcher(bot *tgbotapi.BotAPI) {
	BotUserID, err := strconv.ParseInt(BotUser, 10, 64)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	GpioPinID, err := strconv.ParseInt(GpioPin, 10, 8)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	watcher := gpio.NewWatcher()
	watcher.AddPin(uint(GpioPinID))
	defer watcher.Close()
	for {
		pin, value := watcher.Watch()
		if value == 1 {
			if Watchdog == true {
				t1 := time.Now().Format("15-04-05")
				cmd := exec.Command("fswebcam", "-S", "5", t1)
				cmd.Run()
				fmt.Println("Motion detected, send messege", pin)
				bot.Send(tgbotapi.NewMessage(BotUserID, "Motion detected"))
				bot.Send(tgbotapi.NewPhotoUpload(BotUserID, t1))
			}

		}

	}

}

func main() {
	BotUserID, err := strconv.ParseInt(BotUser, 10, 64)
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bot.Debug = false

	fmt.Println("Authorized on account %s", bot.Self.UserName)
	bot.Send(tgbotapi.NewMessage(BotUserID, "Start bot"))
	go watcher(bot)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Text == "/watchdog" {
			if strconv.Itoa(update.Message.From.ID) == BotUser {
				if Watchdog == false {
					Watchdog = true
					bot.Send(tgbotapi.NewMessage(BotUserID, "Watchdog ON"))
				} else {
					Watchdog = false
					bot.Send(tgbotapi.NewMessage(BotUserID, "Watchdog OFF"))
				}
			} else {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Permission denied"))
			}
		}
		if update.Message.Text == "/shot" {
			if strconv.Itoa(update.Message.From.ID) == BotUser {
				t2 := time.Now().Format("15-04-05")
				cmd := exec.Command("fswebcam", "-S", "5", t2)
				cmd.Run()
				bot.Send(tgbotapi.NewPhotoUpload(BotUserID, t2))
			} else {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Permission denied"))
			}
		}
	}

}
