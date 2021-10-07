
const express = require('express')
const redis = require('redis');
var crypto = require('crypto');

const app = express()

app.use(express.urlencoded());
app.use(express.json());

const port = 3000


const client = redis.createClient({
  host: 'redis',
  port: 6379,
});

client.on('error', err => {
  console.log('Redis Error ' + err);
});


app.post('/node/sha256', (req, res) => {
  if (!("Input" in req.body)) {
    res.status(422);
    res.send({ Result: "Bad Parameter" })
  } else if (req.body.Input.length < 8) {
    res.status(422);

    res.send({ Result: "At least 8 chars required" })
  }
  else {
    const input = req.body.Input;
    const sha = crypto.createHash('sha256').update(input).digest('hex');

    client.SET(input, sha)
    client.SET(sha, input)
    res.status(200);
    res.send({ Input: input, SHA: sha })
  }
})

app.get('/node/sha256', (req, res) => {
  if (!("sha" in req.query)) {
    res.status(422);
    res.send({ Result: "Bad Parameter" })
  } else {
    const sha = String(req.query.sha);

    client.GET(sha, function (err, reply) {
      if (err || reply == null) {
        res.status(404)
        res.send({ "Result": "Not available in table" })
      }
      else {
        res.status(200);
        res.send({ Input: reply, SHA: sha })
      }
    })
  }
})

app.listen(port, () => {
  console.log("Node Started at port " + port)
})
