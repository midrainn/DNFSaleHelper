import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

/**
 * 保存配置信息
 */

 export const Config=new Vuex.Store({
    strict:true,
    state:{
        APIhost:"localhost",
        APIport:"8090"
    },
    mutations:{

    }
 })