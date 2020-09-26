const puppeteer = require('puppeteer');
const Redis = require('ioredis');
const connection = new Redis(6379, 'redis');

const crawl = async (username) => {
    // initialize
    const browser = await puppeteer.launch({
        executablePath: 'google-chrome-unstable',
        headless: true,
        args: [
            '--no-sandbox',
            '--disable-background-networking',
            '--disk-cache-dir=/dev/null',
            '--disable-default-apps',
            '--disable-extensions',
            '--disable-gpu',
            '--disable-sync',
            '--disable-translate',
            '--hide-scrollbars',
            '--metrics-recording-only',
            '--mute-audio',
            '--no-first-run',
            '--safebrowsing-disable-auto-update',
        ],
    });
    const page = await browser.newPage();

    // set cookie
    await page.setCookie({
        name: 'flag',
        value: process.env.FLAG,
        domain: process.env.DOMAIN,
        expires: Date.now() / 1000 + 10,
    });

    // access
    const url = `${process.env.SCHEME}://${process.env.DOMAIN}/?username=${encodeURIComponent(username)}`;
    try {
        await page.goto(url, {
            waitUntil: 'networkidle0',
            timeout: 5000,
        });
    } catch (err) {
        console.log(err);
    }

    // finalize
    await page.close();
    await browser.close();
};

(async () => {
    while (true) {
        console.log("[*] waiting new query ...");
        await connection.blpop("query", 0).then(v => {
            const username = v[1];
            console.log(`[*] started: ${username}`);
            return crawl(username);
        }).then(() => {
            console.log(`[*] finished.`)
            return connection.incr("proceeded_count");
        }).catch(e => {
            console.log(e);
        });
    };
})();
