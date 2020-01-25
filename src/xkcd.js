var https = require('https');
// actions accept a single parameter, which must be a JSON object.
// accept url, posts data to slack
function main(params) {
  const source_url = params.location + 'info.0.json';
  console.log('source url is: ', source_url);

  return new Promise(function(resolve, reject) {
    https.get(source_url, (resp) => {
      let data = '';
      // chunk of data has been recieved.
      resp.on('data', (chunk) => {
        data += chunk;
      });

      let xkcd = {};
      // whole response has been received. Print out the result.
      resp.on('end', () => {
        //console.log('data is: ', data);
        let xkcd = JSON.parse(data);
        let imgUrl = xkcd.img;
        let imgTitle = xkcd.title;
        //console.log('imgUrl is: ', imgUrl);
        //console.log('imgTitle is: ', imgTitle);

        // setup post to be title + url
        let post = imgTitle + " " + imgUrl;
        //console.log(post);

        var postData = JSON.stringify({
          text: post
        });

        // Slack setup #spicy_memes
        var options = {
          hostname: 'hooks.slack.com',
          path: params.SLACK_HOOK_PATH_MEMES,
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            'Content-Length': Buffer.byteLength(postData)
          }
        };
        // console.log(postData);

        // setup the request
        var req = https.request(options, (res) => {
          console.log('statusCode:', res.statusCode);
          console.log('headers:', res.headers);
        });

        // Handle Errors
        req.on('error', (e) => {
          console.error(e);
        });
        req.end(postData);

        resolve({
          done: true
        });
      });
    }).on("error", (err) => {
      console.log("Error: " + err.message);
    });
  });
}
