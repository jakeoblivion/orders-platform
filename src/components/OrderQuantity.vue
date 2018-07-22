<template>
  <div class="orderQuantity">
    <h1>{{ platform }}</h1>
    <p>
      {{ quantity }}
    </p>
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
.orderQuantity {
  display: inline-block;
  width: 25%;
}
</style>
