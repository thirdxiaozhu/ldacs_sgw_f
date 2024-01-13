// 自动生成模板AccontFlight
package ldacs_sgw_forward

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 航班 结构体  AccontFlight
type AccontFlight struct {
      global.GVA_MODEL
      Flight  string `json:"flight" form:"flight" gorm:"column:flight;comment:;"binding:"required"`  //航班号 
}


// TableName 航班 AccontFlight自定义表名 account_flight
func (AccontFlight) TableName() string {
  return "account_flight"
}

