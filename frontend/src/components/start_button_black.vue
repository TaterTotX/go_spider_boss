<script setup>

</script>

<template>
  <button class="button">
    启动
  </button>
</template>

<style scoped>
.button {
  font-size: 20px;
  --black: #000000;
  --ch-black: #141414;
  --eer-black: #1b1b1b;
  --night-rider: #2e2e2e;
  --white: #ffffff;
  --af-white: #f3f3f3;
  --ch-white: #e1e1e1;
  border-radius: 8px;
  width: 70px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  padding: 0px 15px;
  border: none;
  color: white;
  position: relative;
  cursor: pointer;
  font-weight: 700;
  transition-duration: .2s;
  background-color: var(--ch-black);
}

.button:before, .button:after {
  content: '';
  position: absolute;
  left: -2px;
  top: -2px;
  border-radius: 10px;
  background: linear-gradient(45deg,
  var(--ch-black), var(--eer-black),
  var(--night-rider), var(--ch-white), var(--night-rider),
  var(--eer-black), var(--ch-black),var(--ch-black));
  background-size: 400%;
  width: calc(100% + 4px);
  height: calc(100% + 4px);
  z-index: -1;
  animation: steam 20s linear infinite;
}

@keyframes steam {
  0% {
    background-position: 0 0;
  }

  50% {
    background-position: 400% 0;
  }

  100% {
    background-position: 0 0;
  }
}

.button:after {
  filter: blur(50px);
}


</style>