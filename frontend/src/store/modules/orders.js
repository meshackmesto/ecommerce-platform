// File: frontend/src/store/modules/orders.js
import { defineStore } from 'pinia';
import api from '../../services/api';

export const useOrderStore = defineStore('orders', {
  state: () => ({
    orders: [], // For the current user's orders
    allOrders: [], // For the admin panel
    loading: false,
    error: null,
  }),
  actions: {
    async createOrder(orderData) {
      const response = await api.post('/orders', orderData);
      return response.data;
    },
    async fetchOrders() { // For regular users
      this.loading = true;
      this.error = null;
      try {
        const response = await api.get('/orders');
        if (response.data.success) {
          this.orders = response.data.data;
        }
      } catch (error) {
        this.error = error.message || 'Failed to fetch orders.';
      } finally {
        this.loading = false;
      }
    },
    // NEW ACTION FOR ADMINS
    async fetchAllOrders() {
      this.loading = true;
      this.error = null;
      try {
        // This endpoint is defined in your backend main.go for admins
        const response = await api.get('/admin/orders');
        if (response.data.success) {
          this.allOrders = response.data.data;
        }
      } catch (error) {
        this.error = error.message || 'Failed to fetch all orders.';
      } finally {
        this.loading = false;
      }
    }
  },
});