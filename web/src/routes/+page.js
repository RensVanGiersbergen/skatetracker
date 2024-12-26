// src/routes/+page.js
import api from '$lib/axios';
import { showToast } from '$lib/stores/feedbackStore';

export async function load() {
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