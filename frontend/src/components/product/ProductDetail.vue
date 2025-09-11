// File: frontend/src/components/product/ProductDetail.vue
<template>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-12 items-start">
        <div>
            <img :src="primaryImage" :alt="product.name" class="w-full rounded-lg shadow-lg">
            </div>

        <div>
            <p v-if="product.brand" class="text-sm text-gray-500 mb-2 font-semibold tracking-wider uppercase">{{ product.brand }}</p>
            <h1 class="text-4xl font-bold text-gray-800">{{ product.name }}</h1>
            <p class="text-3xl font-light text-teal-600 my-4">{{ formatCurrency(product.price) }}</p>
            <p class="text-gray-600 leading-relaxed mt-4">{{ product.description }}</p>

            <div class="mt-6 border-t pt-4">
                <p class="text-sm text-gray-700">
                    <span class="font-bold">SKU:</span> {{ product.sku || 'N/A' }}
                </p>
                 <p class="text-sm text-gray-700 mt-1">
                    <span class="font-bold">Stock:</span> {{ product.stock_quantity > 0 ? `${product.stock_quantity} available` : 'Out of Stock' }}
                </p>
            </div>

            <div class="mt-8">
                <button 
                    @click="handleAddToCart" 
                    :disabled="product.stock_quantity === 0"
                    class="w-full bg-gray-800 text-white font-semibold py-3 px-8 rounded-lg hover:bg-gray-900 transition duration-300 disabled:bg-gray-400 disabled:cursor-not-allowed">
                    {{ product.stock_quantity > 0 ? 'Add to Cart' : 'Out of Stock' }}
                </button>
            </div>
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
    }
});

const cartStore = useCartStore();

const primaryImage = computed(() => {
  try {
    const images = JSON.parse(props.product.images);
    return images && images.length > 0 ? images[0] : 'https://via.placeholder.com/600';
  } catch (e) {
    return 'https://via.placeholder.com/600';
  }
});

const handleAddToCart = () => {
    cartStore.addToCart(props.product);
    alert(`${props.product.name} added to cart!`);
}
</script>