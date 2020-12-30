const createClient = require('hafas-client')
const vbbProfile = require('hafas-client/p/insa')

const client = createClient(vbbProfile, 'my-awesome-program')
client.departures('461', {duration: 3}).then(console.log).catch(console.error)
