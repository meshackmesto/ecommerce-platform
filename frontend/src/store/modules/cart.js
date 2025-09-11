// File: frontend/src/store/modules/cart.js
import { defineStore } from 'pinia';

export const useCartStore = defineStore('cart', {
  state: () => ({
    items: JSON.parse(localStorage.getItem('cart_items')) || [],
  }),
  getters: {
    cartItemCount: (state) => state.items.reduce((total, item) => total + item.quantity, 0),
    cartTotal: (state) => {
      return state.items.reduce((total, item) => {
        // Price from backend is a string, so we must parse it
        const price = parseFloat(item.price);
        return total + price * item.quantity;
      }, 0);
    },
  },
  actions: {
    addToCart(product) {
      const existingItem = this.items.find((item) => item.id === product.id);
      if (existingItem) {
        existingItem.quantity++;
      } else {
        this.items.push({ ...product, quantity: 1 });
      }
      this.persistCart();
    },
    removeFromCart(productId) {
      this.items = this.items.filter((item) => item.id !== productId);
      this.persistCart();
    },
    clearCart() {
      this.items = [];
      this.persistCart();
    },
    persistCart() {
      localStorage.setItem('cart_items', JSON.stringify(this.items));
    },
  },
});