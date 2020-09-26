// set up redis
const Redis = require('ioredis');
const connection = new Redis(6379, 'redis');
connection.set('queued_count', 0);
connection.set('proceeded_count', 0);

// set up express
const app = require('express')();
const port = 8080;

const bodyParser = require('body-parser');
app.use(bodyParser.urlencoded({ extended: false }))

app.post('/inquiry', (req, res) => {
    try {
        if (req.body.username !== undefined && req.body.username != '') {
            connection.rpush('query', req.body.username).then(() => {
                connection.incr('queued_count')
            }).then(() => {
                console.log(`[*] Queried: ${req.body.username}`);
                res.send('Okay! I got it :-)');
            });
        } else {
            throw Error;
        }
    } catch (e) {
        res.send('Umm, there is something wrong ...');
    }
});

app.listen(port, () => {
    console.log(`Publisher is listening on port ${port}!`);
});
