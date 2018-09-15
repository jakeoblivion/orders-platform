import { shallowMount } from "@vue/test-utils";
import OrderQuantity from "@/components/OrderQuantity.vue";

describe("OrderQuantity.vue", () => {
  it("renders props.type when passed", () => {
    const type = "etsy";
    const wrapper = shallowMount(OrderQuantity, {
      propsData: { platform: type, quantity: 1 }
    });
    expect(wrapper.text()).toMatch("1");
  });
});
