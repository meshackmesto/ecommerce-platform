<template>
  <div v-if="productStore.loading" class="flex justify-center items-center h-64">
    <Loading message="Fetching our latest products..." />
  </div>
  <div v-else-if="productStore.error" class="text-center text-red-500">
    <p>{{ productStore.error }}</p>
  </div>
  <div v-else-if="productStore.products.length === 0" class="text-center text-gray-500">
    <p>No products available at the moment.</p>
  </div>
  <div v-else class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-8">
    <ProductCard v-for="product in productStore.products" :key="product.id" :product="product" />
  </div>
</template>

<script setup>
import { onMounted } from 'vue';
import { useProductStore } from '../../store/modules/products';
import ProductCard from './ProductCard.vue';
import Loading from '../common/Loading.vue';

const productStore = useProductStore();

onMounted(() => {
  productStore.fetchProducts();
});
</script>