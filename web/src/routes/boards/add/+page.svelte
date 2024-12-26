<script>
    import api from "$lib/axios.js";
    import { goto } from "$app/navigation";
    import { loadingController } from "ionic-svelte";
    import { showToast } from "$lib/stores/feedbackStore";
    import { arrowBack } from "ionicons/icons";
    import Header from "$lib/components/Header.svelte";

    let nickname = "";
    let brand = "";
    let primary = false;

    async function handleBoard() {
        let loading;

        loading = await loadingController.create({
            message: "Adding board...",
            spinner: "crescent",
            duration: 10000,
        });
        await loading.present();

        // Call the API
        try {
            await api.post("/board/add", { nickname, brand, primary });
            showToast({
                color: "success",
                message: "Added board successfully",
                duration: 2000,
            });
            goto("/boards");
        } catch (error) {
            showToast({
                color: "danger",
                message: error.response.data.error,
                duration: 5000,
            });
        } finally {
            await loading.dismiss();
        }
    }
</script>

<svelte:head>
    <title>Add board - Skatetracker</title>
</svelte:head>

<Header title="Add board" />

<ion-content fullscreen class="ion-padding">
    <ion-card>
        <ion-card-header>
            <ion-card-title>Enter your board's details</ion-card-title>
        </ion-card-header>

        <ion-card-content>
            <ion-item>
                <ion-input
                    name="nickname"
                    label="Nickname"
                    label-placement="floating"
                    value={nickname}
                    on:ionInput={(e) => (nickname = e.target.value)}
                    type="text"
                />
            </ion-item>

            <ion-item>
                <ion-input
                    name="brand"
                    label="Brand"
                    label-placement="floating"
                    value={brand}
                    on:ionInput={(e) => (brand = e.target.value)}
                    type="text"
                />
            </ion-item>

            <ion-item>
                <!-- svelte-ignore a11y-click-events-have-key-events -->
                <!-- svelte-ignore a11y-no-static-element-interactions -->
                <ion-checkbox
                    value="primary"
                    on:ionChange={(e) => (primary = e.detail.checked)}
                    color="secondary"
                    slot="end"
                />
                <!-- make label uninteractable-->
                <ion-label style="pointer-events: none;"
                    >Primary board</ion-label
                >
            </ion-item>

            <br />
            <!-- svelte-ignore a11y-click-events-have-key-events -->
            <!-- svelte-ignore a11y-no-static-element-interactions -->
            <ion-button
                expand="block"
                color="primary"
                on:click={() => handleBoard()}>Add board</ion-button
            >
        </ion-card-content>
    </ion-card>
</ion-content>

<style>
    ion-content {
        --background: var(--ion-color-light);
    }
</style>
