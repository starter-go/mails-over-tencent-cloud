package test4motc
import (
    pd671d76a1 "github.com/starter-go/mails"
    p330ce58af "github.com/starter-go/mails-over-tencent-cloud/src/test/golang/unit"
     "github.com/starter-go/application"
)

// type p330ce58af.SendMailExample in package:github.com/starter-go/mails-over-tencent-cloud/src/test/golang/unit
//
// id:com-330ce58af7e5859c-unit-SendMailExample
// class:
// alias:
// scope:singleton
//
type p330ce58af7_unit_SendMailExample struct {
}

func (inst* p330ce58af7_unit_SendMailExample) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-330ce58af7e5859c-unit-SendMailExample"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p330ce58af7_unit_SendMailExample) new() any {
    return &p330ce58af.SendMailExample{}
}

func (inst* p330ce58af7_unit_SendMailExample) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p330ce58af.SendMailExample)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Address1 = inst.getAddress1(ie)
    com.Address2 = inst.getAddress2(ie)


    return nil
}


func (inst*p330ce58af7_unit_SendMailExample) getSender(ie application.InjectionExt)pd671d76a1.Service{
    return ie.GetComponent("#alias-d671d76a169061f84f6814f84b98af24-Service").(pd671d76a1.Service)
}


func (inst*p330ce58af7_unit_SendMailExample) getAddress1(ie application.InjectionExt)string{
    return ie.GetString("${test.send-mail.address-from}")
}


func (inst*p330ce58af7_unit_SendMailExample) getAddress2(ie application.InjectionExt)string{
    return ie.GetString("${test.send-mail.address-to}")
}


