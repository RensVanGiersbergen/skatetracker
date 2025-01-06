// src/routes/add/+page.js
import api from '$lib/axios';
import { showToast } from '$lib/stores/feedbackStore';
import { goto } from '$app/navigation';

export async function load() {
    try {
        const response = await api.get('/board/mine')
        const boards = await response.data

        if (boards.length === 0) {
            showToast({
                color: "warning",
                message: "You don't have any boards yet. Create one first!",
                duration: 5000,
            });
            goto('/boards')
        }
        return { boards }
    } catch (error) {
        showToast({
            color: "danger",
            message: error.response?.data?.error,
            duration: 5000,
        });
    }
}