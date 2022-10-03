<style lang="scss" scoped>
@import "./styles/FiltersField";
</style>
<template>
  <div class="container">
    <FilterSelect
      :array="producersArray"
      :label="labelFilterProducers"
      @update="producer = $event"
    />
    <BtnLoadSales
      class="btn-send"
      :labelBtn="labelBtnLoadSales"
      @clickBtn="getSales()"
    />
  </div>
</template>
<script>
import FilterSelect from "../shared/FilterSelect";
import { mapActions, mapGetters } from "vuex";
import BtnLoadSales from "@/components/home/BtnLoadSales";
import axios from "axios";

export default {
  name: "FilterComponent",
  components: {
    FilterSelect,
    BtnLoadSales,
  },
  data() {
    return {
      labelBtnLoadSales: "Get Sales",
      labelFilterProducers: "Producer: ",
      producersArray: ["No Producers"],
      producer: "No Producers",
    };
  },
  computed: {
    ...mapGetters({
      data: "getData",
      sales: "getSales",
      producerStore: "getProducer",
    }),
  },
  methods: {
    ...mapActions([
      "setLoading",
      "setSales",
      "setShowTable",
      "setData",
      "setProducer",
    ]),
    ...mapGetters(["getSales", "getData"]),
    getSales() {
      this.setLoading();
      axios
        .get("/api/v1/sales")
        .then((response) => {
          const results = response.data;
          const producersWithDuplicates = results.map(
            (result) => result.producer
          );
          this.producersArray = [...new Set(producersWithDuplicates)];
          this.producer = this.producersArray[0];
          this.setShowTable();
          this.setData(results);
          this.setSales(results);
          this.setProducer(this.producer);
          this.setLoading();
        })
        .catch(function (r) {
          console.log(r);
          this.setLoading();
          const { response } = r;
          alert(response.data.message);
        });
    },
  },
  watch: {
    producer() {
      this.setProducer(this.producer);
      if (this.data.length > 0) {
        const sales = this.data.filter((e) => e.producer === this.producer);
        this.setSales(sales);
      }
    },
  },
};
</script>
