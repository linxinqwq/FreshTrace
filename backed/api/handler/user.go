package handler

import (
	"backed/app/core/consts"
	model "backed/app/models"
	"backed/app/utils"
	"backed/internal/database"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// UserRegister 用户注册
func UserRegister(ctx *gin.Context) {
	// 获取请求参数
	var requestData model.UserInfo
	err := ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	// 检查请求参数是否合理
	if len(requestData.Username) < 6 || len(requestData.Password) < 6 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户输入参数不合理",
		})
		return
	}

	userDao := database.UserInfo

	// 查看用户名是否被占用
	findUserInfo, err := userDao.WithContext(ctx).
		Where(
			userDao.Username.Eq(requestData.Username),
			userDao.Disabled.Is(false),
		).
		First()
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": "获取用户名是否存在失败: " + err.Error(),
			})
			return
		}
	}

	if findUserInfo != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户名已存在",
		})
		return
	}

	// 创建用户
	if requestData.Identity != "用户" {
		requestData.Disabled = true
	}
	err = userDao.WithContext(ctx).Create(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "创建用户失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建用户成功",
	})
}

// UserLogin 用户登录
func UserLogin(ctx *gin.Context) {
	// 获取请求参数
	var requestData model.UserInfo
	err := ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	// 检查请求参数是否合理
	if len(requestData.Username) < 6 || len(requestData.Password) < 6 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户输入参数不合理",
		})
		return
	}

	userDao := database.UserInfo

	// 获取用户信息
	findUserInfo, err := userDao.WithContext(ctx).
		Where(
			userDao.Username.Eq(requestData.Username),
			userDao.Disabled.Is(false)).
		First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败: " + err.Error(),
		})
		return
	}

	if findUserInfo.Password != requestData.Password {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "密码错误",
		})
		return
	}

	// 生成token
	token, err := utils.GenToken(findUserInfo.Username, consts.TokenExpireDuration, []byte(consts.TokenSecret))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "生成token失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":     200,
		"message":  "登录成功",
		"token":    token,
		"username": findUserInfo.Username,
		"identity": findUserInfo.Identity,
	})
}

// UserRetrieve 用户找回
func UserRetrieve(ctx *gin.Context) {
	// 获取请求参数
	var requestData model.UserInfo
	err := ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	// 检查请求参数是否合理
	if len(requestData.Username) < 6 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户输入参数不合理",
		})
		return
	}

	userDao := database.UserInfo

	// 获取用户信息
	findUserInfo, err := userDao.WithContext(ctx).
		Where(
			userDao.Username.Eq(requestData.Username),
			userDao.Disabled.Is(false)).
		First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":     200,
		"message":  "获取密码成功",
		"question": findUserInfo.Question,
	})
}

// UserRetrieveResult 用户重置结果
func UserRetrieveResult(ctx *gin.Context) {
	// 获取请求参数
	var requestData model.UserInfo
	err := ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	// 检查请求参数是否合理
	if len(requestData.Username) < 6 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户输入参数不合理",
		})
		return
	}

	userDao := database.UserInfo

	// 获取用户信息
	findUserInfo, err := userDao.WithContext(ctx).
		Where(
			userDao.Username.Eq(requestData.Username),
			userDao.Disabled.Is(false)).
		First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败: " + err.Error(),
		})
		return
	}

	if findUserInfo.Question == requestData.Question && findUserInfo.Answer == requestData.Answer {
		ctx.JSON(http.StatusOK, gin.H{
			"code":     200,
			"message":  "获取密码成功",
			"password": findUserInfo.Password,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "密保错误",
		})
		return
	}
}

// UserUpdateInfo 更新用户信息
func UserUpdateInfo(ctx *gin.Context) {
	// 获取请求参数
	var requestData model.UserInfo
	err := ctx.ShouldBindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "绑定请求参数失败: " + err.Error(),
		})
		return
	}

	// 检查请求参数是否合理
	if len(requestData.Username) < 6 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "用户输入参数不合理",
		})
		return
	}

	userDao := database.UserInfo

	// 获取用户信息
	findUserInfo, err := userDao.WithContext(ctx).
		Where(
			userDao.Username.Eq(requestData.Username),
			userDao.Disabled.Is(false)).
		First()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "获取用户信息失败: " + err.Error(),
		})
		return
	}

	if len(requestData.Password) != 0 {
		findUserInfo.Password = requestData.Password
	}

	if len(requestData.Question) != 0 {
		findUserInfo.Question = requestData.Question
		findUserInfo.Answer = requestData.Answer
	}

	err = userDao.WithContext(ctx).Save(findUserInfo)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "更新用户信息失败: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新用户信息成功",
	})
}
