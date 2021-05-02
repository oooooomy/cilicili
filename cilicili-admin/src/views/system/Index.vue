<template>
  <div class="system">
    <iframe src="http://localhost:8848/nacos/index.html"/>
  </div>
</template>

<script>
export default {
  name: "Index",
}
</script>

<style scoped>
.system iframe {
  width: 100%;
  height: 100vh;
  border: none;
}
</style>