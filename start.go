package main

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"

	"github.com/Vivekkumar-IN/EditguardianBot/config"
)

func start(b *gotgbot.Bot, ctx *ext.Context) error {
	chat := ctx.EffectiveChat.Type

	if chat == "private" {
		file := gotgbot.InputFileByURL(config.StartImage)

		caption := fmt.Sprintf(
			`Hello %s 👋, I'm your %s, here to maintain a secure environment for our discussions.

🚫 𝗘𝗱𝗶𝘁𝗲𝗱 𝗠𝗲𝘀𝘀𝗮𝗴𝗲 𝗗𝗲𝗹𝗲𝘁𝗶𝗼𝗻: 𝗜'𝗹𝗹 𝗿𝗲𝗺𝗼𝘃𝗲 𝗲𝗱𝗶𝘁𝗲𝗱 𝗺𝗲𝘀𝘀𝗮𝗴𝗲𝘀 𝘁𝗼 𝗺𝗮𝗶𝗻𝘁𝗮𝗶𝗻 𝘁𝗿𝗮𝗻𝘀𝗽𝗮𝗿𝗲𝗻𝗰𝘆.

📣 𝗡𝗼𝘁𝗶𝗳𝗶𝗰𝗮𝘁𝗶𝗼𝗻𝘀: 𝗬𝗼𝘂'𝗹𝗹 𝗯𝗲 𝗶𝗻𝗳𝗼𝗿𝗺𝗲𝗱 𝗲𝗮𝗰𝗵 𝘁𝗶𝗺𝗲 𝗮 𝗺𝗲𝘀𝘀𝗮𝗴𝗲 𝗶𝘀 𝗱𝗲𝗹𝗲𝘁𝗲𝗱.

🌟 𝗚𝗲𝘁 𝗦𝘁𝗮𝗿𝘁𝗲𝗱:
1. Add me to your group.
2. I'll start protecting instantly.

➡️ Click on 𝗔𝗱𝗱 𝗚𝗿𝗼𝘂𝗽 to add me and keep our group safe!`,
			ctx.EffectiveUser.FirstName,
			b.User.Username,
		)

		keyboard := gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
				{
					{
						Text: "🔄 Update Channel",
						Url:  "https://t.me/Dns_Official_Channel",
					},
					{
						Text: "💬 Update Group",
						Url:  "https://t.me/dns_support_group",
					},
				},
				{
					{
						Text:         "❓ Help & Commands",
						CallbackData: "help_callback",
					},
				},
				{
					{
						Text: "➕ Add me to Your Group",
						Url: fmt.Sprintf(
							"https://t.me/%s?startgroup=s&admin=delete_messages+invite_users",
							b.User.Username,
						),
					},
				},
			},
		}
		_, err := b.SendPhoto(
			ctx.EffectiveChat.Id,
			file,
			&gotgbot.SendPhotoOpts{
				Caption:        caption,
				ProtectContent: true,
				ParseMode:      "HTML",
				ReplyMarkup:    keyboard,
			},
		)
		if err != nil {
			return fmt.Errorf("failed to send photo: %w", err)
		}

		logStr := fmt.Sprintf(
			`<a href="tg://user?id=%d">%s</a> has started the bot.

<b>User ID:</b> <code>%d</code>
<b>User Name:</b> %s %s`,
			ctx.EffectiveUser.Id,
			ctx.EffectiveUser.FirstName,
			ctx.EffectiveUser.Id,
			ctx.EffectiveUser.FirstName,
			ctx.EffectiveUser.LastName,
		)
		b.SendMessage(
			config.LoggerId,
			logStr,
			&gotgbot.SendMessageOpts{ParseMode: "HTML"},
		)
	} else if chat == "group" {
		message := `⚠️ Warning: I can't function in a basic group!

To use my features, please upgrade this group to a supergroup.

✅ How to upgrade:
1. Go to Group Settings.
2. Tap on "Chat History" and set it to "Visible".
3. Re-add me, and I'll be ready to help!`

		ctx.EffectiveMessage.Reply(b, message, nil)
		ctx.EffectiveChat.Leave(b, nil)
	} else if chat == "supergroup" {
		ctx.EffectiveMessage.Reply(b, "✅ I am active and ready to protect this supergroup!", nil)
	}
	return ext.EndGroups
}
