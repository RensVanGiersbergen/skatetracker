// src/routes/boards/+page.js
import api from '$lib/axios';
import { showToast } from '$lib/stores/feedbackStore';

export async function load() {
    try {
        const response = await api.get('/board/mine')
        const boards = await response.data
        return { boards }
    } catch (error) {
        showToast({
            color: "danger",
            message: error.response?.data?.error,
            duration: 5000,
        });
    }
}