// File: frontend/src/store/modules/products.js
import { defineStore } from 'pinia';
import api from '../../services/api';

export const useProductStore = defineStore('products', {
  state: () => ({
    products: [],
    currentProduct: null,
    loading: false,
    error: null,
  }),
  actions: {
    async fetchProducts() {
      this.loading = true;
      this.error = null;
      try {
        const response = await api.get('/products');
        if (response.data.success) {
          this.products = response.data.data;
        }
      } catch (error) {
        this.error = error.message || 'Failed to fetch products.';
        console.error(error);
      } finally {
        this.loading = false;
      }
    },
    async fetchProduct(id) {
      this.loading = true;
      this.error = null;
      this.currentProduct = null;
      try {
        const response = await api.get(`/products/${id}`);
        if (response.data.success) {
          this.currentProduct = response.data.data;
        }
      } catch (error) {
        this.error = error.message || `Failed to fetch product ${id}.`;
        console.error(error);
      } finally {
        this.loading = false;
      }
    },
    async createProduct(productData) {
      // productData should be FormData from the component
      const response = await api.post('/admin/products', productData, {
        headers: { 'Content-Type': 'multipart/form-data' },
      });
      if (response.data.success) {
        this.products.unshift(response.data.data); // Add new product to the start of the list
      }
      return response;
    }
  },
});