/**
 *
 * main() will be run when you invoke this action
 *
 * @param Cloud Functions actions accept a single parameter, which must be a JSON object.
 *
 * @return The output of this action, which must be a JSON object.
 *
 */
var https = require('https');

function main(params) {

  return new Promise(function(resolve, reject) {
    https.get("https://c.xkcd.com/random/comic/", (resp) => {
      let data = '';

      // A chunk of data has been recieved.
      resp.on('data', (chunk) => {
        data += chunk;
      });

      resp.on('end', () => {

        console.log('headers are: ', resp.headers);
        const location = resp.headers.location


        resolve({
          location: location
        });
      });


    }).on("error", (err) => {
      console.log("Error: " + err.message);
    });

  });
}
