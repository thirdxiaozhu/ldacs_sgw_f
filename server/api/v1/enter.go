package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/ldacs_sgw_forward"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup            system.ApiGroup
	ExampleApiGroup           example.ApiGroup
	Ldacs_sgw_forwardApiGroup ldacs_sgw_forward.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
