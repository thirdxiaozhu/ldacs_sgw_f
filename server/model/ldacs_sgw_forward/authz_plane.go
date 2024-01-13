// 自动生成模板AuthzPlane
package ldacs_sgw_forward

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 飞机业务授权 结构体  AuthzPlane
type AuthzPlane struct {
	global.GVA_MODEL
	Authz_planeId int          `json:"authz_PlaneId" form:"authz_PlaneId" gorm:"column:authz_plane_id;comment:;"` //被授权飞机
	Authz_flight  int          `json:"authz_flight" form:"authz_flight" gorm:"column:authz_flight;comment:;"`     //被授权航班
	Authz_authz   int          `json:"authz_autz" form:"authz_autz" gorm:"column:authz_autz;comment:;"`           //权限
	Authz_state   int          `json:"authz_state" form:"authz_state" gorm:"column:authz_state;comment:;"`        //授权状态
	Plane_id      AccountPlane `json:"plane_id" form:"plane_id" gorm:"foreignKey:Authz_planeId"`
	Flight        AccontFlight `json:"flight" form:"flight" gorm:"foreignKey:Authz_flight"`
	Authz         AccountAuthz `json:"authz" form:"authz" gorm:"foreignKey:Authz_authz"`
}

// TableName 飞机业务授权 AuthzPlane自定义表名 authz_plane
func (AuthzPlane) TableName() string {
	return "authz_plane"
}
