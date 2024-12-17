// src/routes/boards/+page.js
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
        const response = await api.get('/board/mine')
        const boards = await response.data
        return { boards }
    } catch (error) {
        showToast({
            color: "danger",
            message: error.response.data.error,
            duration: 5000,
        });
    }
}