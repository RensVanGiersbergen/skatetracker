<script>
    import { onMount } from "svelte";
    import { goto } from "$app/navigation";
    import { page } from "$app/stores";
    import { writable } from "svelte/store";
    import { showToast } from "$lib/stores/feedbackStore";
    import Header from "$lib/components/Header.svelte";
    import L from "leaflet";
    import "leaflet/dist/leaflet.css";
    import api from "$lib/axios.js";

    // Capacitor plugins
    import { Geolocation } from "@capacitor/geolocation";
    import { Motion } from "@capacitor/motion";
    import { KeepAwake } from "@capacitor-community/keep-awake";

    // Tracking variables
    let isTracking = false;
    let rideName = "";
    let rideDescription = "";
    let selectedBoard = $page.data.boards[0];
    let startTime;
    let elapsed;
    let interval;

    // Variables
    let isModalOpen = true;
    let locations = writable([]);
    let map = null;
    let watchId;
    let shakinessAverage = null;
    let measurementCount = 0;

    // Function to start the timer when tracking starts
    function startTimer() {
        startTime = Date.now(); // Set the start time when tracking begins
        elapsed = 0; // Reset elapsed time
        interval = setInterval(() => {
            elapsed = Math.floor((Date.now() - startTime) / 1000); // Update elapsed time every second
        }, 1000);
    }

    // Function to stop the timer
    function stopTimer() {
        clearInterval(interval); // Stop the timer
    }

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

    async function handlePermission() {
        // Check if location permission is granted
        const permission = await Geolocation.checkPermissions();
        if (permission.location === "granted") {
            return;
        } else {
            // Request location permission
            const permission = await Geolocation.requestPermissions();
            if (permission.location === "granted") {
                return;
            } else {
                console.error("Location permission denied");
                showToast({
                    color: "danger",
                    message: "Location permission is required for tracking",
                    duration: 5000,
                });
                goto("/");
            }
        }
    }
    async function startTracking() {
        isTracking = true;
        isModalOpen = false;

        // Keep screen awake
        await KeepAwake.keepAwake();

        // Add listener for accelerometer data
        await Motion.addListener("accel", (event) => {
            let shakiness = Math.sqrt(
                event.acceleration.x * event.acceleration.x +
                    event.acceleration.y * event.acceleration.y +
                    event.acceleration.z * event.acceleration.z,
            );

            measurementCount++;

            if (shakinessAverage == null) {
                shakinessAverage = shakiness;
            } else {
                shakinessAverage =
                    (shakinessAverage * (measurementCount - 1) + shakiness) /
                    measurementCount;
            }
        });

        // Start watching the location
        try {
            watchId = await Geolocation.watchPosition(
                {
                    // watch options
                    enableHighAccuracy: true,
                    minimumUpdateInterval: 1000,
                },
                (position, err) => {
                    if (err) {
                        console.error(err);
                        return;
                    }

                    // Add location to locations array and use capacitor motion to determine shakiness factor
                    if (position) {
                        locations.update((locs) => {
                            locs.push({
                                tracking_time: new Date(
                                    position.timestamp,
                                ).toISOString(),
                                latitude: position.coords.latitude,
                                longitude: position.coords.longitude,
                                speed: position.coords.speed,
                                shakiness: Number(
                                    (shakinessAverage ?? 0).toFixed(2),
                                ),
                            });
                            return locs;
                        });

                        // Reset shakiness average
                        shakinessAverage = null;
                        measurementCount = 0;

                        // Start moving the map and draw polyline if there are at least two points
                        if (map != null && $locations.length > 1) {
                            map.setView([
                                position.coords.latitude,
                                position.coords.longitude,
                            ]);

                            // Draw polyline between last two points
                            let last = $locations[$locations.length - 2];
                            let current = $locations[$locations.length - 1];

                            // Draw polyline and set color according to speed
                            L.polyline(
                                [
                                    [last.latitude, last.longitude],
                                    [current.latitude, current.longitude],
                                ],
                                {
                                    color:
                                        current.speed > 13.89
                                            ? "#f56042"
                                            : current.speed > 11.11
                                              ? "#f5b642"
                                              : current.speed > 8.33
                                                ? "#f5e942"
                                                : current.speed > 5.56
                                                  ? "#c8f542"
                                                  : "#72f542",
                                },
                            ).addTo(map);
                        }
                    }
                },
            );
        } catch (e) {
            console.error("Error starting location tracking:", e);
        }

        // Start timer
        startTimer();
    }

    async function stopTracking() {
        // Upload ride data
        try {
            const response = await api.post("/ride/add", {
                title: rideName,
                description: rideDescription,
                board_id: selectedBoard.board_id,
                trackings: $locations,
            });

            isTracking = false;
            isModalOpen = false;

            // Stop timer
            stopTimer();

            // Stop watching location and reset variables
            await Geolocation.clearWatch({ id: watchId });
            $locations = [];
            watchId = null;

            // Stop watching motion and reset variables
            await Motion.removeAllListeners();
            shakinessAverage = null;
            measurementCount = 0;

            // Allow screen to sleep
            await KeepAwake.allowSleep();

            showToast({
                color: "success",
                message: "Ride saved successfully",
                duration: 3000,
            });

            goto(`/details/${response.data.ride_id}`);
        } catch (error) {
            showToast({
                color: "danger",
                message: error.response.data.error,
                duration: 5000,
            });
        }
    }

    onMount(async () => {
        // Request location permission
        await handlePermission();

        let loc = await Geolocation.getCurrentPosition({
            enableHighAccuracy: true,
            minimumUpdateInterval: 1,
        });

        // Initialize map
        map = L.map("map", {
            zoomControl: false,
            attributionControl: false,
            boxZoom: false,
            touchZoom: true,
            center: [52.1436278895767, 5.543447534013997],
        });

        // Set the view to the current location if available
        if (loc != null) {
            map.setView([loc.coords.latitude, loc.coords.longitude], 16);
        }

        // Add the tile layer
        L.tileLayer(
            "https://{s}.basemaps.cartocdn.com/rastertiles/voyager/{z}/{x}/{y}{r}.png",
            {
                maxZoom: 18,
                minZoom: 16,
            },
        ).addTo(map);
    });

    function calculateDistance(locations) {
        if (locations.length < 2) return 0;

        const haversine = (lat1, lon1, lat2, lon2) => {
            const R = 6371; // Earth radius in km
            const φ1 = lat1 * (Math.PI / 180); // Convert degrees to radians
            const φ2 = lat2 * (Math.PI / 180);
            const Δφ = (lat2 - lat1) * (Math.PI / 180);
            const Δλ = (lon2 - lon1) * (Math.PI / 180);

            const a =
                Math.sin(Δφ / 2) * Math.sin(Δφ / 2) +
                Math.cos(φ1) *
                    Math.cos(φ2) *
                    Math.sin(Δλ / 2) *
                    Math.sin(Δλ / 2);
            const c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a));

            return R * c; // Returns the distance in kilometers
        };

        return locations.reduce((acc, curr, idx, arr) => {
            if (idx === 0) return acc; // Skip the first point (no previous point to calculate distance)
            const prev = arr[idx - 1];
            const distance = haversine(
                prev.latitude,
                prev.longitude,
                curr.latitude,
                curr.longitude,
            );
            return acc + distance;
        }, 0);
    }
</script>

<svelte:head>
    <title>Add Ride - Skatetracker</title>
</svelte:head>

<Header title="Ride" disableBack={isTracking} />

<ion-content fullscreen>
    <div class="map_section" id="mapContainer">
        <div id="map"></div>
    </div>
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <!-- svelte-ignore a11y-no-static-element-interactions -->
    <ion-button
        class="ion-margin-top"
        color="primary"
        expand="block"
        on:click={() => {
            isModalOpen = true;
        }}
    >
        Open Details
    </ion-button>

    <!-- Make model set is-open to false on dismiss -->
    <ion-modal
        is-open={isModalOpen}
        initial-breakpoint={isTracking ? 0.3 : 0.5}
        breakpoints={[0.3, 0.5, 0.7]}
        can-dismiss={true}
        on:ionModalDidDismiss={() => {
            isModalOpen = false;
        }}
    >
        <ion-content>
            {#if isTracking}
                <div class="stats">
                    <ion-item>
                        <ion-label>Ride time:</ion-label>
                        <ion-note slot="end">
                            {#if startTime}
                                {formatRideDuration(elapsed)}
                            {:else}
                                00:00:00
                            {/if}
                        </ion-note>
                    </ion-item>
                    <ion-item>
                        <ion-label>Average Speed:</ion-label>
                        <ion-note slot="end">
                            {$locations.length > 0
                                ? (
                                      ($locations.reduce(
                                          (sum, loc) => sum + loc.speed,
                                          0,
                                      ) /
                                          $locations.length) *
                                      3.6
                                  ).toFixed(2)
                                : "0"} km/h</ion-note
                        >
                    </ion-item>
                    <ion-item>
                        <ion-label>Top Speed:</ion-label>
                        <ion-note slot="end"
                            >{$locations.length > 0
                                ? (
                                      Math.max(
                                          ...$locations.map((loc) => loc.speed),
                                      ) * 3.6
                                  ).toFixed(2)
                                : "0"} km/h</ion-note
                        >
                    </ion-item>
                    <ion-item>
                        <ion-label>Total distance:</ion-label>
                        <ion-note slot="end">
                            {calculateDistance($locations).toFixed(2)} km</ion-note
                        >
                    </ion-item>
                </div>
                <br />
            {/if}
            <div class="forms">
                <ion-item>
                    <ion-input
                        label="Title:"
                        value={rideName}
                        on:ionInput={(e) => (rideName = e.target.value)}
                        type="text"
                        name="text"
                    />
                </ion-item>

                <ion-item>
                    <ion-input
                        label="Description:"
                        value={rideDescription}
                        on:ionInput={(e) => (rideDescription = e.target.value)}
                        type="text"
                        name="text"
                    >
                    </ion-input>
                </ion-item>
                {#if $page.data.boards.length == 1 || isTracking}
                    <ion-item>
                        <ion-input
                            label="Board:"
                            value={selectedBoard.nickname}
                            type="text"
                            name="text"
                            disabled="true"
                        >
                        </ion-input></ion-item
                    >
                {/if}
            </div>
            <!-- svelte-ignore a11y-click-events-have-key-events -->
            <!-- svelte-ignore a11y-no-static-element-interactions -->
            {#if $page.data.boards.length > 1 && !isTracking}
                <ion-picker
                    on:ionChange={(e) => {
                        selectedBoard = e.detail.value;
                    }}
                >
                    <ion-picker-column value={selectedBoard}>
                        {#each $page.data.boards as board}
                            <ion-picker-column-option value={board}
                                >{board.nickname}</ion-picker-column-option
                            >
                        {/each}
                    </ion-picker-column>
                </ion-picker>
            {/if}
            <br />
            <!-- svelte-ignore a11y-click-events-have-key-events -->
            <!-- svelte-ignore a11y-no-static-element-interactions -->
            <ion-button
                class="track_button"
                expand="block"
                color={isTracking ? "danger" : "success"}
                on:click={() => {
                    if (isTracking) {
                        stopTracking();
                    } else {
                        startTracking();
                    }
                }}
            >
                {isTracking ? "Stop Tracking" : "Start Tracking"}
            </ion-button>
        </ion-content>
    </ion-modal>
</ion-content>

<style>
    ion-content {
        --background: var(--ion-color-light);
    }

    #map {
        height: calc(90vh - 56px);
        z-index: 1000;
    }
</style>
