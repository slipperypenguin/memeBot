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
    https.get("https://www.reddit.com/r/ProgrammerHumor.json", (resp) => {
      let data = '';

      // A chunk of data has been recieved.
      resp.on('data', (chunk) => {
        data += chunk;
      });

      let reddit = {};

      // The whole response has been received. Print out the result.
      resp.on('end', () => {
        // console.log(data);
        let myUrl = "";
        let reddit = JSON.parse(data);
        let fileExt = "";
        let x = 0;
        myUrl = "";

        // Grab a random top post
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

        // Slack Setup #spicy_memes
        var options = {
          hostname: 'hooks.slack.com',
          path: '/services/T71A2UYUT/BCNKD9RUM/AUNrgK8NzEwtwiI9EHSGBNAw',
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            'Content-Length': Buffer.byteLength(postData)
          }
        };
        // console.log(postData);

        // Setup the request
        var req = https.request(options, (res) => {
          console.log('statusCode:', res.statusCode);
          console.log('headers:', res.headers);
        });

        // Handle Errors
        req.on('error', (e) => {
          console.error(e);
        });

        req.end(postData);


        // ********************
        // This section is for @pmay's side channel
        // Should probably move this into it's own action
        // If we keep it here, get rid of duplicate code
        // ********************
        fileExt = "";
        myUrl = "";
        while (fileExt != 'png' && fileExt != 'jpg') {
          x = Math.floor(Math.random() * 25);
          let url = reddit.data.children[x].data.url;

          myUrl = url;
          let str = url.split('.');
          fileExt = str[str.length - 1];
        }
        // console.log('MYURL=' + myUrl)

        // setup post to be title + url
        post = reddit.data.children[x].data.title + " " + myUrl
        console.log(post)

        postData = JSON.stringify({
          text: post
        })

        // Slack Setup
        options = {
          hostname: 'hooks.slack.com',
          path: '/services/T812P9HB3/BCN2NNC65/RZ2JkT4qLTFAYllsoazZCp9U',
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            'Content-Length': Buffer.byteLength(postData)
          }
        };
        // Replacement paths for env_vars
        // path: params.SLACK_HOOK_PATH_MEMES,
        // path: params.SLACK_HOOK_PATH_PETER,

        // Setup the request
        req = https.request(options, (res) => {
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
