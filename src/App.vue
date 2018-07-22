<script>
import OrderQuantity from "./components/OrderQuantity.vue";

export default {
  name: "app",
  components: {
    OrderQuantity
  }
};
</script>

<template>
  <div id="app">
    <nav class="logoContainer">
      <img class="logo" src="./assets/logo.jpg">
    </nav>
    <main class="orderQuantityContainer">
      <OrderQuantity type="Ebay"/>
      <OrderQuantity type="Etsy"/>
      <OrderQuantity type="Woo"/>
      <TotalOrders
    </main>
  </div>
</template>
