package driver4motc
import (
    pc107889e2 "github.com/starter-go/mails-over-tencent-cloud/src/driver/golang/driver"
     "github.com/starter-go/application"
)

// type pc107889e2.MOTCDriverRegistry in package:github.com/starter-go/mails-over-tencent-cloud/src/driver/golang/driver
//
// id:com-c107889e269e779a-driver-MOTCDriverRegistry
// class:class-d671d76a169061f84f6814f84b98af24-DriverRegistry
// alias:
// scope:singleton
//
type pc107889e26_driver_MOTCDriverRegistry struct {
}

func (inst* pc107889e26_driver_MOTCDriverRegistry) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-c107889e269e779a-driver-MOTCDriverRegistry"
	r.Classes = "class-d671d76a169061f84f6814f84b98af24-DriverRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pc107889e26_driver_MOTCDriverRegistry) new() any {
    return &pc107889e2.MOTCDriverRegistry{}
}

func (inst* pc107889e26_driver_MOTCDriverRegistry) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pc107889e2.MOTCDriverRegistry)
	nop(ie, com)

	
    com.CloudSecretID = inst.getCloudSecretID(ie)
    com.CloudSecretKey = inst.getCloudSecretKey(ie)
    com.CloudRegion = inst.getCloudRegion(ie)
    com.SendMailTemplateID = inst.getSendMailTemplateID(ie)


    return nil
}


func (inst*pc107889e26_driver_MOTCDriverRegistry) getCloudSecretID(ie application.InjectionExt)string{
    return ie.GetString("${cloud.tencent.secret-id}")
}


func (inst*pc107889e26_driver_MOTCDriverRegistry) getCloudSecretKey(ie application.InjectionExt)string{
    return ie.GetString("${cloud.tencent.secret-key}")
}


func (inst*pc107889e26_driver_MOTCDriverRegistry) getCloudRegion(ie application.InjectionExt)string{
    return ie.GetString("${cloud.tencent.region}")
}


func (inst*pc107889e26_driver_MOTCDriverRegistry) getSendMailTemplateID(ie application.InjectionExt)uint64{
    return ie.GetUint64("${cloud.tencent.send-mail.template-id}")
}


