<script setup>

</script>

<template>
  <input  class="input" name="text" type="text"  id="name" v-model="data.ws" autocomplete="off"  placeholder="请输入ws的值">
  <input  class="input" name="text" type="text"  id="name" v-model="data.helloText" autocomplete="off"  placeholder="请输入打招呼语">
  <input  class="input" name="text" type="text"  id="name" v-model="data.regexp_text" autocomplete="off"  placeholder="(?i)数据分析|数据分析师">
</template>

<style scoped>
.input {
  background-color: #212121;
  max-width: 190px;
  height: 20px;
  padding: 10px;
  /* text-align: center; */
  border: 2px solid white;
  border-radius: 5px;
  /* box-shadow: 3px 3px 2px rgb(249, 255, 85); */
}

.input:focus {
  color: rgb(0, 255, 255);
  background-color: #212121;
  outline-color: rgb(0, 255, 255);
  box-shadow: -3px -3px 15px rgb(0, 255, 255);
  transition: .1s;
  transition-property: box-shadow;
}
</style>