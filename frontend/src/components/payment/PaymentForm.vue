// File: frontend/src/components/payment/PaymentForm.vue
<template>
  <form @submit.prevent="submit" class="space-y-6">
    <div class="bg-gray-50 p-6 rounded-lg">
        <h3 class="text-lg font-semibold mb-4">Shipping Address</h3>
        <div class="grid grid-cols-1 gap-4">
            <div>
                <label for="street" class="form-label">Street Address</label>
                <input id="street" v-model="form.shipping_address.street" type="text" required class="form-input">
            </div>
             <div>
                <label for="city" class="form-label">City</label>
                <input id="city" v-model="form.shipping_address.city" type="text" required class="form-input">
            </div>
            <div>
                <label for="country" class="form-label">Country</label>
                <input id="country" v-model="form.shipping_address.country" type="text" required class="form-input">
            </div>
        </div>
    </div>
    <div class="bg-gray-50 p-6 rounded-lg">
        <h3 class="text-lg font-semibold mb-4">Payment Information</h3>
         <div>
            <label for="phone" class="form-label">M-Pesa Phone Number</label>
            <input id="phone" v-model="form.phone_number" type="tel" required placeholder="2547..." class="form-input">
        </div>
    </div>
    <button type="submit" class="w-full text-center inline-block bg-gray-800 text-white font-semibold py-3 px-8 rounded-lg hover:bg-gray-900 transition duration-300">
        Place Order & Pay
    </button>
  </form>
</template>

<script setup>
import { ref } from 'vue';
import { useAuthStore } from '../../store/modules/auth';

const emit = defineEmits(['submit-payment']);
const authStore = useAuthStore();

const form = ref({
    shipping_address: {
        street: '',
        city: '',
        country: '',
        zip_code: '00100', // Default or make it an input
    },
    billing_address: {}, // Will copy from shipping
    phone_number: authStore.user?.phone || ''
});

const submit = () => {
    // Copy shipping to billing address for simplicity
    form.value.billing_address = form.value.shipping_address;
    emit('submit-payment', form.value);
};
</script>

<style scoped>
.form-label {
    @apply block text-sm font-medium text-gray-700 mb-1;
}
.form-input {
    @apply block w-full rounded-md border-gray-300 shadow-sm focus:border-teal-500 focus:ring-teal-500 sm:text-sm;
}
</style>