// File: frontend/src/components/admin/ProductManagement.vue
<template>
    <div>
        <div class="flex justify-between items-center mb-8">
            <h1 class="text-4xl font-bold text-foreground">Products</h1>
            <button class="flex items-center bg-primary text-primary-foreground font-semibold py-2 px-4 rounded-lg hover:bg-primary-hover transition-colors">
                <PlusIcon class="h-5 w-5 mr-2" />
                Add New Product
            </button>
        </div>

        <div class="space-y-4">
            <div v-if="productStore.loading"><Loading /></div>
            <div v-else v-for="product in productStore.products" :key="product.id" class="bg-card p-4 rounded-lg shadow-custom-md flex items-center justify-between">
                <div class="flex items-center">
                    <img :src="getPrimaryImage(product)" class="h-12 w-12 rounded-md object-cover mr-4">
                    <div>
                        <p class="text-sm font-semibold text-primary uppercase">{{ product.category || 'Uncategorized' }}</p>
                        <p class="text-lg font-bold text-foreground">{{ product.name }}</p>
                        <p class="text-sm text-muted-foreground">{{ product.description }}</p>
                    </div>
                </div>
                <div class="flex items-center space-x-8 text-sm text-center">
                    <div>
                        <p class="text-muted-foreground font-medium">PRICE</p>
                        <p class="text-foreground font-semibold">{{ formatCurrency(product.price) }}</p>
                    </div>
                     <div>
                        <p class="text-muted-foreground font-medium">STOCK</p>
                        <p class="text-foreground font-semibold">{{ product.stock_quantity }}</p>
                    </div>
                    <div>
                        <p class="text-muted-foreground font-medium">STATUS</p>
                        <p class="flex items-center text-foreground font-semibold">
                            <span class="h-2 w-2 rounded-full mr-2" :class="product.is_active ? 'bg-success' : 'bg-destructive'"></span>
                            {{ product.is_active ? 'Active' : 'Draft' }}
                        </p>
                    </div>
                </div>
                <div class="flex items-center space-x-4">
                    <button class="text-muted-foreground hover:text-foreground"><PencilIcon class="h-5 w-5" /></button>
                    <button class="text-muted-foreground hover:text-destructive"><TrashIcon class="h-5 w-5" /></button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { onMounted } from 'vue';
import { useProductStore } from '../../store/modules/products';
import { formatCurrency } from '../../utils/helpers';
import { PlusIcon, PencilIcon, TrashIcon } from '@heroicons/vue/24/outline';
import Loading from '../common/Loading.vue';

const productStore = useProductStore();

onMounted(() => {
    productStore.fetchProducts();
});

const getPrimaryImage = (product) => {
  try {
    const images = JSON.parse(product.images);
    return images && images.length > 0 ? images[0] : 'https://via.placeholder.com/150';
  } catch (e) {
    return 'https://via.placeholder.com/150';
  }
};
</script>