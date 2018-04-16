<html>
<center>
<head>
       <title>Upload CSV, Download XML</title>
</head>
<body>
<h1>Upload a CSV file, exported from generations</h1>
<h2>You will get a XML file back, import that into EZ Claim<h2>
<form enctype="multipart/form-data" action="/upload" method="post">
    <input type="file" name="uploadfile" />
    <input type="hidden" name="token" value="{{.}}"/>
    <input type="submit" value="upload" />
</form>
</body>
</center>
</html>
