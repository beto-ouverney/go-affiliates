<style lang="scss" scoped>
@import "./styles/FiltersField";
</style>
<template>
  <button type="button" @click="$emit('clickBtn')">
    {{ labelBtn }}
  </button>
</template>

<script>
export default {
  name: "BtnLoadSales",
  props: {
    labelBtn: String,
  },
};
</script>
