import axios from 'axios';
import { goto } from '$app/navigation';

// Axios configuration
const api = axios.create({
  baseURL: 'http://192.168.2.9:8080' /* http://localhost:8080, http://172.27.176.1:8080, http://192.168.1.143:8080, http://192.168.2.8:8080 */,
  headers: {
    'Content-Type': 'application/json'
  }
});

// Add interceptors for authentication
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('authToken');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

// Add response interceptor to catch only network errors
api.interceptors.response.use(
  (response) => response, // Pass successful responses through
  (error) => {
    if (error.request && !error.response) {
      goto('/offline');
    }

    // Pass all other errors (with a response) to the calling code
    return Promise.reject(error);
  }
);

export default api;