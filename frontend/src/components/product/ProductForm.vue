// File: frontend/src/components/product/ProductForm.vue
<template>
    <form @submit.prevent="submitForm" class="space-y-6">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
                <label for="name" class="form-label">Product Name</label>
                <input id="name" type="text" v-model="form.name" required class="form-input">
            </div>
            <div>
                <label for="brand" class="form-label">Brand</label>
                <input id="brand" type="text" v-model="form.brand" class="form-input">
            </div>
            <div>
                <label for="price" class="form-label">Price (KES)</label>
                <input id="price" type="number" step="0.01" v-model="form.price" required class="form-input">
            </div>
            <div>
                <label for="stock" class="form-label">Stock Quantity</label>
                <input id="stock" type="number" v-model="form.stock_quantity" required class="form-input">
            </div>
             <div>
                <label for="category" class="form-label">Category</label>
                <input id="category" type="text" v-model="form.category" class="form-input">
            </div>
             <div>
                <label for="sku" class="form-label">SKU</label>
                <input id="sku" type="text" v-model="form.sku" class="form-input">
            </div>
             <div class="md:col-span-2">
                <label for="description" class="form-label">Description</label>
                <textarea id="description" v-model="form.description" rows="4" class="form-input"></textarea>
            </div>
             <div class="md:col-span-2">
                <label for="image" class="form-label">Product Image</label>
                <input id="image" type="file" @change="onFileChange" class="block w-full text-sm text-gray-500 file:mr-4 file:py-2 file:px-4 file:rounded-full file:border-0 file:text-sm file:font-semibold file:bg-teal-50 file:text-teal-700 hover:file:bg-teal-100">
            </div>
        </div>
        <div class="flex justify-end space-x-4 pt-4">
            <button type="button" @click="$emit('cancel')" class="px-6 py-2 rounded-lg border border-gray-300 text-gray-700 hover:bg-gray-100">Cancel</button>
            <button type="submit" class="px-6 py-2 rounded-lg bg-gray-800 text-white hover:bg-gray-900">Save Product</button>
        </div>
    </form>
</template>

<script setup>
import { ref } from 'vue';

const emit = defineEmits(['submit', 'cancel']);

const form = ref({
    name: '',
    description: '',
    price: 0,
    stock_quantity: 0,
    category: '',
    brand: '',
    sku: '',
});

const imageFile = ref(null);

const onFileChange = (e) => {
    imageFile.value = e.target.files[0];
};

const submitForm = () => {
    const formData = new FormData();
    // Append all form fields to FormData
    Object.keys(form.value).forEach(key => {
        formData.append(key, form.value[key]);
    });
    
    if (imageFile.value) {
        formData.append('image', imageFile.value);
    }

    emit('submit', formData);
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