<template>
    <div class="app">
        <div class="relative flex items-center min-h-screen px-4 py-12 sm:px-6 lg:px-8">
            <div class="mx-auto w-full max-w-sm space-y-4" style="background-color:#fff; border:1px solid #fff; border-radius:10px;">
                <div class="text-center mb-8">
                    <img src="../../assets/logo.png" alt="FreshTrace" class="mx-auto w-48 mb-4">
                    <h1 class="text-2xl font-bold text-gray-800">欢迎来到 FreshTrace</h1>
                    <p class="text-gray-600 mt-2">您的水果溯源专家</p>
                </div>
                <div class="p-4 grid gap-2 rounded-lg border border-gray-200 dark:border-gray-800">
                    <form class="space-y-2">
                        <p class="text-sm text-gray-500 dark:text-gray-400" style="margin-bottom:20px">
                            找回账号
                        </p>
                        <div class="grid gap-2">
                            <div class="grid gap-2"><label
                                    class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                                    for="username">用户名</label><input
                                    class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                                    id="username" placeholder="输入你的用户名" style="margin-bottom: 10px;" v-model="userName"
                                    :disabled="question.length != 0">
                            </div>
                            <div class="grid gap-2" v-show="question.length != 0"><label
                                    class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                                    for="username">问题</label><input
                                    class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                                    id="username" :placeholder="question" disabled style="margin-bottom: 10px;"></div>
                            <div class="grid gap-2" v-show="question.length != 0"><label
                                    class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                                    for="username">答案</label><input
                                    class="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                                    id="username" placeholder="请输入答案" style="margin-bottom: 10px;" v-model="answer"></div>
                        </div><button
                            class="inline-flex items-center justify-center bg-[#f8b811] text-white rounded-md text-sm font-medium h-10 px-4 py-2 w-full hover:opacity-90 transition-opacity"
                            type="button" style="margin-bottom: 10px;" @click="retrieveByUsername">
                            找回账号
                        </button>
                    </form>
                </div>
                <div class="space-y-2 text-center" @click="gotoLogin"><a class="font-semibold underline"
                        style="cursor: pointer;">
                        登录账号
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
    <div class="modal" role="dialog" id="my_modal_8">
        <div class="modal-box">
            <h3 class="font-bold text-lg">提示</h3>
            <p class="py-4">请保存密码：</p>
            <p class="py-4" v-text="password"></p>
            <div class="modal-action">
                <a href="#" class="btn">确定</a>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "retrieve",
    data() {
        return {
            userName: '',
            question: '',
            answer: '',
            password: ''
        }
    },
    methods: {
        gotoLogin() {
            this.$router.push('/');
        },
        retrieveByUsername() {
            if (this.userName.length < 6) {
                this.$message.error("用户名不合法")
                return
            }

            if (this.question.length != 0) {
                this.retrieve()
                return
            }
            this.$axios.post('/user/retrieve', {
                username: this.userName
            }).then((res) => {
                console.log(res);
                if (res.data.code == 200) {
                    this.question = res.data.question
                } else {
                    this.$message.error(res.data.message)
                }
            }).catch((res) => {
                console.log(res)
            })
        },
        retrieve() {
            if (this.userName.length < 6 || this.question.length == 0) {
                this.$message.error("用户名不合法")
                return
            }
            this.$axios.post('/user/retrieve_result', {
                username: this.userName,
                question: this.question,
                answer: this.answer
            }).then((res) => {
                if (res.data.code == 200) {
                    this.password = res.data.password
                    window.location.href = '#my_modal_8';
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

/* 输入框焦点样式 */
input:focus {
  border-color: #f8b811;
  box-shadow: 0 0 0 2px rgba(248, 184, 17, 0.1);
  outline: none;
}

/* 模态框按钮样式 */
.modal-action .btn {
  background-color: #f8b811;
  color: white;
  transition: opacity 0.3s;
}

.modal-action .btn:hover {
  opacity: 0.9;
}

/* 链接悬停样式 */
.underline:hover {
  color: #f8b811;
}
</style>