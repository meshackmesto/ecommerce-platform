// File: frontend/src/components/admin/OrderManagement.vue
<template>
    <div>
        <h1 class="text-3xl font-bold text-gray-800 mb-6">Order Management</h1>

        <div class="bg-white rounded-lg shadow overflow-x-auto">
             <div v-if="orderStore.loading">
                <Loading message="Fetching all orders..." />
            </div>
            <table v-else class="min-w-full divide-y divide-gray-200">
                <thead class="bg-gray-50">
                    <tr>
                        <th class="th">Order #</th>
                        <th class="th">Date</th>
                        <th class="th">Customer ID</th>
                        <th class="th">Total</th>
                        <th class="th">Status</th>
                        <th class="th">Actions</th>
                    </tr>
                </thead>
                <tbody class="bg-white divide-y divide-gray-200">
                    <tr v-if="orderStore.allOrders.length === 0">
                        <td colspan="6" class="py-12 text-center text-gray-500">No orders found.</td>
                    </tr>
                    <tr v-for="order in orderStore.allOrders" :key="order.id">
                        <td class="td font-mono">{{ order.order_number }}</td>
                        <td class="td">{{ new Date(order.created_at).toLocaleString() }}</td>
                        <td class="td">{{ order.user_id }}</td>
                        <td class="td">{{ formatCurrency(order.total_amount) }}</td>
                        <td class="td"><OrderStatus :status="order.status" /></td>
                        <td class="td">
                            <button class="text-teal-600 hover:text-teal-900">View</button>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>

<script setup>
import { onMounted } from 'vue';
import { useOrderStore } from '../../store/modules/orders';
import { formatCurrency } from '../../utils/helpers';
import Loading from '../common/Loading.vue';
import OrderStatus from '../order/OrderStatus.vue';

const orderStore = useOrderStore();

onMounted(() => {
    orderStore.fetchAllOrders();
});
</script>

<style scoped>
.th {
    @apply px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider;
}
.td {
    @apply px-6 py-4 whitespace-nowrap text-sm text-gray-700;
}
</style>