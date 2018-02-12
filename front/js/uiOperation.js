mytap = window.ontouchstart===null?"touchstart":"click";
mode=0;
from=0;
//to=0;
timeTable={};
fromId=[["センター","駅前"],["学内","路線","シャトル"]];
//ToId=["直行","青年の家","鳶尾団地"];

document.getElementById("go").addEventListener(mytap, function(){
    mode=0;
    document.getElementById("mode").innerHTML = "行き";
    nextStep();
}, false);

document.getElementById("return").addEventListener(mytap, function(){
    mode=1;
    document.getElementById("mode").innerHTML = "帰り";
    nextStep();
    //document.getElementById("to").style.display = "none";
}, false);


function nextStep(){
    document.getElementById("first").style.display = "none";
    getTimeData();
}

function firstView(){
    var tmpH=timeTable.fast[0]%60;
    var tmpM=timeTable.fast[0]-tmpH*60;
    var tmpTime=getTime();
    tmpTime = timeTable.fast[0] - tmpTime;
    displayTime(tmpH,tmpM,tmpTime,[timeTable.fast[1]]);
}

function displayTime(h,m,inmin,id){
    if(checkExistBus(h,m)){
        document.getElementById("time").style.fontSize = "9vh";
        document.getElementById("time").style.marginTop = "4vh";
        document.getElementById("stop").style.marginTop = "2vh";
        document.getElementById("time").innerHTML = h + ":" + m;
        document.getElementById("inmin").innerHTML = "約" + inmin + "分後";
    }
    document.getElementById("stop").innerHTML = "乗り場:" + fromId[mode][id];
}

function checkExistBus(h,m){
    if(h===0 && m===0){
        document.getElementById("time").style.fontSize = "8vh";
        document.getElementById("time").style.marginTop = "6vh";
        document.getElementById("stop").style.marginTop = "4vh";
        document.getElementById("time").innerHTML = "バス終わり";
        document.getElementById("inmin").innerHTML = "";
        return false
    }
    return true
}

function getTime(){
    var time = new Date();
    return time.getHours()*60+time.getMinutes();
}

document.getElementById("from").addEventListener(mytap, function(){
    //to = 0;
    var tmpH=Math.floor(timeTable.from[from]/60);
    var tmpM=timeTable.from[from]-tmpH*60;
    var tmpTime=getTime();
    tmpTime = timeTable.from[from] - tmpTime;

    displayTime(tmpH,tmpM,tmpTime,from);
    document.getElementById("title").innerHTML = fromId[mode][from]+"発最速"

    if(mode === 0 && from === 0){
        document.getElementById("from").style.backgroundColor = "rgba(189,189,189 ,1)"
        from++;
    }else if(mode === 0 && from === 1){
        document.getElementById("from").style.backgroundColor = "rgba(158,158,158 ,1)"
        from = 0;
    }else if(mode === 1 && from < 2) {
        if(from===0){
            document.getElementById("from").style.backgroundColor = "rgba(224,224,224 ,1)"
        }else if(from===1){
            document.getElementById("from").style.backgroundColor = "rgba(189,189,189 ,1)"
        }
        from++;
    }else if(mode === 1 && from === 2){
        document.getElementById("from").style.backgroundColor = "rgba(158,158,158 ,1)"
        from = 0;
    }
}, false);

/*
document.getElementById("to").addEventListener(mytap, function(){
    from = 0;
    var tmpH=floor(timeTable.to[to]/60);
    var tmpM=timeTable.to[to]-tmpH*60;
    var tmpTime=getTime();
    tmpTime -= timeTable.to[to];

    displayTime(tmpH,tmpM,tmpTime,to);

    if(to < 2) {
        to++;
    }else if(to === 2){
        to = 0;
    }
}, false);
*/
