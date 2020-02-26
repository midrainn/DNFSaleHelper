import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter)

import index from './components/index.vue'
// import login from './components/login.vue'

const routes=[
    {path:'/',component:index},
]

const router=new VueRouter({
    routes
})

export default router
