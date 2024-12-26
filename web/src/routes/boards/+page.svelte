<script>
    import { page } from "$app/stores";
    import { goto } from "$app/navigation";
    import { add, compass, speedometer, stopwatch } from "ionicons/icons";
    import dayjs from "dayjs";
    import { alertController } from "ionic-svelte";

    function formatRideDuration(duration) {
        // Convert to hours, minutes, and seconds
        const totalSeconds = duration;
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
    <title>Boards - Skatetracker</title>
</svelte:head>

<ion-content fullscreen class="ion-padding">
    {#if $page.data === undefined}
        <ion-spinner name="crescent"></ion-spinner>
    {:else if $page.data.boards === null}
        <ion-card>
            <ion-card-header>
                <ion-card-subtitle>No boards found</ion-card-subtitle>
                <ion-card-title>¯\_(ツ)_/¯</ion-card-title>
            </ion-card-header>

            <ion-card-content>
                <p>Add your first board with the orange button below :D</p>
            </ion-card-content>
        </ion-card>
    {:else}
        {#each $page.data.boards as board}
            <ion-card>
                <ion-card-header>
                    <ion-card-subtitle
                        >{dayjs(board.created_at).format(
                            "HH:mm - D MMMM YYYY",
                        )}</ion-card-subtitle
                    >
                    <ion-text class="title">{board.nickname}</ion-text>
                </ion-card-header>

                <ion-card-content>
                    <ion-grid>
                        <ion-row>
                            <ion-text color="secondary">{board.brand}</ion-text>
                        </ion-row>
                        <ion-row>
                            <!-- svelte-ignore a11y-click-events-have-key-events -->
                            <!-- svelte-ignore a11y-no-static-element-interactions -->
                            <ion-chip
                                color="secondary"
                                on:click={showAlert({
                                    header: "Distance",
                                    message:
                                        "Your board went " +
                                        (board.total_distance / 1000).toFixed(
                                            2,
                                        ) +
                                        " km far in total.",
                                    buttons: [`Let's go!`],
                                })}
                            >
                                <ion-icon icon={compass}></ion-icon>
                                <ion-label
                                    >{(board.total_distance / 1000).toFixed(2)} km</ion-label
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
                                        "Your board's top speed was " +
                                        (board.top_speed * 3.6).toFixed(2) +
                                        " km/h.",
                                    buttons: [`Let's go!`],
                                })}
                            >
                                <ion-icon icon={speedometer}></ion-icon>
                                <ion-label
                                    >{(board.top_speed * 3.6).toFixed(2)} km/h</ion-label
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
                                        "Your board's total ride time is " +
                                        formatRideDuration(
                                            board.total_ridetime,
                                        ) +
                                        " in total.",
                                    buttons: [`Let's go!`],
                                })}
                            >
                                <ion-icon icon={stopwatch}></ion-icon>
                                <ion-label
                                    >{formatRideDuration(
                                        board.total_ridetime,
                                    )}</ion-label
                                >
                            </ion-chip>
                        </ion-row>
                    </ion-grid>
                </ion-card-content>
            </ion-card>
        {/each}
    {/if}

    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <!-- svelte-ignore a11y-no-static-element-interactions -->
    <ion-fab
        vertical="bottom"
        horizontal="end"
        slot="fixed"
        on:click={() => goto("boards/add")}
    >
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
