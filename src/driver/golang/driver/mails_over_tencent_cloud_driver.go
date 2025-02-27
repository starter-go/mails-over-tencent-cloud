package driver

import (
	"context"
	"fmt"
	"strings"

	"github.com/starter-go/mails"
	"github.com/starter-go/vlog"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	ses "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ses/v20201002"
)

const theDriverName = "tencentcloud"

////////////////////////////////////////////////////////////////////////////////

// MOTCDriverRegistry ...
type MOTCDriverRegistry struct {

	//starter:component

	_as func(mails.DriverRegistry) //starter:as(".")

	CloudSecretID  string //starter:inject("${cloud.tencent.secret-id}")
	CloudSecretKey string //starter:inject("${cloud.tencent.secret-key}")
	CloudRegion    string //starter:inject("${cloud.tencent.region}")

	SendMailTemplateID uint64 //starter:inject("${cloud.tencent.send-mail.template-id}")

}

func (inst *MOTCDriverRegistry) _impl() mails.DriverRegistry {
	return inst
}

// ListRegistrations ...
func (inst *MOTCDriverRegistry) ListRegistrations() []*mails.DriverRegistration {

	driver := &motcDriver{
		parent: inst,
	}

	reg := &mails.DriverRegistration{
		Name:    theDriverName,
		Enabled: true,
		Driver:  driver,
	}

	return []*mails.DriverRegistration{reg}
}

////////////////////////////////////////////////////////////////////////////////

// MOTCDriver ...
type motcDriver struct {
	parent *MOTCDriverRegistry
}

func (inst *motcDriver) _impl() mails.Driver {
	return inst
}

func (inst *motcDriver) Accept(cfg *mails.Configuration) bool {
	if cfg == nil {
		return false
	}
	return cfg.Driver == theDriverName
}

func (inst *motcDriver) CreateDispatcher(cfg *mails.Configuration) (mails.Dispatcher, error) {
	if !inst.Accept(cfg) {
		return nil, fmt.Errorf("unsupported mails configuration")
	}
	disp := &motcDispatcher{
		parent: inst.parent,
	}
	disp.config = *cfg
	return disp, nil
}

////////////////////////////////////////////////////////////////////////////////

type motcDispatcher struct {
	config mails.Configuration
	parent *MOTCDriverRegistry
}

func (inst *motcDispatcher) _impl() mails.Dispatcher {
	return inst
}

func (inst *motcDispatcher) Accept(c context.Context, msg *mails.Message) bool {
	a1 := inst.config.SenderAddress
	a2 := msg.FromAddress
	str := a1.String()
	return a1 == a2 && strings.ContainsRune(str, '@')
}

func (inst *motcDispatcher) Send(ctx context.Context, msg *mails.Message) error {

	client, err := inst.getClient()
	if err != nil {
		return err
	}

	request := ses.NewSendEmailRequest()
	err = inst.prepareSendEmailRequest(request, msg)
	if err != nil {
		return err
	}

	response, err := client.SendEmailWithContext(ctx, request)
	if err != nil {
		return err
	}

	respMsgID := response.Response.MessageId
	vlog.Info("send email ok, resp=" + *respMsgID)
	return nil
}

func (inst *motcDispatcher) getClient() (*ses.Client, error) {

	secretID := inst.parent.CloudSecretID
	secretKey := inst.parent.CloudSecretKey

	credential := common.NewCredential(secretID, secretKey)
	region := inst.getRegion()
	clientProfile := profile.NewClientProfile()

	return ses.NewClient(credential, region, clientProfile)
}

func (inst *motcDispatcher) getRegion() string {

	region1 := inst.parent.CloudRegion
	region2 := strings.ToLower("ap-" + region1)
	have := ""

	all := []string{
		regions.Beijing,
		regions.Shanghai,
		regions.Guangzhou,
		regions.ShenzhenFSI,
		regions.HongKong,
	}

	for _, item := range all {
		if item == region2 {
			return item // 全等，直接返回结果
		}
		if strings.Contains(item, region2) {
			have = item // 相似, 暂存待用
		}
	}

	if have == "" {
		return "bad-tencentcloud-region:" + region1
	}
	return have
}

func (inst *motcDispatcher) getAddress2(msg *mails.Message) (mails.Address, error) {

	result := mails.Address("")
	count := 0
	list := msg.ToAddresses
	for _, item := range list {
		result = item
		count++
	}

	if count == 1 && result != "" {
		return result, nil
	}

	return "", fmt.Errorf("必须有且只能有一个目标邮件地址")
}

func (inst *motcDispatcher) prepareSendEmailRequest(req *ses.SendEmailRequest, msg *mails.Message) error {

	address2, err := inst.getAddress2(msg)
	if err != nil {
		return err
	}

	title := msg.Title                           // @param
	addr1 := msg.FromAddress.String()            // @param
	addr2 := address2.String()                   // @param
	templateID := inst.parent.SendMailTemplateID // @param
	contentType := msg.ContentType               // @param
	content := msg.Content                       // @param

	// 注意：传入的内容必须是JSON格式 (application/json)
	// 并且符合 tencentcloud/ses 模板参数的形式, 如： {"p1":"value1"}
	if contentType != "application/json" {
		return fmt.Errorf("通过 mails.Message 传入的内容必须是JSON格式 (application/json), 并且符合 tencentcloud/ses 模板参数的形式")
	}

	templateData := string(content)
	templ := &ses.Template{}
	templ.TemplateID = &templateID
	templ.TemplateData = &templateData

	triggerType := uint64(1)

	req.Subject = &title
	req.Template = templ
	req.FromEmailAddress = &addr1
	req.Destination = []*string{&addr2}
	req.TriggerType = &triggerType

	return nil
}

////////////////////////////////////////////////////////////////////////////////
