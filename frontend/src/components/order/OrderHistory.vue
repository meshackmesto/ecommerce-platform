// File: frontend/src/components/order/OrderHistory.vue
<template>
    <div class="bg-white shadow-lg rounded-lg p-8">
        <h2 class="text-2xl font-semibold mb-6">Your Order History</h2>
        <div v-if="orderStore.loading">
            <Loading message="Fetching your orders..." />
        </div>
        <div v-else-if="orderStore.orders.length === 0" class="text-center text-gray-500 py-8">
            <p>You haven't placed any orders yet.</p>
        </div>
        <div v-else class="space-y-4">
            <div v-for="order in orderStore.orders" :key="order.id" class="border rounded-lg p-4 flex justify-between items-center">
                <div>
                    <p class="font-bold text-gray-800">Order #{{ order.order_number }}</p>
                    <p class="text-sm text-gray-500">Placed on {{ new Date(order.created_at).toLocaleDateString() }}</p>
                </div>
                <div class="text-right">
                    <p class="font-semibold">{{ formatCurrency(order.total_amount) }}</p>
                    <OrderStatus :status="order.status" />
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { onMounted } from 'vue';
import { useOrderStore } from '../../store/modules/orders';
import { formatCurrency } from '../../utils/helpers';
import OrderStatus from './OrderStatus.vue';
import Loading from '../common/Loading.vue';

const orderStore = useOrderStore();

onMounted(() => {
    orderStore.fetchOrders();
});
</script>