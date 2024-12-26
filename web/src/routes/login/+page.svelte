<script>
	import api from "$lib/axios.js";
	import { loadingController } from "ionic-svelte";
	import { showToast } from "$lib/stores/feedbackStore";
	import { mailOutline, lockClosedOutline } from "ionicons/icons";
	import { goto } from "$app/navigation";

	let email = "";
	let password = "";

	async function handleLogin() {
		let loading;

		loading = await loadingController.create({
			message: "Logging in...",
			spinner: "crescent",
			duration: 10000,
		});
		await loading.present();

		try {
			// Call the API
			let response = await api.post("/account/login", {
				email,
				password,
			});
			showToast({
				color: "success",
				message: response.data.message,
				duration: 2000,
			});
			localStorage.setItem("authToken", response.data.token);
			goto("/");
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
</script>

<svelte:head>
	<title>Login - Skatetracker</title>
</svelte:head>

<ion-content fullscreen class="ion-padding">
	<div class="form-container">
		<ion-img src="/images/skatetrackerlogo.png" alt="logo" class="logo"
		></ion-img>
		<br />
		<div class="forms">
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
				on:click={handleLogin}
				on:keydown={(e) => e.key === "Enter" && handleLogin()}
				role="button">Login</ion-button
			>
			<br />
			<ion-label
				>Create an account <a href="/register">here</a>...</ion-label
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
