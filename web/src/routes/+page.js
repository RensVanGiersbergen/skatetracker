// src/routes/+page.js
import { goto } from '$app/navigation';
import api from '$lib/axios';
import { showToast } from '$lib/stores/feedbackStore';

export async function load() {
    const token = localStorage.getItem('authToken');
    const isLoggedIn = !!token;

    // Redirect to login page if user is not logged in
    if (!isLoggedIn) {
        goto('/login');
    }

    try {
        const response = await api.get('/ride/all?page=1&limit=10')
        const rides = await response.data
        return { rides }
    } catch (error) {
        showToast({
            color: "danger",
            message: error.response.data.error,
            duration: 5000,
        });
    }
}