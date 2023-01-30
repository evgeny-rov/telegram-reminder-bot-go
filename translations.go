package main

type translation struct {
	start               string
	help                string
	badParams           string
	outOfRange          string
	alertWithMessage    string
	alertWithoutMessage string
	created             string
	cancelled           string
	noReminders         string
}

type translations struct {
	en translation
	ru translation
}

var Translations = translations{
	en: translation{
		start: `Hey!
I can create reminder messages for you that will ping you after a certain amount of time, so that you don't forget anything!

/remind <hours-minutes> <message> - create a reminder. 
/cancel - cancel your current reminder.
/help - get detailed information on how to use the commands.`,
		help: `How to use /remind command

<hours-minutes> - Represents a certain amount of time after which I will send you a reminder message.
If you specify only one number, it will be considered as minutes.
Maximum is 24 hours, minimum is 1 minute.

<message> - Non-required message that will be sent with your reminder.

Example:
/remind 1-10 The cake is ready
`,
		badParams:           "You sent something I don't understand. Try again - specify the time and message if necessary, for example: /remind 1-10 The cake is ready",
		outOfRange:          "I can create reminders ranging from 1 minute to 24 hours.",
		alertWithMessage:    "Reminder:",
		alertWithoutMessage: "Ding dong! This is a reminder you asked me to create.",
		created:             "Okay. I'll send you a reminder at the right time.",
		cancelled:           "Ok. I forgot your current reminder.",
		noReminders:         "You don't have active reminders yet.",
	},
	ru: translation{
		start: `Привет!
Я могу создавать для тебя сообщения-напоминания, которые будут приходить через определенное время, чтобы ты ничего не забыл!
		
/remind <часы-минуты> <сообщение> - создать напоминание.
/cancel - отменить текущее напоминание.
/help - получить подробную информацию о том, как использовать команды.`,
		help: `Как использовать команду /remind

<часы-минуты> - Определенное количество времени, по истечении которого я отправлю тебе сообщение с напоминанием.
Если указать только одно число, оно будет считаться минутами.
Максимальное значение 24 часа, минимальное 1 минута.

<сообщение> - Необязательное сообщение которое будет отправлено с напоминанием.

Пример:
/remind 1-10 Пирог готов
`,
		badParams:           "Ты отправил что-то непонятное. Попробуй еще раз - время и сообщение если нужно, например: /remind 1-10 Пирог готов",
		outOfRange:          "Я могу создавать напоминания в пределах от 1 минуты до 24 часов.",
		alertWithMessage:    "Напоминание:",
		alertWithoutMessage: "Динг донг! Это напоминание, которое ты просил создать.",
		created:             "Окей. я напишу тебе в нужное время.",
		cancelled:           "Окей. я забыл твое текущее напоминание.",
		noReminders:         "У тебя пока нет активных напоминаний.",
	},
}
