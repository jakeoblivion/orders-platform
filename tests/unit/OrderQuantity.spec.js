import { shallowMount } from "@vue/test-utils";
import OrderQuantity from "@/components/OrderQuantity.vue";

describe("OrderQuantity.vue", () => {
  it("renders props.type when passed", () => {
    const type = "Etsy";
    const wrapper = shallowMount(OrderQuantity, {
      propsData: { type }
    });
    expect(wrapper.text()).toMatch(type);
  });
});
