document.addEventListener('DOMContentLoaded', function () {
  const stationsDataElement = document.getElementById('station-data')
  const stationsData = stationsDataElement.getAttribute('data-stations')

  const stations = JSON.parse(stationsData)

  const initialMapPosition = stations[0]
    ? [stations[0].lat, stations[0].lon]
    : [51.5074, -0.1278]

  const map = L.map('map').setView(initialMapPosition, 12)

  L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    maxZoom: 19,
    attribution:
      '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>',
  }).addTo(map)

  stations.forEach((station) => {
    L.marker([station.lat, station.lon]).addTo(map).bindPopup(station.name)
  })
})
