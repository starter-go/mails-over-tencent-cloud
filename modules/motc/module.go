package motc

import (
	"github.com/starter-go/application"
	motc1 "github.com/starter-go/mails-over-tencent-cloud"
	"github.com/starter-go/mails-over-tencent-cloud/gen/driver4motc"
	"github.com/starter-go/mails-over-tencent-cloud/gen/test4motc"
	"github.com/starter-go/mails/modules/mails"
	"github.com/starter-go/starter"
)

// ModuleForDriver ...
func ModuleForDriver() application.Module {
	mb := motc1.NewDriverModule()
	mb.Components(driver4motc.ExportComponents)
	mb.Depend(mails.LibModule())
	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {
	mb := motc1.NewTestModule()
	mb.Components(test4motc.ExportComponents)
	mb.Depend(ModuleForDriver())
	mb.Depend(starter.Module())
	return mb.Create()
}
