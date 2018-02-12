var mode=0;
var mytap = window.ontouchstart===null?"touchstart":"click";
var timeTable;
var fromId=[["センター","駅前"],["学内","路線","シャトル"]];
var ToId=["直行","青年の家","鳶尾団地"];

document.getElementById("go").addEventListener(mytap, function(){
    nextStep(0);
}, false);

document.getElementById("return").addEventListener(mytap, function(){
    nextStep(1);
    document.getElementById("to").style.display = "none";
}, false);


function nextStep(data){
    mode = data;
    getTimeData();
    document.getElementById("first").style.display = "none";
    document.getElementById("last").style.display = "block";
}

function displayTime(){
    var tmpH=timeTable.fast[0]%60;
    var tmpM=timeTable.fast[0]-tmpH*60;
    document.getElementById("time").innerHTML = tmpH + ":" + tmpM;
    var tmpTime=getHours()*60+getMinutes();
    document.getElementById("inmin").innerHTML = tmpTime - timeTable.fast[0];
    document.getElementById("stop").innerHTML = fromId[mode][timeTable.fast[1]]; //goかreturnをmodeで判断し添字はjsonより//
}
