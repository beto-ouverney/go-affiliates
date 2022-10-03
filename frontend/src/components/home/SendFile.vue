<style lang="scss" scoped>
@import "./styles/SendFile";
</style>
<template>
  <div class="container">
    <h2>Send Sales File</h2>

    <input type="file" @change="handleFileUpload($event)" />

    <button class="btn-send" v-on:click="submitFile()">Submit</button>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "SendFile",
  data() {
    return {
      file: "",
    };
  },

  methods: {
    handleFileUpload(event) {
      this.file = event.target.files[0];
    },

    submitFile() {
      let formData = new FormData();

      formData.append("file", this.file);

      axios
        .post("/api/v1/sales/upload", formData, {
          headers: {
            "Content-Type": "multipart/form-data",
          },
        })
        .then((response) => {
          const results = response.data;
          alert(results.message);
        })
        .catch(function (r) {
          const { response } = r;
          alert(response.data.message);
        });
    },
  },
};
</script>
