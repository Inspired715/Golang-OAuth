<!DOCTYPE html>
<html>
<head>
<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.7.1/jquery.min.js"></script>
</head>
<body>
<script>
    $(document).ready(function(){
      $.ajax({
            url: 'http://localhost:8000/getList',
            type: 'GET',
            dataType: 'json',
            headers: {
                'x-api-key': localStorage.getItem('token'),
            },
            success: function (data) {
                if(data.code == 200){
                  var newDiv = document.createElement("div");
                  newDiv.id = "dynamicDiv";
                  newDiv.className = "dynamic-class";
                  var newText = document.createTextNode("This is a dynamically created div.");

                  temp = data.data.results;
                  temp.forEach(element => {
                    var newParagraph = document.createElement("p");
                    var newText = document.createTextNode(element.Description + " " + element.Title);
                    newParagraph.appendChild(newText);
                    newDiv.appendChild(newParagraph);
                  });

                  document.body.appendChild(newDiv);
                }
            },
            error: function(xhr, textStatus, errorThrown) {
                var errorResponse = JSON.parse(xhr.responseText);
                console.log(errorResponse.code, ' : ', errorResponse.msg);          
            }
        });
    });
</script>
</body>
</html>