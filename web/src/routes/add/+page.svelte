<script>
    import { onMount } from "svelte";
    import Header from "$lib/components/Header.svelte";
    import L from "leaflet";
    import "leaflet/dist/leaflet.css";

    import { Geolocation } from "@capacitor/geolocation";

    var loc = null;

    async function getCurrentPosition() {
        try {
            const res = await Geolocation.getCurrentPosition({
                enableHighAccuracy: true,
            });
            loc = res;
        } catch (e) {
            console.log(e);
        }
    }

    onMount(async () => {
        await getCurrentPosition();

        var map = L.map("map", {
            zoomControl: false,
            attributionControl: false,
            boxZoom: false,
            touchZoom: true,
            center: [52.1436278895767, 5.543447534013997],
        });

        if (loc != null) {
            map.setView([loc.coords.latitude, loc.coords.longitude], 13);
        }
        L.tileLayer(
            "https://{s}.basemaps.cartocdn.com/rastertiles/voyager/{z}/{x}/{y}{r}.png",
            {
                maxZoom: 18,
                minZoom: 16,
            },
        ).addTo(map);
    });
</script>

<svelte:head>
    <title>Add Ride - Skatetracker</title>
</svelte:head>

<Header title="Ride" />

<ion-content fullscreen>
    <div id="mapContainer">
        <div id="map"></div>
    </div>
</ion-content>

<style>
    ion-content {
        --background: var(--ion-color-light);
    }

    #map {
        height: calc(100vh - 100px);
    }
</style>
