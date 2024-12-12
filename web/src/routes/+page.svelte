<script>
    import { page } from "$app/stores";
    import { add, compass, speedometer, stopwatch } from "ionicons/icons";
    import dayjs from "dayjs";
    import { alertController } from "ionic-svelte";

    function formatRideDuration(startTime, endTime) {
        // Parse ISO strings into Date objects
        const start = new Date(startTime);
        const end = new Date(endTime);

        // Calculate the difference in milliseconds
        const durationMs = end - start;

        // Convert to hours, minutes, and seconds
        const totalSeconds = Math.floor(durationMs / 1000);
        const hours = Math.floor(totalSeconds / 3600);
        const minutes = Math.floor((totalSeconds % 3600) / 60);
        const seconds = totalSeconds % 60;

        // Format as "HH:mm:ss"
        const formattedHours = String(hours);
        const formattedMinutes = String(minutes).padStart(2, "0");
        const formattedSeconds = String(seconds).padStart(2, "0");

        return `${formattedHours}:${formattedMinutes}:${formattedSeconds}`;
    }

    const showAlert = async (options) => {
        const alert = await alertController.create(options);
        alert.present();
    };
</script>

<svelte:head>
    <title>Rides - Skatetracker</title>
</svelte:head>

<ion-content fullscreen class="ion-padding">
    {#if $page.data.rides === null}
        <ion-card>
            <ion-card-header>
                <ion-card-subtitle>No rides found</ion-card-subtitle>
                <ion-card-title>¯\_(ツ)_/¯</ion-card-title>
            </ion-card-header>

            <ion-card-content>
                <p>Start your first ride with the orange button below :D</p>
            </ion-card-content>
        </ion-card>
    {:else}
        {#each $page.data.rides as ride}
            <ion-card>
                <ion-card-header>
                    <ion-card-subtitle
                        >{dayjs(ride.start_time).format(
                            "HH:mm - D MMMM YYYY",
                        )}</ion-card-subtitle
                    >
                    <ion-text class="title">{ride.title}</ion-text>
                </ion-card-header>

                <ion-card-content>
                    <ion-grid>
                        <ion-row>
                            <ion-text color="secondary"
                                >{ride.description}</ion-text
                            >
                        </ion-row>
                        <ion-row>
                            <!-- svelte-ignore a11y-click-events-have-key-events -->
                            <!-- svelte-ignore a11y-no-static-element-interactions -->
                            <ion-chip
                                color="secondary"
                                on:click={showAlert({
                                    header: "Distance",
                                    message:
                                        "Your ride was " +
                                        (ride.distance / 1000).toFixed(2) +
                                        " km long.",
                                    buttons: [`Let's go!`],
                                })}
                            >
                                <ion-icon icon={compass}></ion-icon>
                                <ion-label
                                    >{(ride.distance / 1000).toFixed(2)} km</ion-label
                                >
                            </ion-chip>
                        </ion-row>
                        <ion-row>
                            <!-- svelte-ignore a11y-click-events-have-key-events -->
                            <!-- svelte-ignore a11y-no-static-element-interactions -->
                            <ion-chip
                                color="secondary"
                                on:click={showAlert({
                                    header: "Top speed",
                                    message:
                                        "Your top speed was " +
                                        (ride.top_speed * 3.6).toFixed(2) +
                                        " km/h.",
                                    buttons: [`Let's go!`],
                                })}
                            >
                                <ion-icon icon={speedometer}></ion-icon>
                                <ion-label
                                    >{(ride.top_speed * 3.6).toFixed(2)} km/h</ion-label
                                >
                            </ion-chip>
                        </ion-row>
                        <ion-row>
                            <!-- svelte-ignore a11y-click-events-have-key-events -->
                            <!-- svelte-ignore a11y-no-static-element-interactions -->
                            <ion-chip
                                color="secondary"
                                on:click={showAlert({
                                    header: "Ride time",
                                    message:
                                        "Your were riding for " +
                                        formatRideDuration(
                                            ride.start_time,
                                            ride.end_time,
                                        ) +
                                        " in total.",
                                    buttons: [`Let's go!`],
                                })}
                            >
                                <ion-icon icon={stopwatch}></ion-icon>
                                <ion-label
                                    >{formatRideDuration(
                                        ride.start_time,
                                        ride.end_time,
                                    )}</ion-label
                                >
                            </ion-chip>
                        </ion-row>
                    </ion-grid>
                </ion-card-content>
            </ion-card>
        {/each}
    {/if}

    <ion-fab vertical="bottom" horizontal="end" slot="fixed">
        <ion-fab-button>
            <ion-icon icon={add}></ion-icon>
        </ion-fab-button>
    </ion-fab>
</ion-content>

<style>
    ion-content {
        --background: var(--ion-color-light);
    }

    ion-card {
        background-image: linear-gradient(
            130deg,
            var(--ion-color-primary-tint),
            rgb(244, 161, 88),
            var(--ion-color-primary-tint)
        );
    }

    .title {
        color: var(--ion-color-secondary);
        font-weight: 700;
        font-size: large;
        margin-top: 0.5rem;
    }
</style>
