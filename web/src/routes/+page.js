// src/routes/+page.js
import { goto } from '$app/navigation';
import { showToast } from '$lib/stores/feedbackStore';
import api from '$lib/axios';

export async function load() {
    const token = localStorage.getItem('authToken');
    const isLoggedIn = !!token;

    // Redirect to login page if user is not logged in
    if (!isLoggedIn) {
        goto('/login');
    }

    // Verify the token on the server
    try {
        await api.get('/account/verify');
    } catch (error) {
        showToast({
            color: 'danger',
            message: error.response.data.error,
            duration: 5000
        });
        localStorage.removeItem('authToken');
        goto('/login');
    }
}