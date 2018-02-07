function getTimeData(data){
    var req = new XMLHttpRequest();

    req.onreadystatechange = function() {
      if (req.readyState == 4) { // 通信の完了時
        if (req.status == 200) { // 通信の成功時
            displayTime();
            document.getElementById("loading").style.display = "none";
        }
      }else{
          document.getElementById("loading").style.display = "block";
      }
    }

    //dataはtrueがgo,falseがreturn
    req.open('GET', 'localhost:8080?data=' + data , true);
    req.send(null);
    timeTable = eval('(' + req.responseText + ')');;
}
