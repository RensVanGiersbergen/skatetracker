// src/routes/login/+page.js
import { goto } from '$app/navigation';

export async function load() {
    const token = localStorage.getItem('authToken');
    const isLoggedIn = !!token;

    if (isLoggedIn) {
        goto('/');
    }
}