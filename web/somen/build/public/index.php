<?php
$nonce = base64_encode(random_bytes(20));
header("Content-Security-Policy: default-src 'none'; script-src 'nonce-${nonce}' 'strict-dynamic' 'sha256-nus+LGcHkEgf6BITG7CKrSgUIb1qMexlF8e5Iwx1L2A='");
?>

<head>
    <title>Best somen for <?= isset($_GET["username"]) ? $_GET["username"] : "You" ?></title>

    <script src="/security.js" integrity="sha256-nus+LGcHkEgf6BITG7CKrSgUIb1qMexlF8e5Iwx1L2A="></script>
    <script nonce="<?= $nonce ?>">
        const choice = l => l[Math.floor(Math.random() * l.length)];

        window.onload = () => {
            const username = new URL(location).searchParams.get("username");
            const adjective = choice(["Nagashi", "Hiyashi"]);
            if (username !== null)
                document.getElementById("message").innerHTML = `${username}, I recommend ${adjective} somen for you.`;
        }
    </script>
</head>

<body>
    <h1>Best somen for You</h1>

    <p>Please input your name. You can use only alphabets and digits.</p>
    <p>This page works fine with latest Google Chrome / Chromium. We won't support other browsers :P</p>
    <p id="message"></p>
    <form action="/" method="GET">
        <input type="text" name="username" place="Your name"></input>
        <button type="submit">Ask</button>
    </form>
    <hr>

    <p> If your name causes suspicious behavior, please tell me that from the following form. Admin will acceess /?username=${encodeURIComponent(your input)} and see what happens.</p>
    <form action="/inquiry" method="POST">
        <input type="text" name="username" place="Your name"></input>
        <button type="submit">Ask</button>
    </form>

</body>
