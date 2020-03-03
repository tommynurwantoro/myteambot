package method

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bot/myteambot/app/utility"
	"github.com/bot/myteambot/app/utility/repository"
	tb "gopkg.in/tucnak/telebot.v2"
)

// TitipReview _
func (m *Method) TitipReview(message *tb.Message) {
	invalid := ValidateGroup(message)

	if invalid != "" {
		m.Bot.Send(message.Chat, invalid)
		return
	}

	if message.Payload == "" {
		m.Bot.Send(message.Chat, utility.InvalidParameter())
		return
	}

	split := strings.Split(message.Payload, "#")

	if len(split) < 3 {
		m.Bot.Send(message.Chat, utility.InvalidParameter())
		return
	}

	repository.InsertReview(split[0], split[1], split[2], message.Chat.ID)

	m.Bot.Send(message.Chat, utility.SuccessInsertData())
}

// AntrianReview _
func (m *Method) AntrianReview(message *tb.Message) {
	invalid := ValidateGroup(message)

	if invalid != "" {
		m.Bot.Send(message.Chat, invalid)
		return
	}

	reviews := repository.GetAllNeedReview(message.Chat.ID)

	if len(reviews) == 0 {
		m.Bot.Send(message.Chat, "Gak ada antrian review nih 👍🏻")
		return
	}

	m.Bot.Send(message.Chat, fmt.Sprintf("Ini antrian review tim kamu:\n%s", utility.GenerateHTMLReview(reviews)), tb.ModeHTML, tb.NoPreview)
}

// SudahDireview _
func (m *Method) SudahDireview(message *tb.Message) {
	invalid := ValidateGroup(message)

	if invalid != "" {
		m.Bot.Send(message.Chat, invalid)
		return
	}

	if message.Payload == "" {
		m.Bot.Send(message.Chat, utility.InvalidParameter())
		return
	}

	sequences := strings.Split(message.Payload, " ")
	success := repository.UpdateToDoneReview(sequences, message.Chat.ID, fmt.Sprintf("@%s", message.Sender.Username), false)

	if success {
		m.Bot.Send(message.Chat, utility.SuccessUpdateData())
		m.AntrianReview(message)
		return
	}

	m.Bot.Send(message.Chat, utility.InvalidSequece())
}

// SudahDireviewSemua _
func (m *Method) SudahDireviewSemua(message *tb.Message) {
	invalid := ValidateGroup(message)

	if invalid != "" {
		m.Bot.Send(message.Chat, invalid)
		return
	}

	if message.Payload == "" {
		m.Bot.Send(message.Chat, utility.InvalidParameter())
		return
	}

	sequences := strings.Split(message.Payload, " ")
	success := repository.UpdateToDoneReview(sequences, message.Chat.ID, fmt.Sprintf("@%s", message.Sender.Username), true)

	if success {
		m.Bot.Send(message.Chat, utility.SuccessUpdateData())
		m.AntrianReview(message)
		return
	}

	m.Bot.Send(message.Chat, utility.InvalidSequece())
}

// TambahUserReview _
func (m *Method) TambahUserReview(message *tb.Message) {
	invalid := ValidateGroup(message)

	if invalid != "" {
		m.Bot.Send(message.Chat, invalid)
		return
	}

	if message.Payload == "" {
		m.Bot.Send(message.Chat, utility.InvalidParameter())
		return
	}

	split := strings.Split(message.Payload, "#")

	sequence, err := strconv.Atoi(split[0])

	if len(split) < 2 || err != nil {
		m.Bot.Send(message.Chat, utility.InvalidParameter())
		return
	}

	reviews := repository.GetAllNeedReview(message.Chat.ID)

	for i, review := range reviews {
		if i+1 == sequence {
			repository.UpdateReview(review.ID, review.Title, review.URL, fmt.Sprintf("%s %s", review.Users, split[1]))
			m.Bot.Send(message.Chat, utility.SuccessUpdateData())
			m.AntrianReview(message)
			return
		}
	}

	m.Bot.Send(message.Chat, utility.InvalidSequece())
}

// HapusReview _
func (m *Method) HapusReview(message *tb.Message) {
	invalid := ValidateGroup(message)

	if invalid != "" {
		m.Bot.Send(message.Chat, invalid)
		return
	}

	if message.Payload == "" {
		m.Bot.Send(message.Chat, utility.InvalidParameter())
		return
	}

	sequences := strings.Split(message.Payload, " ")
	success := repository.DeleteReview(sequences, message.Chat.ID, fmt.Sprintf("@%s", message.Sender.Username))

	if success {
		m.Bot.Send(message.Chat, utility.SuccessUpdateData())
		m.AntrianReview(message)
		return
	}

	m.Bot.Send(message.Chat, utility.InvalidSequece())
}

// SiapQA _
func (m *Method) SiapQA(message *tb.Message) {
	invalid := ValidateGroup(message)

	if invalid != "" {
		m.Bot.Send(message.Chat, invalid)
		return
	}

	if message.Payload == "" {
		m.Bot.Send(message.Chat, utility.InvalidParameter(), tb.ModeHTML, tb.NoPreview)
		return
	}

	sequences := strings.Split(message.Payload, " ")
	success := repository.UpdateToReadyQA(sequences, message.Chat.ID)

	if success {
		m.Bot.Send(message.Chat, utility.SuccessUpdateData(), tb.ModeHTML, tb.NoPreview)
		m.AntrianReview(message)
		return
	}

	m.Bot.Send(message.Chat, utility.InvalidSequece(), tb.ModeHTML, tb.NoPreview)
}

// AntrianQA _
func (m *Method) AntrianQA(message *tb.Message) {
	invalid := ValidateGroup(message)

	if invalid != "" {
		m.Bot.Send(message.Chat, invalid)
		return
	}

	reviews := repository.GetAllNeedQA(message.Chat.ID)

	if len(reviews) == 0 {
		m.Bot.Send(message.Chat, "Gak ada antrian QA nih 👍🏻", tb.ModeHTML, tb.NoPreview)
		return
	}

	m.Bot.Send(message.Chat, fmt.Sprintf("Ini antrian QA tim kamu:\n%s", utility.GenerateHTMLReview(reviews)), tb.ModeHTML, tb.NoPreview)
}

// SudahDites _
func (m *Method) SudahDites(message *tb.Message) {
	invalid := ValidateGroup(message)

	if invalid != "" {
		m.Bot.Send(message.Chat, invalid)
		return
	}

	if message.Payload == "" {
		m.Bot.Send(message.Chat, utility.InvalidParameter(), tb.ModeHTML, tb.NoPreview)
		return
	}

	sequences := strings.Split(message.Payload, " ")
	success := repository.UpdateToDoneQA(sequences, message.Chat.ID)

	if success {
		m.Bot.Send(message.Chat, utility.SuccessUpdateData(), tb.ModeHTML, tb.NoPreview)
		m.AntrianQA(message)
		return
	}

	m.Bot.Send(message.Chat, utility.InvalidSequece(), tb.ModeHTML, tb.NoPreview)
}
