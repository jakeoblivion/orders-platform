
<template>
    <div class="md-layout-item md-medium-size-25 md-small-size-50 md-xsmall-size-50">
    <md-card name="orderQuantityContainer" class="orderQuantityContainer center greybg cursor">
      <md-card-header>
        <div class="md-title">
          <img class="platformLogo" :src="'/static/'+platform+'-logo.png'" >
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
  name: "OrderQuantity",
  props: {
    platform: String,
    quantity: Number
  },
};
</script>

<style scoped>
.orderQuantityContainer {
  margin-top: 8px;
}

.orderQuantityNumber {
  font-size: 40px;
}

.platformLogo {
  width: 100px;
  padding: 5px;
  max-width: 100%;
}
</style>
