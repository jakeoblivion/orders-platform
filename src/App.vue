<script>
import OrderQuantity from "./components/OrderQuantity.vue";
import TotalOrders from "./components/TotalOrders.vue";
import PlatformConstants from './platformConstants';
import Vue from 'vue'
import VueMaterial from 'vue-material'
import axios from 'axios'
import 'vue-material/dist/vue-material.min.css'

Vue.use(VueMaterial)

export default {
  name: "app",
  components: {
    OrderQuantity,
    TotalOrders
  },
  methods: {
    calculateTotal() {
      return this.platforms.reduce((sum, item) => { return sum + item.quantity}, 0);
    }
  },
  data() {
    return {
      platforms : [
        {platform: 'etsy', quantity: 1},
        {platform: 'ebay', quantity: 3},
        {platform: 'woo', quantity: 6},
      ],
      response: null
    }
  },
  mounted() {
    axios
      .get('http://localhost:3000/get-orders')
      .then(response => this.response = response.data)
      .catch(err => console.log(err))
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
        <TotalOrders :quantity=calculateTotal() />
        <OrderQuantity v-for="item in platforms" :platform="item.platform" :quantity="item.quantity" />
      </div>
      <div>{{response}}</div>
    </main>
  </div>
</template>
