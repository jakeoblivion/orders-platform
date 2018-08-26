<template xmlns:src="http://www.w3.org/1999/xhtml">
  <div class="orderContainer md-layout-item md-small-size-100">
    <md-card class="">
      <md-card-content class="md-layout p0">
        <div class="md-layout-item md-size-66">
          <div class="md-layout" v-for="orderItem of orderDetails.orderItems">
            <div class="md-layout-item md-size-40">
              <img class="orderItemImg" :src="orderItem.imageUrl" />
            </div>

            <div class="md-layout-item orderContent">
                <div class="orderTitle">{{orderItem.itemName.length > 60 ? orderItem.itemName.slice(0, 60) + '...' : orderItem.itemName}}</div>
                <div>£{{orderItem.itemPrice}}</div>
                <!--<div class="orderedDate">Ordered: {{ orderDetails.orderDate | moment("ddd, MMMM Do") }}</div>-->
            </div>

          </div>
        </div>

        <div class="md-layout-item md-size-33 shippingAddress">
            <div class="platformImg">
              <img :src="'/static/'+orderDetails.platform+'-logo.jpg'" />
              <span class="total">£{{orderDetails.orderTotal}}</span>
            </div>

            <div>{{orderDetails.shippingAddress.name}}</div>
            <div>{{orderDetails.shippingAddress.addressLine}}, {{orderDetails.shippingAddress.city}}</div>
            <div>{{orderDetails.shippingAddress.postCode}}, {{orderDetails.shippingAddress.country}}</div>
            <div class="shipByDate" v-bind:class="{ 'orange': isOverdue(orderDetails.shipByDate) }">Ship by: {{ orderDetails.shipByDate | moment("ddd, MMM Do") }}</div>

        </div>
      </md-card-content>
    </md-card>
  </div>
</template>

<script>
export default {
  name: "Order",
  props: {
    orderDetails: {}
  },
  methods: {
    isOverdue: function (date) {
      return new Date(date) < new Date()
    }
  }
}


</script>

<style scoped>
    .orderContainer {
        padding-top: 8px;
        padding-bottom: 8px;
    }

.orderTitle {
  font-size: 16px;
  font-weight: bold;
}

.total {
  color: #9bb5ab;
  font-weight: 800;
  font-size: 20px;
  top: 4px;
    left: 4px;
  padding-right: 4px;
  position: relative;
}

.shipByDate {
    padding-top: 8px;
    font-weight: bold;
}

.orange {
    color: orangered;
}

.orderedDate {
  font-weight: bold;
}

.platformImg {
  border: solid 1px lightgrey;
  border-radius: 4px;
  padding: 5px;
  margin-bottom: 5px;
    margin-left: -5px;
  background: white;
  display: inline-block;
}

.platformImg img {
  width: 100%;
  max-width: 75px;
}

.orderItemImg {
  width: 100%;
}

.orderContent {
  padding: 16px 8px;
  margin-right: 4%;
  border-right: dashed 2px #f1f1f1;
}

.shippingAddress {
  padding: 16px 8px;
}
</style>
