package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/ldacs_sgw_forward"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup            system.ServiceGroup
	ExampleServiceGroup           example.ServiceGroup
	Ldacs_sgw_forwardServiceGroup ldacs_sgw_forward.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
