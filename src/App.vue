<script>
import OrderQuantity from "./components/OrderQuantity.vue";
import TotalOrders from "./components/TotalOrders.vue";
import PlatformConstants from './platformConstants';
import Vue from 'vue'
import VueMaterial from 'vue-material'
import 'vue-material/dist/vue-material.min.css'

Vue.use(VueMaterial)

export default {
  //props: {quantity: 123},
  name: "app",
  components: {
    OrderQuantity,
    TotalOrders
  },
  data() {
    return {
      platforms : [
        {platform: 'etsy', quantity: 1},
        {platform: 'ebay', quantity: 3},
        {platform: 'woo', quantity: 6},
      ],
      quantity: 6
    }
  }
};
</script>

<template>
  <div id="app">
    <nav class="logoContainer">
      <img class="logo" src="./assets/logo.jpg">
    </nav>
    <main class="orderQuantityContainer">
      <div class="md-layout md-gutter md-alignment-center">
        <TotalOrders :quantity=quantity />
        <OrderQuantity v-for="item in platforms" :platform="item.platform" :quantity="item.quantity" />
      </div>
      <!-- <span v-for="item in platforms">
        {{v-bind:qty="item.qty"}}
      </span> -->


    </main>
  </div>
</template>
