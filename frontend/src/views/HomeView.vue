<style lang="scss" scoped>
@import "./styles/HomeView.scss";
</style>
<template>
  <div class="top-home">
    <h1>Go Affiliates</h1>
    <div class="navigation">
      <SendFile />
      <FilterComponent />
    </div>
    <TableSales :show="showTable" />
    <LoadingComponent :show="isLoading" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import FilterComponent from "@/components/home/FilterComponent.vue";
import SendFile from "@/components/home/SendFile.vue";
import TableSales from "@/components/home/TableSales.vue";
import { mapState } from "vuex";
import LoadingComponent from "@/components/shared/LoadingComponent.vue";

export default defineComponent({
  name: "HomeView",
  components: {
    FilterComponent,
    SendFile,
    LoadingComponent,
    TableSales,
  },
  computed: {
    ...mapState(["isLoading", "showTable"]),
  },
});
</script>
