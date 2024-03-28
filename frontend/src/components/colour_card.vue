<script setup>

</script>

<template>
  <div class="card" > <h2>BOSS Robot</h2> </div>

</template>

<style scoped>
.card {
  width: 250px;
  height: 150px;
  background: #020b17;
  position: relative;
  display: flex;
  place-content: center;
  place-items: center;
  overflow: hidden;
  border-radius: 20px;
  margin: auto;

  transform: translateY(10px); /* 将卡片向下移动 10 像素 */
}

.card h2 {
  z-index: 1;
  color: white;
  font-size: 2em;
}

.card::before {
  content: '';
  position: absolute;
  width: 100px;
  background-image: linear-gradient(180deg, rgb(0, 183, 255), rgb(255, 48, 255));
  height: 200%;
  animation: rotBGimg 3s linear infinite;
  transition: all 0.2s linear;
}

@keyframes rotBGimg {
  from {
    transform: rotate(0deg);
  }

  to {
    transform: rotate(360deg);
  }
}

.card::after {
  content: '';
  position: absolute;
  background: #000000;
;
  inset: 5px;
  border-radius: 15px;
}


</style>