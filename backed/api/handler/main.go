package handler

import (
	"backed/app/core/variable"
	"backed/app/models"
	"backed/app/utils"
	Marketplace "backed/internal/contract"
	"backed/internal/database"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/big"
	"net/http"
	"strconv"
	"time"
	"fmt"
)

// FileUpload 文件上传接口
func FileUpload(ctx *gin.Context) {
	var requestData models.File
	// 获取请求参数
	err := ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	// 创建uuid
	UUID, err := uuid.NewUUID()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "创建uuid失败: " + err.Error(),
		})
		return
	}

	requestData.Id = UUID.String()

	// 创建文件对象
	fileDao := database.File
	err = fileDao.WithContext(ctx).Create(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "创建文件对象失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "文件上传成功",
		"id":      requestData.Id,
	})
}

// AddGoods 供应商添加商品接口
func AddGoods(ctx *gin.Context) {
	// 获取用户ID
	username, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户姓名失败: " + err.Error(),
		})
		return
	}

	// 检查用户权限
	userDao := database.UserInfo
	userInfo, err := userDao.WithContext(ctx).
		Where(userDao.Username.Eq(username)).First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败: " + err.Error(),
		})
		return
	}

	if userInfo.Identity == "用户" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户无权添加商品",
		})
		return
	}

	// 获取用户调用的请求参数
	var requestData models.AddGoodsRequest
	err = ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	_, receipt, err := variable.ContractSession.AddProduct(
		requestData.Images,
		big.NewInt(requestData.Store),
		big.NewInt(requestData.Price),
		requestData.Name,
		requestData.Uint,
		username,
		requestData.Desc,
		requestData.Origin,
		requestData.Kind,
	)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "添加商品失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":            200,
		"message":         "添加商品成功",
		"transactionHash": receipt.TransactionHash,
	})
}

// GetUserManageGoods 获取商品列表
func GetUserManageGoods(ctx *gin.Context) {
	// 获取用户ID
	username, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户姓名失败: " + err.Error(),
		})
		return
	}

	// 检查用户权限
	userDao := database.UserInfo
	userInfo, err := userDao.WithContext(ctx).
		Where(userDao.Username.Eq(username)).First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败: " + err.Error(),
		})
		return
	}

	if userInfo.Identity == "用户" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户无权���加商品",
		})
		return
	}

	// 获取页面
	pageString := ctx.Query("page")
	page, err := strconv.ParseInt(pageString, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	// 获取商品列表
	products, err := variable.ContractSession.GetSellerListings(username)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取商品列表失败: " + err.Error(),
		})
		return
	}
	var filteredProducts []Marketplace.MarketplaceProductResult

	zero := big.NewInt(0) // 用于比较的0值
	// 删除商品中
	for _, product := range products {
		if product.Id.Cmp(zero) != 0 {
			filteredProducts = append(filteredProducts, product)
		}
	}

	if len(filteredProducts) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":     200,
			"message":  "获取商品列表成功",
			"products": []Marketplace.MarketplaceProduct{},
		})
		return
	}

	// 假设page为当前页码，pageSize为每页大小（本例中为10）
	pageSize := int64(6)

	// 计算总页数
	totalPages := (int64(len(filteredProducts)) + pageSize - 1) / pageSize

	// 确保请求的页码在有效范围内
	if page < 1 {
		page = 1
	}
	if page > int64(totalPages) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":     200,
			"message":  "获取商品列表成功",
			"products": []Marketplace.MarketplaceProduct{},
		})
		return
	}

	// 计算当前页的起始和结束索引
	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize
	if endIndex > int64(len(filteredProducts)) {
		endIndex = int64(len(filteredProducts))
	}

	// 获取当前页的产品切片
	var currentPageProducts []Marketplace.MarketplaceProductResult // 假设Product是你的产品类型
	if startIndex < int64(len(filteredProducts)) {
		currentPageProducts = filteredProducts[startIndex:endIndex]
	} else {
		// 如果startIndex超出了products的长度，返回一个空切片
		currentPageProducts = []Marketplace.MarketplaceProductResult{}
	}

	var fileIds []string
	for _, product := range currentPageProducts {
		fileIds = append(fileIds, product.ImageId)
		for _, element := range product.Process {
			fileIds = append(fileIds, element.ImageId)
		}
	}

	fileDao := database.File
	find, err := fileDao.WithContext(ctx).Where(fileDao.Id.In(fileIds...)).Find()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取商品图片失败: " + err.Error(),
		})
		return
	}

	fileMap := make(map[string]string) // 假设Id是string类型，FileData也是string
	for _, file := range find {
		fileMap[file.Id] = file.FileData
	}

	for i, product := range currentPageProducts {
		if newData, ok := fileMap[product.ImageId]; ok {
			currentPageProducts[i].ImageId = newData // 确保对原切片进行修改
			for j, element := range product.Process {
				if newData, ok := fileMap[element.ImageId]; ok {
					currentPageProducts[i].Process[j].ImageId = newData // 确保对原切片进行修改
				}
			}
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":     200,
		"message":  "获取商品列表成功",
		"products": currentPageProducts,
	})

}

// UpdateGoodsInfo 修改商品信息
func UpdateGoodsInfo(ctx *gin.Context) {
	// 获取用户ID
	username, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户姓名失败: " + err.Error(),
		})
		return
	}

	// 获取请求参数
	var requestData models.UpdateGoodsInfoRequest
	err = ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	_, receipt, err := variable.ContractSession.UpdateProduct(big.NewInt(requestData.Id), big.NewInt(requestData.Store), big.NewInt(requestData.Price), requestData.Desc, requestData.Name, requestData.Images, username, requestData.Origin, requestData.Kind, requestData.Uint)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "修改产品失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":            200,
		"message":         "商品修改成功",
		"transactionHash": receipt.TransactionHash,
	})

}

// RemoveGoods 商品下架
func RemoveGoods(ctx *gin.Context) {
	// 获取请求参数
	var requestData models.RemoveGoodsRequest
	err := ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	_, receipt, err := variable.ContractSession.ReviewProduct(big.NewInt(requestData.Id), uint8(requestData.Status))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "删除商品失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":            200,
		"message":         "商品删除成功",
		"transactionHash": receipt.TransactionHash,
	})
}

// GetSalesTopThree 获取销售榜前三
func GetSalesTopThree(ctx *gin.Context) {
	products, err := variable.ContractSession.GetTopSellingProducts()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取销售榜失败: " + err.Error(),
		})
		return
	}

	var fileIds []string
	for _, product := range products {
		fileIds = append(fileIds, product.ImageId)
	}

	fileDao := database.File
	find, err := fileDao.WithContext(ctx).Where(fileDao.Id.In(fileIds...)).Find()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取商品图片失败: " + err.Error(),
		})
		return
	}

	fileMap := make(map[string]string) // 假设Id是string类型，FileData也是string
	for _, file := range find {
		fileMap[file.Id] = file.FileData
	}

	for i, product := range products {
		if newData, ok := fileMap[product.ImageId]; ok {
			products[i].ImageId = newData // 确保对原切片进行修改
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":     200,
		"message":  "获取销售榜成功",
		"products": products,
	})
}

// GetGoodInfo 获取商品详情
func GetGoodInfo(ctx *gin.Context) {
	// 获取请求参数
	idString := ctx.Query("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	product, err := variable.ContractSession.GetProductDetail(big.NewInt(id))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取商品详情失败: " + err.Error(),
		})
		return
	}

	fileDao := database.File
	find, err := fileDao.WithContext(ctx).Where(fileDao.Id.Eq(product.ImageId)).First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取商品图片失败: " + err.Error(),
		})
		return
	}

	product.ImageId = find.FileData

	for i := range product.Process {
		element := &product.Process[i] // 获取指向切片元素的指针
		find, err := fileDao.WithContext(ctx).Where(fileDao.Id.Eq(element.ImageId)).First()
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "获取商品图片失败: " + err.Error(),
			})
			return
		}
		element.ImageId = find.FileData
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取商品详情成功",
		"product": product,
	})
}

// BuyGoods 购买商品信息
func BuyGoods(ctx *gin.Context) {
	// 获取用户ID
	username, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户姓名失败: " + err.Error(),
		})
		return
	}

	// 获取请求参数
	var requestData models.BuyGoodsRequest
	err = ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	_, receipt, err := variable.ContractSession.BuyProduct(big.NewInt(requestData.Id), big.NewInt(requestData.Store), username, requestData.Address, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "购买商品失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":            200,
		"message":         "商品购买成功",
		"transactionHash": receipt.TransactionHash,
	})
}

// GetUserOrders 获取用户订单
func GetUserOrders(ctx *gin.Context) {
	// 获取用户ID
	username, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户姓名失败: " + err.Error(),
		})
		return
	}

	orders, err := variable.ContractSession.GetUserPurchases(username)

	// 获取页面
	pageString := ctx.Query("page")
	page, err := strconv.ParseInt(pageString, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	if len(orders) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":     200,
			"message":  "获取商品列表成功",
			"products": []models.GetUserOrdersResponse{},
		})
		return
	}

	// 假设page为当前页码，pageSize为每页大小（本例中为10）
	pageSize := int64(6)

	// 计算总页数
	totalPages := (int64(len(orders)) + pageSize - 1) / pageSize

	// 确保请求的页码在有效范围内
	if page < 1 {
		page = 1
	}
	if page > int64(totalPages) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":     200,
			"message":  "获取商品列表成功",
			"products": []models.GetUserOrdersResponse{},
		})
		return
	}

	// 计算当前页的起始和结束索引
	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize
	if endIndex > int64(len(orders)) {
		endIndex = int64(len(orders))
	}

	// 获取当前页的产品切片
	var currentPageProducts []Marketplace.MarketplacePurchaseOrder // 假设Product是你的产品类型
	if startIndex < int64(len(orders)) {
		currentPageProducts = orders[startIndex:endIndex]
	} else {
		// 如果startIndex超出了products的长度，返回一个空切片
		currentPageProducts = []Marketplace.MarketplacePurchaseOrder{}
	}

	fileDao := database.File
	// 获取商品ID
	var productIds []models.GetUserOrdersResponse
	for _, element := range currentPageProducts {
		product, err := variable.ContractSession.GetProductDetail(element.ProductId)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "获取商品详情失败: " + err.Error(),
			})
			return
		}

		find, err := fileDao.WithContext(ctx).Where(fileDao.Id.Eq(product.ImageId)).First()
		if err != nil {
			continue
		}
		if find == nil {
			find.Id = ""
		}
		var a int64 = 0
		if element.Status {
			a = 1
		}
		productIds = append(productIds, models.GetUserOrdersResponse{
			Id:          element.ProductId.Int64(),
			Number:      element.Quantity.Int64(),
			Name:        product.Name,
			Senders:     product.Seller,
			BlockNumber: element.BlockNumber.Int64(),
			Image:       find.FileData,
			Price:       product.Price.Int64(),
			TotalPrice:  element.Quantity.Int64() * product.Price.Int64(),
			CreateTime:  element.CreatedTime,
			Address:     element.UserAddress,
			Status:      a,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":     200,
		"message":  "获取商品列表成功",
		"products": productIds,
	})
}

// GetGoodsList 商品列表
func GetGoodsList(ctx *gin.Context) {
	// 首先获取上次最后的商品ID
	lastIdString := ctx.Query("lastId")
	// 获取水果种类
	kindString := ctx.Query("kind")
	lastId, err := strconv.ParseInt(lastIdString, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	id, err := variable.ContractSession.NextProductId()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取商品ID失败: " + err.Error(),
		})
		return
	}

	// 说明还有下一页
	if lastId < id.Int64()-1 {
		var result []Marketplace.MarketplaceProductResult
		for {
			lastId += 1
			product, err := variable.ContractSession.GetProductDetail(big.NewInt(lastId))
			if err != nil {
				break
			}

			if product.Status != 2 {
				continue
			}

			if kindString != "全部" && product.Kind != kindString {
				continue
			}

			fileDao := database.File
			find, err := fileDao.WithContext(ctx).Where(fileDao.Id.Eq(product.ImageId)).First()
			if err != nil {
				continue
			}

			product.ImageId = find.FileData
			result = append(result, product)
			if len(result) == 8 {
				break
			}

			if lastId == id.Int64()-1 {
				break
			}
		}

		ctx.JSON(http.StatusOK, gin.H{
			"code":     200,
			"message":  "获取商品列表成功",
			"products": result,
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "没有下一页",
		})
	}
}

// GetUncheckedGoods 获取未审核的商品
func GetUncheckedGoods(ctx *gin.Context) {
	// 获取用户ID
	username, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户姓名失败: " + err.Error(),
		})
		return
	}

	// 检查用户权限
	userDao := database.UserInfo
	userInfo, err := userDao.WithContext(ctx).
		Where(userDao.Username.Eq(username)).First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败: " + err.Error(),
		})
		return
	}

	if userInfo.Identity == "用户" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户无权添加商品",
		})
		return
	}

	// 获取页面
	pageString := ctx.Query("page")
	page, err := strconv.ParseInt(pageString, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	// 获取商品列表
	products, err := variable.ContractSession.GetPendingProducts()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取商品列表失败: " + err.Error(),
		})
		return
	}

	var filteredProducts []Marketplace.MarketplaceProductResult

	zero := big.NewInt(0) // 用于比较的0值
	// 删除商品中
	for _, product := range products {
		if product.Id.Cmp(zero) != 0 {
			filteredProducts = append(filteredProducts, product)
		}
	}

	if len(filteredProducts) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":     200,
			"message":  "获取商品列表成功",
			"products": []Marketplace.MarketplaceProductResult{},
		})
		return
	}

	// 假设page为当前页码，pageSize为每页大小（本例中为10）
	pageSize := int64(6)

	// 计算总页数
	totalPages := (int64(len(filteredProducts)) + pageSize - 1) / pageSize

	// 确保请求的页码在有效范围内
	if page < 1 {
		page = 1
	}
	if page > int64(totalPages) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":     200,
			"message":  "获取商品列表成功",
			"products": []Marketplace.MarketplaceProduct{},
		})
		return
	}

	// 计算当前页的起始和结束索引
	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize
	if endIndex > int64(len(filteredProducts)) {
		endIndex = int64(len(filteredProducts))
	}

	// 获取当前页的产品切片
	var currentPageProducts []Marketplace.MarketplaceProductResult // 假设Product是你的产品类型
	if startIndex < int64(len(filteredProducts)) {
		currentPageProducts = filteredProducts[startIndex:endIndex]
	} else {
		// 如果startIndex超出了products的长度，返回一个空切片
		currentPageProducts = []Marketplace.MarketplaceProductResult{}
	}

	var fileIds []string
	for _, product := range currentPageProducts {
		fileIds = append(fileIds, product.ImageId)
		for _, element := range product.Process {
			fileIds = append(fileIds, element.ImageId)
		}
	}

	fileDao := database.File
	find, err := fileDao.WithContext(ctx).Where(fileDao.Id.In(fileIds...)).Find()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取商品图片失败: " + err.Error(),
		})
		return
	}

	fileMap := make(map[string]string) // 假设Id是string类型，FileData也是string
	for _, file := range find {
		fileMap[file.Id] = file.FileData
	}

	for i, product := range currentPageProducts {
		if newData, ok := fileMap[product.ImageId]; ok {
			currentPageProducts[i].ImageId = newData // 确保对原切片进行修改
			for j, element := range product.Process {
				if newData, ok := fileMap[element.ImageId]; ok {
					currentPageProducts[i].Process[j].ImageId = newData // 确保对原切片进行修改
				}
			}
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":     200,
		"message":  "获取商品列表成功",
		"products": currentPageProducts,
	})
}

// UpdateCheckStatus 修改审核状态
func UpdateCheckStatus(ctx *gin.Context) {
	// 获取用户ID
	username, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户姓名失败: " + err.Error(),
		})
		return
	}

	// 判断用户权限
	userDao := database.UserInfo
	userInfo, err := userDao.WithContext(ctx).Where(userDao.Username.Eq(username)).First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败:" + err.Error(),
		})
		return
	}

	if userInfo.Identity != "管理员" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户非审核员，无法进行下列操作",
		})
		return
	}

	// 获取请求参数
	var requestData models.UpdateCheckStatusRequest
	err = ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	_, receipt, err := variable.ContractSession.ReviewProduct(big.NewInt(requestData.Id), uint8(requestData.Status))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "修改审核状态失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":            200,
		"message":         "修改审核状态成功",
		"transactionHash": receipt.TransactionHash,
	})
}

// UpdateGoodsProcess 更新商家进程
func UpdateGoodsProcess(ctx *gin.Context) {
	// 获取用户ID
	username, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户姓名失败: " + err.Error(),
		})
		return
	}

	// 获取请求参数
	var requestData models.UpdateGoodsProcessRequest
	err = ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	// 获取当前时间日期string
	timeStr := time.Now().Format("2006-01-02 15:04:05")

	_, receipt, err := variable.ContractSession.AddProductionCondition(username, big.NewInt(requestData.Id), timeStr, requestData.Operator, requestData.ImageId, uint8(requestData.Status))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "修改商品进程失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":            200,
		"message":         "修改商品进程成功",
		"transactionHash": receipt.TransactionHash,
	})
}

// AddToCart 加入购物车
func AddToCart(ctx *gin.Context) {
	// 获取用户名
	username, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户名失败: " + err.Error(),
		})
		return
	}

	// 获取请求参数
	var requestData models.GoodsCard
	err = ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	// 设置用户名
	requestData.UserName = username

	// 创建购物车
	cartDao := database.GoodsCard

	// 首先确保购物车的商品数量少于30
	count, err := cartDao.WithContext(ctx).Where(cartDao.UserName.Eq(username)).Count()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "加入购物车失败: " + err.Error(),
		})
		return
	}

	if count > 30 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "购物车最多容纳30个",
		})
		return
	}

	err = cartDao.WithContext(ctx).Create(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "加入购物车失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "加入购物车成功",
	})
}

// GetUserCart 获取用户购物车
func GetUserCart(ctx *gin.Context) {
	// 获取用户名
	username, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户名失败: " + err.Error(),
		})
	}

	// 获取购物车
	cartDao := database.GoodsCard
	carts, err := cartDao.WithContext(ctx).Where(cartDao.UserName.Eq(username)).Find()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取购物车失败: " + err.Error(),
		})
	}
	fileDao := database.File

	// 获取购物车中商品信息
	var product []models.GoodsCardDo
	for _, cart := range carts {
		productInfo, err := variable.ContractSession.GetProductDetail(big.NewInt(cart.GoodsId))
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "获取商品信息失败: " + err.Error(),
			})
		}

		// 获取文件信息
		find, err := fileDao.WithContext(ctx).Where(fileDao.Id.Eq(productInfo.ImageId)).First()
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "获取商品图片失败: " + err.Error(),
			})
			return
		}

		product = append(product, models.GoodsCardDo{
			Id:      int64(cart.ID),
			GoodsId: productInfo.Id.Int64(),
			Name:    productInfo.Name,
			Desc:    productInfo.Description,
			Price:   productInfo.Price.Int64(),
			Store:   cart.Number,
			Uint:    productInfo.UintGoods,
			Origin:  productInfo.Origin,
			Time:    cart.CreatedAt.Format("2006-01-02 15:04:05"),
			Images:  find.FileData,
			Check:   false,
		})
	}

	// 获取商品信息

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取购物车成功",
		"carts":   product,
	})
}

// RemoveFromCart 移除购物车
func RemoveFromCart(ctx *gin.Context) {
	// 获取用户名
	username, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户名失败: " + err.Error(),
		})
	}

	// 获取请求参数
	var requestData models.GoodsCard
	err = ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
	}

	// 移除购物车
	cartDao := database.GoodsCard
	_, err = cartDao.WithContext(ctx).Where(cartDao.ID.Eq(uint(requestData.GoodsId)), cartDao.UserName.Eq(username)).Delete()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "移除购物车失败: " + err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "移除购物车成功",
	})
}

// BuyFromCart 购买购物车
func BuyFromCart(ctx *gin.Context) {
	// 获取用户名
	username, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户名失败: " + err.Error(),
		})
	}

	// 获取请求参数
	var requestData models.BuyFromCartRequest
	err = ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
	}

	// 获取购物车
	cartDao := database.GoodsCard
	carts, err := cartDao.WithContext(ctx).Where(cartDao.UserName.Eq(username), cartDao.ID.In(requestData.Ids...)).Find()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取购物车失败: " + err.Error(),
		})
	}

	// 购买购物车中商品
	for _, cart := range carts {
		_, _, err = variable.ContractSession.BuyProduct(big.NewInt(cart.GoodsId), big.NewInt(cart.Number), username, requestData.Address, time.Now().Format("2006-01-02 15:04:05"))
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "购买商品失败: " + err.Error(),
			})
		}
	}

	// 清空购物车
	_, err = cartDao.WithContext(ctx).Where(cartDao.UserName.Eq(username), cartDao.ID.In(requestData.Ids...)).Delete()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "清空购物车失败: " + err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "购买购物车成功",
	})
}

// GetAllUserInfo 获取所有用户信息
func GetAllUserInfo(ctx *gin.Context) {
	// 获取用户ID
	username, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户姓名失败: " + err.Error(),
		})
		return
	}

	// 获取用户��息
	userDao := database.UserInfo
	userInfo, err := userDao.WithContext(ctx).Where(userDao.Username.Eq(username)).First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败: " + err.Error(),
		})
		return
	}

	if userInfo.Identity != "管理员" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户无权查看用户信息",
		})
		return
	}

	// 获取页面
	pageString := ctx.Query("page")
	page, err := strconv.ParseInt(pageString, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	// 获取用户信息
	users, err := userDao.WithContext(ctx).Order(userDao.CreatedAt.Desc()).Find()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败: " + err.Error(),
		})
		return
	}

	if len(users) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "获取用户信息成功",
			"users":   []models.UserInfo{},
		})
		return
	}

	// 假设page为当前页码，pageSize为每页大小（本例中为10）
	pageSize := int64(6)

	// 计算总页数
	totalPages := (int64(len(users)) + pageSize - 1) / pageSize

	// 确保请求的页码在有效范围内
	if page < 1 {
		page = 1
	}
	if page > int64(totalPages) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "获取用户信息成功",
			"users":   []models.UserInfo{},
		})
		return
	}

	// 计算当前页的起始和结束索引
	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize
	if endIndex > int64(len(users)) {
		endIndex = int64(len(users))
	}

	// 获取当前页的产品切片
	var userInfos []*models.UserInfo
	if startIndex < int64(len(users)) {
		userInfos = users[startIndex:endIndex]
	} else {
		// 如果startIndex超出了products的长度，返回一个空切片
		userInfos = []*models.UserInfo{}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取用户信息成功",
		"users":   userInfos,
	})
}

// UpdateUserStatus 修改用户状态
func UpdateUserStatus(ctx *gin.Context) {
	// 获取用户ID
	username, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户姓名失败: " + err.Error(),
		})
		return
	}

	// 获取请求参数
	var requestData models.UpdateUserStatusRequest
	err = ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	// 获取用户信息
	userDao := database.UserInfo
	userInfo, err := userDao.WithContext(ctx).Where(userDao.Username.Eq(username)).First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败: " + err.Error(),
		})
		return
	}

	if userInfo.Identity != "管理员" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户无权修改用户信息",
		})
		return
	}

	// 修改用户状态
	_, err = userDao.WithContext(ctx).Unscoped().
		Where(userDao.Username.Eq(requestData.UserName)).
		Update(userDao.Disabled, requestData.Status)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "修改用户状态失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "修改用户状态成功",
	})
}

// GetFruitKind 获取水果种类
func GetFruitKind(ctx *gin.Context) {
	// 获取用户ID
	username, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户姓名失败: " + err.Error(),
		})
		return
	}

	// 获取用户信息
	userDao := database.UserInfo
	userInfo, err := userDao.WithContext(ctx).Where(userDao.Username.Eq(username)).First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败: " + err.Error(),
		})
		return
	}

	if userInfo.Identity != "管理员" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户无权修改用户信息",
		})
		return
	}

	// 获取页面
	pageString := ctx.Query("page")
	page, err := strconv.ParseInt(pageString, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	// 获取水果种类
	fruitDao := database.FruitKind
	fruits, err := fruitDao.WithContext(ctx).Order(fruitDao.CreatedAt.Desc()).Find()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取水果种类失败: " + err.Error(),
		})
		return
	}

	if len(fruits) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "获取水��种类成功",
			"users":   []models.FruitKind{},
		})
		return
	}

	// 假设page为当前页码，pageSize为每页大小（本例中为10）
	pageSize := int64(6)

	// 计算总页数
	totalPages := (int64(len(fruits)) + pageSize - 1) / pageSize

	// 确保请求的页码在有效范围内
	if page < 1 {
		page = 1
	}
	if page > int64(totalPages) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "获取水果种类成功",
			"users":   []models.FruitKind{},
		})
		return
	}

	// 计算当前页的起始和结束索引
	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize
	if endIndex > int64(len(fruits)) {
		endIndex = int64(len(fruits))
	}

	// 获取当前页的产品切片
	var fruitKind []*models.FruitKind
	if startIndex < int64(len(fruits)) {
		fruitKind = fruits[startIndex:endIndex]
	} else {
		// 如果startIndex超出了products的长度，返回一个空切片
		fruitKind = []*models.FruitKind{}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取水果种类成功",
		"fruits":  fruitKind,
	})
}

// AddFruitKind 添加水果种类
func AddFruitKind(ctx *gin.Context) {
	// 获取用户ID
	username, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户姓名失败: " + err.Error(),
		})
		return
	}

	// 获取用户信息
	userDao := database.UserInfo
	userInfo, err := userDao.WithContext(ctx).Where(userDao.Username.Eq(username)).First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败: " + err.Error(),
		})
		return
	}

	if userInfo.Identity != "管理员" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户无权修改用户��息",
		})
		return
	}

	// 获取请求参数
	var requestData models.FruitKind
	err = ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	// 检查种类名是否存在
	fruitDao := database.FruitKind
	count, err := fruitDao.WithContext(ctx).Where(fruitDao.Kind.Eq(requestData.Kind)).Count()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "查询水果种类失败: " + err.Error(),
			})
			return
		}
	}

	if count != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "水果种类已存在",
		})
		return
	}

	// 添加水果种类
	err = fruitDao.WithContext(ctx).Create(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "添加水果种类失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "添加水果种类成功",
	})
}

// RemoveFruitKind 删除水果种类
func RemoveFruitKind(ctx *gin.Context) {
	// 获取用户ID
	username, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户姓名失败: " + err.Error(),
		})
		return
	}

	// 获取用户信息
	userDao := database.UserInfo
	userInfo, err := userDao.WithContext(ctx).Where(userDao.Username.Eq(username)).First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败: " + err.Error(),
		})
		return
	}

	if userInfo.Identity != "管理员" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户无权修改用户信息",
		})
		return
	}

	// 获取请求参数
	var requestData models.FruitKind
	err = ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	// 删除水果种类
	fruitDao := database.FruitKind
	_, err = fruitDao.WithContext(ctx).Where(fruitDao.ID.Eq(requestData.ID)).Delete()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "删除水果种类失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除水果种类成功",
	})
}

// GetAllFruitKind 获取全部水果种类
func GetAllFruitKind(ctx *gin.Context) {
	// 获取水果种类
	fruitDao := database.FruitKind
	fruits, err := fruitDao.WithContext(ctx).Find()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取水果种类失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取水果种类成功",
		"fruits":  fruits,
	})
}

// GetAllOrder 获取所有订单
func GetAllOrder(ctx *gin.Context) {
	// 获取用户ID
	username, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户姓名失败: " + err.Error(),
		})
		return
	}

	// 获取用户信息
	userDao := database.UserInfo
	userInfo, err := userDao.WithContext(ctx).Where(userDao.Username.Eq(username)).First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败: " + err.Error(),
		})
		return
	}

	if userInfo.Identity != "管理员" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户无权查看订单信息",
		})
		return
	}

	// 获取页面
	pageString := ctx.Query("page")
	page, err := strconv.ParseInt(pageString, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	// 获取用户下一订单id
	id, err := variable.ContractSession.NextOrderId()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	fileDao := database.File

	var productIds []models.GetUserOrdersResponse
	for i := 6 * (page - 1); i < id.Int64(); i++ {
		element, err := variable.ContractSession.TotalOrders(big.NewInt(int64(i)))
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "获取订单信息失败: " + err.Error(),
			})
			return
		}

		product, err := variable.ContractSession.GetProductDetail(element.ProductId)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "获取商品详情失败: " + err.Error(),
			})
			return
		}

		find, err := fileDao.WithContext(ctx).Where(fileDao.Id.Eq(product.ImageId)).First()
		if err != nil {
			continue
		}
		if find == nil {
			find.Id = ""
		}

		var a int = 0
		if element.Status {
			a = 1
		}
		productIds = append(productIds, models.GetUserOrdersResponse{
			Id:          element.Id.Int64(),
			Number:      element.Quantity.Int64(),
			Name:        product.Name,
			Senders:     product.Seller,
			BlockNumber: element.BlockNumber.Int64(),
			Image:       find.FileData,
			Price:       product.Price.Int64(),
			TotalPrice:  element.Quantity.Int64() * product.Price.Int64(),
			CreateTime:  element.CreatedTime,
			Address:     element.UserAddress,
			Status:      int64(a),
		})

		if len(productIds) >= 6 {
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取订单信息成功",
		"orders":  productIds,
	})
}

// CheckOrder 审核订单
func CheckOrder(ctx *gin.Context) {
	// 获取用户ID
	username, err := utils.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户姓名失败: " + err.Error(),
		})
		return
	}

	// 获取用户信息
	userDao := database.UserInfo
	userInfo, err := userDao.WithContext(ctx).Where(userDao.Username.Eq(username)).First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败: " + err.Error(),
		})
		return
	}

	if userInfo.Identity != "管理员" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户无权查看订单信息",
		})
		return
	}

	// 获取请求参数
	var requestData models.CheckOrderRequest
	err = ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}
	
	fmt.Println(requestData.Id)

	_, receipt, err := variable.ContractSession.SetOrderStatus(big.NewInt(requestData.Id), true)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "审核订单失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":            200,
		"message":         "审核订单成功",
		"transactionHash": receipt.TransactionHash,
	})
}
