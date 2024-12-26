<script>
	import { setupIonicBase } from "ionic-svelte";

	/* Call Ionic's setup routine. */
	setupIonicBase();

	/* Import all components. (You can selectively import components instead; see below.) */
	import "ionic-svelte/components/all";

	/* Theme variables */
	import "../theme/variables.css";

	/*
		This part - import 'ionic-svelte/components/all'; - loads all components at once. Importing this way adds 80 components and >800kb (uncompressed) to your bundle.

		Alternately, you can choose to import only the components you want to use.

		Doing selective imports in this file is recommended because you only have to do such imports once.
		If you like to code-split differently, you are free to import whereever you like.

		Example: If you replace the line import 'ionic-svelte/components/all'; with the imports below, the resulting bundle becomes much smaller.

		import 'ionic-svelte/components/ion-app';
		import 'ionic-svelte/components/ion-card';
		import 'ionic-svelte/components/ion-card-title';
		import 'ionic-svelte/components/ion-card-subtitle';
		import 'ionic-svelte/components/ion-card-header';
		import 'ionic-svelte/components/ion-card-content';
		import 'ionic-svelte/components/ion-button';
		import 'ionic-svelte/components/ion-item';
		import 'ionic-svelte/components/ion-label';

		To see the full list of possible imports, click ionic-svelte-components-all-import above.

		When you decide to do selective imports, ion-app must be imported in this file like this:

	    import 'ionic-svelte/components/ion-app';

		Report issues here - https://github.com/Tommertom/svelte-ionic-npm/issues
		Want to know more about what is happening? Follow me on X! - https://x.com/Tommertomm
		Discord channel on Ionic server - https://discordapp.com/channels/520266681499779082/1049388501629681675
	*/

	// Import the icons
	import { rocket, person, cog } from "ionicons/icons";

	import { toastController } from "ionic-svelte";
	import { toastParams } from "$lib/stores/feedbackStore";

	// Listen for changes to toastParams
	$: if (toastParams) {
		toastParams.subscribe(async (params) => {
			if (params) {
				const toast = await toastController.create({
					position: "top",
					color: params.color || "danger",
					duration: params.duration || 4000,
					message: params.message || "Default toast message",
					showCloseButton: false,
				});
				await toast.present();
				toastParams.set(null); // Reset after showing
			}
		});
	}

	// Verify token before accessing the app
	import { showToast } from "$lib/stores/feedbackStore";
	import api from "$lib/axios";
	(async () => {
		if (!localStorage.getItem("authToken")) {
			goto("/login");
			return;
		}

		try {
			await api.get("/account/verify");
		} catch (error) {
			showToast({
				color: "danger",
				message: error.response.data.error,
				duration: 5000,
			});
			localStorage.removeItem("authToken");
			goto("/login");
		}
	})();

	import { page } from "$app/stores";
	import { goto } from "$app/navigation";

	const myTabs = [
		/* {
			title: "Trophies",
			url: "/trophies",
			icon: trophy,
		}, */
		{
			title: "Rides",
			url: "/",
			icon: rocket,
		},
		{
			title: "Boards",
			url: "/boards",
			icon: cog,
		},
		{
			title: "Profile",
			url: "/profile",
			icon: person,
		},
	];

	function navigateHandler(page) {
		goto(page);
	}
</script>

<ion-app>
	<!-- If page is any page not in the myTabs show no bottom nav-->
	{#if $page.url && myTabs.find((tab) => tab.url === $page.url.pathname)}
		<div id="content">
			<slot />
		</div>
		<ion-tabs>
			<ion-tab-bar slot="bottom">
				{#each myTabs as tab}
					<!-- svelte-ignore a11y-click-events-have-key-events -->
					<!-- svelte-ignore a11y-no-static-element-interactions -->
					<ion-tab-button
						type="button"
						on:click={() => {
							navigateHandler(tab.url);
						}}
					>
						<ion-icon
							icon={tab.icon}
							color={$page.url.pathname === tab.url
								? "primary"
								: "medium"}
						></ion-icon>
						<ion-label
							color={$page.url.pathname === tab.url
								? "primary"
								: "medium"}>{tab.title}</ion-label
						>
					</ion-tab-button>
				{/each}
			</ion-tab-bar>
		</ion-tabs>
	{:else}
		<ion-app>
			<slot />
		</ion-app>
	{/if}
</ion-app>

<style>
	ion-app {
		position: relative;
		height: 100%;
		display: flex;
		flex-direction: column;
	}

	#content {
		padding-bottom: 57px; /* Reserve space for the tab bar */
		overflow-y: auto; /* Enable scrolling for long content */
		flex: 1; /* Allows content to expand dynamically */
	}

	ion-tabs {
		top: auto;
		height: 57px;
	}

	ion-tab-bar {
		position: absolute;
		bottom: 0;
		height: 57px; /* Ensure tab bar height is flexible */
		width: 100%;
		border-top: 1px solid var(--ion-color-primary);
	}
</style>
