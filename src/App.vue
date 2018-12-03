<!--suppress UnterminatedStatementJS -->
<script>
import OrderQuantity from "./components/OrderQuantity.vue";
import TotalOrders from "./components/TotalOrders.vue";
import Order from "./components/Order.vue";
import ShippingOptions from "./components/ShippingOptions.vue";
import PlatformConstants from "./platformConstants";
import Vue from "vue";
import VueMaterial from "vue-material";
import Moment from "vue-moment";
import axios from "axios";
import "vue-material/dist/vue-material.min.css";

Vue.use(VueMaterial);
Vue.use(Moment);

export default {
  name: "app",
  components: {
    OrderQuantity,
    TotalOrders,
    Order,
    ShippingOptions,
  },
  methods: {
    calculateTotalOrders() {
      return this.orders.reduce(sum => {
        return ++sum;
      }, 0);
    },
    getOrderQuantityForPlatform(platform) {
      return this.orders.reduce((sum, order) => {
        return order.platform === platform ? ++sum : sum;
      }, 0);
    },
    hidePlatformOrders(platformClicked) {
      this.platforms.map(item => {
        item.platform === platformClicked ? item.hidden = false : item.hidden = true
      })
    },
    showAllOrders() {
        this.platforms.map(item => {
            item.hidden = false
        })
    },
      sortOrdersByShipByDate(orders) {
      return orders.sort((a,b) => {
              return new Date(a.shipByDate) - new Date(b.shipByDate)
          })
      }
  },
  data() {
    return {
      platforms: [
        { hidden: false, platform: PlatformConstants.ETSY},
        { hidden: false, platform: PlatformConstants.EBAY},
        { hidden: false, platform: PlatformConstants.WOO },
      ],
      orders: [],
    };
  },
  mounted() {
    axios
      .get("/get-orders")
      .then(response => (this.orders = this.sortOrdersByShipByDate(response.data) || []))
      .catch(err => console.log(err))
  }
}
</script>

<template>
  <div id="app">
    <nav class="logoContainer">
      <img class="logo" src="./assets/logo.jpg">
    </nav>
    <main>
      <div class="orderQuantityContainer md-layout md-gutter md-alignment-center">
        <TotalOrders @click.native=showAllOrders :quantity=calculateTotalOrders() />
        <OrderQuantity @click.native=hidePlatformOrders(item.platform) v-for="item in platforms" :platform="item.platform" :quantity=getOrderQuantityForPlatform(item.platform) />
      </div>
      <h1 class="ordersHeading">ORDERS</h1>
      <div class="ordersContainer md-layout md-gutter">
        <Order :class="[{'hidden': platforms.find(item => item.platform === order.platform).hidden}]" v-for="order in orders" :orderDetails="order" />
      </div>
      <ShippingOptions v-if="!!orders.length" />

    </main>
  </div>
</template>
