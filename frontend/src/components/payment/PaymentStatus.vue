// File: frontend/src/components/payment/PaymentStatus.vue
<template>
  <div class="p-6 rounded-lg text-center" :class="statusInfo.bgClass">
    <div class="flex justify-center items-center mb-4">
        <component :is="statusInfo.icon" class="h-12 w-12" :class="statusInfo.textClass" />
    </div>
    <h3 class="text-xl font-bold" :class="statusInfo.textClass">{{ statusInfo.title }}</h3>
    <p class="mt-2" :class="statusInfo.subtextClass">{{ statusInfo.message }}</p>
  </div>
</template>

<script setup>
import { computed } from 'vue';
import { ClockIcon, CheckCircleIcon, XCircleIcon } from '@heroicons/vue/24/solid';

const props = defineProps({
  status: {
    type: String, // 'pending', 'success', 'failed'
    required: true,
  },
});

const statusInfo = computed(() => {
  switch (props.status) {
    case 'success':
      return {
        title: 'Payment Successful!',
        message: 'Your order is confirmed. Redirecting...',
        icon: CheckCircleIcon,
        bgClass: 'bg-green-50',
        textClass: 'text-green-600',
        subtextClass: 'text-green-800'
      };
    case 'failed':
      return {
        title: 'Payment Failed',
        message: 'Please try again or contact support.',
        icon: XCircleIcon,
        bgClass: 'bg-red-50',
        textClass: 'text-red-600',
        subtextClass: 'text-red-800'
      };
    case 'pending':
    default:
      return {
        title: 'Awaiting Payment...',
        message: 'Please check your phone and enter your M-Pesa PIN to complete the payment.',
        icon: ClockIcon,
        bgClass: 'bg-yellow-50',
        textClass: 'text-yellow-600',
        subtextClass: 'text-yellow-800'
      };
  }
});
</script>