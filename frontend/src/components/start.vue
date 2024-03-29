<script setup>
import {reactive} from 'vue'
import {Greet} from '../../wailsjs/go/main/App'
import Start_button_black from "./start_button_black.vue";
const data = reactive({
  name: "",
  resultText: "使用前一定要看说明哦 点击显示说明",
  helloText:"",
  regexp_text: "",
})

function greet() {
  // 确保Greet函数可以接受三个参数，这里我们假设它是Greet(name, ws, helloText)
  Greet(data.name, data.helloText,data.regexp_text).then(result => {
    data.resultText = result;
  });
}

</script>

<template>
  <main>

    <div id="result" class="result animate__animated animate__jackInTheBox">
      <a href="https://github.com/zypotato/go_spider_boss" style="color: inherit; text-decoration: none;">{{ data.resultText }}</a>
    </div>    <div id="input" class="input-box">

      <input id="name" v-model="data.helloText" autocomplete="off" class="input" type="text" placeholder="请输入打招呼语"/>
      <br>
      <br>
      <input id="name" v-model="data.regexp_text" autocomplete="off" class="input" type="text" placeholder="(?i)java开发|算法"/>
      <br>
      <br>
      <br>
      <start_button_black style="display: block; margin-left: auto; margin-right: auto; width: 150px;" @click="greet"></start_button_black>
    </div>
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}




.input {
  background-color: #212121;
  max-width: 190px;
  height: 20px;
  padding: 10px;
  /* text-align: center; */
  border: 2px solid white;
  border-radius: 5px;
  /* box-shadow: 3px 3px 2px rgb(249, 255, 85); */
  color: white; /* 设置文字颜色为白色 */
}

.input:focus {
  color: rgb(0, 255, 255);
  background-color: #212121;
  outline-color: rgb(255, 255, 255);
  box-shadow: -3px -3px 15px rgb(0, 255, 255);
  transition: .1s;
  transition-property: box-shadow;
}

</style>




