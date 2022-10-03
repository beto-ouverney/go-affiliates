<style lang="scss" scoped>
@import "./styles/TableSales";
</style>
<template>
  <div v-show="show" class="div-plans">
    <h1>Sales Report</h1>
    <table class="table-plans">
      <thead>
        <tr>
          <th v-for="tHead in tHeads" :key="tHead">{{ tHead }}</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in arrayTable" :key="item.date">
          <td>{{ item.product }}</td>
          <td>{{ item.seller }}</td>
          <td>{{ item.value }}</td>
          <td>{{ item.commission }}</td>
          <td>{{ item.date }}</td>
          <td>{{ item.total }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
<script>
import { mapGetters } from "vuex";

export default {
  name: "TableSales",
  props: {
    show: Boolean,
  },
  data() {
    return {
      tHeads: [
        "Product",
        "Producer/Affiliate",
        "Value",
        "Commission",
        "Date",
        "Total",
      ],
      arrayTable: [],
    };
  },
  computed: {
    ...mapGetters({
      data: "getData",
      sales: "getSales",
      producer: "getProducer",
    }),
  },
  watch: {
    sales() {
      const array = [];
      for (let i = 0; i < this.sales.length; i++) {
        let seller = "";
        if (this.sales[i].affiliate !== "") {
          seller = this.sales[i].affiliate;
        } else {
          seller = this.sales[i].producer;
        }
        const total = this.sales[i].value - this.sales[i].commission;
        array.push({
          product: this.sales[i].product,
          seller: seller,
          value: this.sales[i].value,
          commission: this.sales[i].commission,
          date: this.sales[i].date,
          total: total,
        });
        this.arrayTable = array;
      }
    },
  },
};
</script>
