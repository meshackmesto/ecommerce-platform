// File: frontend/src/components/order/OrderStatus.vue
<template>
  <span :class="statusClass" class="px-2 py-1 text-xs font-semibold rounded-full capitalize">
    {{ status }}
  </span>
</template>

<script setup>
import { computed } from 'vue';

const props = defineProps({
  status: {
    type: String,
    required: true,
  },
});

const statusClass = computed(() => {
  switch (props.status.toLowerCase()) {
    case 'delivered':
    case 'completed':
      return 'bg-green-100 text-green-800';
    case 'shipped':
    case 'processing':
      return 'bg-blue-100 text-blue-800';
    case 'cancelled':
    case 'failed':
      return 'bg-red-100 text-red-800';
    case 'pending':
    default:
      return 'bg-yellow-100 text-yellow-800';
  }
});
</script>