<template>
  <header class="flex items-center h-14 px-4 border-b lg:h-20 xl:px-6">
    <!-- Logo部分 -->
    <div class="logo flex items-center mr-4 md:mr-6 lg:mr-8">
      <a href="#" title="水果溯源系统" class="flex items-center gap-3">
        <img :src="logo" alt="水果溯源系统" class="h-12 md:h-14 lg:h-16" />
        <span class="logo-text hidden sm:block">FreshTrace</span>
      </a>
    </div>
    
    <!-- 导航链接部分 -->
    <div class="nav-outer flex-grow">
      <nav class="main-menu hidden md:flex items-center justify-center">
        <div class="navbar-collapse show collapse clearfix">
          <ul class="navigation clearfix flex items-center space-x-8">
            <li>  
              <RouterLink to="/home" 
                        class="text-gray-600 hover:text-[#f8b811] transition-colors" 
                        style="font-size: 25px;font-weight: 500;"
                        :class="{ 'nav-active': $route.path === '/home' }">
                首页
              </RouterLink>
            </li>
            <li>
              <RouterLink to="/fruit" 
                        class="text-gray-600 hover:text-[#f8b811] transition-colors" 
                        style="font-size: 25px;font-weight: 500;"
                        :class="{ 'nav-active': $route.path.includes('/fruit') || $route.path.includes('/buy-detail') }">
                水果
              </RouterLink>
            </li>
            <li>
              <RouterLink to="/order" 
                        class="text-gray-600 hover:text-[#f8b811] transition-colors" 
                        style="font-size: 25px;font-weight: 500;"
                        :class="{ 'nav-active': $route.path === '/order' }">
                订单
              </RouterLink>
            </li>
            <li>
              <RouterLink to="/cart" 
                        class="text-gray-600 hover:text-[#f8b811] transition-colors" 
                        style="font-size: 25px;font-weight: 500;"
                        :class="{ 'nav-active': $route.path === '/cart' }">
                购物车
              </RouterLink>
            </li>
            <li>
              <RouterLink to="/about" 
                        class="text-gray-600 hover:text-[#f8b811] transition-colors" 
                        style="font-size: 25px;font-weight: 500;"
                        :class="{ 'nav-active': $route.path === '/about' }">
                关于我们
              </RouterLink>
            </li>
            <li v-show="identity == '供应商'">
              <RouterLink to="/manage" 
                        class="text-gray-600 hover:text-[#f8b811] transition-colors" 
                        style="font-size: 25px;font-weight: 500;"
                        :class="{ 'nav-active': $route.path === '/manage' }">
                商品管理
              </RouterLink>
            </li>
          </ul>
        </div>
      </nav>

      <!-- 移动端菜单按钮 -->
      <div class="mobile-nav-toggler block md:hidden">
        <span class="icon lnr-icon-bars"></span>
      </div>
    </div>

    <!-- 用户菜单部分 -->
    <div class="dropdown dropdown-end ml-auto">
      <div tabindex="0" role="button" class="btn btn-ghost rounded-btn" v-text="username"></div>
      <ul tabindex="0" class="menu dropdown-content z-[1] p-2 shadow bg-base-100 rounded-box w-52 mt-4">
        <li @click="openModel"><a>修改用户信息</a></li>
        <li @click="logout"><a>退出登录</a></li>
      </ul>
    </div>
  </header>

  <!-- 修改用户信息模态框 -->
  <div class="modal" role="dialog" id="my_modal_8">
    <div class="modal-box">
      <h3 class="font-bold text-lg">修改用户信息</h3>
      <div class="grid gap-2" style="margin-top:20px">
        <div class="grid gap-2">
          <label class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70" for="username">用户名</label>
          <input class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
            id="username" placeholder="输入你的用户名" style="margin-bottom: 10px;" v-model="localUsername" disabled/>
        </div>
        <div class="grid gap-2">
          <label class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70" for="password">密码</label>
          <input class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
            id="password" placeholder="输入你的密码" type="password" v-model="password"/>
        </div>
        <div class="grid gap-2">
          <label class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70" for="security-question">问题</label>
          <input class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
            id="security-question" v-model="question" style="margin-bottom: 10px;" placeholder="请输入密保"/>
        </div>
        <div class="grid gap-2">
          <label class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70" for="answer">答案</label>
          <input class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
            id="answer" placeholder="请输入答案" style="margin-bottom: 10px;" v-model="answer"/>
        </div>
      </div>
      <div class="modal-action">
        <a class="btn" @click="updateUserInfo">确定</a>
      </div>
    </div>
  </div>
</template>

<script>
import logo from '../assets/logo.png'

export default {
  name: 'NavHeader',
  data() {
    return {
      logo,
      localUsername: '',
      password: '',
      question: '',
      answer: ''
    }
  },
  props: {
    username: {
      type: String,
      required: true
    },
    identity: {
      type: String,
      required: true
    }
  },
  watch: {
    username: {
      immediate: true,
      handler(newVal) {
        this.localUsername = newVal
      }
    }
  },
  methods: {
    openModel() {
      window.location.href = '#my_modal_8';
    },
    updateUserInfo() {
      this.$axios.post('/user/update_user_info', {
        username: this.localUsername,
        password: this.password,
        question: this.question,
        answer: this.answer
      }).then((res) => {
        if (res.data.code === 200) {
          this.$message.success('修改成功')
        } else {
          this.$message.error(res.data.message)
        }
        console.log(res);
      }).catch((res) => {
        this.$message.error('修改失败')
        console.log(res);
      })
      window.location.href = '#';
      this.password = ''
      this.question = ''
      this.answer = ''
    },
    logout() {
      localStorage.removeItem("auth_token")
      localStorage.removeItem("username")
      localStorage.removeItem("identity")
      this.$message.success('退出登录成功')
      setTimeout(() => {
        this.$router.push('/')
      }, 1000)
    }
  }
}
</script>

<style scoped>
.sticky-header {
  position: sticky;
  top: 0;
  z-index: 100;
  background: white;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.logo {
  min-width: 120px;
  @media (min-width: 640px) {
    min-width: 160px;
  }
  @media (min-width: 1024px) {
    min-width: 200px;
  }
}

.logo-text {
  font-size: clamp(20px, 2vw, 28px);
  font-weight: bold;
  color: #f8b811;
  font-family: 'Arial', sans-serif;
  white-space: nowrap;
}

/* 添加移动端菜单样式 */
@media (max-width: 768px) {
  .navigation {
    display: none;
  }
  
  .mobile-nav-toggler {
    display: block;
  }
}

@media (min-width: 769px) {
  .mobile-nav-toggler {
    display: none;
  }
}

.nav-active {
  border-bottom: 8px solid #f8b811;
}
</style> 