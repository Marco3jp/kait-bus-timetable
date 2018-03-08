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
}, false);


function nextStep(){
    document.getElementById("first").style.display = "none";
    getTimeData();
}

function firstView(){
    let tmpH=Math.floor(timeTable.fast[0]/60);
    let tmpM=timeTable.fast[0]-tmpH*60;
    let tmpTime=getTime();
    tmpTime = timeTable.fast[0] - tmpTime;
    displayTime(tmpH,tmpM,tmpTime,[timeTable.fast[1]]);
}

function displayTime(h,m,inmin,id){
    let docStop = document.getElementById("stop");
    if(checkExistBus(h,m)){
        let docTime = document.getElementById("time");
        docTime.style.fontSize = "9vh";
        docTime.style.marginTop = "4vh";
        docStop.style.marginTop = "2vh";
        if (m < 10) {
            docTime.textContent = h + ":0" + m;
        }else {
            docTime.textContent = h + ":" + m;
        }
        document.getElementById("inmin").textContent = "約" + inmin + "分後";
    }
    docStop.textContent = "乗り場:" + fromId[mode][id];
}

function checkExistBus(h,m){
    if(h===0 && m===0){
        let docTime = document.getElementById("time");
        docTime.style.fontSize = "8vh";
        docTime.style.marginTop = "6vh";
        document.getElementById("stop").style.marginTop = "4vh";
        docTime.textContent = "バス終わり";
        document.getElementById("inmin").textContent = "";
        return false
    }
    return true
}

function getTime(){
    let time = new Date();
    return time.getHours()*60+time.getMinutes();
}

document.getElementById("from").addEventListener(mytap, function(){
    let docFrom = document.getElementById("from");
    let tmpH=Math.floor(timeTable.from[from]/60);
    let tmpM=timeTable.from[from]-tmpH*60;
    let tmpTime=getTime();
    tmpTime = timeTable.from[from] - tmpTime;

    displayTime(tmpH,tmpM,tmpTime,from);
    document.getElementById("title").innerHTML = fromId[mode][from]+"発最速"

    if(mode === 0 && from === 0){
        docFrom.style.backgroundColor = "rgba(189,189,189 ,1)"
        from++;
    }else if(mode === 0 && from === 1){
        docFrom.style.backgroundColor = "rgba(158,158,158 ,1)"
        from = 0;
    }else if(mode === 1 && from < 2) {
        if(from===0){
            docFrom.style.backgroundColor = "rgba(224,224,224 ,1)"
        }else if(from===1){
            docFrom.style.backgroundColor = "rgba(189,189,189 ,1)"
        }
        from++;
    }else if(mode === 1 && from === 2){
        docFrom.style.backgroundColor = "rgba(158,158,158 ,1)"
        from = 0;
    }
}, false);
