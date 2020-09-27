
package notification

type GateName string

const (
	GateNameMainTelegram GateName = "main_telegram"
	GateNameTechTelegram GateName = "tech_telegram"
)

const (
	MainTelegramConfigPathToken  = "telegram.main.token"
	MainTelegramConfigPathChatID = "telegram.main.chat"
	TechTelegramConfigPathToken  = "telegram.tech.token"
	TechTelegramConfigPathChatID = "telegram.tech.chat"
)
