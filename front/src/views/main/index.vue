<template>
  <div class="manage-page">
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
          <div class="grid gap-2">
            <label class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70" for="username">用户名</label>
            <input class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
              id="username" placeholder="输入你的用户名" style="margin-bottom: 10px;" v-model="username" disabled>
          </div>
          <div class="grid gap-2">
            <label class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70" for="password">密码</label>
            <input class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
              id="password" placeholder="输入你的密码" type="password" v-model="password">
          </div>
          <div class="grid gap-2">
            <label class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70" for="username">问题</label>
            <input class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
              id="username" v-model="question" style="margin-bottom: 10px;" placeholder="请输入密保">
          </div>
          <div class="grid gap-2">
            <label class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70" for="username">答案</label>
            <input class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
              id="username" placeholder="请输入答案" style="margin-bottom: 10px;" v-model="answer">
          </div>
        </div>
        <div class="modal-action">
          <a class="btn" @click="updateUserInfo">确定</a>
        </div>
      </div>
    </div>

    <div class="app">
      <div class="flex flex-col h-screen">
        <main className="flex-1 flex items-center justify-center flex-col gap-4 p-4 lg:gap-8 md:p-6">
          <div class="hero-section" style="position: relative; width: 100%; height: 110vh;">
            <!-- 背景图片 -->
            <div class="bg-image" style="position: absolute; width: 100%; height: 100%; z-index: 0;">
              <img src="@/assets/images/main-slider/bg-1.jpg" alt="Hero Background" style="width: 100%; height: 100%; object-fit: cover;">
            </div>
            
            <!-- 内容层 -->
            <div class="content-layer" style="position: relative; z-index: 1; height: 100%; display: flex; align-items: center; padding: 0 4rem;">
              <div class="text-content">
                <h1 class="slogan" >
                  为您送来<br>最新鲜的有机水果！
                </h1>
                <p class="slogan-description">
                  Organic Mix of fruits, Perfect for weekly cooking & Snacking!
                </p>
              </div>
            </div>
          </div>

          <div class="features-container">
            <div class="feature-grid">
              <!-- 100% 有机 -->
              <div class="feature-block">
                <div class="icon-wrapper">
                  <img src="@/assets/images/resource/feature-icon-1.png" alt="organic" class="feature-icon">
                </div>
                <div class="content">
                  <h3 class="title">100% 有机</h3>
                  <p class="description">只吃健康食品</p>
                </div>
              </div>

              <!-- 每日新鲜 -->
              <div class="feature-block">
                <div class="icon-wrapper">
                  <img src="@/assets/images/resource/feature-icon-2.png" alt="fresh" class="feature-icon">
                </div>
                <div class="content">
                  <h3 class="title">每日新鲜</h3>
                  <p class="description">每天新鲜食材</p>
                </div>
              </div>

              <!-- 安全支付 -->
              <div class="feature-block">
                <div class="icon-wrapper">
                  <img src="@/assets/images/resource/feature-icon-3.png" alt="payment" class="feature-icon">
                </div>
                <div class="content">
                  <h3 class="title">安全支付</h3>
                  <p class="description">安全可靠的支付</p>
                </div>
              </div>

              <!-- 免费配送 -->
              <div class="feature-block">
                <div class="icon-wrapper">
                  <img src="@/assets/images/resource/feature-icon-4.png" alt="shipping" class="feature-icon">
                </div>
                <div class="content">
                  <h3 class="title">免费配送</h3>
                  <p class="description">订单满60元免运费</p>
                </div>
              </div>
            </div>
          </div>

          <div className="best-selling-container">
            <h1 class="best-selling-title">最畅销</h1>
            <p class="best-selling-description">
              发现我们的客户评价的最受欢迎的产品
            </p>
          </div>
          
          <div class="product-grid">
            <div class="product-item" v-for="(product, index) in products" :key="index" @click="navigateToProductDetail(product.Id)">
              <img :alt="`${product.Name} image`" :src="product.ImageId">
              <div class="product-info">
                <p class="product-title">商品名称: {{ product.Name }}</p>
                <p class="product-price">商品价格：￥ {{ product.Price }}</p>
                <p class="product-sales">销售量： <span>{{ product.Sales }}</span></p>
              </div>
            </div>
          </div>
        </main>
        <footer class="footer">
          <div class="footer-content">
            <!-- Logo部分 -->
            <div class="footer-logo">
              <img src="@/assets/logo.png" alt="FreshTrace" class="logo">
            </div>

            <!-- 联系方式部分 -->
            <div class="footer-contact">
              <h3>Contact</h3>
              <p class="address">上海市嘉定区<br>曹安公路4800号同济大学</p>
              <p class="email"><i class="fas fa-envelope"></i>@2250591.tongji.edu.cn</p>
              <!--<p class="phone"><i class="fas fa-phone"></i>+86 13800000000</p>-->
            </div>
          </div>

          <!-- 版权信息 -->
          <div class="footer-bottom">
            <p>© Copyright Reserved by FreshTrace</p>
          </div>
        </footer>
      </div>
    </div>
  </div>
</template>

<script>
import NavHeader from '@/components/NavHeader.vue'

export default {
  name: 'index',
  components: {
    NavHeader
  },
  data() {
    return {
      username: '',
      password: '',
      question: '',
      answer: '',
      identity: '',
      products: [],
      currentSlide: 'slide1',
      slidesOrder: ['slide1', 'slide2', 'slide3', 'slide4'],
      autoPlayInterval: null
    }
  },
  created() {
    let imgs = [
      "src/assets/carousel-1.jpg",
      "src/assets/carousel-2.png",
      "src/assets/carousel-3.png",
      "src/assets/carousel-4.png",
    ]

    for (let img of imgs) {
      let image = new Image()
      image.src = img
      image.onload = () => {
        this.count++
      }
    }
    this.username = localStorage.getItem('username')
    this.identity = localStorage.getItem('identity')
    this.$axios.get('/main/get_sales_top_three').then((res) => {
      this.products = res.data.products
      console.log(res);
    }).catch((res) => {
      console.log(res);
    })
  },
  mounted() {
    this.startAutoPlay();
  },
  beforeUnmount() {
    this.stopAutoPlay();
  },
  methods: {
    changeSlide(slideId) {
      this.currentSlide = slideId;
      this.resetAutoPlay();
    },
    startAutoPlay() {
      this.autoPlayInterval = setInterval(() => {
        const currentIndex = this.slidesOrder.indexOf(this.currentSlide);
        const nextIndex = (currentIndex + 1) % this.slidesOrder.length;
        this.currentSlide = this.slidesOrder[nextIndex];
      }, 3000); // 每 3 秒切换一次幻灯片
    },
    stopAutoPlay() {
      clearInterval(this.autoPlayInterval);
    },
    resetAutoPlay() {
      this.stopAutoPlay();
      this.startAutoPlay();
    },
    // 退出登录
    logout() {
      // 清除本地存储
      localStorage.removeItem("auth_token")
      localStorage.removeItem("username")
      localStorage.removeItem("identity")
      
      // 显示成功消息
      this.$message.success('退出登录成功')
      
      // 延迟跳转
      setTimeout(() => {
        this.$router.push('/')
      }, 1000)
    },
    navigateToProductDetail(productId) {
      this.$router.push({path: `/buy-detail`, query: {id: productId}});
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
  }
}

</script>

<style scoped>
.hero-section {
  background-color: #f8f8f8;
  min-height: 600px;
  position: relative;
  overflow: hidden;
}

.carousel-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 40px 20px;
}

.slide {
  position: relative;
}

.content-wrapper {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 40px;
}

.text-content {
  flex: 1;
  max-width: 650px;
  margin-left: 100px;
}

.slogan{
  font-size: 65px;
  font-weight: 540;
  color: #555555;
  line-height: 1.6;
  margin-bottom: 40px;
}

.slogan-description {
  font-size: 30px;
  color: #666;
  line-height: 1.6;
  margin-bottom: 40px;
}

.button-group {
  display: flex;
  gap: 40px;
}

.home-btn {
  padding: 20px 60px;
  height: 80px;
  border-radius: 30px;
  font-weight: 600;
  font-size: 25px;
  color: #fff;
  text-align: center;
  transition: all 0.3s ease;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  line-height: 1;
}

.btn {
  padding: 15px 35px;
  border-radius: 30px;
  font-weight: 600;
  font-size: 16px;
  transition: all 0.3s ease;
  cursor: pointer;
}

.image-content {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
}

.hero-image {
  max-width: 100%;
  height: auto;
}

/* 过渡动画 */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

/* 响应式设计 */
@media (max-width: 968px) {
  .content-wrapper {
    flex-direction: column;
    text-align: center;
  }

  .title {
    font-size: 48px;
  }

  .button-group {
    justify-content: center;
  }
}

@media (max-width: 640px) {
  .title {
    font-size: 36px;
  }

  .description {
    font-size: 18px;
  }

  .button-group {
    flex-direction: column;
    align-items: center;
  }
}

.features-container {
  padding: 4rem 2rem;
  background: #ffffff;
  width: 100%;
  margin: 0 auto;
}

.feature-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr); /* 2列布局 */
  gap: 3rem;
  margin: 0 auto;
}

.feature-block {
  display: flex;
  align-items: center;
  padding: 2rem;
  background: #f8f8f8;
  border-radius: 12px;
  transition: all 0.3s ease;
}

.feature-block:hover {
  transform: translateY(-5px);
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
}

.icon-wrapper {
  margin-right: 2rem;
  background: #f8b811;
  padding: 1rem;
  border-radius: 50%;
  width: 120px;
  height: 120px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.feature-icon {
  width: 80px;
  height: 80px;
  object-fit: contain;
}

@media (max-width: 768px) {
  .feature-grid {
    grid-template-columns: 1fr;
    gap: 2rem;
  }
  
  .features-container {
    padding: 2rem 1rem;
  }
  
  .feature-block {
    padding: 1.5rem;
  }
  
  .icon-wrapper {
    width: 60px;
    height: 60px;
    margin-right: 1.5rem;
  }
  
  .feature-icon {
    width: 30px;
    height: 30px;
  }
}

.content {
  display: flex;
  flex-direction: column;
}

.title {
  font-size: 1.2rem;
  font-weight: 600;
  margin: 0;
  color: #333;
}

.description {
  font-size: 0.9rem;
  color: #666;
  margin: 0.25rem 0 0 0;
}

.best-selling-title{
  font-size: 40px;
  font-weight: 700;
  text-align: center;
  color:#515151;
}

.best-selling-description{
  font-size: 20px;
  font-weight: 400;
  color:#515151;

}

.footer {
  background-color: #f7c74c;
  color: #ffffff;
  padding: 60px 0 20px;
  position: relative;
  margin-top: 100px;  
}

.footer-content {
  max-width: 1200px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 40px;
  padding: 0 20px;
}

.footer-logo {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.logo {
  max-width: 180px;
}

.social-links {
  display: flex;
  gap: 15px;
}

.social-link {
  width: 40px;
  height: 40px;
  background-color: rgba(255, 255, 255, 0.1);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  transition: all 0.3s ease;
}

.social-link:hover {
  background-color: #f8b811;
}

h3 {
  color: #ffffff;
  font-size: 24px;
  margin-bottom: 25px;
  font-weight: 600;
}

.footer-contact p {
  color: #ffffff;
  margin-bottom: 15px;
  display: flex;
  align-items: center;
  gap: 10px;
}

.footer-bottom {
  text-align: center;
  padding-top: 30px;
  margin-top: 50px;
  border-top: 1px solid rgba(255, 255, 255, 0.3);
  color: #ffffff;
}

/* 响应式设计 */
@media (max-width: 1024px) {
  .footer-content {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 640px) {
  .footer-content {
    grid-template-columns: 1fr;
  }
  
  .gallery-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

.product-grid {
  display: grid;
  gap: 3em;
  width: 90%;
}

@media (min-width: 768px) {
  .product-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (min-width: 1024px) {
  .product-grid {
    grid-template-columns: repeat(3, 1fr);
  }
}

.product-item {
  cursor: pointer;
}

.product-item img {
  aspect-ratio: 1;
  border-radius: 0.5rem;
  object-fit: cover;
  width: 80%;
  height:80%;
}

.product-info {
  display: grid;
  gap: 0.375rem;
}

.product-title {
  margin-top: 10px;
  font-size: 20px;
  font-weight: 500;
}

.product-price {
  font-size: 18px;
  font-weight: 500;
  letter-spacing: -0.025em;
  color: red;
}

.product-sales {
  font-size: 18px;
  font-weight: 500;
  letter-spacing: -0.025em;
}

.product-sales span {
  color: green;
}

/* 确保导航栏和模态框在最上层 */
.manage-page {
  position: relative;
}

/* 导航栏容器样式 */
:deep(.nav-header) {
  position: relative;
  z-index: 1000; /* 提高 z-index 确保在最顶层 */
  pointer-events: auto;
}

/* 模态框样式 */
.modal {
  z-index: 2000; /* 确保模态框在最顶层 */
}

/* 确保背景图片不会遮挡点击事件 */
.bg-image {
  position: absolute;
  width: 100%;
  height: 100%;
  z-index: 1;
  pointer-events: none;
}

/* 修改内容层的 z-index */
.content-layer {
  position: relative;
  z-index: 2;
  height: 100%;
  display: flex;
  align-items: center;
  padding: 0 4rem;
  pointer-events: none; /* 防止内容层阻挡点击事件 */
}

/* 恢复文本内容的点击事件 */
.text-content {
  pointer-events: auto;
}

/* 确保整个页面容器正确定位 */
.manage-page {
  position: relative;
  min-height: 100vh;
}

</style>
