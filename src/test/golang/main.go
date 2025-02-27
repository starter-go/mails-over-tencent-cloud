package main

import (
	"os"

	"github.com/starter-go/mails-over-tencent-cloud/modules/motc"
	"github.com/starter-go/starter"
)

func main() {
	m := motc.ModuleForTest()
	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
