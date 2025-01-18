<template>
  <div class="scrollable-container" style="overflow: visible; position: relative;">
    <div class="modal" role="dialog" id="lookProcess">
      <div class="modal-box">
        <div class="modal-header">
          <h3 class="font-bold text-lg" style="color: #333;">商品溯源步骤</h3>
          <a href="#" class="close-btn">&times;</a>
        </div>
        
        <div class="process" style="margin-top: 20px;margin-bottom: 20px">
          <div class="steps-container">
            <div class="step-item">
              <div class="step-circle">1</div>
              <div class="step-content">
                <h4>商品上链</h4>
                <p class="step-date">初始步骤</p>
              </div>
            </div>
            <template v-for="(step, index) in product.Process" :key="index">
              <div class="step-connector"></div>
              <div class="step-item" @click="openPopup(index)">
                <div class="step-circle">{{ index + 2 }}</div>
                <div class="step-content">
                  <h4>{{ step.Operation }}</h4>
                  <p class="step-date">{{ step.Date }}</p>
                </div>
              </div>
            </template>
          </div>
        </div>
        <div class="modal-action">
          <a href="#" class="btn" style="background-color: #f8b811; color: white;">取消</a>
          <a href="#" class="btn" style="background-color: #f8b811; color: white;" @click="updateProcess">确定</a>
        </div>
      </div>
    </div>
  </div>
  <div class="app">
    <div v-if="showPopup" class="popup-container">
      <h2 style="font-weight: bolder">操作: {{ PopupData.Operation }}</h2>
      <h2 style="font-weight: bolder">时间: {{ PopupData.Date }}</h2>
      <img class="mask mask-squircle"
           :src="PopupData.ImageId">
      <button class="btn" @click="closePopup">关闭</button>
    </div>
    <div class="flex flex-col h-screen">
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

      <main class="main-container">
        <div class="product-layout">
          <!-- 左侧商品图片 -->
          <div class="product-image">
            <img :src="product.ImageId" :alt="product.Name">
          </div>

          <!-- 右侧商品信息 -->
          <div class="product-content">

            <!-- 商品名称和查看过程按钮 -->
            <div class="title-process-wrapper">
              <h2 class="product-title">{{ product.Name }}</h2>
              <button class="process-btn" @click="openProcess">查看过程</button>
            </div>

            <!-- 商品描述 -->
            <p class="product-description">{{ product.Description }}</p>

            <!-- 商品信息 -->
            <div class="product-info">
              <p class="ref-number">商品编号: {{ product.Id }}</p>
            </div>

            <!-- 商品详细信息 -->
            <div class="product-details">
              <p>发布者: {{ product.Seller }}</p>
              <p>单位: {{ product.UintGoods }}</p>
              <p>产地: {{ product.Origin }}</p>
              <p>水果类型: {{ product.Kind }}</p>
            </div>

            <!-- 数量选择 -->
            <div class="quantity-section">
              <span class="quantity-label">选择数量</span>
              <div class="quantity-controls">
                <select v-model.number="number" class="quantity-select">
                  <option v-for="n in 5" :key="n" :value="n">{{ n }}</option>
                </select>
              </div>
            </div>

            <!-- 地址输入 -->
            <div class="address-section">
              <span class="address-label">配送地址</span>
              <input type="text" v-model="address" placeholder="请输入配送地址" class="address-input"/>
            </div>

            <!-- 按钮组 -->
            <div class="action-buttons">
              <button class="add-cart-btn" @click="buyGoods">立即购买</button>
              <button class="wishlist-btn" @click="addCard">加入购物车</button>
            </div>
          </div>
        </div>
      </main>
      
      <Footer />
    </div>
  </div>
</template>

<script>
import NavHeader from '@/components/NavHeader.vue'
import Footer from '@/components/Footer.vue'

export default {
  name: 'buy-detail',
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
      product: {},
      number: 1,
      showPopup: false,
      PopupData: {},
      address: ''
    }
  },
  created() {
    this.username = localStorage.getItem('username')
    this.identity = localStorage.getItem('identity')
    this.$axios.get("/main/get_goods_detail", {
      params: {
        id: this.$route.query.id,
      }
    }).then((res) => {
      if (res.data.code == 200) {
        this.product = res.data.product
      } else {
        this.$message.error("获取商品信息失败: " + res.data.message)
      }
      console.log(res);
    }).catch((res) => {
      console.log(res);
    })
  },
  methods: {
    // 加入购物车
    addCard() {
      if (this.number === 0 || this.number > this.product.Stock) {
        this.$message.error("请填写正常的数量！！！")
        return
      }
      this.$axios.post("/main/add_to_cart", {
        goods_id: parseInt(this.$route.query.id, 10),
        number: this.number
      }).then((res) => {
        if (res.data.code === 200) {
          this.$message.success(res.data.message);
        } else {
          this.$message.error(res.data.message)
        }
      }).catch((res) => {
        console.log(res)
      })
    },
    // 触发
    openPopup(index) {
      this.PopupData = this.product.Process[index]
      this.showPopup = true;
    },
    closePopup() {
      this.showPopup = false;
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
    openProcess() {
      window.location.href = '#lookProcess';
    },
    openModel() {
      window.location.href = '#my_modal_8';
    },
    buyGoods() {
      if (this.number > this.product.Stock) {
        this.$message.error("库存不足")
      } else {
        this.$axios.post("/main/buy_goods", {
          id: parseInt(this.$route.query.id, 10),
          store: this.number,
          address: this.address
        }).then((res) => {
          if (res.data.code == 200) {
            this.$message.success("商品购买成功: " + res.data.transactionHash)
          } else {
            this.$message.error("商品购买失败: " + res.data.message)
          }
          console.log(res);
        }).catch((res) => {
          console.log(res);
        })
      }
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
  }
}
</script>

<style scoped>
.scrollable-container {
  display: flex; /* 使用 Flexbox 布局 */
  justify-content: center; /* 在水平方向上居中 */
  width: 100%; /* 容器宽度 */
  overflow-x: auto; /* 允许水平滚动 */
  white-space: nowrap; /* 防止子元素换行 */
}

.modal {
  vertical-align: top; /* 顶部对齐 */
}

.modal-box {
  width: 500px; /* 设定 modal-box 的宽度 */
  margin-right: 10px; /* 可以根据需要调整 */
  /* 其他样式 */
}

.process {
  overflow-x: auto; /* 允许水平滚动 */
  white-space: nowrap; /* 防止步骤换行 */
  display: flex; /* 使用 Flex 布局 */
  justify-content: space-between; /* 使步骤之间保持平均距离 */
}

.popup-container {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background-color: white;
  padding: 20px;
  border: 2px solid #f8b811;
  border-radius: 5px;
  box-shadow: 0 4px 20px rgba(248, 184, 17, 0.1);
  z-index: 9999;
}

.popup-container button {
  background-color: #f8b811;
  color: white;
  border: none;
  padding: 8px 16px;
  border-radius: 4px;
  cursor: pointer;
  transition: opacity 0.3s;
}

.popup-container button:hover {
  opacity: 0.9;
}

.main-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem 1rem;
}

.product-layout {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 4rem;
  align-items: start;
}

.product-image {
  width: 100%;
  background: #fff;
  padding: 1rem;
  border-radius: 8px;
  box-shadow: 0 4px 20px rgba(248, 184, 17, 0.1);
}

.product-image img {
  width: 100%;
  height: auto;
  object-fit: cover;
  border-radius: 4px;
}

.product-content {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.product-name{
  font-size: 28px;
  font-weight: 550;
  color:#585757;
  border-bottom: 3px solid #f8b811;
}

.rating-box {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.stars {
  color: #f8b811;
  font-size: 1.2rem;
}

.star {
  margin-right: 2px;
}

.review-count {
  color: #666;
  font-size: 0.9rem;
}

.product-description {
  color: #666;
  line-height: 1.6;
  font-size: 0.9rem;
}

.product-info {
  display: grid;
  gap: 0.5rem;
}

.ref-number, .availability {
  color: #666;
  font-size: 0.9rem;
}

.product-details {
  display: grid;
  gap: 0.75rem;
  color: #333;
}

.quantity-section, .address-section {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.quantity-label, .address-label {
  font-weight: 500;
  color: #333;
}

.quantity-select {
  width: 120px;
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  background: white;
}

.quantity-select:focus,
.address-input:focus {
  outline: none;
  border-color: #f8b811;
  box-shadow: 0 0 0 2px rgba(248, 184, 17, 0.1);
}

.address-input {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.action-buttons {
  display: flex;
  gap: 1rem;
  margin-top: 1rem;
}

.add-cart-btn, .wishlist-btn {
  padding: 1rem 2rem;
  border-radius: 30px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
}

.add-cart-btn {
  background-color: #f8b811;
  color: white;
  border: none;
  flex: 1;
}

.wishlist-btn {
  background-color: white;
  color: #f8b811;
  border: 2px solid #f8b811;
  flex: 1;
}

.process-btn {
  width: 20%;
  padding: 0.60rem;
  background-color: #f8b811;
  border: none;
  border-radius: 20px;
  color: #f9f9f9;
  font-size: 12px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.process-btn:hover,
.add-cart-btn:hover,
.wishlist-btn:hover {
  opacity: 0.9;
}

@media (max-width: 768px) {
  .product-layout {
    grid-template-columns: 1fr;
    gap: 2rem;
  }

  .action-buttons {
    flex-direction: column;
  }
}

.title-process-wrapper {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 70px;
  gap: 1rem;
  margin-bottom: 1rem;
  border-bottom: 2px solid #ebebeb;
}

.product-title {
  font-size: 1.8rem;
  font-weight: 600;
  color: #333;
  margin: 0;
}

.steps-container {
  padding: 20px;
  margin: 20px 0;
}

.step-item {
  display: flex;
  align-items: flex-start;
  position: relative;
  padding: 10px 0;
  cursor: pointer;
}

.step-connector {
  width: 2px;
  height: 40px;
  background-color: #f8b811;
  margin-left: 19px;
}

.step-circle {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background-color: #f8b811;
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  margin-right: 15px;
  flex-shrink: 0;
  border: 2px solid #f8b811;
  transition: all 0.3s ease;
}

.step-content {
  flex-grow: 1;
  padding: 5px 15px;
  background-color: #f8f9fa;
  border-radius: 8px;
  margin-top: 5px;
  transition: all 0.3s ease;
}

.step-content h4 {
  margin: 0;
  color: #333;
  font-size: 1.1em;
  font-weight: 600;
}

.step-date {
  margin: 5px 0 0;
  color: #666;
  font-size: 0.9em;
}

.step-item:hover .step-circle {
  transform: scale(1.1);
  box-shadow: 0 0 0 4px rgba(248, 184, 17, 0.2);
}

.step-item:hover .step-content {
  background-color: rgba(248, 184, 17, 0.1);
  transform: translateX(5px);
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-bottom: 15px;
  border-bottom: 1px solid #eee;
  margin-bottom: 20px;
}

.close-btn {
  font-size: 24px;
  color: #666;
  text-decoration: none;
  transition: color 0.3s;
}

.close-btn:hover {
  color: #f8b811;
}

.modal-box .btn {
  transition: opacity 0.3s;
}

.modal-box .btn:hover {
  opacity: 0.9;
}

@media (max-width: 640px) {
  .title-process-wrapper {
    flex-direction: column;
    align-items: flex-start;
    gap: 0.5rem;
  }
}
</style>
