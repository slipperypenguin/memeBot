/**
 *
 * main() will be run when you invoke this action
 *
 * @param Cloud Functions actions accept a single parameter, which must be a JSON object.
 *
 * @jcates (ooo, return 10/1) The output of this action, which must be a JSON object.
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

      let reddit = {}

      // The whole response has been received. Print out the result.
      resp.on('end', () => {
        // console.log(data);
        let myUrl = ""
        let reddit = JSON.parse(data)
        let fileExt = ""
        let x = 0
        myUrl = ""
        while (fileExt != 'png' && fileExt != 'jpg') {

          x = Math.floor(Math.random() * 25)
          let url = reddit.data.children[x].data.url

          myUrl = url
          let str = url.split('.')
          fileExt = str[str.length - 1]
        }

        console.log('MYURL=' + myUrl)

        let post = reddit.data.children[x].data.title + " " + myUrl
        console.log(post)

        var postData = JSON.stringify({
          text: post
        })

        var options = {
          hostname: 'hooks.slack.com',
          path: '/services/T71A2UYUT/BCNKD9RUM/AUNrgK8NzEwtwiI9EHSGBNAw',
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            'Content-Length': Buffer.byteLength(postData)
          }
        };

        // console.log(postData)

        var req = https.request(options, (res) => {
          console.log('statusCode:', res.statusCode);
          console.log('headers:', res.headers);
        });

        req.on('error', (e) => {
          console.error(e);
        });

        req.end(postData)

        fileExt = ""
        myUrl = ""
        while (fileExt != 'png' && fileExt != 'jpg') {

          x = Math.floor(Math.random() * 25)
          let url = reddit.data.children[x].data.url

          myUrl = url
          let str = url.split('.')
          fileExt = str[str.length - 1]

        }

        // console.log('MYURL=' + myUrl)

        post = reddit.data.children[x].data.title + " " + myUrl
        console.log(post)

        postData = JSON.stringify({
          text: post
        })

        options = {
          hostname: 'hooks.slack.com',
          path: '/services/T812P9HB3/BCN2NNC65/RZ2JkT4qLTFAYllsoazZCp9U',
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            'Content-Length': Buffer.byteLength(postData)
          }
        };

        req = https.request(options, (res) => {
          console.log('statusCode:', res.statusCode);
          console.log('headers:', res.headers);
        });

        req.on('error', (e) => {
          console.error(e);
        });

        req.end(postData)


        resolve({
          done: true
        });
      });


    }).on("error", (err) => {
      console.log("Error: " + err.message);
    });
  });
}
