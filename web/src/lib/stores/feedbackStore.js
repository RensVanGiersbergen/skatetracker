
import { writable } from 'svelte/store';

// Writable store to manage feedbacks
export const toastParams = writable(null);

// Function to show a toast
export const showToast = async (options) => {
	toastParams.set(options);
};