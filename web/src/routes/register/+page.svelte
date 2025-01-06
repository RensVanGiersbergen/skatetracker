<script>
	import api from "$lib/axios.js";
	import { loadingController } from "ionic-svelte";
	import { showToast } from "$lib/stores/feedbackStore";
	import {
		personCircleOutline,
		mailOutline,
		lockClosedOutline,
	} from "ionicons/icons";
	import { goto } from "$app/navigation";

	let username = "";
	let email = "";
	let password = "";

	async function handleRegister() {
		let loading;

		// Show spinner
		loading = await loadingController.create({
			message: "Creating account...",
			spinner: "crescent",
			duration: 10000,
		});
		await loading.present();

		try {
			// Call the API
			let response = await api.post("/account/register", {
				username,
				email,
				password,
			});
			showToast({
				color: "success",
				message: response.data.message,
				duration: 2000,
			});
			goto("/login");
		} catch (error) {
			showToast({
				color: "danger",
				message: error.response?.data?.error || "An error occurred",
				duration: 5000,
			});
		} finally {
			await loading.dismiss();
		}
	}

	function handleCancel() {
		goto("/login");
	}
</script>

<svelte:head>
	<title>Register - Skatetracker</title>
</svelte:head>

<ion-content fullscreen class="ion-padding">
	<div class="form-container">
		<ion-img src="/images/skatetrackerlogo.png" alt="logo" class="logo"
		></ion-img>
		<br />
		<div class="forms">
			<ion-item>
				<ion-icon slot="start" icon={personCircleOutline}></ion-icon>
				<ion-input
					placeholder="Username"
					value={username}
					on:ionInput={(e) => (username = e.target.value)}
					type="text"
					name="text"
				/>
			</ion-item>

			<ion-item>
				<ion-icon slot="start" icon={mailOutline}></ion-icon>
				<ion-input
					placeholder="Email"
					value={email}
					on:ionInput={(e) => (email = e.target.value)}
					type="email"
					name="email"
				/>
			</ion-item>

			<ion-item>
				<ion-icon slot="start" icon={lockClosedOutline}></ion-icon>
				<ion-input
					placeholder="Password"
					value={password}
					on:ionInput={(e) => (password = e.target.value)}
					type="password"
					name="password"
				>
					<ion-input-password-toggle slot="end" color="secondary"
					></ion-input-password-toggle>
				</ion-input>
			</ion-item>

			<!-- svelte-ignore a11y-interactive-supports-focus -->
			<ion-button
				expand="full"
				on:click={handleRegister}
				on:keydown={(e) => e.key === "Enter" && handleRegister()}
				role="button">Register</ion-button
			>

			<!-- svelte-ignore a11y-interactive-supports-focus -->
			<ion-button
				expand="full"
				color="danger"
				on:click={handleCancel}
				on:keydown={(e) => e.key === "Enter" && handleCancel()}
				role="button">Cancel</ion-button
			>
		</div>
	</div>
</ion-content>

<style>
	ion-content {
		--background: var(--ion-color-light);
	}

	.forms {
		width: 100%;
		max-width: 500px;
	}

	.form-container {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		height: 100%;
		min-width: 100%;
	}

	ion-icon {
		color: var(--ion-color-secondary);
	}

	ion-img {
		width: 100%;
		max-width: 200px;
		margin-bottom: 20px;
	}
</style>
