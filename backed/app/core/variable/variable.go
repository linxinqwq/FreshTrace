package variable

import (
	Marketplace "backed/internal/contract"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	Viper           *viper.Viper // 读取配置文件包
	Gorm            *gorm.DB     // 存储mysql连接
	ContractAddress *common.Address
	FiscoBcos       *client.Client // fisco bcos客户端

	ContractSession *Marketplace.MarketplaceSession
	ContractAbi     string
)
