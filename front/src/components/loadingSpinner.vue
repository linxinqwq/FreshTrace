<template>
  <div v-if="isLoading" class="loading-overlay">
    <img src="../assets/preloader.gif" alt="Loading..." class="spinner" />
    <p class="loading-text">Loading...</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      isLoading: true,  // 控制 loading 显示
    };
  },
  methods: {
    hideLoading() {
      this.isLoading = false;  // 页面加载完成时调用此方法
    },
  },
};
</script>

<style scoped>
/* Loading overlay 样式 */
.loading-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(255, 255, 255, 0.8); /* 背景半透明 */
  display: flex;
  flex-direction: column; /* 使内容垂直排列 */
  justify-content: center;
  align-items: center;
  z-index: 9999; /* 确保在其他内容之上 */
}

/* Spinner 样式 */
.spinner {
  width: 70px; /* 根据需要调整 GIF 的大小 */
  height: 70px;
}

/* Loading 文本样式 */
.loading-text {
  margin-top: 0; /* 与 GIF 之间的间距 */
  font-size: 1rem; /* 字体大小 */
  color: #333; /* 字体颜色 */
}
</style>
  