<template>
  <div class="about-page">
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
          <label
              class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
              for="username">
            用户名
          </label>
          <input
              class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
              id="username" placeholder="输入你的用户名" style="margin-bottom: 10px;" v-model="username" disabled/>
        </div>
        <div class="grid gap-2">
          <label
              class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
              for="password">
            密码
          </label>
          <input
              class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
              id="password" placeholder="输入你的密码" type="password" v-model="password"/>
        </div>
        <div class="grid gap-2">
          <label
              class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
              for="security-question">
            问题
          </label>
          <input
              class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
              id="security-question" v-model="question" style="margin-bottom: 10px;" placeholder="请输入密保"/>
        </div>
        <div class="grid gap-2">
          <label
              class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
              for="answer">
            答案
          </label>
          <input
              class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
              id="answer" placeholder="请输入答案" style="margin-bottom: 10px;" v-model="answer"/>
        </div>
      </div>
      <div class="modal-action">
        <a class="btn" @click="updateUserInfo">确定</a>
      </div>
    </div>
  </div>

    <div class="header-section" style="background: linear-gradient(to right, rgba(248, 184, 17, 0.8), rgba(248, 184, 17, 0.4)); height: 200px; margin-bottom: 30px;">
      <div style="height: 100%; display: flex; flex-direction: column; justify-content: center; padding: 0 40px;">
        <h1 style="font-size: 36px; font-weight: bold; margin-bottom: 10px; color: #333;">关于我们</h1>
        <p style="color: #666; font-size: 16px;">了解更多关于水果溯源系统的信息</p>
      </div>
    </div>

    <div class="app">
      <div class="flex flex-col min-h-screen">
        <section class="w-full py-16 lg:py-32 px-4">
          <div class="container grid gap-4 text-center md:gap-8 lg:grid-cols-2 lg:text-left w-full">
            <div class="space-y-4 md:space-y-6">
              <div class="space-y-3">
                <h2 class="text-4xl font-bold tracking-tighter sm:text-5xl" style="color: #f8b811;">我们的使命</h2>
                <p class="mx-auto max-w-[700px] text-gray-600 text-lg md:text-xl leading-relaxed">
                  我们致力于提供关于水果从农场到餐桌的透明和可信赖的信息。从果园的栽培、采摘，到运输、分拣，再到最终送达您手中的每一颗水果，我们追踪每个步骤，确保您能够放心享用新鲜、高质量的水果。
                </p>
              </div>
            </div>

            <div class="mx-auto grid max-w-[500px] gap-4 lg:max-w-[500px] lg:gap-8">
              <div class="mx-auto grid max-w-[350px] gap-6 lg:max-w-[500px]">
                <div class="space-y-4">
                  <h3 class="text-3xl font-bold" style="color: #f8b811;">我们的愿景</h3>
                  <p class="text-lg text-gray-600 leading-relaxed">
                    通过开放和可访问的溯源系统，赋予消费者有关于水果起源和质量的知识。我们追求将这些信息以清晰易懂的方式呈现给您，让您更好地了解您所选择的水果的来龙去脉。
                  </p>
                </div>
                <div class="space-y-4">
                  <h3 class="text-3xl font-bold" style="color: #f8b811;">我们的价值观</h3>
                  <p class="text-lg text-gray-600 leading-relaxed">
                    在这个世界上，每个人都能信任他们食用水果的来源，支持可持续的做法和当地社区的福祉。我们期望通过推动溯源技术的发展，为全球提供更多的水果信息透明度。
                  </p>
                </div>
              </div>
            </div>
          </div>
        </section>

        <section class="w-full py-16 bg-[#f8b811] bg-opacity-10 rounded-lg px-4">
          <div class="container text-center w-full">
            <h2 class="text-4xl font-bold mb-8" style="color: #f8b811;">为什么选择我们的水果溯源服务？</h2>
            <p class="text-xl leading-relaxed text-gray-600 max-w-4xl mx-auto px-6">
              我们采用基于FISCO-BCOS搭建的许可链平台，对进入区块链的每个节点进行审核和证书颁布，确保水果信息的真实性和可追溯性。通过我们的平台，您可以不仅仅了解水果的产地和生长过程，还能获取农场的实时数据和可化报告。
            </p>
          </div>
        </section>

        <Footer />
      </div>
    </div>
  </div>

</template>


<script>
import NavHeader from '@/components/NavHeader.vue'
import Footer from '@/components/Footer.vue'

export default {
  name: 'about',
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
      identity: ''
    }
  },
  created() {
    this.username = localStorage.getItem('username')
    this.identity = localStorage.getItem('identity')
  },
  methods: {
    // 退出登录
    logout() {
      localStorage.removeItem("auth_token")
      localStorage.removeItem("username")
      localStorage.removeItem("identity")
      this.$message.success('退出登录成功')
      setTimeout(() => {
        this.$router.push('/')
      }, 1000); // 1秒后行
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
  }
}
</script>
<style scoped>
  .app {
    width: 100%;
    margin: 0 auto;
    flex: 1;
  }
  
  .about-page {
    min-height: 100vh;
    display: flex;
    flex-direction: column;
  }
  
  .header-section {
    margin-bottom: 80px;
  }
  
  section {
    margin-bottom: 80px;
  }
  
  h2, h3 {
    margin-bottom: 2rem;
  }
  
  p {
    line-height: 2;
  }
  
  .container {
    max-width: none;
    width: 100%;
  }
</style>