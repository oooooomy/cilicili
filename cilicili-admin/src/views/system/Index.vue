<template>
  <div class="system">
    <iframe src="http://localhost:8761" />
  </div>
</template>

<script>
export default {
  name: "Index",
};
</script>

<style scoped>
.system iframe {
  width: 100%;
  height: 100vh;
  border: none;
}
</style>