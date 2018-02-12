function getTimeData(){
    var req = new XMLHttpRequest();

    req.onreadystatechange = function() {
      if (req.readyState == 4) { // 通信の完了時
        if (req.status == 200) { // 通信の成功時
            timeTable = req.responseText;
            displayTime();
            document.getElementById("loading").style.display = "none";
        }
      }else{
          document.getElementById("loading").style.display = "block";
      }
    }

    //dataはtrueがgo,falseがreturn
    switch (mode) {
        case 0:
        req.open('GET', 'localhost:7650/api/goFull' , false);
            break;
        case 1:
        req.open('GET', 'localhost:7650/api/returnFull' , false);
            break;:
    }
    req.send(null);
}
