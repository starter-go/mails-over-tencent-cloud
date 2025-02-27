package unit

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/starter-go/application"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/mails"
)

// SendMailExample ...
type SendMailExample struct {

	//starter:component

	Sender mails.Service //starter:inject("#")

	Address1 string //starter:inject("${test.send-mail.address-from}")
	Address2 string //starter:inject("${test.send-mail.address-to}")

}

func (inst *SendMailExample) _impl() application.Lifecycle {
	return inst
}

// Life ...
func (inst *SendMailExample) Life() *application.Life {
	l := &application.Life{
		OnLoop: inst.run,
	}
	return l
}

func (inst *SendMailExample) run() error {

	now := lang.Now()
	params := map[string]string{}
	params["code"] = strconv.FormatInt(now.Int()/1000, 10)

	addr1 := inst.Address1
	addr2 := inst.Address2
	content, err := json.Marshal(params)
	if err != nil {
		return err
	}

	ctx := context.Background()
	msg := &mails.Message{
		FromAddress: mails.Address(addr1),
		ToAddresses: []mails.Address{mails.Address(addr2)},
		Title:       "a demo mail (SendMailExample)",
		ContentType: "application/json",
		Content:     content,
	}

	return inst.Sender.Send(ctx, msg)
}
