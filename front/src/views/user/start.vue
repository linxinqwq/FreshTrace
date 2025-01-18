<template>
  <div class="start">
    <div class="content-wrapper">
      <div class="welcome-text">
        <h1 class="title">欢迎来到 FreshTrace</h1>
        <p class="subtitle">您的水果溯源专家</p>
      </div>
      <img :class="['logo', { 'animate': isAnimated }]" src="../../assets/logo.png" alt="Orgonik Logo">
      <div class="button-container">
        <button class="btn-primary" @click="gotoRegister">注册</button>
        <button class="btn-secondary" @click="gotoLogin">登录</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "start",
  data() {
    return {
      isAnimated: false // 控制动画的状态
    };
  },
  mounted() {
    this.isAnimated = true; // 页面加载时启动动画
    setTimeout(() => {
      this.isAnimated = false; // 动画结束后移除动画类
    }, 1500); // 动画持续时间
  },
  methods: {
    gotoRegister() {
      this.$router.push("/register");
    },
    gotoLogin() {
      this.$router.push("/login");
    }
  }
}
</script>

<style scoped>
.start {
  background-image: url('../../assets/bg-2.jpg');
  background-size: cover;
  background-position: center;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}

.content-wrapper {
  text-align: center;
  background: rgba(255, 255, 255, 0.95);
  padding: 30px 60px;
  border-radius: 20px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  width: 600px;
}

.welcome-text {
  margin-bottom: 20px;
}

.title {
  font-size: 2.5rem;
  font-weight: bold;
  color: #333;
  margin-bottom: 10px;
}

.subtitle {
  font-size: 1.2rem;
  color: #666;
}

.logo {
  width: 200px;
  height: 200px;
  object-fit: contain;
  margin-bottom: 80px;
  transition: transform 1.5s ease-in-out;
  display: block;
  margin: 0 auto 40px;
}

.logo.animate {
  animation: logoAnimation 1.5s ease-in-out;
}

@keyframes logoAnimation {
  0%, 100% {
    transform: scale(1);
  }
  50% {
    transform: scale(1.1);
  }
}

.button-container {
  display: flex;
  gap: 30px;
  justify-content: center;
  margin-top: 30px;
}

.btn-primary,
.btn-secondary {
  padding: 15px 30px;
  width: 150px;
  border: none;
  border-radius: 20px;
  font-size: 1.2rem;
  cursor: pointer;
  transition: all 0.3s ease;
}

.btn-primary {
  background-color: #f8b811;
  color: white;
}

.btn-secondary {
  background-color: #FF6E2F;
  color: white;
}

.btn-primary:hover,
.btn-secondary:hover {
  transform: translateY(-2px);
  opacity: 0.9;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}
</style> 