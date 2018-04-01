# Webcambot :camera:

![scheme](https://raw.githubusercontent.com/silentbay/webcambot/master/scheme.png)

## En

Create a bot using [@BotFather](https://t.me/botfather).

Find out your ID using [@userinfobot](https://t.me/userinfobot)

Compile or [download](https://github.com/silentbay/webcambot/releases) webcambot.

Create configuration file **botconfig.json**:

    {
    "bottoken": "TOKEN",
    "botuser": "UserID",
    "gpiopin": "17"
    }

Prepare your Raspberry Pi:
1. Install OS.
2. Install fswebcam.
3. Connect the webcam and the PIR-sensor.
4. Copy bot and configuration.
5. Start bot and enjoy.

You can use supervisor to control bot. Example of config:

    ; Supervisor configuration file. /etc/supervisor/conf.d
    [program:webcambot]
    directory=/home/pi/webcambot
    command=/home/pi/webcambot/webcambot-arm6
    autostart=true
    autorestart=true
    stderr_logfile=/home/pi/webcambot/bot.err.log
    stdout_logfile=/home/pi/webcambot/bot.out.log

## Ru

Создай бота используя [@BotFather](https://t.me/botfather).

Определи свой UserID используя [@userinfobot](https://t.me/userinfobot)

Скомпилируй или [скачай](https://github.com/silentbay/webcambot/releases) бота.

Создай файл конфигурации **botconfig.json**:

    {
    "bottoken": "TOKEN",
    "botuser": "UserID",
    "gpiopin": "17"
    }

Подготовь свой Raspberry Pi:
1. Установи операционную систему.
2. Установи fswebcam.
3. Подключи вебкамеру и датчик движения.
4. Скопируй бота и файл конфигурации.
5. Запускай бота и используй.

Можно использовать supervisor для контроля бота. Пример конфигурации:

    ; Supervisor configuration file. /etc/supervisor/conf.d
    [program:webcambot]
    directory=/home/pi/webcambot
    command=/home/pi/webcambot/webcambot-arm6
    autostart=true
    autorestart=true
    stderr_logfile=/home/pi/webcambot/bot.err.log
    stdout_logfile=/home/pi/webcambot/bot.out.log