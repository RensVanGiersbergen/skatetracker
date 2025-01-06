// src/routes/details/[ride]/+page.js
import api from '$lib/axios.js';
import { showToast } from '$lib/stores/feedbackStore';

export async function load({ params }) {
    try {
        const response = await api.get('/ride/' + params.ride)
        const ride = await response.data

        return { ride }
    } catch (error) {
        console.log(error);
        showToast({
            color: "danger",
            message: error.response?.data?.error,
            duration: 5000,
        });
    }
}