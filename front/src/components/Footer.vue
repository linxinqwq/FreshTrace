<template>
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
</template>

<script>
export default {
  name: 'Footer'
}
</script>

<style scoped>
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

@media (max-width: 1024px) {
  .footer-content {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 640px) {
  .footer-content {
    grid-template-columns: 1fr;
  }
}
</style> 