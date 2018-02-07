var mode=true;
var mytap = window.ontouchstart===null?"touchstart":"click";

document.getElementById("go").addEventListener(mytap, function(){
    nextStep(true);
}, false);

document.getElementById("return").addEventListener(mytap, function(){
    nextStep(false);
}, false);

function nextStep(data){
    mode = data;
    document.getElementById("first").style.display = "none";
    document.getElementById("last").style.display = "block";
}
