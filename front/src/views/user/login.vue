<template>
  <div class="app">
    <div class="relative flex items-center min-h-screen px-4 py-12 sm:px-6 lg:px-8">
      <div class="mx-auto w-full space-y-4 max-w-lg"
           style="background-color:#fff; border:1px solid #fff; border-radius:10px;">
        <div class="text-center mb-8">
          <img src="../../assets/logo.png" alt="FreshTrace" class="mx-auto w-48 mb-4">
          <h1 class="text-2xl font-bold text-gray-800">欢迎来到 FreshTrace</h1>
          <p class="text-gray-600 mt-2">您的水果溯源专家</p>
        </div>
        <div class="p-4 grid gap-2 rounded-lg border border-gray-200 dark:border-gray-800">
          <form class="space-y-2">
            <p class="text-lg font-semibold text-center text-gray-700">登录</p>
            <div class="grid gap-2">
              <div class="grid gap-2"><label
                  class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                  for="username">用户名</label><input
                  class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                  id="username" placeholder="输入你的用户名" v-model="username"></div>
              <div class="grid gap-2"><label
                  class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                  for="password">密码</label><input
                  class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                  id="password" placeholder="输入你的密码" type="password" v-model="password"></div>
            </div>
            <button
                class="inline-flex items-center justify-center bg-[#f8b811] text-white rounded-md text-sm font-medium h-10 px-4 py-2 w-full hover:opacity-90 transition-opacity"
                type="button" @click="login">
              登录
            </button>
          </form>
        </div>
        <div class="space-y-2 text-center"><a class="font-semibold underline" style="cursor: pointer;"
                                              @click="gotoRegister">
          创建账号
        </a>&nbsp <a class="font-semibold underline" style="cursor: pointer;" @click="gotoRetrieve">
          找回账号
        </a></div>
      </div>
      <div class="absolute inset-0 flex items-center justify-center pointer-events-none">
        <div
            class="grid w-full max-w-2xl gap-1.5 text-center text-sm text-gray-900 divide-y divide-gray-200 dark:text-gray-100 dark:divide-gray-800">
          <div class="flex items-center justify-center h-12"><span class="animate-pulse"></span></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "login",
  data() {
    return {
      username: '',
      password: ''
    }
  },
  methods: {
    gotoRegister() {
      this.$router.push('/register');
    },
    gotoRetrieve() {
      this.$router.push('/retrieve');
    },
    login() {
      if (this.username.length < 6 || this.password.length < 6) {
        this.$message.error("请输入正确的用户名或密码")
        return
      }

      this.$axios.post("/user/login", {
        username: this.username,
        password: this.password
      }).then((res) => {
        if (res.data.code === 200) {
          localStorage.setItem("auth_token", res.data.token)
          localStorage.setItem("username", res.data.username)
          localStorage.setItem("identity", res.data.identity)
          this.$message.success(res.data.message)
          // 在这里做判断
          if (res.data.identity !== "管理员") {
            setTimeout(() => {
              this.$router.push('/home')
            }, 1000); // 1秒后执行
          } else {
            setTimeout(() => {
              this.$router.push('/main')
            }, 1000); // 1秒后执行
          }
          // 跳转到首页
        } else {
          this.$message.error(res.data.message)
        }
        console.log(res);
      }).catch((res) => {
        console.log(res);
      })
    }
  }
}
</script>
<style scoped>
.app {
  /* 设置背景图片 */
  background-image: url('../../assets/bg-2.jpg');
  /* 确保背景图片覆盖整个屏幕 */
  background-position: center;
  background-size: cover;
  background-repeat: no-repeat;
  /* 设置背景图片填充整个屏幕 */
  height: 100vh;
  width: 100%;
}

/* 添加输入框焦点样式 */
.input-field:focus {
  border-color: #f8b811;
  box-shadow: 0 0 0 2px rgba(248, 184, 17, 0.1);
  outline: none;
}

/* 链接悬停样式 */
.underline:hover {
  color: #f8b811;
}
</style>


