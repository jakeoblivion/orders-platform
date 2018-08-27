<template>
  <div class="md-layout-item md-medium-size-25 md-small-size-50 md-xsmall-size-50">
    <md-card class="orderQuantityContainer center primaryBackground cursor">
      <md-card-header>
        <div class="md-title">
          Total Orders
        </div>
      </md-card-header>
      <md-card-content class=orderQuantityNumber>
        {{ quantity }}
      </md-card-content>
    </md-card>
  </div>
</template>

<script>
export default {
  name: "TotalOrders",
  props: {
    quantity: Number
  }
};
</script>

<style scoped>
.orderQuantityContainer {
  margin-top: 8px;
  padding: 12px;
}

.orderQuantityNumber {
  font-size: 40px;
  padding-top: 5px;
}
</style>
