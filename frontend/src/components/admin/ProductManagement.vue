// File: frontend/src/components/admin/ProductManagement.vue
<template>
    <div>
        <div class="flex justify-between items-center mb-6">
            <h1 class="text-3xl font-bold text-gray-800">Product Management</h1>
            <button @click="openModal" class="bg-gray-800 text-white font-semibold py-2 px-5 rounded-lg hover:bg-gray-900 transition duration-300">
                Add New Product
            </button>
        </div>

        <div class="bg-white p-6 rounded-lg shadow">
            <ProductList />
        </div>

        <Modal :isOpen="isModalOpen" @close="closeModal">
            <template #title>Add New Product</template>
            <template #content>
                <ProductForm @submit="handleCreateProduct" @cancel="closeModal" />
            </template>
        </Modal>
    </div>
</template>

<script setup>
import { ref } from 'vue';
import { useProductStore } from '../../store/modules/products';
import ProductList from '../product/ProductList.vue';
import ProductForm from '../product/ProductForm.vue';
import Modal from '../common/Modal.vue';

const isModalOpen = ref(false);
const productStore = useProductStore();

const openModal = () => isModalOpen.value = true;
const closeModal = () => isModalOpen.value = false;

const handleCreateProduct = async (formData) => {
    try {
        await productStore.createProduct(formData);
        alert('Product created successfully!');
        closeModal();
    } catch (error) {
        console.error("Failed to create product:", error);
        alert('Error: ' + error.message);
    }
};
</script>