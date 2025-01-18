package bootstrap

import (
	"backed/app/core/consts"
	"backed/app/core/variable"
	"backed/app/models"
	Marketplace "backed/internal/contract"
	"backed/internal/database"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
)

func init() {
	// 读取配置文件
	newViper := viper.New()
	// 设置配置文件存在目录
	newViper.SetConfigFile(consts.ConfigFilePath)
	// 读取配置文件内容
	err := newViper.ReadInConfig()
	if err != nil {
		log.Println("读取配置文件出错: ", err.Error())
	}
	// 初始化化viper
	variable.Viper = newViper

	// 配置日志输出到文件
	logFile, err := os.OpenFile(consts.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		// 设置日志输出到文件
		logrus.SetOutput(logFile)
	} else {
		logrus.Info("设置日志文件成功...")
	}

	// 设置日志级别 日志等级设置为最低
	logrus.SetLevel(logrus.TraceLevel)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		newViper.GetString("mysql.username"),
		newViper.GetString("mysql.password"),
		newViper.GetString("mysql.host"),
		newViper.GetString("mysql.port"),
		newViper.GetString("mysql.database"))

	variable.Gorm, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Info("连接mysql出错: ", err.Error())
		os.Exit(http.StatusBadRequest)
	}

	// 设置Gorm的日志级别为Error，这样就不会打印查询语句了
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目的地）
		logger.Config{
			LogLevel: logger.Error, // 只有错误信息会被打印
			// 其他日志配置...
		},
	)

	variable.Gorm.Logger = newLogger

	// 解析配置文件
	configs, err := conf.ParseConfigFile(consts.FiscoBcosFilePath)
	if err != nil {
		log.Fatal(err)
	}
	config := &configs[0]
	// 连接webase客户端
	client, err := client.Dial(config)
	if err != nil {
		log.Fatal("连接webase客户端失败: ", err.Error())
	}

	variable.FiscoBcos = client

	// 初始化合约对象
	ContractAddress := common.HexToAddress(variable.Viper.GetString("contract_address"))
	variable.ContractAddress = &ContractAddress

	// 创建合约实例
	instance, err := Marketplace.NewMarketplace(*variable.ContractAddress, variable.FiscoBcos)
	if err != nil {
		logrus.Error("创建合约实例失败: ", err.Error())
	}

	// 获取合约session
	session := &Marketplace.MarketplaceSession{Contract: instance, CallOpts: *variable.FiscoBcos.GetCallOpts(), TransactOpts: *variable.FiscoBcos.GetTransactOpts()}
	variable.ContractSession = session

	// 迁移数据库
	variable.Gorm.AutoMigrate(&models.UserInfo{})
	variable.Gorm.AutoMigrate(&models.File{})
	variable.Gorm.AutoMigrate(&models.GoodsCard{})
	variable.Gorm.AutoMigrate(&models.FruitKind{})

	// 设置默认初始化数据库
	database.SetDefault(variable.Gorm)
}
