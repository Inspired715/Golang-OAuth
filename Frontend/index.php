<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
</head>

<body>
	<form method="post" action="http://localhost:8000/connect" id="google_form">
    	<img src="./google-plus.png" onclick="google_connection()"/>
	</form>
	<a href="welcome.php">Welcome page</a>
</body>
<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.3/jquery.min.js"></script>
<script>
	const hash = window.location.hash.substring(1);
	const searchparam = new URLSearchParams(hash);
	const token = searchparam.get('access_token');
	if(token)
		localStorage.setItem('token', token);

	function google_connection() {
		$('#google_form').submit();
	}

</script>
</html>