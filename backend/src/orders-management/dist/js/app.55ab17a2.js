(function(t){function e(e){for(var i,n,o=e[0],l=e[1],d=e[2],u=0,m=[];u<o.length;u++)n=o[u],s[n]&&m.push(s[n][0]),s[n]=0;for(i in l)Object.prototype.hasOwnProperty.call(l,i)&&(t[i]=l[i]);c&&c(e);while(m.length)m.shift()();return a.push.apply(a,d||[]),r()}function r(){for(var t,e=0;e<a.length;e++){for(var r=a[e],i=!0,o=1;o<r.length;o++){var l=r[o];0!==s[l]&&(i=!1)}i&&(a.splice(e--,1),t=n(n.s=r[0]))}return t}var i={},s={0:0},a=[];function n(e){if(i[e])return i[e].exports;var r=i[e]={i:e,l:!1,exports:{}};return t[e].call(r.exports,r,r.exports,n),r.l=!0,r.exports}n.m=t,n.c=i,n.d=function(t,e,r){n.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:r})},n.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},n.t=function(t,e){if(1&e&&(t=n(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var r=Object.create(null);if(n.r(r),Object.defineProperty(r,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var i in t)n.d(r,i,function(e){return t[e]}.bind(null,i));return r},n.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return n.d(e,"a",e),e},n.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},n.p="/";var o=window["webpackJsonp"]=window["webpackJsonp"]||[],l=o.push.bind(o);o.push=e,o=o.slice();for(var d=0;d<o.length;d++)e(o[d]);var c=l;a.push([0,1]),r()})({0:function(t,e,r){t.exports=r("Vtdi")},"0aNq":function(t,e,r){"use strict";var i=r("biQ0"),s=r.n(i);s.a},NN4i:function(t,e,r){},OMi8:function(t,e,r){},VrzN:function(t,e,r){"use strict";var i=r("vBkZ"),s=r.n(i);s.a},Vtdi:function(t,e,r){"use strict";r.r(e);r("yt8O"),r("VRzm");var i=r("Kw5r"),s=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{attrs:{id:"app"}},[t._m(0),r("main",[r("div",{staticClass:"orderQuantityContainer md-layout md-gutter md-alignment-center"},[r("TotalOrders",{attrs:{quantity:t.calculateTotalOrders()},nativeOn:{click:function(e){return t.showAllOrders(e)}}}),t._l(t.platforms,function(e){return r("OrderQuantity",{attrs:{platform:e.platform,quantity:t.getOrderQuantityForPlatform(e.platform)},nativeOn:{click:function(r){t.hidePlatformOrders(e.platform)}}})})],2),r("h1",{staticClass:"ordersHeading"},[t._v("ORDERS")]),r("div",{staticClass:"ordersContainer md-layout md-gutter"},t._l(t.orders,function(e){return r("Order",{class:[{hidden:t.platforms.find(function(t){return t.platform===e.platform}).hidden}],attrs:{orderDetails:e}})})),t.orders.length?r("ShippingOptions"):t._e()],1)])},a=[function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("nav",{staticClass:"logoContainer"},[i("img",{staticClass:"logo",attrs:{src:r("tkAB")}})])}],n=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"md-layout-item md-medium-size-25 md-small-size-50 md-xsmall-size-50"},[r("md-card",{staticClass:"orderQuantityContainer center greybg cursor",attrs:{name:"orderQuantityContainer"}},[r("md-card-header",[r("div",{staticClass:"md-title"},[r("img",{staticClass:"platformLogo",attrs:{src:"/static/"+t.platform+"-logo.png"}})])]),r("md-card-content",{staticClass:"orderQuantityNumber"},[t._v("\n      "+t._s(t.quantity)+"\n    ")])],1)],1)},o=[],l=(r("xfY5"),{name:"OrderQuantity",props:{platform:String,quantity:Number}}),d=l,c=(r("uvrk"),r("KHd+")),u=Object(c["a"])(d,n,o,!1,null,"06393f21",null),m=u.exports,p=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"md-layout-item md-medium-size-25 md-small-size-50 md-xsmall-size-50"},[r("md-card",{staticClass:"orderQuantityContainer center primaryBackground cursor"},[r("md-card-header",[r("div",{staticClass:"md-title"},[t._v("\n        Total Orders\n      ")])]),r("md-card-content",{staticClass:"orderQuantityNumber"},[t._v("\n      "+t._s(t.quantity)+"\n    ")])],1)],1)},f=[],g={name:"TotalOrders",props:{quantity:Number}},h=g,v=(r("jrW0"),Object(c["a"])(h,p,f,!1,null,"3e20ee9a",null)),y=v.exports,_=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"order md-layout-item md-large-size-50 md-medium-size-50 md-small-size-100",class:t.orderDetails.platform},[r("md-card",{staticClass:"lightgreybg"},[r("md-card-content",{staticClass:"md-layout p0"},[r("div",{staticClass:"md-layout-item md-size-66"},t._l(t.orderDetails.orderItems,function(e){return r("div",{staticClass:"md-layout"},[r("div",{staticClass:"md-layout-item md-size-40"},[r("img",{staticClass:"orderItemImg",attrs:{src:e.imageUrl}})]),r("div",{staticClass:"md-layout-item orderContent"},[r("div",{staticClass:"orderTitle"},[t._v(t._s(e.itemName.length>60?e.itemName.slice(0,60)+"...":e.itemName))]),r("div",[t._v("£"+t._s(e.itemPrice))])])])})),r("div",{staticClass:"md-layout-item md-size-33 shippingAddress"},[r("div",{staticClass:"platformImg"},[r("img",{attrs:{src:"/static/"+t.orderDetails.platform+"-logo.png"}}),r("span",{staticClass:"total"},[t._v("£"+t._s(t.orderDetails.orderTotal))])]),r("div",[t._v(t._s(t.orderDetails.shippingAddress.name))]),r("div",[t._v(t._s(t.orderDetails.shippingAddress.addressLine.length>20?t.orderDetails.shippingAddress.addressLine.slice(0,20)+"...":t.orderDetails.shippingAddress.addressLine)+", "+t._s(t.orderDetails.shippingAddress.city.length>12?t.orderDetails.shippingAddress.city.slice(0,12)+"...":t.orderDetails.shippingAddress.city))]),r("div",[t._v(t._s(t.orderDetails.shippingAddress.postCode)+", "+t._s(t.orderDetails.shippingAddress.country))]),r("div",{staticClass:"shipByDate",class:{orange:t.isOverdue(t.orderDetails.shipByDate)}},[t._v("Ship by: "+t._s(t._f("moment")(t.orderDetails.shipByDate,"ddd, MMM Do")))])])])],1)],1)},O=[],b={name:"Order",props:{orderDetails:{}},methods:{isOverdue:function(t){return new Date(t)<new Date}}},C=b,w=(r("0aNq"),Object(c["a"])(C,_,O,!1,null,"f92c2684",null)),D=w.exports,z=function(){var t=this,e=t.$createElement;t._self._c;return t._m(0)},j=[function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"md-layout"},[r("div",{staticClass:"shippingOption md-layout-item md-large-size-33 md-medium-size-33 md-small-size-100"},[r("a",{attrs:{href:"https://business.parcel.royalmail.com/orders/",target:"_blank"}},[r("div",{staticClass:"shipping-link"},[r("p",[t._v("Print Shipping Labels")]),r("img",{attrs:{src:"https://parcel.royalmail.com/Themes/RoyalMail/Images/royal-mail-logo-large.png"}})])])]),r("div",{staticClass:"shippingOption md-layout-item md-large-size-33 md-medium-size-33 md-small-size-100"},[r("a",{attrs:{href:"https://www.paypal.com/businessexp/transactions?tab=activity&transactiontype=PAYMENTS_RECEIVED",target:"_blank"}},[r("div",{staticClass:"shipping-link"},[r("p",[t._v("Print Shipping Labels")]),r("img",{attrs:{src:"https://www.paypalobjects.com/webstatic/i/logo/rebrand/ppcom-white.svg"}})])])]),r("div",{staticClass:"shippingOption md-layout-item md-large-size-33 md-medium-size-33 md-small-size-100"},[r("a",{attrs:{href:"https://www.parcel2go.com/myaccount/uploadeditems",target:"_blank"}},[r("div",{staticClass:"shipping-link"},[r("p",[t._v("Print Shipping Labels")]),r("img",{staticStyle:{"padding-top":"20px"},attrs:{src:"https://cdn.parcel2go.com/42c5534b-0f47-4342-80f0-8f21203f0669/layout/p2g_logo_new.svg"}})])])])])}],S={name:"ShippingOptions"},k=S,x=(r("VrzN"),Object(c["a"])(k,z,j,!1,null,"3f8ba0ce",null)),A=x.exports,E=(r("DW2E"),Object.freeze({EBAY:"ebay",ETSY:"etsy",WOO:"woo"})),N=r("Q/mQ"),P=r.n(N),Q=r("Lq01"),T=r.n(Q),q=r("vDqi"),M=r.n(q);r("Ud79");i["default"].use(P.a),i["default"].use(T.a);var B={name:"app",components:{OrderQuantity:m,TotalOrders:y,Order:D,ShippingOptions:A},methods:{calculateTotalOrders:function(){return this.orders.reduce(function(t){return++t},0)},getOrderQuantityForPlatform:function(t){return this.orders.reduce(function(e,r){return r.platform===t?++e:e},0)},hidePlatformOrders:function(t){this.platforms.map(function(e){e.platform===t?e.hidden=!1:e.hidden=!0})},showAllOrders:function(){this.platforms.map(function(t){t.hidden=!1})}},data:function(){return{platforms:[{hidden:!1,platform:E.ETSY},{hidden:!1,platform:E.EBAY},{hidden:!1,platform:E.WOO}],orders:[]}},mounted:function(){var t=this;M.a.get("http://localhost:3001/get-orders").then(function(e){return t.orders=e.data}).catch(function(t){return console.log(t)})}},L=B,$=Object(c["a"])(L,s,a,!1,null,null,null),I=$.exports;r("OMi8");i["default"].config.productionTip=!1,new i["default"]({render:function(t){return t(I)}}).$mount("#app")},biQ0:function(t,e,r){},jrW0:function(t,e,r){"use strict";var i=r("q59+"),s=r.n(i);s.a},"q59+":function(t,e,r){},tkAB:function(t,e,r){t.exports=r.p+"img/logo.d0d78388.jpg"},uvrk:function(t,e,r){"use strict";var i=r("NN4i"),s=r.n(i);s.a},vBkZ:function(t,e,r){}});
//# sourceMappingURL=app.55ab17a2.js.map