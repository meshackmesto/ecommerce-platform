// File: frontend/src/components/order/OrderHistory.vue
<template>
    <div class="bg-card shadow-custom-md rounded-lg">
        <div class="p-6 border-b border-border">
            <h2 class="text-2xl font-semibold text-card-foreground">Order History</h2>
        </div>
        <div v-if="orderStore.loading">
            <Loading message="Fetching your orders..." />
        </div>
        <div v-else-if="orderStore.orders.length === 0" class="text-center text-muted-foreground py-16 px-6">
            <p class="mb-4">You haven't placed any orders yet.</p>
            <router-link to="/products" class="btn-primary">
                Start Shopping
            </router-link>
        </div>
        <div v-else class="divide-y divide-border">
            <div v-for="order in orderStore.orders" :key="order.id" class="p-6 hover:bg-secondary transition-colors">
                <div class="flex flex-col sm:flex-row justify-between sm:items-center">
                    <div>
                        <p class="font-bold text-primary text-lg">Order #{{ order.order_number }}</p>
                        <p class="text-sm text-muted-foreground mt-1">Placed on {{ new Date(order.created_at).toLocaleDateString() }}</p>
                    </div>
                    <div class="mt-4 sm:mt-0 flex items-center space-x-6">
                         <div class="text-right">
                            <p class="font-semibold text-foreground">{{ formatCurrency(order.total_amount) }}</p>
                            <OrderStatus :status="order.status" class="mt-1" />
                        </div>
                        <button class="btn-secondary !py-1 !px-3 text-sm">View Details</button>
                    </div>
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