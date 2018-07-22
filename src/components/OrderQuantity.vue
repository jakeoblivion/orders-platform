<template>
  <div class="orderQuantityContainer">
    <div class="orderQuantity">
    <img class="platformLogo" src="../assets/ebay-logo.jpg">
    <p>
      {{ quantity }}
    </p>
  </div>
</div>
</template>

<script>
export default {
  name: "OrderQuantity",
  props: {
    platform: String,
    quantity: String
  }
};
</script>

<style scoped>
.orderQuantityContainer {
  display: inline-block;
  width: 25%;
  padding: 15px;
}
.orderQuantity {
  border: solid 1px grey;
}

.platformLogo {
  width: 100px;
  padding: 5px;
  max-width: 100%;
}
</style>
