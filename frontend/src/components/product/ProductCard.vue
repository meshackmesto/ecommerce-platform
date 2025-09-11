// File: frontend/src/components/product/ProductCard.vue
<template>
  <div class="bg-white rounded-lg shadow-md overflow-hidden transform hover:-translate-y-1 transition-transform duration-300 group">
    <router-link :to="`/products/${product.id}`">
      <div class="relative h-64 overflow-hidden">
        <img :src="primaryImage" :alt="product.name" class="w-full h-full object-cover transform group-hover:scale-110 transition-transform duration-500">
      </div>
    </router-link>
    <div class="p-4">
      <p class="text-sm text-gray-500 mb-1">{{ product.brand || 'Generic' }}</p>
      <h3 class="font-semibold text-lg truncate h-7">{{ product.name }}</h3>
      <p class="text-gray-800 font-bold mt-2">{{ formatCurrency(product.price) }}</p>
      <button @click="handleAddToCart" class="mt-4 w-full bg-gray-800 text-white py-2 rounded-lg hover:bg-gray-900 transition duration-300">
        Add to Cart
      </button>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue';
import { useCartStore } from '../../store/modules/cart';
import { formatCurrency } from '../../utils/helpers';

const props = defineProps({
  product: {
    type: Object,
    required: true,
  },
});

const cartStore = useCartStore();

const primaryImage = computed(() => {
  try {
    // Backend 'images' field is a JSON string array, e.g., "[\"url1\", \"url2\"]"
    const images = JSON.parse(props.product.images);
    return images && images.length > 0 ? images[0] : 'https://via.placeholder.com/400';
  } catch (e) {
    return 'https://via.placeholder.com/400';
  }
});

const handleAddToCart = () => {
  cartStore.addToCart(props.product);
  // You could add a more sophisticated notification here
  alert(`${props.product.name} added to cart!`);
};
</script>