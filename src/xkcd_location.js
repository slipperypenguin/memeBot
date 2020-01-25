var https = require('https');
// actions accept a single parameter, which must be a JSON object.
// returns a JSON object containing url to be passed to next action.
function main(params) {
  return new Promise(function(resolve, reject) {
    https.get("https://c.xkcd.com/random/comic/", (resp) => {
      let data = '';

      // chunk of data has been recieved.
      resp.on('data', (chunk) => {
        data += chunk;
      });

      resp.on('end', () => {
        console.log('headers are: ', resp.headers);
        const location = resp.headers.location;

        resolve({
          location: location
        });
      });
    }).on("error", (err) => {
      console.log("Error: " + err.message);
    });
  });
}
