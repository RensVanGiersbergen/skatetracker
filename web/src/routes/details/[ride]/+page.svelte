<script>
    import { page } from "$app/stores";
    import Header from "$lib/components/Header.svelte";
    import { onMount } from "svelte";
    import L from "leaflet";
    import "leaflet/dist/leaflet.css";
    import "leaflet.heat";
    import dayjs from "dayjs";

    let segment = "speed";
    let speedMap = null;
    let shakinessMap = null;

    onMount(async () => {
        await loadSpeedMap();
    });

    async function loadSpeedMap() {
        setTimeout(() => {
            speedMap.invalidateSize();
        }, 100);
        // Create speed map
        speedMap = L.map("speedMap", {
            zoomControl: false,
            attributionControl: false,
            boxZoom: false,
            touchZoom: true,
        });

        // Add the tile layer
        L.tileLayer(
            "https://{s}.basemaps.cartocdn.com/rastertiles/voyager/{z}/{x}/{y}{r}.png",
            {
                maxZoom: 20,
                minZoom: 12,
            },
        ).addTo(speedMap);

        // Set bounds around ride
        speedMap.fitBounds(await getBounds($page.data.ride.trackings));

        // Set max bounds
        let bounds = await getBounds($page.data.ride.trackings);
        // Extend bounds to allow for panning
        bounds[0][0] -= 0.01;
        bounds[0][1] -= 0.01;
        bounds[1][0] += 0.01;
        bounds[1][1] += 0.01;
        speedMap.setMaxBounds(bounds);

        // Add speed to map using polyline
        $page.data.ride.trackings.forEach((tracking, index) => {
            if (index === 0) {
                return;
            }

            let previousTracking = $page.data.ride.trackings[index - 1];
            let latlngs = [
                [previousTracking.latitude, previousTracking.longitude],
                [tracking.latitude, tracking.longitude],
            ];

            L.polyline(latlngs, {
                color:
                    tracking.speed > 13.89
                        ? "#f56042"
                        : tracking.speed > 11.11
                          ? "#f5b642"
                          : tracking.speed > 8.33
                            ? "#f5e942"
                            : tracking.speed > 5.56
                              ? "#c8f542"
                              : "#72f542",
            }).addTo(speedMap);
        });

        // Add start marker
        let startMarker = L.marker(
            [
                $page.data.ride.trackings[0].latitude,
                $page.data.ride.trackings[0].longitude,
            ],
            {
                icon: L.icon({
                    iconUrl: "/icons/startflag.svg",
                    iconSize: [30, 30],
                    iconAnchor: [7.5, 30],
                }),
            },
        ).addTo(speedMap);

        // Add popup to start marker
        startMarker.bindPopup(
            `<b>Started</b><br>time: ${dayjs($page.data.ride.start_time).format(
                "HH:mm",
            )}`,
        );

        // Add finish marker
        let finishMarker = L.marker(
            [
                $page.data.ride.trackings[$page.data.ride.trackings.length - 1]
                    .latitude,
                $page.data.ride.trackings[$page.data.ride.trackings.length - 1]
                    .longitude,
            ],
            {
                icon: L.icon({
                    iconUrl: "/icons/finishflag.svg",
                    iconSize: [30, 30],
                    iconAnchor: [5, 27.5],
                }),
            },
        ).addTo(speedMap);

        // Add popup to start marker
        finishMarker.bindPopup(
            `<b>Finished</b><br>time: ${dayjs($page.data.ride.end_time).format(
                "HH:mm",
            )}`,
        );

        // Add fastest point marker
        let fastestTracking = $page.data.ride.trackings.reduce((a, b) =>
            a.speed > b.speed ? a : b,
        );
        let fastestMarker = L.marker(
            [fastestTracking.latitude, fastestTracking.longitude],
            {
                icon: L.icon({
                    iconUrl: "/icons/lightning.svg",
                    iconSize: [40, 40],
                }),
            },
        ).addTo(speedMap);

        // Add popup to fastest point marker
        fastestMarker.bindPopup(
            `<b>Fastest point</b><br>Speed: ${(
                fastestTracking.speed * 3.6
            ).toFixed(2)} km/h`,
        );
    }

    async function loadShakinessMap() {
        setTimeout(() => {
            shakinessMap.invalidateSize();
        }, 100);
        // Create speed map
        shakinessMap = L.map("shakinessMap", {
            zoomControl: false,
            attributionControl: false,
            boxZoom: false,
            touchZoom: true,
        });

        // Add the tile layer
        L.tileLayer(
            "https://{s}.basemaps.cartocdn.com/rastertiles/voyager/{z}/{x}/{y}{r}.png",
            {
                maxZoom: 20,
                minZoom: 12,
            },
        ).addTo(shakinessMap);

        // Set bounds around ride
        shakinessMap.fitBounds(await getBounds($page.data.ride.trackings));

        // Set max bounds
        let bounds = await getBounds($page.data.ride.trackings);
        // Extend bounds to allow for panning
        bounds[0][0] -= 0.01;
        bounds[0][1] -= 0.01;
        bounds[1][0] += 0.01;
        bounds[1][1] += 0.01;
        shakinessMap.setMaxBounds(bounds);

        L.heatLayer(
            $page.data.ride.trackings.map((tracking) => [
                tracking.latitude,
                tracking.longitude,
                tracking.shakiness,
            ]),
            {
                radius: 40,
                gradient: {
                    0.2: "blue",
                    0.4: "lime",
                    0.6: "yellow",
                    0.8: "orange",
                    1: "red",
                },
            },
        ).addTo(shakinessMap);
    }
    async function getBounds(trackings) {
        let bottom;
        let top;
        let left;
        let right;

        for (let tracking of trackings) {
            if (!top || tracking.latitude > top) {
                top = tracking.latitude;
            }
            if (!bottom || tracking.latitude < bottom) {
                bottom = tracking.latitude;
            }
            if (!left || tracking.longitude < left) {
                left = tracking.longitude;
            }
            if (!right || tracking.longitude > right) {
                right = tracking.longitude;
            }
        }

        return [
            [bottom, left],
            [top, right],
        ];
    }
</script>

<svelte:head>
    <title>{$page.data.ride.title} - Skatetracker</title>
</svelte:head>

<Header title={$page.data.ride.title} goHome="true" />

<ion-content fullscreen>
    <ion-segment
        value={segment}
        on:ionChange={(e) => {
            segment = e.detail.value;
            setTimeout(() => {
                if (segment === "speed") {
                    loadSpeedMap();
                } else {
                    loadShakinessMap();
                }
            }, 100);
        }}
    >
        <ion-segment-button value="speed" checked="true">
            <ion-label>Speed</ion-label>
        </ion-segment-button>
        <ion-segment-button value="shakiness">
            <ion-label>Shakiness</ion-label>
        </ion-segment-button>
    </ion-segment>

    <div class="map_section" id="mapContainer">
        {#if segment === "speed"}
            <div class="map" id="speedMap"></div>
        {/if}
        {#if segment === "shakiness"}
            <div class="map" id="shakinessMap"></div>
        {/if}
    </div>

    <ion-card>
        <ion-card-header>
            <ion-card-subtitle>Top speed</ion-card-subtitle>
            <ion-card-title
                >{($page.data.ride.top_speed * 3.6).toFixed(2)} km/h</ion-card-title
            >
        </ion-card-header>
        <ion-card-content>
            <p>
                The highest speed reached during the ride. This is the highest
                speed of all tracking points.
            </p>
        </ion-card-content>
    </ion-card>

    <!--Average speed using all trackings to calculate-->
    <ion-card>
        <ion-card-header>
            <ion-card-subtitle>Average speed</ion-card-subtitle>
            <ion-card-title>
                {(
                    ($page.data.ride.trackings.reduce(
                        (sum, tracking) => sum + tracking.speed,
                        0,
                    ) /
                        $page.data.ride.trackings.length) *
                    3.6
                ).toFixed(2)} km/h
            </ion-card-title>
        </ion-card-header>
        <ion-card-content>
            <p>
                The average speed of the ride. This is calculated using all
                tracking points.
            </p>
        </ion-card-content>
    </ion-card>

    <ion-card>
        <ion-card-header>
            <ion-card-subtitle>Distance</ion-card-subtitle>
            <ion-card-title
                >{($page.data.ride.distance / 1000).toFixed(2)} km</ion-card-title
            >
        </ion-card-header>
        <ion-card-content>
            <p>
                The total distance of the ride. This is the distance between the
                first and last tracking point.
            </p>
        </ion-card-content>
    </ion-card>

    <ion-card>
        <ion-card-header>
            <ion-card-subtitle>Duration</ion-card-subtitle>
            <ion-card-title
                >{dayjs($page.data.ride.end_time).diff(
                    $page.data.ride.start_time,
                    "minute",
                )} minutes</ion-card-title
            >
        </ion-card-header>
        <ion-card-content>
            <p>
                The total duration of the ride. This is the time between the
                first and last tracking point.
            </p>
        </ion-card-content>
    </ion-card>
</ion-content>

<style>
    ion-content {
        --background: var(--ion-color-light);
    }

    .map {
        height: calc(70vh - 56px);
        z-index: 1000;
    }
</style>
