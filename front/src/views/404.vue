<template>
    <div class="app">
        <!--
// v0 by Vercel.
// https://v0.dev/t/SJaojmyUce3
-->
        <div class="flex flex-col items-center justify-center min-h-[400px] space-y-4">
            <div class="flex flex-col items-center space-y-2">
                <h1 class="text-4xl font-bold tracking-tighter sm:text-6xl">
                    404 Not Found
                </h1>
                <p class="max-w-[600px] text-center text-gray-500 md:text-xl/relaxed dark:text-gray-400">
                    页面找不到了
                </p>
            </div><a
                class="inline-flex h-10 items-center rounded-md border border-gray-200 border-gray-200 bg-white px-8 text-sm font-medium shadow-sm transition-colors hover:bg-gray-100 hover:text-gray-900 focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-gray-950 dark:border-gray-800 dark:border-gray-800 dark:bg-gray-950 dark:hover:bg-gray-800 dark:hover:text-gray-50 dark:focus-visible:ring-gray-300"
                @click="gotoLogin">
                回到主页
            </a>
        </div>
    </div>
</template>

<script>
export default {
    name: "404",
    methods: {
        gotoLogin() {
            this.$router.push('/home');
        },
    }
}
</script>

<style></style>