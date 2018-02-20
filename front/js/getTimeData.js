function getTimeData(){
    let req = new XMLHttpRequest();

    req.onreadystatechange = function() {
      if (req.readyState == 4) { // 通信の完了時
        if (req.status == 200) { // 通信の成功時
            timeTable = JSON.parse(req.responseText);
            firstView();
            document.getElementById("loading").style.display = "none";
            document.getElementById("last").style.display = "block";
        }else if(req.status == 404){
            //not found
        }
      }else{
          document.getElementById("loading").style.display = "block";
      }
    }

    //dataはtrueがgo,falseがreturn
    switch (mode) {
        case 0:
        req.open('GET', 'http://marco.plus:7650/api/goFull' , true);
            break;
        case 1:
        req.open('GET', 'http://marco.plus:7650/api/returnFull' , true);
            break;
    }
    req.send(null);
}
