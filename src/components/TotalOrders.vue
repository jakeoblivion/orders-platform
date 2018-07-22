<template>
  <div class="totalOrders">
    <h1>Total Orders</h1>
    <p>
      {{ quantity }}
    </p>
  </div>
</template>

<script>
export default {
  name: "TotalOrders",
  props: {
    quantity: String
  }
};
</script>

<style scoped>
.totalOrders {
  display: inline-block;
  width: 25%;
}
</style>
