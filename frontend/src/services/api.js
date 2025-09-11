// File: frontend/src/services/api.js
import axios from 'axios';
import { useAuthStore } from '../store/modules/auth';

const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

export const setupInterceptors = (router) => {
  // This is a "request interceptor". It runs BEFORE each request is sent.
  apiClient.interceptors.request.use(
    (config) => {
      const authStore = useAuthStore();
      const token = authStore.token;
      if (token) {
        // If a token exists, add it to the Authorization header.
        config.headers.Authorization = `Bearer ${token}`;
      }
      return config;
    },
    (error) => Promise.reject(error)
  );

  // This is a "response interceptor". It runs AFTER a response is received.
  apiClient.interceptors.response.use(
    (response) => response,
    (error) => {
      // If we get a 401 Unauthorized error, it means the token is bad.
      // We automatically log the user out and redirect them to the login page.
      if (error.response && error.response.status === 401) {
        const authStore = useAuthStore();
        authStore.logout();
        router.push('/login');
      }
      return Promise.reject(error.response?.data || error);
    }
  );
};

export default apiClient;