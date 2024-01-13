package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/ldacs_sgw_forward"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System            system.RouterGroup
	Example           example.RouterGroup
	Ldacs_sgw_forward ldacs_sgw_forward.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
