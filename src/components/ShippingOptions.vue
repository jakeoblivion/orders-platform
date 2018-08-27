<template>
  <div class="md-layout shippingOptions">
    <div class="md-layout-item md-large-size-33 md-medium-size-33 md-small-size-100">
      <a href="https://business.parcel.royalmail.com/orders/" target="_blank">
        <div class="shipping-link">
          <p>Print Shipping Labels</p>
          <img src="https://parcel.royalmail.com/Themes/RoyalMail/Images/royal-mail-logo-large.png">
        </div>
      </a>
    </div>
    <div class="md-layout-item md-large-size-33 md-medium-size-33 md-small-size-100">
      <a href="https://www.paypal.com/businessexp/transactions?tab=activity&transactiontype=PAYMENTS_RECEIVED" target="_blank">
        <div class="shipping-link">
          <p>Print Shipping Labels</p>
          <img src="https://www.paypalobjects.com/webstatic/i/logo/rebrand/ppcom-white.svg">
        </div>
      </a>
    </div>
    <div class="md-layout-item md-large-size-33 md-medium-size-33 md-small-size-100">
      <a href="https://www.parcel2go.com/myaccount/uploadeditems" target="_blank">
        <div class="shipping-link">
          <p>Print Shipping Labels</p>
          <img style="padding-top: 20px" src="https://cdn.parcel2go.com/42c5534b-0f47-4342-80f0-8f21203f0669/layout/p2g_logo_new.svg">
        </div>
      </a>
    </div>
  </div>
</template>

<script>
export default {
  name: "ShippingOptions",
};
</script>

<style scoped>
  .shippingOptions {
    margin-top: 24px;
  }

  .shipping-link {
    border: solid 1px rgb(191, 191, 191);
    border-radius: 20px;
    background: rgb(155, 181, 172);
    position: relative;
    display: inline-block;
    width: calc(100% - 30px);
    text-align: center;
    min-height: 175px;
  }

  .shipping-link img {
    margin: 0 auto;
    padding: 10px 25px;
    max-width: 100%;
    max-height: 80px;
    display: block;
  }

  .shipping-link {
    font-size: 28px;
    color: #2c3e50;
  }

  .shipping-link:hover {
    color: white;
  }

</style>
