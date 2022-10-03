<template>
  <div v-show="show">
    <p>Loading...</p>
  </div>
</template>
<script>
export default {
  props: {
    show: Boolean,
  },
};
</script>
