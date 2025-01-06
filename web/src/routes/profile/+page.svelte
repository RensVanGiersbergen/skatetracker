<script>
    import { goto } from "$app/navigation";

    function handleLogout() {
        localStorage.removeItem("authToken");
        goto("/login");
    }
</script>

<svelte:head>
    <title>Profile - Skatetracker</title>
</svelte:head>

<ion-content fullscreen class="ion-padding">
    <ion-title>Profile</ion-title>
    <!-- svelte-ignore a11y-interactive-supports-focus -->
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <ion-button
        expand="full"
        color="danger"
        on:click={handleLogout}
        role="button">Logout</ion-button
    >
</ion-content>

<style>
    ion-content {
        --background: var(--ion-color-light);
    }
</style>
