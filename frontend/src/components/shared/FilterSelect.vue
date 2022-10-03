<template>
  <div>
    <label
      >{{ label
      }}<select v-model="value" @change="handleChange">
        <option v-for="item in array" :key="item" :value="item">
          {{ item }}
        </option>
      </select>
    </label>
  </div>
</template>

<script>
export default {
  name: "FilterSelect",
  props: {
    label: String,
    array: Array,
  },
  data() {
    return {
      value: this.array[0],
    };
  },
  watch: {
    array() {
      this.value = this.array[0];
    },
  },
  methods: {
    handleChange() {
      console.log(this.value);
      this.$emit("update", this.value);
    },
  },
};
</script>
