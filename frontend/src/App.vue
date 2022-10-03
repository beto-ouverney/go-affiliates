<template>
  <router-view />
</template>

<style lang="scss">
* {
  font-family: "Poppins", sans-serif;
}

nav {
  display: flex;
  flex-direction: row;
  height: 5%;
  justify-content: center;

  a {
    color: black;
    font-weight: bold;
    margin: 5px;
    padding: 10px;

    &.router-link-exact-active {
      border: 1px solid rgb(8, 8, 209);
      border-radius: 10px;
      color: rgb(8, 8, 209);
    }
  }
}
</style>
