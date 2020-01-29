var https = require('https');
// actions accept a single parameter, which must be a JSON object.
// grab a random top meme and post data to slack
function main(params) {
  return new Promise(function(resolve, reject) {
    https.get("https://www.reddit.com/r/ProgrammerHumor.json", (resp) => {
      let data = '';
      // chunk of data has been recieved.
      resp.on('data', (chunk) => {
        data += chunk;
      });

      let reddit = {};
      // whole response has been received. Print out the result.
      resp.on('end', () => {
        // console.log(data);
        let myUrl = "";
        let reddit = JSON.parse(data);
        let fileExt = "";
        let x = 0;
        myUrl = "";

        // grab a random top post
        while (fileExt != 'png' && fileExt != 'jpg') {
          x = Math.floor(Math.random() * 25);
          let url = reddit.data.children[x].data.url;

          myUrl = url;
          let str = url.split('.');
          fileExt = str[str.length - 1];
        }
        console.log('MYURL=' + myUrl);

        // setup post to be title + url
        let post = reddit.data.children[x].data.title + " " + myUrl;
        console.log(post);

        var postData = JSON.stringify({
          text: post
        });

        // Slack setup #spicy_memes
        var options = {
          hostname: 'hooks.slack.com',
          path: env.SLACK_HOOK_PATH_MEMES,
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
