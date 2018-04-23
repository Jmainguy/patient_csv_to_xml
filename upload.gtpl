<html>
<center>
<head>
       <title>Upload CSV, Download XML</title>
</head>
<body>
<h1>Upload a CSV file, exported from generations</h1>
<h2>You will get a XML file back, import that into EZ Claim<h2>
<form enctype="multipart/form-data" action="/upload" method="post">
    <h3>Choose an office</h3>
    <select name="patClass">
      <option value="AHHCDUNN">AHHCDUNN</option>
      <option value="AHHCSMF">AHHCSMF</option>
      <option value="AHHCVB">AHHCVB</option>
      <option value="AHHCWBG">AHHCWBG</option>
      <option value="AHHCRICH">AHHCRICH</option>
      <option value="AHHCALEX">AHHCALEX</option>
      <option value="AHHCHCS">AHHCHCS</option>
      <option value="AHHCAUGE">AHHCAUGE</option>
    </select>
    <br><br>
    <h3>Choose a csv to upload</h3>
    <input type="file" name="uploadfile" />
    <input type="hidden" name="token" value="{{.}}"/>
    <input type="submit" value="upload" />
</form>
</body>
</center>
</html>
