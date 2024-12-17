<script>
    import { arrowBack } from "ionicons/icons";

    export let title = "Header title";
    export let color = "primary";
</script>

<ion-header>
    <ion-toolbar {color}>
        <ion-buttons slot="start">
            <!-- svelte-ignore a11y-click-events-have-key-events -->
            <!-- svelte-ignore a11y-no-static-element-interactions -->
            <ion-button on:click={() => window.history.back()} fill="clear">
                <ion-icon slot="icon-only" icon={arrowBack} />
            </ion-button>
        </ion-buttons>
        <ion-title>{title}</ion-title>
    </ion-toolbar>
</ion-header>
