const puppeteer = require('puppeteer');

/* ... ... */

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
// username is the input value of players
const url = `https://somen.quals.beginners.seccon.jp/?username=${encodeURIComponent(username)}`;
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

/* ... ... */
