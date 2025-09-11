// File: frontend/src/components/order/OrderSummary.vue
<template>
  <div class="bg-gray-50 rounded-lg p-6">
    <h3 class="text-xl font-semibold mb-4">Summary</h3>
    <div class="space-y-2">
      <div class="flex justify-between">
        <span class="text-gray-600">Subtotal</span>
        <span class="font-medium">{{ formatCurrency(cartStore.cartTotal) }}</span>
      </div>
      <div class="flex justify-between">
        <span class="text-gray-600">Shipping</span>
        <span class="font-medium">TBD</span>
      </div>
      <div class="flex justify-between">
        <span class="text-gray-600">Taxes</span>
        <span class="font-medium">TBD</span>
      </div>
    </div>
    <div class="border-t my-4"></div>
    <div class="flex justify-between text-lg font-bold">
      <span>Total</span>
      <span>{{ formatCurrency(cartStore.cartTotal) }}</span>
    </div>
  </div>
</template>

<script setup>
import { useCartStore } from '../../store/modules/cart';
import { formatCurrency } from '../../utils/helpers';

const cartStore = useCartStore();
</script>