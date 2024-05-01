document.addEventListener('DOMContentLoaded', function () {
  const stationsDataElement = document.getElementById('station-data')
  const stationsData = stationsDataElement.getAttribute('data-stations')
  const lineStringsData = stationsDataElement.getAttribute('data-linestrings')

  const stations = JSON.parse(stationsData)
  let lineStrings = JSON.parse(lineStringsData)

  lineStrings = lineStrings.map((line) =>
    JSON.parse(line)
      .flat()
      .map(([lon, lat]) => [lat, lon])
  )

  const initialMapPosition = stations[0]
    ? [stations[0].lat, stations[0].lon]
    : [51.5074, -0.1278]

  const map = L.map('map').setView(initialMapPosition, 12)

  L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    maxZoom: 19,
    attribution:
      '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>',
  }).addTo(map)

  const polyline = L.polyline(lineStrings, { color: 'red' }).addTo(map)
  map.fitBounds(polyline.getBounds())

  stations.forEach((station) => {
    L.marker([station.lat, station.lon]).addTo(map).bindPopup(station.name)
  })
})
