import { shallowMount } from "@vue/test-utils";
import OrderQuantity from "@/components/OrderQuantity.vue";
import PlatformConstants from "src/platformConstants";

describe("OrderQuantity.vue", () => {
  it("renders props.type when passed", () => {
    const type = PlatformConstants.ETSY;
    const wrapper = shallowMount(OrderQuantity, {
      propsData: { type }
    });
    expect(wrapper.text()).toMatch(type);
  });
});
