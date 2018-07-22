<script>
import OrderQuantity from "./components/OrderQuantity.vue";
import TotalOrders from "./components/TotalOrders.vue";

export default {
  name: "app",
  components: {
    OrderQuantity,
    TotalOrders
  }
};
</script>

<template>
  <div id="app">
    <nav class="logoContainer">
      <img class="logo" src="./assets/logo.jpg">
    </nav>
    <main class="orderQuantityContainer">
      <OrderQuantity platform="Ebay" quantity=1 />
      <OrderQuantity platform="Etsy" quantity=2 />
      <OrderQuantity platform="Woo" quantity=3 />
      <TotalOrders quantity=6 />
    </main>
  </div>
</template>
