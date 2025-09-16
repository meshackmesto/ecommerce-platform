// File: frontend/src/components/product/ProductCard.vue
<template>
  <div class="bg-card rounded-lg shadow-custom-md overflow-hidden transform hover:shadow-custom-xl hover:-translate-y-1 transition-all duration-300 group">
    <router-link :to="`/products/${product.id}`">
      <div class="relative h-64 overflow-hidden">
        <img :src="primaryImage" :alt="product.name" class="w-full h-full object-cover transform group-hover:scale-110 transition-transform duration-500">
      </div>
    </router-link>
    <div class="p-4">
      <p class="text-sm text-muted-foreground mb-1">{{ product.brand || 'Generic' }}</p>
      <h3 class="font-semibold text-lg text-card-foreground truncate h-7">{{ product.name }}</h3>
      <p class="text-primary font-bold mt-2">{{ formatCurrency(product.price) }}</p>
      <button @click="handleAddToCart" class="mt-4 w-full btn-primary !py-2">
        Add to Cart
      </button>
    </div>
  </div>
</template>

<script setup>
// Script content remains the same as before
import { computed } from 'vue';
import { useCartStore } from '../../store/modules/cart';
import { formatCurrency } from '../../utils/helpers';

const props = defineProps({
  product: { type: Object, required: true },
});
const cartStore = useCartStore();
const primaryImage = computed(() => {
  try {
    const images = JSON.parse(props.product.images);
    return images && images.length > 0 ? images[0] : 'https://via.placeholder.com/400';
  } catch (e) {
    return 'https://via.placeholder.com/400';
  }
});
const handleAddToCart = () => {
  cartStore.addToCart(props.product);
  alert(`${props.product.name} added to cart!`);
};
</script>