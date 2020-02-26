<template>
  <div class="mainstyle">
    <el-row :gutter="20">
      <el-col :offset="7" :span="5">
        <h1>登</h1>
      </el-col>
      <el-col :span="3">
        <h1>录</h1>
      </el-col>
      <br />
      <br />
      <br />
      <br />
      <br />
      <br />
    </el-row>
    <el-row :gutter="20">
      <el-col :span="13" :offset="5">
        <el-input placeholder="账号" v-model="username">
          <template slot="prepend">账号</template>
        </el-input>
      </el-col>
    </el-row>
    <br />
    <el-row :gutter="20">
      <el-col :span="13" :offset="5">
        <el-input placeholder="密码" v-model="password" show-password>
          <template slot="prepend">密码</template>
        </el-input>
      </el-col>
    </el-row>
    <br />
    <br />
    <br />
    <br />
    <el-row :gutter="20">
      <el-col :span="7" :offset="5">
        <el-button @click="login()" type="success" round>登录</el-button>
      </el-col>
      <el-col :span="5">
        <el-button @click="cancel()" type="danger" round>取消</el-button>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import {Config} from './Config.js'
export default {
  name: "login",
  data: function() {
    return {
      username: "",
      password: ""
    };
  },
  methods: {
      /**
     * 关闭登录界面
     * @author mid
     * @date 2020年2月26日03:53:05
     * 
     */
    cancel: function() {
      const that = this;
      that.$emit("closelogin");
    },
      /**
     * 解析服务器返回的CODE
     * @author mid
     * @date 2020年2月26日03:53:05
     * 
     */
    parseServerMsg(msg1) {
      switch (msg1) {
        case "CODE1000":
          return { msg: "登录成功", flag: true };
        case "CODE1002":
          return { msg: "数据提交成功", flag: true };
        case "CODE2001":
          return { msg: "尚未登录", flag: false };
        case "CODE2002":
          return { msg: "账号或密码错误", flag: false };
        case "CODE2003":
          return { msg: "数据提交失败", flag: false };
        default:
          return {
            msg: "未知返回值,请与管理员联系,错误值:" + msg1,
            flag: false
          };
      }
    },
      /**
     * 提交登录请求
     * @author mid
     * @date 2020年2月26日03:53:05
     * 
     */
    login: function() {
      const that = this;
      that
        .axios({
          headers: {
            "Content-Type": "application/x-www-form-urlencoded;charset=UTF-8"
          },
          method: "post",
          url: "http://"+Config.state.APIhost+":"+Config.state.APIport+"/api/dnflogin",
          data: {
            username: that.username,
            password: that.password,
            flag: "ccc",
            sfsdf: "sss"
          },
          withCredentials: true
        })
        .then(Response => {
         let result = that.parseServerMsg(Response.data);
          if (result.flag) {
            that.$message({
              message: result.msg,
              type: "success"
            });
            that.$emit("closelogin");
            that.$emit("loginsuccess");
          } else {
            that.$message({
              message: result.msg,
              type: "error"
            });
          }
        })
        .catch(error => {
          that.$message({
            message: error,
            type: "error"
          });
        });
    }
  }
};
</script>
<style scoped>
.mainstyle {
  position: absolute;
  transform: translate(-50%, -50%);
  left: 50%;
  top: 50%;
  width: 800px;
  height: 400px;
  background-color: rgba(255, 255, 255, 0.8);
  border: 1px solid rgba(0, 0, 0, 0.5);
}
</style>