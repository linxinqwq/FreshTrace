<template>
  <div class="fruit-page">
    <NavHeader 
      :username="username"
      :identity="identity"
      @open-model="openModel"
      @logout="logout"
    />
    <div class="modal" role="dialog" id="my_modal_8">
      <div class="modal-box">
        <h3 class="font-bold text-lg">修改用户信息</h3>
        <div class="grid gap-2" style="margin-top:20px">
          <div class="grid gap-2"><label
              class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
              for="username">用户名</label><input
              class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
              id="username" placeholder="输入你的用户名" style="margin-bottom: 10px;" v-model="username" disabled>
          </div>
          <div class="grid gap-2"><label
              class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
              for="password">密码</label><input
              class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
              id="password" placeholder="输入你的密码" type="password" v-model="password"></div>
          <div class="grid gap-2"><label
              class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
              for="username">问题</label><input
              class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
              id="username" v-model="question" style="margin-bottom: 10px;" placeholder="请输入密保"></div>
          <div class="grid gap-2"><label
              class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
              for="username">答案</label><input
              class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
              id="username" placeholder="请输入答案" style="margin-bottom: 10px;" v-model="answer"></div>
        </div>
        <div class="modal-action">
          <a class="btn" @click="updateUserInfo">确定</a>
        </div>
      </div>
    </div>
<!-- 添加 header section -->
<div class="header-section" style="background: linear-gradient(to right, rgba(248, 184, 17, 0.8), rgba(248, 184, 17, 0.4)); height: 200px; margin-bottom: 30px;">
  <div style="height: 100%; display: flex; flex-direction: column; justify-content: center; padding: 0 40px;">
    <h1 style="font-size: 36px; font-weight: bold; margin-bottom: 10px; color: #333;">水果商城</h1>
    <p style="color: #666; font-size: 16px;">品质保证，溯源可查，新鲜直达</p>
  </div>
</div>

    <div class="app">
      <div class="flex flex-col h-screen">
        <div class="mytab">
          <a v-for="(tab, index) in kind" :key="index" class="tab-item"
             :class="{ 'tab-active': index === activeTabIndex }" @click="setActiveTab(index)">
            
            {{ tab.kind }}
            <!-- <span v-if="index === activeTabIndex" class="tab-indicator"></span> -->
          </a>
        </div>
        <main class="flex-1 p-4 grid grid-cols-4 gap-4 items-start">
          <div class="product-card" data-v0-t="card"
               v-for="(product, index) in products" :key="index" @click="navigateToProductDetail(product.Id)">
            <div class="aspect-square overflow-hidden rounded-t-lg flex items-center justify-center">
              <img :src="product.ImageId" width="200" height="200" alt="Cover image"
                   class="aspect-square object-cover"/>
            </div>
            <div class="p-4 flex flex-col gap-2">
              <h3 class="font-semibold">商品名称: {{ product.Name }}</h3>
              <p class="text-lg font-bold tracking-tighter" style="color: #FF6E2F">
                商品价格：
                <span style="font-size: 24px;">￥{{ product.Price }}</span>
              </p>
            </div>
            <button class="add-to-cart">查看详情</button>
          </div>
        </main>
        <div class="join" style="margin-left:45%;margin-right:45%;padding-bottom:20px">
          <button class="join-item btn" @click="prevPage">«</button>
          <button class="join-item btn">Page {{ page }}</button>
          <button class="join-item btn" @click="nextPage">»</button>
        </div>
        <Footer />
      </div>
    </div>
  </div>
</template>

<script>
import NavHeader from '@/components/NavHeader.vue'
import Footer from '@/components/Footer.vue'

export default {
  name: 'fruit',
  components: {
    NavHeader,
    Footer
  },
  data() {
    return {
      username: '',
      password: '',
      question: '',
      answer: '',
      identity: '',
      page: 1,
      latestId: 0,
      products: [],
      showPopup: false,
      kind: [],
      activeTabIndex: 0 // 默认选中第一个选项卡
    }
  },
  created() {
    this.username = localStorage.getItem('username')
    this.identity = localStorage.getItem('identity')
    this.$axios.get("/main/get_all_fruit_kind").then((res) => {
      if (res.data.code === 200) {
        res.data.fruits.unshift({"kind": "全部"});
        this.kind = res.data.fruits;
      } else {
        this.$message.error("获取商品种类失败: " + res.data.message)
      }
      console.log(res)
    }).catch((res) => {
      console.log(res)
    })

    this.$axios.get("/main/get_goods_list", {
      params: {
        lastId: this.latestId,
        kind: "全部"
      }
    }).then((res) => {
      if (res.data.code == 200) {
        this.products = res.data.products
        this.latestId = this.products[this.products.length - 1].Id
      } else {
        this.$message.error('获取商品列表失败: ' + res.data.message)
      }
      console.log(res);
    }).catch((res) => {
      console.log(res);
    })
  },
  methods: {
    navigateToProductDetail(productId) {
      this.$router.push({path: `/buy-detail`, query: {id: productId}});
    },
    prevPage() {
      if (this.page > 1) {
        this.page -= 1;
        this.latestId = this.products[0].Id - 9
        // 获取当前种类
        this.$axios.get("/main/get_goods_list", {
          params: {
            lastId: this.latestId,
            kind: this.kind[this.activeTabIndex].kind
          }
        }).then((res) => {
          if (res.data.code == 200) {
            this.products = res.data.products
            this.latestId = this.products[this.products.length - 1].Id
          } else {
            this.$message.error("获取商品列表失败: " + res.data.message)
          }
          console.log(res);
        }).catch((res) => {
          console.log(res);
        })
      } else {
        this.$message.error('已经是第一页了')
        return
      }
    },
    nextPage() {
      // 这里您可能需要一个逻辑来检查是否已经是最后一页
      if (this.products.length == 8) {
        this.page += 1;
      } else {
        this.$message.error('已经是最后一页了')
        return
      }
      this.$axios.get("/main/get_goods_list", {
        params: {
          lastId: this.latestId,
          kind: this.kind[this.activeTabIndex].kind
        }
      }).then((res) => {
        if (res.data.code == 200) {
          this.products = res.data.products
          this.latestId = this.products[this.products.length - 1].Id
        } else {
          this.$message.error("获取商品列表失败: " + res.data.message)
        }
        console.log(res);
      }).catch((res) => {
        console.log(res);
      })
    },
    // 退出登录
    logout() {
      localStorage.removeItem("auth_token")
      localStorage.removeItem("username")
      localStorage.removeItem("identity")
      this.$message.success('退出登录成功')
      setTimeout(() => {
        this.$router.push('/')
      }, 1000); // 1秒后执行
      // 跳转到首页
    },
    openModel() {
      window.location.href = '#my_modal_8';
    },
    updateUserInfo() {
      this.$axios.post('/user/update_user_info', {
        username: this.username,
        password: this.password,
        question: this.question,
        answer: this.answer
      }).then((res) => {
        console.log(res);
      }).catch((res) => {
        console.log(res);
      })
      window.location.href = '#';
      this.password = ''
      this.question = ''
      this.answer = ''
    },
    setActiveTab(index) {
      this.activeTabIndex = index;
      this.page = 1;
      this.latestId = 0;
      this.products = [];
      this.$axios.get("/main/get_goods_list", {
        params: {
          lastId: this.latestId,
          kind: this.kind[this.activeTabIndex].kind
        }
      }).then((res) => {
        if (res.data.code == 200) {
          this.products = res.data.products
          this.latestId = this.products[this.products.length - 1].Id
        } else {
          this.$message.error("获取商品列表失败: " + res.data.message)
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
.mytab{
  margin-top: 5px;
  display: flex;
  justify-content: center; /* 使���项卡列表居中 */
  border: none; /* 消去边框 */
}

.tab-item {
  position: relative; /* 使下划线定位相对于选项卡 */
  text-align: center; /* 文本居中显示 */
  white-space: nowrap; /* 文本不换行 */
  overflow: hidden; /* 超出部分隐藏 */
  text-overflow: ellipsis; /* 超出部分显示省略号 */
  min-width: 50px; /* 设置最小宽度为50px */
  min-height: 30px;
  flex-grow: 1; /* 允许在需要时拉伸 */
  background-color: #fff; /* 默认背景颜色为fff */
  color: black; /* 默认字体颜色为black */
  margin: 0 10px; /* 在每个选项卡之间设定间距 */
  border: none; /* 消去选项卡边框 */
  transition: color 0.3s; /* 添加过渡效果 */
}

.tab-item:hover,
.tab-active {
  font-weight: bold;
  color: #fff;
  background-color: rgba(248, 184, 17, 0.8); /* #f8b811 with 0.8 opacity */
}
.tab-item::after {
  content: '';
  position: absolute;
  left: 50%;
  bottom: -5px;
  width: 0;
  height: 2px;
  background-color: #f8b811;
  transition: width 0.3s, left 0.3s;
}

.tab-item:hover::after,
.tab-active::after {
  width: 100%; /* 鼠标悬浮或选中时下划线宽度为100% */
  left: 0; /* 使下划线从左侧开始 */
}
.tab-indicator {
  position: absolute; /* 绝对定位 */
  bottom: -5px; /* 距离底部5px */
  left: 50%; /* 水平居中 */
  transform: translateX(-50%); /* 使下划线居中 */
  width: 50%; /* 下划线宽度 */
  height: 2px; /* 下划线高度 */
  background-color: #80BD44; /* 下划线颜色 */
  display: block;
  transition: width 0.3s; /* 添加过渡效果 */
}
.product-card {
  background-color: rgba(248, 184, 17, 0.1); /* 淡黄色背景 */
  position: relative;
  width: auto;
  margin: 8px;
  transition: transform 0.3s, box-shadow 0.3s;
  overflow: hidden;
}

.product-card:hover {
  transform: scale(1.05);
  box-shadow: 0 4px 20px rgba(248, 184, 17, 0.2); /* 阴影也改为淡黄色 */
}

/* 默认样式 */
.product-image {
  width: 100%; /* 图片宽度为100% */
  height: auto; /* 高度自适应 */
  transition: transform 0.3s; /* 添加过渡效果 */
}

.add-to-cart {
  display: none;
  position: absolute;
  bottom: 10px;
  left: 50%;
  transform: translateX(-50%);
  background-color: #f8b811; /* 按钮改为黄色 */
  color: #fff;
  padding: 10px 20px;
  border: none;
  border-radius: 15px;
  cursor: pointer;
  transition: display 0.3s;
}

.add-to-cart:hover {
  opacity: 0.9;
}

.product-card:hover .add-to-cart {
  display: block; /* 鼠标悬浮时显示按钮 */
}

.product-card:hover .product-image {
  transform: translateY(-30px); /* 鼠标悬浮时向上移动图片 */
}

.product-card:hover .p-4 {
  transform: translateY(-50px); /* 鼠标悬浮时向上移动文字 */
  transition: transform 0.3s; /* 添加过渡效果 */
}

/* 修改分页按钮样式 */
.join .btn {
  background-color: #FF6E2F;
  color: white;
  border: none;
}

.join .btn:hover {
  opacity: 0.9;
}
</style>
