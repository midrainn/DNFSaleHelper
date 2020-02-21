<template>
  <div>
    <h1 style="text-align:center">DNF小账本</h1>
    <h4>{{version}}</h4>
    <el-container style="border: 1px solid black">
      <el-container>
        <el-main>
          <el-table :data="ItemsTable" height="400" :row-class-name="tableRowClassName">
            <el-table-column prop="name" label="物品名称"></el-table-column>
            <el-table-column prop="num" label="当前库存"></el-table-column>
            <el-table-column prop="rate" label="当前税率"></el-table-column>
            <el-table-column prop="buyprice" label="进价"></el-table-column>
            <el-table-column prop="saleprice" label="售价"></el-table-column>
            <el-table-column prop="afterrate" label="税后售价"></el-table-column>
            <el-table-column prop="profit" label="单个利润"></el-table-column>
            <el-table-column prop="change" label="涨幅"></el-table-column>
            <el-table-column prop="income" label="总收入"></el-table-column>
            <el-table-column label="操作" width="250">
              <template slot-scope="scope">
                <el-button @click="addStore(scope.row.name)" size="mini" type="success" round>加仓</el-button>
                <el-button
                  @click="ModifyItemFun(scope.row.name)"
                  size="mini"
                  type="warning"
                  round
                >修改</el-button>
                <el-button @click="DelItem(scope.row.name)" size="mini" type="danger" round>删除</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-main>
      </el-container>
    </el-container>

    <el-row>
      <el-button @click="FlushDate()" type="primary" round>刷新数据</el-button>
      <el-button @click="Show('additem')" type="success" round>增加物品</el-button>
    </el-row>
    <div style="float:right">
      当前税率:{{rate}}
      <el-checkbox @change="FlushDate()" v-model="isDiscount">是否有驴子</el-checkbox>
    </div>
    <div class="ModifyItem" v-show="isModify">
      <h1 style="text-align:center">修改物品</h1>
      <el-row>
        <div class="fixedinput">
          <el-input placeholder="请输入内容" v-model="ModifyItem.name" :disabled="true">
            <template slot="prepend">物品名称:</template>
          </el-input>
        </div>
        <div class="fixedinput">
          <el-input placeholder="请输入内容" v-model="ModifyItem.num">
            <template slot="prepend">物品数量:</template>
          </el-input>
        </div>

        <div class="fixedinput">
          <el-input placeholder="请输入内容" v-model="ModifyItem.buyprice">
            <template slot="prepend">购买价格:</template>
          </el-input>
        </div>
        <div class="fixedinput">
          <el-input placeholder="请输入内容" v-model="ModifyItem.saleprice">
            <template slot="prepend">出售价格:</template>
          </el-input>
        </div>
        <div class="fixedinput">
          <el-button @click="ModifyItemSuccess()" type="success" round>确认</el-button>
          <el-button @click="Hide('modify')" type="danger" round>取消</el-button>
        </div>
        <br />
      </el-row>

    </div>

    <div class="ModifyItem" v-show="isAddItem">
      <el-row>
        <h1 style="text-align:center">增加物品</h1>
        <div class="fixedinput">
          <el-input placeholder="请输入物品名称" v-model="newItem.name">
            <template slot="prepend">物品名称:</template>
          </el-input>
        </div>
        <div class="fixedinput">
          <el-input placeholder="请输入内容" v-model="newItem.num">
            <template slot="prepend">物品数量:</template>
          </el-input>
        </div>

        <div class="fixedinput">
          <el-input placeholder="请输入内容" v-model="newItem.buyprice">
            <template slot="prepend">购买价格:</template>
          </el-input>
        </div>
        <div class="fixedinput">
          <el-input placeholder="请输入内容" v-model="newItem.saleprice">
            <template slot="prepend">出售价格:</template>
          </el-input>
        </div>
        <div class="fixedinput">
          <el-button @click="addItem()" type="success" round>确认</el-button>
          <el-button @click="Hide('additem')" type="danger" round>取消</el-button>
        </div>
        <br />
      </el-row>
    </div>
    <div class="ModifyItem" v-show="isAddStore">
      <h1 style="text-align:center">加仓</h1>
      <el-row>
        <div class="fixedinput">
          <el-input placeholder="请输入内容" v-model="ModifyItem.name" :disabled="true">
            <template slot="prepend">物品名称:</template>
          </el-input>
        </div>
        <div class="fixedinput">
          <el-input placeholder="请输入内容" v-model="ModifyItem.num">
            <template slot="prepend">加仓数量:</template>
          </el-input>
        </div>
        <div class="fixedinput">
          <el-input placeholder="请输入内容" v-model="ModifyItem.buyprice">
            <template slot="prepend">加仓价格:</template>
          </el-input>
        </div>
        <div class="fixedinput">
          <el-button @click="addStoreSuccess()" type="success" round>确认</el-button>
          <el-button @click="Hide('addstore')" type="danger" round>取消</el-button>
        </div>
        <br />
      </el-row>
    </div>
  </div>
</template>

<script>
export default {
  name: "HelloWorld",
  props: {},
  mounted: function() {
    this.LoadSave();
  },
  data: function() {
    return {
      version:"测试版v0.1",
      items: [
        {
          name: "星空深渊原石",
          num: 10,
          buyprice: 46500,
          saleprice: 52500
        }
      ],
      isDiscount: true,
      isModify: false,
      isAddStore: false,
      isAddItem: false,
      rate: 0.97,
      ItemsTable: [],
      Income: 0,
      ModifyItem: {
        name: "未选中",
        num: 0,
        buyprice: 0,
        saleprice: 0
      },
      newItem: {
        name: "",
        num: 0,
        buyprice: 0,
        saleprice: 0
      }
    };
  },
  methods: {
    tableRowClassName({ row, rowIndex }) {
      let flag = parseFloat(row.change.replace("%", ""));
      console.log("" + flag > 0);
      if(flag>0){
        return 'success-row';
      }
      else if(flag<0){
        return 'danger-row';
      }
      return '';
      // if (rowIndex === 1) {
      //   return 'warning-row';
      // } else if (rowIndex === 3) {
      //   return 'success-row';
      // }
      // return '';
    },
    FlushDate() {
      const that = this;
      if (that.isDiscount == true) {
        that.rate = 0.97;
      } else {
        that.rate = 0.95;
      }
      that.ItemsTable = [];
      that.Income = 0;
      that.items.forEach(item => {
        function fixResult(result) {
          let w1 = 10000;
          let w2 = -10000;
          let e1 = 100000000;
          let e2 = -100000000;
          result=parseFloat(result);
          return result >= e1 || result <= e2
            ? (result / e1).toFixed(2) + "亿"
            : result >= w1 || result <= w2
            ? (result / w1).toFixed(2) + "万"
            : (result).toFixed(2);
        }
        let tmpitem = {};

        tmpitem.afterrate = item.saleprice * that.rate;
        tmpitem.profit = tmpitem.afterrate - item.buyprice;
        tmpitem.change =
          ((tmpitem.profit / item.buyprice) * 100).toFixed(2) + "%";

        tmpitem.name = item.name;
        tmpitem.num = fixResult(item.num);
        tmpitem.rate = that.rate;
        tmpitem.buyprice = fixResult(item.buyprice);
        tmpitem.saleprice = fixResult(item.saleprice);
        tmpitem.afterrate = fixResult(tmpitem.afterrate);
        tmpitem.income = fixResult(tmpitem.profit * item.num);
        tmpitem.profit = fixResult(tmpitem.profit);

        that.Income += tmpitem.profit * item.num;
        that.ItemsTable.push(tmpitem);
      });
    },
    LoadSave() {
      const that = this;
      let dataStr = window.localStorage.getItem("Items");
      if (dataStr != null) {
        that.items = JSON.parse(dataStr);
      } else {
        that.items = [];
      }
      that.FlushDate();
    },
    SaveData() {
      const that = this;
      if (that.items != null) {
        window.localStorage.setItem("Items", JSON.stringify(that.items));
      }
    },
    Show(Itype) {
      if (Itype == "modify") {
        this.isModify = true;
      } else if (Itype == "additem") {
        this.isAddItem = true;
      } else if (Itype == "addstore") {
        this.isAddStore = true;
      }
    },
    Hide(Itype) {
      if (Itype == "modify") {
        this.isModify = false;
      } else if (Itype == "additem") {
        this.isAddItem = false;
      } else if (Itype == "addstore") {
        this.isAddStore = false;
      }
    },
    addItem() {
      let isSave = true;
      const that = this;
      if (isNaN(that.newItem.num) || isNaN(that.newItem.buyprice)||isNaN(that.newItem.saleprice)) {
        alert("请检查输入");
      }
      that.items.forEach(item => {
        if (item.name == that.newItem.name) {
          alert("当前物品已存在！");
          isSave = false;
          return;
        }
      });
      if (isSave) {
        let tmpitem=JSON.parse(JSON.stringify(that.newItem));
        that.items.push(tmpitem);
        that.SaveData();
        that.Hide("additem");
        this.$message({
          message: "添加物品" + tmpitem.name + "成功",
          type: "success"
        });
      }
      that.FlushDate();
    },
    addStore(itemname) {
      const that = this;
      that.items.forEach(item => {
        if (item.name == itemname) {
          that.ModifyItem.name = itemname;
          that.ModifyItem.num = 0;
          that.ModifyItem.buyprice = 0;
          return;
        }
      });
      that.Show("addstore");
    },
    addStoreSuccess() {
      const that = this;
      if (isNaN(that.ModifyItem.num) || isNaN(that.ModifyItem.buyprice)) {
        alert("请检查输入");
      }
      that.items.forEach(item => {
        if (item.name == that.ModifyItem.name) {
          let BuySum =
            item.num * item.buyprice +
            that.ModifyItem.buyprice * that.ModifyItem.num;
          item.num = parseInt(item.num) + parseInt(that.ModifyItem.num);
          item.buyprice = BuySum / item.num;
          that.SaveData();
          that.Hide("addstore");
          that.$message({
            message:
              "物品" +
              that.newItem.name +
              "加仓" +
              that.ModifyItem.num +
              "个成功",
            type: "success"
          });
          that.FlushDate();
          return;
        }
      });
    },
    ModifyItemFun(itemname) {
      const that = this;
      that.items.forEach(item => {
        if (item.name == itemname) {
          that.ModifyItem.name = itemname;
          that.ModifyItem.num = item.num;
          that.ModifyItem.buyprice = item.buyprice;
          that.ModifyItem.saleprice = item.saleprice;
          return;
        }
      });
      that.Show("modify");
    },
    ModifyItemSuccess() {
      const that = this;
      that.items.forEach(item => {
        if (item.name == that.ModifyItem.name) {
          if (confirm("确认覆盖" + item.name + "吗")) {
            item.num = that.ModifyItem.num;
            item.buyprice = that.ModifyItem.buyprice;
            item.saleprice = that.ModifyItem.saleprice;
            that.$message({
              showClose: true,
              message: "" + item.name + "已被修改",
              type: "warning"
            });
          }
          that.Hide("modify");
          that.SaveData();
          that.FlushDate();
          return;
        }
      });
    },
    DelItem(itemname) {
      const that = this;
      that.items.forEach((item, index) => {
        if (item.name == itemname) {
          if (confirm("确认删除物品:" + itemname + "吗") == true) {
            that.items.splice(index, 1);
            that.$message({
              showClose: true,
              message: "已删除" + itemname,
              type: "error"
            });
          }
          that.SaveData();
          that.FlushDate();
          return;
        }
      });
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.ModifyItem {
  z-index: 999;
  position: absolute;
  width: 50%;
  text-align: center;
  transform: translate(-50%, -50%);
  left: 50%;
  top: 40%;
  background-color: white;
  border: 1px solid black;
}
.fixedinput {
  width: 50%;
  margin-top: 20px;
  margin-left: 25%;
}
</style>

<style>
.el-table .danger-row {
  background: rgb(250, 22, 26);
}

.el-table .success-row {
  background: #7dfd39;
}
</style>
