<script>
    import { page } from "$app/stores";
</script>

<svelte:head>
    <title>Rides - Skatetracker</title>
</svelte:head>

<ion-content fullscreen class="ion-padding">
    <h1>Details</h1>
    <h2>{$page.params.ride}</h2>
</ion-content>
