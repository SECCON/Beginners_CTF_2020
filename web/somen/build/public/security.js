console.log('!! security.js !!');
const username = new URL(location).searchParams.get("username");
if (username !== null && ! /^[a-zA-Z0-9]*$/.test(username)) {
    document.location = "/error.php";
}