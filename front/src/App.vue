<template>
  <div id="app">
    <!-- <router-link to="/">Home</router-link>
    <router-link to="/about">About</router-link>
    <router-link to="/formulario">About</router-link> -->

    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
/* Estilos para o componente App.vue */
#app {
  font-family: Arial, sans-serif;
  text-align: center;
  margin-top: 20px;
}

router-link {
  margin-right: 10px;
}
</style>

