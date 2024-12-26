<script>
    import { goto } from "$app/navigation";
</script>

<svelte:head>
    <title>Offline - Skatetracker</title>
</svelte:head>

<ion-content fullscreen class="ion-padding">
    <h1>You are offline</h1>
    <p>Check your internet connection and try again</p>
    <ion-button
        expand="block"
        onClick={() => {
            window.history.back();
        }}
    >
        Go back
    </ion-button>
</ion-content>
